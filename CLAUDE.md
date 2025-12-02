# Claude instructions for mdscrape

## Project overview

mdscrape is a Go CLI application that scrapes documentation websites and converts HTML pages to Markdown files. It's designed for collecting documentation for AI agents.

## Project structure

```
mdscrape/
├── main.go                      # Entry point
├── cmd/
│   └── root.go                  # Cobra CLI setup, flags, validation
├── internal/
│   ├── crawler/
│   │   └── crawler.go           # Colly-based web crawler
│   ├── converter/
│   │   └── markdown.go          # HTML to Markdown conversion
│   ├── ui/
│   │   └── progress.go          # Bubble Tea progress UI
│   └── utils/
│       └── paths.go             # URL to filepath conversion
├── go.mod
└── go.sum
```

## Key dependencies

- `github.com/spf13/cobra` - CLI framework
- `github.com/gocolly/colly/v2` - Web scraping with concurrency
- `github.com/PuerkitoBio/goquery` - HTML parsing (jQuery-like)
- `github.com/JohannesKaufmann/html-to-markdown/v2` - HTML to Markdown
- `github.com/charmbracelet/bubbletea` - Terminal UI
- `github.com/charmbracelet/bubbles` - UI components (progress, spinner)

## Build and run

```bash
# Build
go build -o mdscrape .

# Run
./mdscrape -u https://docs.example.com/ -l https://docs.example.com/

# Test with quiet mode
./mdscrape -u https://httpbin.org/html -q
```

## Architecture notes

### Crawler (internal/crawler/crawler.go)
- Uses Colly for concurrent HTTP requests
- Respects depth limits and URL prefix filtering
- Tracks visited URLs to avoid duplicates
- Reports progress via callback function

### Converter (internal/converter/markdown.go)
- Removes unwanted elements (nav, header, footer, scripts)
- Auto-detects main content area using common CSS selectors
- Adds YAML frontmatter with title and source URL
- Uses html-to-markdown library for conversion

### UI (internal/ui/progress.go)
- Bubble Tea model with progress bar and spinner
- Shows real-time stats: files downloaded, size, errors, time
- Supports quiet mode for scripting

### Utils (internal/utils/paths.go)
- Converts URLs to filesystem paths
- Maintains folder hierarchy from URL structure
- Handles path sanitization for filesystem safety

## Common modifications

### Add new content selector
Edit `internal/converter/markdown.go`, add to `contentSelectors` slice in `extractMainContent()`.

### Add new elements to remove
Edit `internal/converter/markdown.go`, add to `selectorsToRemove` slice in `removeUnwantedElements()`.

### Change default settings
Edit `cmd/root.go`, modify flag defaults in `init()`.

### Add new CLI flag
1. Add variable in `cmd/root.go`
2. Add flag in `init()` function
3. Use in `runScrape()` function
4. Pass to crawler config

## Git commits

Do not include Co-Authored-By or "Generated with Claude Code" lines in commit messages.
- do not release until explicitly asked