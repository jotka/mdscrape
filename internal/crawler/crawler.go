package crawler

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"github.com/gocolly/colly/v2"

	"mdscrape/internal/converter"
	"mdscrape/internal/utils"
)

// Crawler constants
const (
	resultChannelBuffer = 1000
	requestTimeout      = 30 * time.Second
	filePermissions     = 0644
	dirPermissions      = 0755
)

// Config holds the crawler configuration
type Config struct {
	StartURL        string
	LimitURL        string
	OutputDir       string
	Threads         int
	MaxDepth        int
	ContentSelector string
	DelayMs         int
	Verbose         bool
	Quiet           bool
	DryRun          bool
	UserAgent       string
	ExcludePatterns []string
}

// Stats holds crawling statistics
type Stats struct {
	TotalDiscovered int64
	TotalDownloaded int64
	TotalBytes      int64
	TotalErrors     int64
	StartTime       time.Time
	ActiveURLs      []string
}

// PageResult represents the result of processing a page
type PageResult struct {
	URL       string
	FilePath  string
	Size      int64
	Title     string
	Success   bool
	Error     error
	Timestamp time.Time
}

// Progress callback function type
type ProgressCallback func(threadID int, url string, status string, size int64)

// Crawler handles web scraping
type Crawler struct {
	config     *Config
	collector  *colly.Collector
	converter  *converter.Converter
	stats      *Stats
	visited    sync.Map
	results    chan PageResult
	progress   ProgressCallback
	activeURLs sync.Map // Track active URLs being processed
	mu         sync.Mutex
}

// NewCrawler creates a new crawler instance
func NewCrawler(config *Config, progressCallback ProgressCallback) *Crawler {
	stats := &Stats{
		StartTime: time.Now(),
	}

	c := &Crawler{
		config:    config,
		converter: converter.NewConverter(config.ContentSelector),
		stats:     stats,
		results:   make(chan PageResult, resultChannelBuffer),
		progress:  progressCallback,
	}

	c.setupCollector()
	return c
}

func (c *Crawler) setupCollector() {
	c.collector = colly.NewCollector(
		colly.MaxDepth(c.config.MaxDepth),
		colly.Async(true),
		colly.UserAgent(c.config.UserAgent),
	)

	// Set request timeout to prevent hanging on unresponsive servers
	c.collector.SetRequestTimeout(requestTimeout)

	// Set up rate limiting
	c.collector.Limit(&colly.LimitRule{
		DomainGlob:  "*",
		Parallelism: c.config.Threads,
		Delay:       time.Duration(c.config.DelayMs) * time.Millisecond,
	})

	// Track request start
	c.collector.OnRequest(func(r *colly.Request) {
		url := r.URL.String()
		c.activeURLs.Store(url, time.Now())
	})

	// Handle HTML responses
	c.collector.OnResponse(func(r *colly.Response) {
		url := r.Request.URL.String()

		// Remove from active URLs
		defer c.activeURLs.Delete(url)

		// Check if within limit
		if !utils.IsURLWithinLimit(url, c.config.LimitURL) {
			return
		}

		// Skip if already visited (normalized)
		normalizedURL := utils.NormalizeURL(url)
		if _, loaded := c.visited.LoadOrStore(normalizedURL, true); loaded {
			return
		}

		atomic.AddInt64(&c.stats.TotalDiscovered, 1)

		// Report progress
		if c.progress != nil {
			c.progress(0, url, "downloading", int64(len(r.Body)))
		}

		// Process the page
		result := c.processPage(r)
		c.results <- result
	})

	// Handle link discovery
	c.collector.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		absLink := e.Request.AbsoluteURL(link)

		if absLink == "" {
			return
		}

		// Check if within limit
		if !utils.IsURLWithinLimit(absLink, c.config.LimitURL) {
			return
		}

		// Check exclude patterns
		if utils.MatchesPattern(absLink, c.config.ExcludePatterns) {
			return
		}

		// Check if already visited
		normalizedURL := utils.NormalizeURL(absLink)
		if _, exists := c.visited.Load(normalizedURL); exists {
			return
		}

		e.Request.Visit(absLink)
	})

	// Handle errors
	c.collector.OnError(func(r *colly.Response, err error) {
		// Ignore "already visited" errors - these are expected from Colly's internal tracking
		if err != nil && strings.Contains(err.Error(), "already visited") {
			return
		}

		var url string
		if r != nil && r.Request != nil && r.Request.URL != nil {
			url = r.Request.URL.String()
			c.activeURLs.Delete(url)
		}
		atomic.AddInt64(&c.stats.TotalErrors, 1)
		if c.progress != nil {
			c.progress(0, url, "error", 0)
		}
		// Send error result to channel
		c.results <- PageResult{
			URL:       url,
			Error:     err,
			Success:   false,
			Timestamp: time.Now(),
		}
	})
}

func (c *Crawler) processPage(r *colly.Response) PageResult {
	url := r.Request.URL.String()
	html := string(r.Body)

	result := PageResult{
		URL:       url,
		Timestamp: time.Now(),
	}

	// Get title
	result.Title = converter.ExtractTitle(html)

	// Generate file path
	filePath, err := utils.URLToFilePath(c.config.LimitURL, url)
	if err != nil {
		result.Error = fmt.Errorf("failed to generate file path: %w", err)
		return result
	}

	result.FilePath = filepath.Join(c.config.OutputDir, filePath)

	// Skip actual file writing in dry-run mode
	if c.config.DryRun {
		result.Success = true
		result.Size = int64(len(r.Body))
		atomic.AddInt64(&c.stats.TotalDownloaded, 1)
		atomic.AddInt64(&c.stats.TotalBytes, result.Size)
		return result
	}

	// Convert to Markdown
	markdown, err := c.converter.Convert(html, url, result.Title)
	if err != nil {
		result.Error = fmt.Errorf("failed to convert to markdown: %w", err)
		return result
	}

	// Ensure directory exists
	if err := utils.EnsureDir(result.FilePath); err != nil {
		result.Error = fmt.Errorf("failed to create directory: %w", err)
		return result
	}

	// Write file
	if err := os.WriteFile(result.FilePath, []byte(markdown), filePermissions); err != nil {
		result.Error = fmt.Errorf("failed to write file: %w", err)
		return result
	}

	result.Size = int64(len(markdown))
	result.Success = true

	atomic.AddInt64(&c.stats.TotalDownloaded, 1)
	atomic.AddInt64(&c.stats.TotalBytes, result.Size)

	if c.progress != nil {
		c.progress(0, url, "done", result.Size)
	}

	return result
}

// Start begins the crawling process
func (c *Crawler) Start() error {
	// Create output directory
	if !c.config.DryRun {
		if err := os.MkdirAll(c.config.OutputDir, dirPermissions); err != nil {
			return fmt.Errorf("failed to create output directory: %w", err)
		}
	}

	// Start crawling
	c.collector.Visit(c.config.StartURL)

	return nil
}

// Wait waits for all crawling to complete
func (c *Crawler) Wait() {
	c.collector.Wait()
	close(c.results)
}

// Results returns the results channel
func (c *Crawler) Results() <-chan PageResult {
	return c.results
}

// GetStats returns current statistics
func (c *Crawler) GetStats() Stats {
	// Collect active URLs
	var activeURLs []string
	c.activeURLs.Range(func(key, value interface{}) bool {
		if url, ok := key.(string); ok {
			activeURLs = append(activeURLs, url)
		}
		return true
	})

	return Stats{
		TotalDiscovered: atomic.LoadInt64(&c.stats.TotalDiscovered),
		TotalDownloaded: atomic.LoadInt64(&c.stats.TotalDownloaded),
		TotalBytes:      atomic.LoadInt64(&c.stats.TotalBytes),
		TotalErrors:     atomic.LoadInt64(&c.stats.TotalErrors),
		StartTime:       c.stats.StartTime,
		ActiveURLs:      activeURLs,
	}
}

// IsRunning checks if the crawler is still running
func (c *Crawler) IsRunning() bool {
	// Check if collector is still active
	return c.collector != nil
}
