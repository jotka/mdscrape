---
title: "Conflicting Public File and Page File"
source: "https://nextjs.org/docs/messages/conflicting-public-file-page"
---

# Conflicting Public File and Page File

Menu

Using App Router

Features available in /app

Latest Version

16.0.5

[Docs](/docs)[Errors](/docs)Conflicting Public File and Page File

# Conflicting Public File and Page File

## [Why This Error Occurred](#why-this-error-occurred)

One of your public files has the same path as a page file which is not supported. Since only one resource can reside at the URL both public files and page files must be unique.

## [Possible Ways to Fix It](#possible-ways-to-fix-it)

Rename either the public file or page file that is causing the conflict.

Example conflict between public file and page file

Folder structure

```
public/
  hello
pages/
  hello.js
```

Non-conflicting public file and page file

Folder structure

```
public/
  hello.txt
pages/
  hello.js
```

## [Useful Links](#useful-links)

- [Static file serving docs](/docs/pages/api-reference/file-conventions/public-folder)

Was this helpful?

supported.

Send
