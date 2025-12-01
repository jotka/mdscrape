---
title: "NoEmptyContinuation"
source: "https://docs.docker.com/reference/build-checks/no-empty-continuation/"
---

# NoEmptyContinuation

Back

Ask AI

Start typing to search or try Ask AI.

[Contact support](https://hub.docker.com/support/contact)

[Reference](https://docs.docker.com/reference/)

- [Get started](/get-started/)
- [Guides](/guides/)
- [Manuals](/manuals/)

# NoEmptyContinuation

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
Empty continuation line found in: RUN apk add     gnupg     curl
```

## [Description](#description)

Support for empty continuation (`/`) lines have been deprecated and will generate errors in future versions of the Dockerfile syntax.

Empty continuation lines are empty lines following a newline escape:

```dockerfile
FROM alpine
RUN apk add \

    gnupg \

    curl
```

Support for such empty lines is deprecated, and a future BuildKit release will remove support for this syntax entirely, causing builds to break. To avoid future errors, remove the empty lines, or add comments, since lines with comments aren't considered empty.

## [Examples](#examples)

❌ Bad: empty continuation line between `EXPOSE` and 80.

```dockerfile
FROM alpine
EXPOSE \

80
```

✅ Good: comments do not count as empty lines.

```dockerfile
FROM alpine
EXPOSE \
# Port
80
```

[Request changes](https://github.com/docker/docs/issues/new?template=doc_issue.yml&location=https%3A%2F%2Fdocs.docker.com%2Freference%2Fbuild-checks%2Fno-empty-continuation%2F&labels=status%2Ftriage)

Table of contents
