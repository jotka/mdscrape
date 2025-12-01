---
title: "RedundantTargetPlatform"
source: "https://docs.docker.com/reference/build-checks/redundant-target-platform/"
---

# RedundantTargetPlatform

Back

Ask AI

Start typing to search or try Ask AI.

[Contact support](https://hub.docker.com/support/contact)

[Reference](https://docs.docker.com/reference/)

- [Get started](/get-started/)
- [Guides](/guides/)
- [Manuals](/manuals/)

# RedundantTargetPlatform

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
Setting platform to predefined $TARGETPLATFORM in FROM is redundant as this is the default behavior
```

## [Description](#description)

A custom platform can be used for a base image. The default platform is the same platform as the target output so setting the platform to `$TARGETPLATFORM` is redundant and unnecessary.

## [Examples](#examples)

❌ Bad: this usage of `--platform` is redundant since `$TARGETPLATFORM` is the default.

```dockerfile
FROM --platform=$TARGETPLATFORM alpine AS builder
RUN apk add --no-cache git
```

✅ Good: omit the `--platform` argument.

```dockerfile
FROM alpine AS builder
RUN apk add --no-cache git
```

[Request changes](https://github.com/docker/docs/issues/new?template=doc_issue.yml&location=https%3A%2F%2Fdocs.docker.com%2Freference%2Fbuild-checks%2Fredundant-target-platform%2F&labels=status%2Ftriage)

Table of contents
