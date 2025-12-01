---
title: "Interface: ExecStreamOptions"
source: "https://docs.docker.com/reference/api/extensions-sdk/ExecStreamOptions/"
---

# Interface: ExecStreamOptions

Back

Ask AI

Start typing to search or try Ask AI.

[Contact support](https://hub.docker.com/support/contact)

[Reference](https://docs.docker.com/reference/)

- [Get started](/get-started/)
- [Guides](/guides/)
- [Manuals](/manuals/)

# Interface: ExecStreamOptions

Page options

Copy page as Markdown for LLMs

View page as plain text

Ask questions with Docs AI

Claude

Open in Claude

Table of contents

* * *

**`Since`**

0.2.2

## [Properties](#properties)

### [onOutput](#onoutput)

• `Optional` **onOutput**: (`data`: { `stdout`: `string` ; `stderr?`: `undefined` } | { `stdout?`: `undefined` ; `stderr`: `string` }) =&gt; `void`

#### [Type declaration](#type-declaration)

▸ (`data`): `void`

Invoked when receiving output from command execution. By default, the output is split into chunks at arbitrary boundaries. If you prefer the output to be split into complete lines, set `splitOutputLines` to true. The callback is then invoked once for each line.

**`Since`**

0.2.0

##### [Parameters](#parameters)

NameTypeDescription`data``{ stdout: string; stderr?: undefined } | { stdout?: undefined; stderr: string }`Output content. Can include either stdout string, or stderr string, one at a time.

##### [Returns](#returns)

`void`

* * *

### [onError](#onerror)

• `Optional` **onError**: (`error`: `any`) =&gt; `void`

#### [Type declaration](#type-declaration-1)

▸ (`error`): `void`

Invoked to report error if the executed command errors.

##### [Parameters](#parameters-1)

NameTypeDescription`error``any`The error happening in the executed command

##### [Returns](#returns-1)

`void`

* * *

### [onClose](#onclose)

• `Optional` **onClose**: (`exitCode`: `number`) =&gt; `void`

#### [Type declaration](#type-declaration-2)

▸ (`exitCode`): `void`

Invoked when process exits.

##### [Parameters](#parameters-2)

NameTypeDescription`exitCode``number`The process exit code

##### [Returns](#returns-2)

`void`

* * *

### [splitOutputLines](#splitoutputlines)

• `Optional` `Readonly` **splitOutputLines**: `boolean`

Specifies the behaviour invoking `onOutput(data)`. Raw output by default, splitting output at any position. If set to true, `onOutput` will be invoked once for each line.

[Edit this page](https://github.com/docker/docs/edit/main/content/reference/api/extensions-sdk/ExecStreamOptions.md)

[Request changes](https://github.com/docker/docs/issues/new?template=doc_issue.yml&location=https%3A%2F%2Fdocs.docker.com%2Freference%2Fapi%2Fextensions-sdk%2FExecStreamOptions%2F&labels=status%2Ftriage)

Table of contents
