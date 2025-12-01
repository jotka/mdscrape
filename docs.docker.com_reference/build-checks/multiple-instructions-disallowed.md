---
title: "MultipleInstructionsDisallowed"
source: "https://docs.docker.com/reference/build-checks/multiple-instructions-disallowed/"
---

# MultipleInstructionsDisallowed

Back

Ask AI

Start typing to search or try Ask AI.

[Contact support](https://hub.docker.com/support/contact)

[Reference](https://docs.docker.com/reference/)

- [Get started](/get-started/)
- [Guides](/guides/)
- [Manuals](/manuals/)

# MultipleInstructionsDisallowed

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
Multiple CMD instructions should not be used in the same stage because only the last one will be used
```

## [Description](#description)

If you have multiple `CMD`, `HEALTHCHECK`, or `ENTRYPOINT` instructions in your Dockerfile, only the last occurrence is used. An image can only ever have one `CMD`, `HEALTHCHECK`, and `ENTRYPOINT`.

## [Examples](#examples)

❌ Bad: Duplicate instructions.

```dockerfile
FROM alpine
ENTRYPOINT ["echo", "Hello, Norway!"]
ENTRYPOINT ["echo", "Hello, Sweden!"]
# Only "Hello, Sweden!" will be printed
```

✅ Good: only one `ENTRYPOINT` instruction.

```dockerfile
FROM alpine
ENTRYPOINT ["echo", "Hello, Norway!\nHello, Sweden!"]
```

You can have both a regular, top-level `CMD` and a separate `CMD` for a `HEALTHCHECK` instruction.

✅ Good: only one top-level `CMD` instruction.

```dockerfile
FROM python:alpine
RUN apk add curl
HEALTHCHECK --interval=1s --timeout=3s \
  CMD ["curl", "-f", "http://localhost:8080"]
CMD ["python", "-m", "http.server", "8080"]
```

[Request changes](https://github.com/docker/docs/issues/new?template=doc_issue.yml&location=https%3A%2F%2Fdocs.docker.com%2Freference%2Fbuild-checks%2Fmultiple-instructions-disallowed%2F&labels=status%2Ftriage)

Table of contents
