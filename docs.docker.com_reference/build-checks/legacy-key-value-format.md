---
title: "LegacyKeyValueFormat"
source: "https://docs.docker.com/reference/build-checks/legacy-key-value-format/"
---

# LegacyKeyValueFormat

Back

Ask AI

Start typing to search or try Ask AI.

[Contact support](https://hub.docker.com/support/contact)

[Reference](https://docs.docker.com/reference/)

- [Get started](/get-started/)
- [Guides](/guides/)
- [Manuals](/manuals/)

# LegacyKeyValueFormat

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
"ENV key=value" should be used instead of legacy "ENV key value" format
```

## [Description](#description)

The correct format for declaring environment variables and build arguments in a Dockerfile is `ENV key=value` and `ARG key=value`, where the variable name (`key`) and value (`value`) are separated by an equals sign (`=`). Historically, Dockerfiles have also supported a space separator between the key and the value (for example, `ARG key value`). This legacy format is deprecated, and you should only use the format with the equals sign.

## [Examples](#examples)

❌ Bad: using a space separator for variable key and value.

```dockerfile
FROM alpine
ARG foo bar
```

✅ Good: use an equals sign to separate key and value.

```dockerfile
FROM alpine
ARG foo=bar
```

❌ Bad: multi-line variable declaration with a space separator.

```dockerfile
ENV DEPS \
    curl \
    git \
    make
```

✅ Good: use an equals sign and wrap the value in quotes.

```dockerfile
ENV DEPS="\
    curl \
    git \
    make"
```

[Request changes](https://github.com/docker/docs/issues/new?template=doc_issue.yml&location=https%3A%2F%2Fdocs.docker.com%2Freference%2Fbuild-checks%2Flegacy-key-value-format%2F&labels=status%2Ftriage)

Table of contents
