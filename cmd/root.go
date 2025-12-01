package cmd

import (
	"fmt"
	"net/url"
	"os"
	"runtime"
	"strings"

	"github.com/spf13/cobra"

	"mdscrape/internal/crawler"
	"mdscrape/internal/ui"
)

// Version information (set via ldflags)
var (
	Version   = "dev"
	GoVersion = runtime.Version()
	Author    = "Jarek Krochmalski"
)

var (
	startURL        string
	limitURL        string
	outputDir       string
	threads         int
	depth           int
	selector        string
	delay           int
	verbose         bool
	quiet           bool
	dryRun          bool
	userAgent       string
	excludePatterns []string
	showVersion     bool
)

var rootCmd = &cobra.Command{
	Use:   "mdscrape <url>",
	Short: "Scrape websites and convert to Markdown files",
	Long: `mdscrape is a CLI tool for scraping documentation websites and converting
them to Markdown files suitable for AI agents like Claude.

Examples:
  # Scrape Docker reference docs
  mdscrape https://docs.docker.com/reference/

  # Scrape with custom output directory and fewer threads
  mdscrape https://docs.docker.com/reference/ -o ./docker-docs -t 5

  # Limit scraping to a specific section
  mdscrape https://docs.docker.com/ -l https://docs.docker.com/reference/

  # Dry run to see what would be scraped
  mdscrape https://docs.docker.com/reference/ --dry-run

  # Scrape with custom content selector
  mdscrape https://docs.docker.com/reference/ -s "article.main-content"`,
	Args: cobra.MaximumNArgs(1),
	RunE: runScrape,
	Version: Version,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.SetVersionTemplate(fmt.Sprintf(`mdscrape %s
Author: %s
Go version: %s
`, Version, Author, GoVersion))

	rootCmd.Flags().StringVarP(&startURL, "url", "u", "", "Starting URL to scrape (can also be passed as argument)")
	rootCmd.Flags().StringVarP(&limitURL, "limit", "l", "", "Limit scraping to URLs matching this prefix (defaults to start URL)")
	rootCmd.Flags().StringVarP(&outputDir, "output", "o", "", "Output directory (defaults to URL-based path)")
	rootCmd.Flags().IntVarP(&threads, "threads", "t", 10, "Number of concurrent threads")
	rootCmd.Flags().IntVarP(&depth, "depth", "d", 50, "Maximum crawl depth")
	rootCmd.Flags().StringVarP(&selector, "selector", "s", "", "CSS selector for main content (e.g., 'article', '.content')")
	rootCmd.Flags().IntVar(&delay, "delay", 100, "Delay between requests in milliseconds")
	rootCmd.Flags().BoolVarP(&verbose, "verbose", "v", false, "Verbose output")
	rootCmd.Flags().BoolVarP(&quiet, "quiet", "q", false, "Quiet mode (minimal output)")
	rootCmd.Flags().BoolVar(&dryRun, "dry-run", false, "Show what would be scraped without downloading")
	rootCmd.Flags().StringVar(&userAgent, "user-agent", "mdscrape/1.0", "User agent string")
	rootCmd.Flags().StringSliceVarP(&excludePatterns, "exclude", "e", []string{}, "URL patterns to exclude (can be specified multiple times)")
}

func runScrape(cmd *cobra.Command, args []string) error {
	// Get URL from argument or flag
	if len(args) > 0 {
		startURL = args[0]
	}

	// Check if URL is provided
	if startURL == "" {
		return fmt.Errorf("URL is required. Usage: mdscrape <url> or mdscrape --url <url>")
	}

	// Validate start URL
	parsedURL, err := url.Parse(startURL)
	if err != nil {
		return fmt.Errorf("invalid URL: %w", err)
	}
	if parsedURL.Scheme == "" || parsedURL.Host == "" {
		return fmt.Errorf("URL must include scheme and host (e.g., https://example.com)")
	}

	// Default limit URL to start URL
	if limitURL == "" {
		limitURL = startURL
	}

	// Validate limit URL
	parsedLimit, err := url.Parse(limitURL)
	if err != nil {
		return fmt.Errorf("invalid limit URL: %w", err)
	}
	if parsedLimit.Scheme == "" || parsedLimit.Host == "" {
		return fmt.Errorf("limit URL must include scheme and host")
	}

	// Generate default output directory from URL
	if outputDir == "" {
		outputDir = generateOutputDir(parsedURL)
	}

	// Create config
	config := &crawler.Config{
		StartURL:        startURL,
		LimitURL:        limitURL,
		OutputDir:       outputDir,
		Threads:         threads,
		MaxDepth:        depth,
		ContentSelector: selector,
		DelayMs:         delay,
		Verbose:         verbose,
		Quiet:           quiet,
		DryRun:          dryRun,
		UserAgent:       userAgent,
		ExcludePatterns: excludePatterns,
	}

	if dryRun {
		fmt.Println("Dry run mode - no files will be written")
		fmt.Printf("Start URL: %s\n", startURL)
		fmt.Printf("Limit URL: %s\n", limitURL)
		fmt.Printf("Output directory: %s\n", outputDir)
		fmt.Printf("Threads: %d\n", threads)
		fmt.Printf("Max depth: %d\n", depth)
		if selector != "" {
			fmt.Printf("Content selector: %s\n", selector)
		}
		fmt.Println()
	}

	// Create and run the UI with crawler
	return ui.Run(config)
}

func generateOutputDir(u *url.URL) string {
	// Create output dir from host and path
	// e.g., docs.docker.com/reference/ -> docs.docker.com_reference
	host := u.Host
	path := strings.Trim(u.Path, "/")

	if path == "" {
		return host
	}

	// Replace path separators with underscores for the base folder name
	path = strings.ReplaceAll(path, "/", "_")
	return fmt.Sprintf("%s_%s", host, path)
}
