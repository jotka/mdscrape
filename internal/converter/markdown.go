package converter

import (
	"bytes"
	"fmt"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
	htmltomarkdown "github.com/JohannesKaufmann/html-to-markdown/v2"
)

// Converter handles HTML to Markdown conversion
type Converter struct {
	contentSelector string
}

// NewConverter creates a new HTML to Markdown converter
func NewConverter(contentSelector string) *Converter {
	return &Converter{
		contentSelector: contentSelector,
	}
}

// Convert converts HTML content to Markdown
func (c *Converter) Convert(html string, pageURL string, title string) (string, error) {
	// Parse HTML
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		return "", fmt.Errorf("failed to parse HTML: %w", err)
	}

	// Remove unwanted elements
	c.removeUnwantedElements(doc)

	// Extract title if not provided
	if title == "" {
		title = doc.Find("title").First().Text()
		title = strings.TrimSpace(title)
	}

	// Select content area
	var contentHTML string
	if c.contentSelector != "" {
		selection := doc.Find(c.contentSelector).First()
		if selection.Length() > 0 {
			contentHTML, _ = selection.Html()
		}
	}

	// Fallback to common content selectors
	if contentHTML == "" {
		contentHTML = c.extractMainContent(doc)
	}

	// If still empty, use body
	if contentHTML == "" {
		contentHTML, _ = doc.Find("body").Html()
	}

	if contentHTML == "" {
		contentHTML = html
	}

	// Convert to Markdown using the package function
	markdown, err := htmltomarkdown.ConvertString(contentHTML)
	if err != nil {
		return "", fmt.Errorf("failed to convert HTML to markdown: %w", err)
	}

	// Post-process markdown
	markdown = c.postProcess(markdown)

	// Build final document with frontmatter
	var buf bytes.Buffer
	buf.WriteString("---\n")
	buf.WriteString("title: \"" + escapeYAML(title) + "\"\n")
	buf.WriteString("source: \"" + pageURL + "\"\n")
	buf.WriteString("---\n\n")

	if title != "" {
		buf.WriteString("# " + title + "\n\n")
	}

	buf.WriteString(markdown)

	return buf.String(), nil
}

// removeUnwantedElements removes navigation, footer, and other non-content elements
func (c *Converter) removeUnwantedElements(doc *goquery.Document) {
	// Common selectors for elements to remove
	selectorsToRemove := []string{
		"nav",
		"header",
		"footer",
		".nav",
		".navbar",
		".navigation",
		".sidebar",
		".side-nav",
		".toc",
		".table-of-contents",
		".breadcrumb",
		".breadcrumbs",
		".pagination",
		".pager",
		".comments",
		".comment-section",
		".social-share",
		".share-buttons",
		".advertisement",
		".ads",
		".ad",
		".cookie-banner",
		".cookie-notice",
		".modal",
		".popup",
		"script",
		"style",
		"noscript",
		"iframe",
		"[role='navigation']",
		"[role='banner']",
		"[role='contentinfo']",
		"[aria-label='breadcrumb']",
	}

	for _, selector := range selectorsToRemove {
		doc.Find(selector).Remove()
	}
}

// extractMainContent tries to find the main content area
func (c *Converter) extractMainContent(doc *goquery.Document) string {
	// Try common content selectors in order of preference
	contentSelectors := []string{
		"main",
		"article",
		"[role='main']",
		".content",
		".main-content",
		".post-content",
		".article-content",
		".entry-content",
		".page-content",
		".markdown-body",
		".documentation",
		".docs-content",
		"#content",
		"#main",
		"#main-content",
	}

	for _, selector := range contentSelectors {
		selection := doc.Find(selector).First()
		if selection.Length() > 0 {
			html, _ := selection.Html()
			if strings.TrimSpace(html) != "" {
				return html
			}
		}
	}

	return ""
}

// postProcess cleans up the converted markdown
func (c *Converter) postProcess(markdown string) string {
	// Remove excessive blank lines (more than 2 in a row)
	re := regexp.MustCompile(`\n{3,}`)
	markdown = re.ReplaceAllString(markdown, "\n\n")

	// Clean up whitespace at beginning and end
	markdown = strings.TrimSpace(markdown)

	// Ensure single newline at end
	markdown = markdown + "\n"

	return markdown
}

// escapeYAML escapes special characters for YAML frontmatter
func escapeYAML(s string) string {
	s = strings.ReplaceAll(s, "\\", "\\\\")
	s = strings.ReplaceAll(s, "\"", "\\\"")
	s = strings.ReplaceAll(s, "\n", " ")
	s = strings.TrimSpace(s)
	return s
}

// ExtractTitle extracts the page title from HTML
func ExtractTitle(html string) string {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		return ""
	}

	// Try various title sources
	if title := doc.Find("h1").First().Text(); title != "" {
		return strings.TrimSpace(title)
	}

	if title := doc.Find("title").First().Text(); title != "" {
		return strings.TrimSpace(title)
	}

	if title, exists := doc.Find("meta[property='og:title']").Attr("content"); exists {
		return strings.TrimSpace(title)
	}

	return ""
}

// ExtractLinks extracts all links from HTML content
func ExtractLinks(html string, baseURL string) []string {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		return nil
	}

	var links []string
	seen := make(map[string]bool)

	doc.Find("a[href]").Each(func(i int, s *goquery.Selection) {
		href, exists := s.Attr("href")
		if !exists || href == "" {
			return
		}

		// Skip anchors, javascript, mailto, etc.
		if strings.HasPrefix(href, "#") ||
			strings.HasPrefix(href, "javascript:") ||
			strings.HasPrefix(href, "mailto:") ||
			strings.HasPrefix(href, "tel:") {
			return
		}

		// Resolve relative URLs
		resolvedURL := resolveURL(baseURL, href)
		if resolvedURL == "" {
			return
		}

		// Deduplicate
		if !seen[resolvedURL] {
			seen[resolvedURL] = true
			links = append(links, resolvedURL)
		}
	})

	return links
}

// resolveURL resolves a potentially relative URL against a base URL
func resolveURL(baseURL, href string) string {
	// Already absolute
	if strings.HasPrefix(href, "http://") || strings.HasPrefix(href, "https://") {
		return href
	}

	// Parse base URL
	if !strings.HasPrefix(baseURL, "http") {
		return ""
	}

	// Simple resolution for common cases
	if strings.HasPrefix(href, "//") {
		// Protocol-relative URL
		if strings.HasPrefix(baseURL, "https://") {
			return "https:" + href
		}
		return "http:" + href
	}

	// Extract base components
	parts := strings.SplitN(baseURL, "://", 2)
	if len(parts) != 2 {
		return ""
	}
	scheme := parts[0]
	rest := parts[1]

	// Find host
	hostEnd := strings.Index(rest, "/")
	var host, basePath string
	if hostEnd == -1 {
		host = rest
		basePath = "/"
	} else {
		host = rest[:hostEnd]
		basePath = rest[hostEnd:]
	}

	if strings.HasPrefix(href, "/") {
		// Absolute path
		return scheme + "://" + host + href
	}

	// Relative path - resolve against base path
	// Remove filename from base path
	lastSlash := strings.LastIndex(basePath, "/")
	if lastSlash != -1 {
		basePath = basePath[:lastSlash+1]
	}

	return scheme + "://" + host + basePath + href
}
