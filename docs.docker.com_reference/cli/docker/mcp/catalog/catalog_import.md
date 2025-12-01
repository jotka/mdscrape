---
title: "docker mcp catalog import"
source: "https://docs.docker.com/reference/cli/docker/mcp/catalog/catalog_import/"
---

# docker mcp catalog import

Back

Ask AI

Start typing to search or try Ask AI.

[Contact support](https://hub.docker.com/support/contact)

[Reference](https://docs.docker.com/reference/)

- [Get started](/get-started/)
- [Guides](/guides/)
- [Manuals](/manuals/)

# docker mcp catalog import

DescriptionImport a catalog from URL or fileUsage`docker mcp catalog import <alias|url|file>`

## [Description](#description)

Import an MCP server catalog from a URL or local file. The catalog will be downloaded and stored locally for use with the MCP gateway.

When --mcp-registry flag is used, the argument must be an existing catalog name, and the command will import servers from the MCP registry URL into that catalog.

## [Options](#options)

OptionDefaultDescription`--dry-run`Show Imported Data but do not update the Catalog`--mcp-registry`Import server from MCP registry URL into existing catalog

## [Examples](#examples)

# [Import from URL](#import-from-url)

docker mcp catalog import [https://example.com/my-catalog.yaml](https://example.com/my-catalog.yaml)

# [Import from local file](#import-from-local-file)

docker mcp catalog import ./shared-catalog.yaml

# [Import from MCP registry URL into existing catalog](#import-from-mcp-registry-url-into-existing-catalog)

docker mcp catalog import my-catalog --mcp-registry [https://registry.example.com/server](https://registry.example.com/server)

Table of contents
