---
title: "Interface: Host"
source: "https://docs.docker.com/reference/api/extensions-sdk/Host/"
---

# Interface: Host

Back

Ask AI

Start typing to search or try Ask AI.

[Contact support](https://hub.docker.com/support/contact)

[Reference](https://docs.docker.com/reference/)

- [Get started](/get-started/)
- [Guides](/guides/)
- [Manuals](/manuals/)

# Interface: Host

Page options

Copy page as Markdown for LLMs

View page as plain text

Ask questions with Docs AI

Claude

Open in Claude

Table of contents

* * *

**`Since`**

0.2.0

## [Methods](#methods)

### [openExternal](#openexternal)

▸ **openExternal**(`url`): `void`

Opens an external URL with the system default browser.

**`Since`**

0.2.0

```typescript
ddClient.host.openExternal("https://docker.com");
```

#### [Parameters](#parameters)

NameTypeDescription`url``string`The URL the browser will open (must have the protocol `http` or `https`).

#### [Returns](#returns)

`void`

## [Properties](#properties)

### [platform](#platform)

• **platform**: `string`

Returns a string identifying the operating system platform. See [https://nodejs.org/api/os.html#osplatform](https://nodejs.org/api/os.html#osplatform)

**`Since`**

0.2.2

* * *

### [arch](#arch)

• **arch**: `string`

Returns the operating system CPU architecture. See [https://nodejs.org/api/os.html#osarch](https://nodejs.org/api/os.html#osarch)

**`Since`**

0.2.2

* * *

### [hostname](#hostname)

• **hostname**: `string`

Returns the host name of the operating system. See [https://nodejs.org/api/os.html#oshostname](https://nodejs.org/api/os.html#oshostname)

**`Since`**

0.2.2

[Edit this page](https://github.com/docker/docs/edit/main/content/reference/api/extensions-sdk/Host.md)

[Request changes](https://github.com/docker/docs/issues/new?template=doc_issue.yml&location=https%3A%2F%2Fdocs.docker.com%2Freference%2Fapi%2Fextensions-sdk%2FHost%2F&labels=status%2Ftriage)

Table of contents
