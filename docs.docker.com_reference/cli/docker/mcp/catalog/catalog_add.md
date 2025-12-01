---
title: "docker mcp catalog add"
source: "https://docs.docker.com/reference/cli/docker/mcp/catalog/catalog_add/"
---

# docker mcp catalog add

Back

Ask AI

Start typing to search or try Ask AI.

[Contact support](https://hub.docker.com/support/contact)

[Reference](https://docs.docker.com/reference/)

- [Get started](/get-started/)
- [Guides](/guides/)
- [Manuals](/manuals/)

# docker mcp catalog add

DescriptionAdd a server to a catalogUsage`docker mcp catalog add <catalog> <server-name> <catalog-file>`

## [Description](#description)

Add an MCP server definition to an existing catalog by copying it from another catalog file. The server definition includes all configuration, tools, and metadata.

## [Options](#options)

OptionDefaultDescription`--force`Overwrite existing server in the catalog

## [Examples](#examples)

# [Add a server from another catalog file](#add-a-server-from-another-catalog-file)

docker mcp catalog add my-catalog github-server ./github-catalog.yaml

# [Add with force to overwrite existing server](#add-with-force-to-overwrite-existing-server)

docker mcp catalog add my-catalog slack-bot ./team-catalog.yaml --force

Table of contents
