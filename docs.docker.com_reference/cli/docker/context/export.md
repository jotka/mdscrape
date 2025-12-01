---
title: "docker context export"
source: "https://docs.docker.com/reference/cli/docker/context/export/"
---

# docker context export

Back

Ask AI

Start typing to search or try Ask AI.

[Contact support](https://hub.docker.com/support/contact)

[Reference](https://docs.docker.com/reference/)

- [Get started](/get-started/)
- [Guides](/guides/)
- [Manuals](/manuals/)

# docker context export

DescriptionExport a context to a tar archive FILE or a tar stream on STDOUT.Usage`docker context export [OPTIONS] CONTEXT [FILE|-]`

## [Description](#description)

Exports a context to a file that can then be used with `docker context import`.

The default output filename is `<CONTEXT>.dockercontext`. To export to `STDOUT`, use `-` as filename, for example:

```console
$ docker context export my-context -
```

Table of contents
