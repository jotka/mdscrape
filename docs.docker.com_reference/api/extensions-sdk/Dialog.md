---
title: "Interface: Dialog"
source: "https://docs.docker.com/reference/api/extensions-sdk/Dialog/"
---

# Interface: Dialog

Back

Ask AI

Start typing to search or try Ask AI.

[Contact support](https://hub.docker.com/support/contact)

[Reference](https://docs.docker.com/reference/)

- [Get started](/get-started/)
- [Guides](/guides/)
- [Manuals](/manuals/)

# Interface: Dialog

Page options

Copy page as Markdown for LLMs

View page as plain text

Ask questions with Docs AI

Claude

Open in Claude

Table of contents

* * *

Allows opening native dialog boxes.

**`Since`**

0.2.3

## [Methods](#methods)

### [showOpenDialog](#showopendialog)

▸ **showOpenDialog**(`dialogProperties`): `Promise`&lt;[`OpenDialogResult`](https://docs.docker.com/reference/api/extensions-sdk/OpenDialogResult/)&gt;

Display a native open dialog. Lets you select a file or a folder.

```typescript
ddClient.desktopUI.dialog.showOpenDialog({properties: ['openFile']});
```

#### [Parameters](#parameters)

NameTypeDescription`dialogProperties``any`Properties to specify the open dialog behaviour, see [https://www.electronjs.org/docs/latest/api/dialog#dialogshowopendialogbrowserwindow-options](https://www.electronjs.org/docs/latest/api/dialog#dialogshowopendialogbrowserwindow-options).

#### [Returns](#returns)

`Promise`&lt;[`OpenDialogResult`](https://docs.docker.com/reference/api/extensions-sdk/OpenDialogResult/)&gt;

[Edit this page](https://github.com/docker/docs/edit/main/content/reference/api/extensions-sdk/Dialog.md)

[Request changes](https://github.com/docker/docs/issues/new?template=doc_issue.yml&location=https%3A%2F%2Fdocs.docker.com%2Freference%2Fapi%2Fextensions-sdk%2FDialog%2F&labels=status%2Ftriage)

Table of contents
