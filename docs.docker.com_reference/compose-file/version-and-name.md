---
title: "Version and name top-level elements"
source: "https://docs.docker.com/reference/compose-file/version-and-name/"
---

# Version and name top-level elements

Back

Ask AI

Start typing to search or try Ask AI.

[Contact support](https://hub.docker.com/support/contact)

[Reference](https://docs.docker.com/reference/)

- [Get started](/get-started/)
- [Guides](/guides/)
- [Manuals](/manuals/)

# Version and name top-level elements

Page options

Copy page as Markdown for LLMs

View page as plain text

Ask questions with Docs AI

Claude

Open in Claude

Table of contents

* * *

## [Version top-level element (obsolete)](#version-top-level-element-obsolete)

> Important
> 
> The top-level `version` property is defined by the Compose Specification for backward compatibility. It is only informative and you'll receive a warning message that it is obsolete if used.

Compose always uses the most recent schema to validate the Compose file, regardless of the `version` field.

Compose validates whether it can fully parse the Compose file. If some fields are unknown, typically because the Compose file was written with fields defined by a newer version of the Specification, you'll receive a warning message.

## [Name top-level element](#name-top-level-element)

The top-level `name` property is defined by the Compose Specification as the project name to be used if you don't set one explicitly.

Compose offers a way for you to override this name, and sets a default project name to be used if the top-level `name` element is not set.

Whenever a project name is defined by top-level `name` or by some custom mechanism, it is exposed for [interpolation](https://docs.docker.com/reference/compose-file/interpolation/) and environment variable resolution as `COMPOSE_PROJECT_NAME`

```yml
name: myapp

services:
  foo:
    image: busybox
    command: echo "I'm running ${COMPOSE_PROJECT_NAME}"
```

For more information on other ways to name Compose projects, see [Specify a project name](https://docs.docker.com/compose/how-tos/project-name/).

[Edit this page](https://github.com/docker/docs/edit/main/content/reference/compose-file/version-and-name.md)

[Request changes](https://github.com/docker/docs/issues/new?template=doc_issue.yml&location=https%3A%2F%2Fdocs.docker.com%2Freference%2Fcompose-file%2Fversion-and-name%2F&labels=status%2Ftriage)

Table of contents
