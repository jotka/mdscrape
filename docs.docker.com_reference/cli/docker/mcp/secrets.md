---
title: "docker mcp secret"
source: "https://docs.docker.com/reference/cli/docker/mcp/secrets/"
---

# docker mcp secret

Back

Ask AI

Start typing to search or try Ask AI.

[Contact support](https://hub.docker.com/support/contact)

[Reference](https://docs.docker.com/reference/)

- [Get started](/get-started/)
- [Guides](/guides/)
- [Manuals](/manuals/)

# docker mcp secret

DescriptionManage secrets

## [Description](#description)

Manage secrets

## [Examples](#examples)

### [Use secrets for postgres password with default policy](#use-secrets-for-postgres-password-with-default-policy)

> docker mcp secret set POSTGRES\_PASSWORD=my-secret-password docker run -d -l x-secret:POSTGRES\_PASSWORD=/pwd.txt -e POSTGRES\_PASSWORD\_FILE=/pwd.txt -p 5432 postgres

### [Pass the secret via STDIN](#pass-the-secret-via-stdin)

> echo my-secret-password &gt; pwd.txt cat pwd.txt | docker mcp secret set POSTGRES\_PASSWORD

## [Subcommands](#subcommands)

CommandDescription[`docker mcp secret export`](https://docs.docker.com/reference/cli/docker/mcp/secrets/secret_export/)Export secrets for the specified servers[`docker mcp secret ls`](https://docs.docker.com/reference/cli/docker/mcp/secrets/secret_ls/)List all secret names in Docker Desktop's secret store[`docker mcp secret rm`](https://docs.docker.com/reference/cli/docker/mcp/secrets/secret_rm/)Remove secrets from Docker Desktop's secret store[`docker mcp secret set`](https://docs.docker.com/reference/cli/docker/mcp/secrets/secret_set/)Set a secret in Docker Desktop's secret store

Table of contents
