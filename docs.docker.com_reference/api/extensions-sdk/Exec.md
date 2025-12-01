---
title: "Interface: Exec"
source: "https://docs.docker.com/reference/api/extensions-sdk/Exec/"
---

# Interface: Exec

Back

Ask AI

Start typing to search or try Ask AI.

[Contact support](https://hub.docker.com/support/contact)

[Reference](https://docs.docker.com/reference/)

- [Get started](/get-started/)
- [Guides](/guides/)
- [Manuals](/manuals/)

# Interface: Exec

Page options

Copy page as Markdown for LLMs

View page as plain text

Ask questions with Docs AI

Claude

Open in Claude

Table of contents

* * *

## [Callable](#callable)

### [Exec](#exec)

▸ **Exec**(`cmd`, `args`, `options?`): `Promise`&lt;[`ExecResult`](https://docs.docker.com/reference/api/extensions-sdk/ExecResult/)&gt;

Executes a command.

**`Since`**

0.2.0

#### [Parameters](#parameters)

NameTypeDescription`cmd``string`The command to execute.`args``string`\[]The arguments of the command to execute.`options?`[`ExecOptions`](https://docs.docker.com/reference/api/extensions-sdk/ExecOptions/)The list of options.

#### [Returns](#returns)

`Promise`&lt;[`ExecResult`](https://docs.docker.com/reference/api/extensions-sdk/ExecResult/)&gt;

A promise that will resolve once the command finishes.

### [Exec](#exec-1)

▸ **Exec**(`cmd`, `args`, `options`): [`ExecProcess`](https://docs.docker.com/reference/api/extensions-sdk/ExecProcess/)

Streams the result of a command if `stream` is specified in the `options` parameter.

Specify the `stream` if the output of your command is too long or if you need to stream things indefinitely (for example container logs).

**`Since`**

0.2.2

#### [Parameters](#parameters-1)

NameTypeDescription`cmd``string`The command to execute.`args``string`\[]The arguments of the command to execute.`options`[`SpawnOptions`](https://docs.docker.com/reference/api/extensions-sdk/SpawnOptions/)The list of options.

#### [Returns](#returns-1)

[`ExecProcess`](https://docs.docker.com/reference/api/extensions-sdk/ExecProcess/)

The spawned process.

[Edit this page](https://github.com/docker/docs/edit/main/content/reference/api/extensions-sdk/Exec.md)

[Request changes](https://github.com/docker/docs/issues/new?template=doc_issue.yml&location=https%3A%2F%2Fdocs.docker.com%2Freference%2Fapi%2Fextensions-sdk%2FExec%2F&labels=status%2Ftriage)

Table of contents
