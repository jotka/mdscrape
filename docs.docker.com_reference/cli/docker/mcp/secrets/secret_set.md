---
title: "docker mcp secret set"
source: "https://docs.docker.com/reference/cli/docker/mcp/secrets/secret_set/"
---

# docker mcp secret set

Back

Ask AI

Start typing to search or try Ask AI.

[Contact support](https://hub.docker.com/support/contact)

[Reference](https://docs.docker.com/reference/)

- [Get started](/get-started/)
- [Guides](/guides/)
- [Manuals](/manuals/)

# docker mcp secret set

DescriptionSet a secret in Docker Desktop's secret storeUsage`docker mcp secret set key[=value]`

## [Description](#description)

Set a secret in Docker Desktop's secret store

## [Options](#options)

OptionDefaultDescription`--provider`Supported: credstore, oauth/

## [Examples](#examples)

### [Use secrets for postgres password with default policy](#use-secrets-for-postgres-password-with-default-policy)

```console
docker mcp secret set POSTGRES_PASSWORD=my-secret-password
docker run -d -l x-secret:POSTGRES_PASSWORD=/pwd.txt -e POSTGRES_PASSWORD_FILE=/pwd.txt -p 5432 postgres
```

### [Pass the secret via STDIN](#pass-the-secret-via-stdin)

```console
echo my-secret-password > pwd.txt
cat pwd.txt | docker mcp secret set POSTGRES_PASSWORD
```

Table of contents
