---
title: "No Document Import in Page"
source: "https://nextjs.org/docs/messages/no-document-import-in-page"
---

# No Document Import in Page

Menu

Using App Router

Features available in /app

Latest Version

16.0.5

[Docs](/docs)[Errors](/docs)No Document Import in Page

# No Document Import in Page

> Prevent importing `next/document` outside of `pages/_document.js`.

## [Why This Error Occurred](#why-this-error-occurred)

`next/document` was imported in a page outside of `pages/_document.js` (or `pages/_document.tsx` if you are using TypeScript). This can cause unexpected issues in your application.

## [Possible Ways to Fix It](#possible-ways-to-fix-it)

Only import and use `next/document` within `pages/_document.js` (or `pages/_document.tsx`) to override the default `Document` component:

pages/\_document.js

```
import Document, { Html, Head, Main, NextScript } from 'next/document'
 
class MyDocument extends Document {
  //...
}
 
export default MyDocument
```

## [Useful Links](#useful-links)

- [Custom Document](/docs/pages/building-your-application/routing/custom-document)

Was this helpful?

supported.

Send
