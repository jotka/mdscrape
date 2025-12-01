---
title: "ExposeProtoCasing"
source: "https://docs.docker.com/reference/build-checks/expose-proto-casing/"
---

# ExposeProtoCasing

Back

Ask AI

Start typing to search or try Ask AI.

[Contact support](https://hub.docker.com/support/contact)

[Reference](https://docs.docker.com/reference/)

- [Get started](/get-started/)
- [Guides](/guides/)
- [Manuals](/manuals/)

# ExposeProtoCasing

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
Defined protocol '80/TcP' in EXPOSE instruction should be lowercase
```

## [Description](#description)

Protocol names in the [`EXPOSE`](https://docs.docker.com/reference/dockerfile/#expose) instruction should be specified in lowercase to maintain consistency and readability. This rule checks for protocols that are not in lowercase and reports them.

## [Examples](#examples)

❌ Bad: protocol is not in lowercase.

```dockerfile
FROM alpine
EXPOSE 80/TcP
```

✅ Good: protocol is in lowercase.

```dockerfile
FROM alpine
EXPOSE 80/tcp
```

[Request changes](https://github.com/docker/docs/issues/new?template=doc_issue.yml&location=https%3A%2F%2Fdocs.docker.com%2Freference%2Fbuild-checks%2Fexpose-proto-casing%2F&labels=status%2Ftriage)

Table of contents
