---
title: "ExposeInvalidFormat"
source: "https://docs.docker.com/reference/build-checks/expose-invalid-format/"
---

# ExposeInvalidFormat

Back

Ask AI

Start typing to search or try Ask AI.

[Contact support](https://hub.docker.com/support/contact)

[Reference](https://docs.docker.com/reference/)

- [Get started](/get-started/)
- [Guides](/guides/)
- [Manuals](/manuals/)

# ExposeInvalidFormat

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
EXPOSE instruction should not define an IP address or host-port mapping, found '127.0.0.1:80:80'
```

## [Description](#description)

The [`EXPOSE`](https://docs.docker.com/reference/dockerfile/#expose) instruction in a Dockerfile is used to indicate which ports the container listens on at runtime. It should not include an IP address or host-port mapping, as this is not the intended use of the `EXPOSE` instruction. Instead, it should only specify the port number and optionally the protocol (TCP or UDP).

> Important
> 
> This will become an error in a future release.

## [Examples](#examples)

❌ Bad: IP address and host-port mapping used.

```dockerfile
FROM alpine
EXPOSE 127.0.0.1:80:80
```

✅ Good: only the port number is specified.

```dockerfile
FROM alpine
EXPOSE 80
```

❌ Bad: Host-port mapping used.

```dockerfile
FROM alpine
EXPOSE 80:80
```

✅ Good: only the port number is specified.

```dockerfile
FROM alpine
EXPOSE 80
```

[Request changes](https://github.com/docker/docs/issues/new?template=doc_issue.yml&location=https%3A%2F%2Fdocs.docker.com%2Freference%2Fbuild-checks%2Fexpose-invalid-format%2F&labels=status%2Ftriage)

Table of contents
