---
title: "DuplicateStageName"
source: "https://docs.docker.com/reference/build-checks/duplicate-stage-name/"
---

# DuplicateStageName

Back

Ask AI

Start typing to search or try Ask AI.

[Contact support](https://hub.docker.com/support/contact)

[Reference](https://docs.docker.com/reference/)

- [Get started](/get-started/)
- [Guides](/guides/)
- [Manuals](/manuals/)

# DuplicateStageName

Page options

Copy page as Markdown for LLMs

View page as plain text

Ask questions with Docs AI

Claude

Open in Claude

Table of contents

* * *

## [Output](#output)

```text
Duplicate stage name 'foo-base', stage names should be unique
```

## [Description](#description)

Defining multiple stages with the same name results in an error because the builder is unable to uniquely resolve the stage name reference.

## [Examples](#examples)

❌ Bad: `builder` is declared as a stage name twice.

```dockerfile
FROM debian:latest AS builder
RUN apt-get update; apt-get install -y curl

FROM golang:latest AS builder
```

✅ Good: stages have unique names.

```dockerfile
FROM debian:latest AS deb-builder
RUN apt-get update; apt-get install -y curl

FROM golang:latest AS go-builder
```

[Request changes](https://github.com/docker/docs/issues/new?template=doc_issue.yml&location=https%3A%2F%2Fdocs.docker.com%2Freference%2Fbuild-checks%2Fduplicate-stage-name%2F&labels=status%2Ftriage)

Table of contents
