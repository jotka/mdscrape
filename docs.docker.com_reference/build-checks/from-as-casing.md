---
title: "FromAsCasing"
source: "https://docs.docker.com/reference/build-checks/from-as-casing/"
---

# FromAsCasing

Back

Ask AI

Start typing to search or try Ask AI.

[Contact support](https://hub.docker.com/support/contact)

[Reference](https://docs.docker.com/reference/)

- [Get started](/get-started/)
- [Guides](/guides/)
- [Manuals](/manuals/)

# FromAsCasing

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
'as' and 'FROM' keywords' casing do not match
```

## [Description](#description)

While Dockerfile keywords can be either uppercase or lowercase, mixing case styles is not recommended for readability. This rule reports violations where mixed case style occurs for a `FROM` instruction with an `AS` keyword declaring a stage name.

## [Examples](#examples)

❌ Bad: `FROM` is uppercase, `AS` is lowercase.

```dockerfile
FROM debian:latest as builder
```

✅ Good: `FROM` and `AS` are both uppercase

```dockerfile
FROM debian:latest AS deb-builder
```

✅ Good: `FROM` and `AS` are both lowercase.

```dockerfile
from debian:latest as deb-builder
```

## [Related errors](#related-errors)

- [`FileConsistentCommandCasing`](https://docs.docker.com/reference/build-checks/consistent-instruction-casing/)

[Request changes](https://github.com/docker/docs/issues/new?template=doc_issue.yml&location=https%3A%2F%2Fdocs.docker.com%2Freference%2Fbuild-checks%2Ffrom-as-casing%2F&labels=status%2Ftriage)

Table of contents
