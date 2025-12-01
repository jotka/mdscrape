---
title: "ReservedStageName"
source: "https://docs.docker.com/reference/build-checks/reserved-stage-name/"
---

# ReservedStageName

Back

Ask AI

Start typing to search or try Ask AI.

[Contact support](https://hub.docker.com/support/contact)

[Reference](https://docs.docker.com/reference/)

- [Get started](/get-started/)
- [Guides](/guides/)
- [Manuals](/manuals/)

# ReservedStageName

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
'scratch' is reserved and should not be used as a stage name
```

## [Description](#description)

Reserved words should not be used as names for stages in multi-stage builds. The reserved words are:

- `context`
- `scratch`

## [Examples](#examples)

❌ Bad: `scratch` and `context` are reserved names.

```dockerfile
FROM alpine AS scratch
FROM alpine AS context
```

✅ Good: the stage name `builder` is not reserved.

```dockerfile
FROM alpine AS builder
```

[Request changes](https://github.com/docker/docs/issues/new?template=doc_issue.yml&location=https%3A%2F%2Fdocs.docker.com%2Freference%2Fbuild-checks%2Freserved-stage-name%2F&labels=status%2Ftriage)

Table of contents
