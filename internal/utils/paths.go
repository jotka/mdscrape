package utils

import (
	"fmt"
	"net/url"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

// URLToFilePath converts a URL to a local file path relative to the base URL
// For example:
//   - baseURL: https://docs.docker.com/reference/
//   - pageURL: https://docs.docker.com/reference/cli/docker-run/
//   - result: cli/docker-run.md
func URLToFilePath(baseURL, pageURL string) (string, error) {
	base, err := url.Parse(baseURL)
	if err != nil {
		return "", err
	}

	page, err := url.Parse(pageURL)
	if err != nil {
		return "", err
	}

	// Get the path relative to the base
	basePath := strings.TrimSuffix(base.Path, "/")
	pagePath := strings.TrimSuffix(page.Path, "/")

	// Calculate relative path
	var relativePath string
	if strings.HasPrefix(pagePath, basePath) {
		relativePath = strings.TrimPrefix(pagePath, basePath)
		relativePath = strings.TrimPrefix(relativePath, "/")
	} else {
		// If not under base path, use full path
		relativePath = strings.TrimPrefix(pagePath, "/")
	}

	// Handle empty path (root of the base URL)
	if relativePath == "" {
		relativePath = "index"
	}

	// Clean up the path for filesystem use
	relativePath = sanitizeFilePath(relativePath)

	// Add .md extension
	if !strings.HasSuffix(relativePath, ".md") {
		relativePath = relativePath + ".md"
	}

	return relativePath, nil
}

// sanitizeFilePath cleans a path for safe filesystem use
func sanitizeFilePath(path string) string {
	// Replace characters that are problematic in filenames
	replacer := strings.NewReplacer(
		":", "_",
		"*", "_",
		"?", "_",
		"\"", "_",
		"<", "_",
		">", "_",
		"|", "_",
	)
	path = replacer.Replace(path)

	// Remove any double slashes
	for strings.Contains(path, "//") {
		path = strings.ReplaceAll(path, "//", "/")
	}

	return path
}

// EnsureDir creates all directories in the path if they don't exist
func EnsureDir(filePath string) error {
	dir := filepath.Dir(filePath)
	return os.MkdirAll(dir, 0755)
}

// IsURLWithinLimit checks if a URL is within the limit URL prefix
func IsURLWithinLimit(targetURL, limitURL string) bool {
	// Normalize URLs
	target := strings.TrimSuffix(targetURL, "/")
	limit := strings.TrimSuffix(limitURL, "/")

	return strings.HasPrefix(target, limit) || target == limit
}

// NormalizeURL normalizes a URL for comparison and deduplication
func NormalizeURL(rawURL string) string {
	u, err := url.Parse(rawURL)
	if err != nil {
		return rawURL
	}

	// Remove fragment
	u.Fragment = ""

	// Remove trailing slash from path
	u.Path = strings.TrimSuffix(u.Path, "/")

	// Sort query parameters for consistent comparison
	// For simplicity, we'll just remove query params for docs
	u.RawQuery = ""

	return u.String()
}

// ExtractFilename extracts a suitable filename from a URL
func ExtractFilename(rawURL string) string {
	u, err := url.Parse(rawURL)
	if err != nil {
		return "unknown"
	}

	path := strings.TrimSuffix(u.Path, "/")
	if path == "" {
		return "index"
	}

	// Get the last segment
	segments := strings.Split(path, "/")
	filename := segments[len(segments)-1]

	if filename == "" {
		filename = "index"
	}

	return sanitizeFilePath(filename)
}

// MatchesPattern checks if a URL matches any of the exclude patterns
func MatchesPattern(urlStr string, patterns []string) bool {
	for _, pattern := range patterns {
		matched, err := regexp.MatchString(pattern, urlStr)
		if err == nil && matched {
			return true
		}
	}
	return false
}

// ValidatePatterns checks if all exclude patterns are valid regex
func ValidatePatterns(patterns []string) error {
	for _, pattern := range patterns {
		_, err := regexp.Compile(pattern)
		if err != nil {
			return fmt.Errorf("invalid exclude pattern %q: %w", pattern, err)
		}
	}
	return nil
}
