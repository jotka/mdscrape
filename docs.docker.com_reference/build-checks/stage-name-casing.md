---
title: "StageNameCasing"
source: "https://docs.docker.com/reference/build-checks/stage-name-casing/"
---

# StageNameCasing

Back

Ask AI

Start typing to search or try Ask AI.

[Contact support](https://hub.docker.com/support/contact)

[Reference](https://docs.docker.com/reference/)

- [Get started](/get-started/)
- [Guides](/guides/)
- [Manuals](/manuals/)

# StageNameCasing

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
Stage name 'BuilderBase' should be lowercase
```

## [Description](#description)

To help distinguish Dockerfile instruction keywords from identifiers, this rule forces names of stages in a multi-stage Dockerfile to be all lowercase.

## [Examples](#examples)

❌ Bad: mixing uppercase and lowercase characters in the stage name.

```dockerfile
FROM alpine AS BuilderBase
```

✅ Good: stage name is all in lowercase.

```dockerfile
FROM alpine AS builder-base
```

[Request changes](https://github.com/docker/docs/issues/new?template=doc_issue.yml&location=https%3A%2F%2Fdocs.docker.com%2Freference%2Fbuild-checks%2Fstage-name-casing%2F&labels=status%2Ftriage)

Table of contents
