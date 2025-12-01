package ui

import (
	"fmt"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/charmbracelet/bubbles/progress"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
	"golang.org/x/term"

	"mdscrape/internal/crawler"
)

var (
	titleStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("170"))

	infoStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("241"))

	successStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("82"))

	errorStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("196"))

	urlStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("117"))

	statsStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("229"))

	spinnerStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("205"))

	threadNumStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("243")).
			Width(4)

	activeURLStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("117"))

	tableHeaderStyle = lipgloss.NewStyle().
				Bold(true).
				Foreground(lipgloss.Color("229"))

	tableCellStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("252"))

	tableBorderStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("238"))
)

// ThreadStatus represents the status of a single thread
type ThreadStatus struct {
	URL     string
	Status  string
	Size    int64
	Spinner spinner.Model
}

// Model is the Bubble Tea model for the progress UI
type Model struct {
	config       *crawler.Config
	crawler      *crawler.Crawler
	threads      []ThreadStatus
	progress     progress.Model
	stats        crawler.Stats
	done         bool
	err          error
	width        int
	mu           sync.RWMutex
	recentFiles  []string
	errorURLs    []string
	maxRecent    int
	quitting     bool
}

// Message types
type tickMsg time.Time
type crawlDoneMsg struct{}
type crawlErrMsg struct{ err error }
type statsUpdateMsg crawler.Stats
type threadUpdateMsg struct {
	threadID int
	url      string
	status   string
	size     int64
}
type resultMsg crawler.PageResult

// NewModel creates a new UI model
func NewModel(config *crawler.Config) *Model {
	// Create spinners for each thread
	threads := make([]ThreadStatus, config.Threads)
	for i := range threads {
		s := spinner.New()
		s.Spinner = spinner.Dot
		s.Style = spinnerStyle

		threads[i] = ThreadStatus{
			Spinner: s,
			Status:  "idle",
		}
	}

	p := progress.New(
		progress.WithDefaultGradient(),
		progress.WithWidth(50),
	)

	return &Model{
		config:      config,
		threads:     threads,
		progress:    p,
		maxRecent:   5,
		recentFiles: make([]string, 0, 5),
	}
}

func (m *Model) Init() tea.Cmd {
	// Create progress callback
	progressCallback := func(threadID int, url string, status string, size int64) {
		// This will be called from the crawler goroutine
	}

	// Create and start crawler
	m.crawler = crawler.NewCrawler(m.config, progressCallback)

	// Start crawler in background and wait for completion
	go func() {
		if err := m.crawler.Start(); err != nil {
			// Handle error
		}
		m.crawler.Wait() // This closes the results channel
	}()

	// Start processing results
	go m.processResults()

	// Initialize spinners
	var cmds []tea.Cmd
	for i := range m.threads {
		cmds = append(cmds, m.threads[i].Spinner.Tick)
	}

	// Add tick command for updates
	cmds = append(cmds, tickCmd())

	return tea.Batch(cmds...)
}

func (m *Model) processResults() {
	for result := range m.crawler.Results() {
		// Update stats
		m.mu.Lock()
		if result.Success {
			// Add to recent files
			shortPath := result.FilePath
			if len(shortPath) > 50 {
				shortPath = "..." + shortPath[len(shortPath)-47:]
			}
			m.recentFiles = append(m.recentFiles, shortPath)
			if len(m.recentFiles) > m.maxRecent {
				m.recentFiles = m.recentFiles[1:]
			}
		} else if result.Error != nil {
			// Track error URLs
			errMsg := fmt.Sprintf("%s: %v", result.URL, result.Error)
			m.errorURLs = append(m.errorURLs, errMsg)
		}
		m.mu.Unlock()
	}

	// Signal completion
	m.done = true
}

func tickCmd() tea.Cmd {
	return tea.Tick(100*time.Millisecond, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			m.quitting = true
			return m, tea.Quit
		}

	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.progress.Width = msg.Width - 20
		if m.progress.Width > 80 {
			m.progress.Width = 80
		}

	case tickMsg:
		if m.done {
			m.quitting = true
			return m, tea.Quit
		}

		// Update stats from crawler
		if m.crawler != nil {
			m.stats = m.crawler.GetStats()
		}

		// Update spinners
		var cmds []tea.Cmd
		for i := range m.threads {
			var cmd tea.Cmd
			m.threads[i].Spinner, cmd = m.threads[i].Spinner.Update(msg)
			cmds = append(cmds, cmd)
		}
		cmds = append(cmds, tickCmd())

		return m, tea.Batch(cmds...)

	case spinner.TickMsg:
		var cmds []tea.Cmd
		for i := range m.threads {
			var cmd tea.Cmd
			m.threads[i].Spinner, cmd = m.threads[i].Spinner.Update(msg)
			cmds = append(cmds, cmd)
		}
		return m, tea.Batch(cmds...)
	}

	return m, nil
}

func (m *Model) View() string {
	if m.quitting {
		return "" // Final stats will be printed after TUI exits
	}

	var b strings.Builder

	// Title
	b.WriteString(titleStyle.Render("mdscrape"))
	b.WriteString(" - ")
	b.WriteString(infoStyle.Render("Web to Markdown scraper"))
	b.WriteString("\n\n")

	// Configuration
	b.WriteString(infoStyle.Render(fmt.Sprintf("URL: %s", m.config.StartURL)))
	b.WriteString("\n")
	b.WriteString(infoStyle.Render(fmt.Sprintf("Output: %s", m.config.OutputDir)))
	b.WriteString("\n\n")

	// Progress bar
	var progressPercent float64
	if m.stats.TotalDiscovered > 0 {
		progressPercent = float64(m.stats.TotalDownloaded) / float64(m.stats.TotalDiscovered)
		if progressPercent > 1 {
			progressPercent = 1
		}
	}
	b.WriteString(m.progress.ViewAs(progressPercent))
	b.WriteString("\n\n")

	// Statistics
	elapsed := time.Since(m.stats.StartTime).Round(time.Second)
	b.WriteString(statsStyle.Render(fmt.Sprintf(
		"Files: %d/%d | Size: %s | Errors: %d | Time: %s",
		m.stats.TotalDownloaded,
		m.stats.TotalDiscovered,
		formatBytes(m.stats.TotalBytes),
		m.stats.TotalErrors,
		elapsed,
	)))
	b.WriteString("\n\n")

	// Recent files
	m.mu.RLock()
	if len(m.recentFiles) > 0 {
		b.WriteString(infoStyle.Render("Recent files:"))
		b.WriteString("\n")
		for _, f := range m.recentFiles {
			b.WriteString(successStyle.Render("  ✓ "))
			b.WriteString(urlStyle.Render(f))
			b.WriteString("\n")
		}
	}
	m.mu.RUnlock()

	// Active threads
	if !m.done && len(m.stats.ActiveURLs) > 0 {
		b.WriteString(infoStyle.Render("Active downloads:"))
		b.WriteString("\n")

		// Show up to config.Threads active URLs
		maxShow := m.config.Threads
		if maxShow > 10 {
			maxShow = 10 // Cap display at 10 for readability
		}

		for i, url := range m.stats.ActiveURLs {
			if i >= maxShow {
				remaining := len(m.stats.ActiveURLs) - maxShow
				if remaining > 0 {
					b.WriteString(infoStyle.Render(fmt.Sprintf("  ... and %d more\n", remaining)))
				}
				break
			}

			// Get spinner for this thread
			threadIdx := i % len(m.threads)
			spinnerView := m.threads[threadIdx].Spinner.View()

			// Truncate URL for display
			displayURL := truncateURL(url, 60)

			b.WriteString(fmt.Sprintf("  %s %s %s\n",
				threadNumStyle.Render(fmt.Sprintf("[%d]", i+1)),
				spinnerView,
				activeURLStyle.Render(displayURL),
			))
		}
		b.WriteString("\n")
	}

	// Status
	if m.done {
		b.WriteString(successStyle.Render("✓ Scraping complete!"))
		b.WriteString("\n")
	} else if len(m.stats.ActiveURLs) == 0 {
		b.WriteString(m.threads[0].Spinner.View())
		b.WriteString(" ")
		if m.stats.TotalDownloaded > 0 && m.stats.TotalDownloaded >= m.stats.TotalDiscovered {
			b.WriteString(infoStyle.Render("Finishing up..."))
		} else if m.stats.TotalDiscovered == 0 {
			b.WriteString(infoStyle.Render("Connecting..."))
		} else {
			b.WriteString(infoStyle.Render("Discovering pages..."))
		}
		b.WriteString("\n")
	}

	b.WriteString("\n")
	b.WriteString(infoStyle.Render("Press q to quit"))

	return b.String()
}

func (m *Model) renderFinalStats() string {
	var b strings.Builder

	elapsed := time.Since(m.stats.StartTime).Round(time.Second)

	b.WriteString("\n")
	b.WriteString(successStyle.Render("✓ Scraping complete!"))
	b.WriteString("\n\n")

	// List errors if any (before the summary table)
	m.mu.RLock()
	if len(m.errorURLs) > 0 {
		b.WriteString(errorStyle.Render("Errors:"))
		b.WriteString("\n")
		maxErrors := 10
		for i, errURL := range m.errorURLs {
			if i >= maxErrors {
				b.WriteString(infoStyle.Render(fmt.Sprintf("  ... and %d more errors\n", len(m.errorURLs)-maxErrors)))
				break
			}
			// Truncate long error messages
			if len(errURL) > 80 {
				errURL = errURL[:77] + "..."
			}
			b.WriteString(fmt.Sprintf("  %s %s\n", errorStyle.Render("✗"), infoStyle.Render(errURL)))
		}
		b.WriteString("\n")
	}
	m.mu.RUnlock()

	// Build table rows
	rows := [][]string{
		{"Files", fmt.Sprintf("%d", m.stats.TotalDownloaded)},
		{"Size", formatBytes(m.stats.TotalBytes)},
		{"Time", elapsed.String()},
	}

	if m.stats.TotalDownloaded > 0 && elapsed.Seconds() > 0 {
		rate := float64(m.stats.TotalDownloaded) / elapsed.Seconds()
		rows = append(rows, []string{"Rate", fmt.Sprintf("%.1f files/sec", rate)})
	}

	if m.stats.TotalErrors > 0 {
		rows = append(rows, []string{"Errors", fmt.Sprintf("%d", m.stats.TotalErrors)})
	}

	// Create table
	t := table.New().
		Border(lipgloss.RoundedBorder()).
		BorderStyle(tableBorderStyle).
		StyleFunc(func(row, col int) lipgloss.Style {
			if col == 0 {
				return tableHeaderStyle
			}
			return tableCellStyle
		}).
		Rows(rows...)

	b.WriteString(t.String())
	b.WriteString("\n")

	b.WriteString(fmt.Sprintf("\n%s %s\n", infoStyle.Render("Output:"), urlStyle.Render(m.config.OutputDir)))

	b.WriteString("\n")
	return b.String()
}

func formatBytes(bytes int64) string {
	const unit = 1024
	if bytes < unit {
		return fmt.Sprintf("%d B", bytes)
	}
	div, exp := int64(unit), 0
	for n := bytes / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %cB", float64(bytes)/float64(div), "KMGTPE"[exp])
}


func truncateURL(url string, maxLen int) string {
	if len(url) <= maxLen {
		return url
	}

	// Try to show the path part
	if idx := strings.Index(url, "://"); idx != -1 {
		afterScheme := url[idx+3:]
		if slashIdx := strings.Index(afterScheme, "/"); slashIdx != -1 {
			path := afterScheme[slashIdx:]
			if len(path) > maxLen-3 {
				return "..." + path[len(path)-(maxLen-3):]
			}
			return "..." + path
		}
	}

	return "..." + url[len(url)-(maxLen-3):]
}

// Run starts the UI and crawler
func Run(config *crawler.Config) error {
	// For quiet mode or non-TTY, run without UI
	if config.Quiet || !term.IsTerminal(int(os.Stdout.Fd())) {
		return runQuiet(config)
	}

	model := NewModel(config)
	p := tea.NewProgram(model)

	finalModel, err := p.Run()
	if err != nil {
		// Fall back to quiet mode if TUI fails
		return runQuiet(config)
	}

	// Print final summary after TUI exits
	if m, ok := finalModel.(*Model); ok {
		fmt.Print(m.renderFinalStats())
	}

	return nil
}

// runQuiet runs the crawler without the fancy UI
func runQuiet(config *crawler.Config) error {
	fmt.Printf("Starting scrape of %s\n", config.StartURL)
	fmt.Printf("Output directory: %s\n", config.OutputDir)

	progressCallback := func(threadID int, url string, status string, size int64) {
		if config.Verbose {
			fmt.Printf("[%s] %s (%s)\n", status, url, formatBytes(size))
		}
	}

	c := crawler.NewCrawler(config, progressCallback)

	if err := c.Start(); err != nil {
		return err
	}

	// Process results
	go func() {
		for result := range c.Results() {
			if result.Error != nil {
				fmt.Printf("Error: %s - %v\n", result.URL, result.Error)
			} else if config.Verbose {
				fmt.Printf("Saved: %s\n", result.FilePath)
			}
		}
	}()

	c.Wait()

	stats := c.GetStats()
	fmt.Printf("\nComplete! Downloaded %d files (%s)\n",
		stats.TotalDownloaded, formatBytes(stats.TotalBytes))

	return nil
}
