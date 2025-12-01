# mdscrape

[![Build](https://github.com/jotka/mdscrape/actions/workflows/build.yml/badge.svg)](https://github.com/jotka/mdscrape/actions/workflows/build.yml)
[![Release](https://github.com/jotka/mdscrape/actions/workflows/release.yml/badge.svg)](https://github.com/jotka/mdscrape/actions/workflows/release.yml)

A fast, concurrent CLI tool for scraping documentation websites and converting them to clean Markdown files. Perfect for building knowledge bases for AI agents.

![mdscrape in action](screenshots/SCR-20251201-hyus.png)

## Features

- **Fast concurrent scraping** - configurable thread pool (default: 10 threads)
- **Smart content extraction** - automatically finds main content, removes nav/footer/ads
- **Preserves formatting** - code blocks with syntax hints, tables, lists, inline code, links, headings
- **Mirror folder structure** - `/docs/api/auth/` becomes `docs/api/auth.md`
- **YAML frontmatter** - includes title and source URL for reference
- **Beautiful progress UI** - real-time stats, per-thread activity display (auto-detects TTY)
- **Flexible filtering** - limit by URL prefix, exclude patterns, set max depth

## Installation

### From source

```bash
git clone https://github.com/yourusername/mdscrape.git
cd mdscrape
go build -o mdscrape .
```

## Quick start

```bash
# Scrape Next.js documentation
mdscrape https://nextjs.org/docs/

# Scrape Docker reference with custom output folder
mdscrape https://docs.docker.com/reference/ -o ./docker-docs

# Limit depth and threads for gentler scraping
mdscrape https://react.dev/reference/ -d 5 -t 3

# Preview what would be scraped (dry run)
mdscrape https://docs.python.org/3/library/ --dry-run
```

## Usage

```
mdscrape <url> [options]
```

### Options

| Option | Short | Default | Description |
|--------|-------|---------|-------------|
| `<url>` | | required | Starting URL to scrape (positional argument) |
| `--limit` | `-l` | start URL | Only scrape URLs with this prefix |
| `--output` | `-o` | auto | Output directory |
| `--threads` | `-t` | 10 | Concurrent download threads |
| `--depth` | `-d` | 50 | Maximum link depth to follow |
| `--selector` | `-s` | auto | CSS selector for content (e.g., `article`) |
| `--exclude` | `-e` | none | URL patterns to skip (repeatable) |
| `--delay` | | 100 | Milliseconds between requests |
| `--dry-run` | | false | Show plan without downloading |
| `--quiet` | `-q` | false | Minimal output, no progress UI |
| `--verbose` | `-v` | false | Show every file downloaded |

### Examples

```bash
# Scrape only the API section
mdscrape https://docs.example.com/api/

# Limit to specific subsection
mdscrape https://docs.example.com/ -l https://docs.example.com/api/

# Use specific content selector
mdscrape https://docs.example.com/ -s "main.docs-content"

# Exclude changelog and blog
mdscrape https://docs.example.com/ -e "/changelog" -e "/blog"

# Gentle scraping with delays
mdscrape https://docs.example.com/ -t 2 --delay 500
```

## Output

### File structure

URLs are converted to a matching folder hierarchy:

```
https://docs.docker.com/reference/cli/docker/run/
                                ↓
output/cli/docker/run.md
```

```
output/
├── index.md
├── getting-started.md
├── cli/
│   └── docker/
│       ├── build.md
│       ├── run.md
│       └── compose/
│           ├── up.md
│           └── down.md
├── api/
│   └── engine/
│       └── index.md
└── reference/
    └── dockerfile.md
```

### Markdown format

Each page includes YAML frontmatter:

```markdown
---
title: "docker run"
source: "https://docs.docker.com/reference/cli/docker/run/"
---

# docker run

Run a command in a new container.

## Usage

\`\`\`bash
docker run [OPTIONS] IMAGE [COMMAND] [ARG...]
\`\`\`

## Options

| Option | Description |
|--------|-------------|
| `-d` | Run in detached mode |
| `-p` | Publish port |
...
```

## How it works

1. **Crawl** - Discovers pages using [Colly](https://github.com/gocolly/colly), respecting depth limits and URL filters
2. **Extract** - Parses HTML with [GoQuery](https://github.com/PuerkitoBio/goquery), removes boilerplate (nav, footer, scripts)
3. **Convert** - Transforms to Markdown using [html-to-markdown](https://github.com/JohannesKaufmann/html-to-markdown)
4. **Save** - Writes files with folder structure matching the URL path

### Smart content detection

The scraper automatically tries these selectors to find main content:

```
main, article, [role="main"], .content, .main-content,
.post-content, .article-content, .markdown-body, .docs-content
```

Override with `--selector` if needed.

### Elements removed

Navigation, headers, footers, sidebars, ads, scripts, and other non-content elements are stripped:

```
nav, header, footer, .sidebar, .toc, .breadcrumb,
.pagination, .comments, .advertisement, script, style
```

## Use cases

- **AI context** - Build knowledge bases for Claude, GPT, or other LLMs
- **Offline docs** - Read documentation without internet
- **Documentation migration** - Convert sites to Markdown for static site generators
- **Archiving** - Preserve documentation snapshots

## Building AI agents with scraped docs

The scraped markdown files are ideal for creating specialized AI coding agents. This approach is often more effective than using MCP (Model Context Protocol) servers or real-time web fetching, which can waste context window space on navigation elements, ads, and irrelevant content with each request. Pre-scraped markdown files are clean, deduplicated, and always available - the AI can reference them instantly without network latency or token overhead from fetching raw HTML.

For example, to build a "Next.js expert" agent:

```bash
# Scrape Next.js documentation
mdscrape https://nextjs.org/docs/ -o nextjs-docs

# Create a specialized agent in .claude/agents/
mkdir -p .claude/agents
cat > .claude/agents/nextjs-expert.md << 'EOF'
You are an expert Next.js developer. When answering questions about Next.js,
refer to the documentation in the `nextjs-docs/` folder for accurate,
up-to-date information about APIs, patterns, and best practices.
EOF
```

Now you can invoke this agent in Claude Code with `/nextjs-expert`. The AI will have access to the complete, current documentation as context, enabling more accurate and framework-specific responses. The markdown format preserves code examples, API references, and structural information that helps the AI understand and apply the documentation correctly.

You can combine multiple documentation sources to create specialized agents:

```bash
mdscrape https://nextjs.org/docs/ -o docs/nextjs
mdscrape https://tailwindcss.com/docs/ -o docs/tailwind
mdscrape https://www.prisma.io/docs/ -o docs/prisma
```

Then reference all three in your project instructions to create a full-stack expert agent.

## License

MIT

## Examples

Folder structure mirroring the documentation site:

![Folder structure](screenshots/SCR-20251201-iaqk.png)

Markdown output with preserved code blocks and formatting:

![Markdown output](screenshots/SCR-20251201-ibct.png)

Using scraped docs with Claude Code agents:

![Claude Code example](screenshots/SCR-20251201-puof.png)
