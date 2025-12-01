---
title: "docker buildx use"
source: "https://docs.docker.com/reference/cli/docker/buildx/use/"
---

# docker buildx use

Back

Ask AI

Start typing to search or try Ask AI.

[Contact support](https://hub.docker.com/support/contact)

[Reference](https://docs.docker.com/reference/)

- [Get started](/get-started/)
- [Guides](/guides/)
- [Manuals](/manuals/)

# docker buildx use

DescriptionSet the current builder instanceUsage`docker buildx use [OPTIONS] NAME`

## [Description](#description)

Switches the current builder instance. Build commands invoked after this command will run on a specified builder. Alternatively, a context name can be used to switch to the default builder of that context.

## [Options](#options)

OptionDefaultDescription`--default`Set builder as default for current context`--global`Builder persists context changes

## [Examples](#examples)

### [Override the configured builder instance (--builder)](#builder)

Same as [`buildx --builder`](/reference/cli/docker/buildx/#builder).

Table of contents
