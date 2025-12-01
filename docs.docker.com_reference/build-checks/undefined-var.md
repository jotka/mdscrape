---
title: "UndefinedVar"
source: "https://docs.docker.com/reference/build-checks/undefined-var/"
---

# UndefinedVar

Back

Ask AI

Start typing to search or try Ask AI.

[Contact support](https://hub.docker.com/support/contact)

[Reference](https://docs.docker.com/reference/)

- [Get started](/get-started/)
- [Guides](/guides/)
- [Manuals](/manuals/)

# UndefinedVar

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
Usage of undefined variable '$foo'
```

## [Description](#description)

This check ensures that environment variables and build arguments are correctly declared before being used. While undeclared variables might not cause an immediate build failure, they can lead to unexpected behavior or errors later in the build process.

This check does not evaluate undefined variables for `RUN`, `CMD`, and `ENTRYPOINT` instructions where you use the [shell form](https://docs.docker.com/reference/dockerfile/#shell-form). That's because when you use shell form, variables are resolved by the command shell.

It also detects common mistakes like typos in variable names. For example, in the following Dockerfile:

```dockerfile
FROM alpine
ENV PATH=$PAHT:/app/bin
```

The check identifies that `$PAHT` is undefined and likely a typo for `$PATH`:

```text
Usage of undefined variable '$PAHT' (did you mean $PATH?)
```

## [Examples](#examples)

❌ Bad: `$foo` is an undefined build argument.

```dockerfile
FROM alpine AS base
COPY $foo .
```

✅ Good: declaring `foo` as a build argument before attempting to access it.

```dockerfile
FROM alpine AS base
ARG foo
COPY $foo .
```

❌ Bad: `$foo` is undefined.

```dockerfile
FROM alpine AS base
ARG VERSION=$foo
```

✅ Good: the base image defines `$PYTHON_VERSION`

```dockerfile
FROM python AS base
ARG VERSION=$PYTHON_VERSION
```

[Request changes](https://github.com/docker/docs/issues/new?template=doc_issue.yml&location=https%3A%2F%2Fdocs.docker.com%2Freference%2Fbuild-checks%2Fundefined-var%2F&labels=status%2Ftriage)

Table of contents
