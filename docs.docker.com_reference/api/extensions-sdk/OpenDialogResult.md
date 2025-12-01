---
title: "Interface: OpenDialogResult"
source: "https://docs.docker.com/reference/api/extensions-sdk/OpenDialogResult/"
---

# Interface: OpenDialogResult

Back

Ask AI

Start typing to search or try Ask AI.

[Contact support](https://hub.docker.com/support/contact)

[Reference](https://docs.docker.com/reference/)

- [Get started](/get-started/)
- [Guides](/guides/)
- [Manuals](/manuals/)

# Interface: OpenDialogResult

Page options

Copy page as Markdown for LLMs

View page as plain text

Ask questions with Docs AI

Claude

Open in Claude

Table of contents

* * *

**`Since`**

0.2.3

## [Properties](#properties)

### [canceled](#canceled)

• `Readonly` **canceled**: `boolean`

Whether the dialog was canceled.

* * *

### [filePaths](#filepaths)

• `Readonly` **filePaths**: `string`\[]

An array of file paths chosen by the user. If the dialog is cancelled this will be an empty array.

* * *

### [bookmarks](#bookmarks)

• `Optional` `Readonly` **bookmarks**: `string`\[]

macOS only. An array matching the `filePaths` array of `base64` encoded strings which contains security scoped bookmark data. `securityScopedBookmarks` must be enabled for this to be populated.

[Edit this page](https://github.com/docker/docs/edit/main/content/reference/api/extensions-sdk/OpenDialogResult.md)

[Request changes](https://github.com/docker/docs/issues/new?template=doc_issue.yml&location=https%3A%2F%2Fdocs.docker.com%2Freference%2Fapi%2Fextensions-sdk%2FOpenDialogResult%2F&labels=status%2Ftriage)

Table of contents
