---
title: "Interface: Toast"
source: "https://docs.docker.com/reference/api/extensions-sdk/Toast/"
---

# Interface: Toast

Back

Ask AI

Start typing to search or try Ask AI.

[Contact support](https://hub.docker.com/support/contact)

[Reference](https://docs.docker.com/reference/)

- [Get started](/get-started/)
- [Guides](/guides/)
- [Manuals](/manuals/)

# Interface: Toast

Page options

Copy page as Markdown for LLMs

View page as plain text

Ask questions with Docs AI

Claude

Open in Claude

Table of contents

* * *

Toasts provide a brief notification to the user. They appear temporarily and shouldn't interrupt the user experience. They also don't require user input to disappear.

**`Since`**

0.2.0

## [Methods](#methods)

### [success](#success)

▸ **success**(`msg`): `void`

Display a toast message of type success.

```typescript
ddClient.desktopUI.toast.success("message");
```

#### [Parameters](#parameters)

NameTypeDescription`msg``string`The message to display in the toast.

#### [Returns](#returns)

`void`

* * *

### [warning](#warning)

▸ **warning**(`msg`): `void`

Display a toast message of type warning.

```typescript
ddClient.desktopUI.toast.warning("message");
```

#### [Parameters](#parameters-1)

NameTypeDescription`msg``string`The message to display in the warning.

#### [Returns](#returns-1)

`void`

* * *

### [error](#error)

▸ **error**(`msg`): `void`

Display a toast message of type error.

```typescript
ddClient.desktopUI.toast.error("message");
```

#### [Parameters](#parameters-2)

NameTypeDescription`msg``string`The message to display in the toast.

#### [Returns](#returns-2)

`void`

[Edit this page](https://github.com/docker/docs/edit/main/content/reference/api/extensions-sdk/Toast.md)

[Request changes](https://github.com/docker/docs/issues/new?template=doc_issue.yml&location=https%3A%2F%2Fdocs.docker.com%2Freference%2Fapi%2Fextensions-sdk%2FToast%2F&labels=status%2Ftriage)

Table of contents
