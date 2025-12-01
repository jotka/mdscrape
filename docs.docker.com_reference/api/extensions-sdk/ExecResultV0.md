---
title: "Interface: ExecResultV0"
source: "https://docs.docker.com/reference/api/extensions-sdk/ExecResultV0/"
---

# Interface: ExecResultV0

Back

Ask AI

Start typing to search or try Ask AI.

[Contact support](https://hub.docker.com/support/contact)

[Reference](https://docs.docker.com/reference/)

- [Get started](/get-started/)
- [Guides](/guides/)
- [Manuals](/manuals/)

# Interface: ExecResultV0

Page options

Copy page as Markdown for LLMs

View page as plain text

Ask questions with Docs AI

Claude

Open in Claude

Table of contents

* * *

## [Properties](#properties)

### [cmd](#cmd)

• `Optional` `Readonly` **cmd**: `string`

* * *

### [killed](#killed)

• `Optional` `Readonly` **killed**: `boolean`

* * *

### [signal](#signal)

• `Optional` `Readonly` **signal**: `string`

* * *

### [code](#code)

• `Optional` `Readonly` **code**: `number`

* * *

### [stdout](#stdout)

• `Readonly` **stdout**: `string`

* * *

### [stderr](#stderr)

• `Readonly` **stderr**: `string`

## [Methods](#methods)

### [lines](#lines)

▸ **lines**(): `string`\[]

Split output lines.

#### [Returns](#returns)

`string`\[]

The list of lines.

* * *

### [parseJsonLines](#parsejsonlines)

▸ **parseJsonLines**(): `any`\[]

Parse each output line as a JSON object.

#### [Returns](#returns-1)

`any`\[]

The list of lines where each line is a JSON object.

* * *

### [parseJsonObject](#parsejsonobject)

▸ **parseJsonObject**(): `any`

Parse a well-formed JSON output.

#### [Returns](#returns-2)

`any`

The JSON object.

[Edit this page](https://github.com/docker/docs/edit/main/content/reference/api/extensions-sdk/ExecResultV0.md)

[Request changes](https://github.com/docker/docs/issues/new?template=doc_issue.yml&location=https%3A%2F%2Fdocs.docker.com%2Freference%2Fapi%2Fextensions-sdk%2FExecResultV0%2F&labels=status%2Ftriage)

Table of contents
