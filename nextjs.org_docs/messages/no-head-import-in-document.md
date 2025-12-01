---
title: "No Head Import in Document"
source: "https://nextjs.org/docs/messages/no-head-import-in-document"
---

# No Head Import in Document

Menu

Using App Router

Features available in /app

Latest Version

16.0.5

[Docs](/docs)[Errors](/docs)No Head Import in Document

# No Head Import in Document

> Prevent usage of `next/head` in `pages/_document.js`.

## [Why This Error Occurred](#why-this-error-occurred)

`next/head` was imported in `pages/_document.js`. This can cause unexpected issues in your application.

## [Possible Ways to Fix It](#possible-ways-to-fix-it)

Only import and use `next/document` within `pages/_document.js` to override the default `Document` component. If you are importing `next/head` to use the `Head` component, import it from `next/document` instead in order to modify `<head>` code across all pages:

pages/\_document.js

```
import Document, { Html, Head, Main, NextScript } from 'next/document'
 
class MyDocument extends Document {
  static async getInitialProps(ctx) {
    //...
  }
 
  render() {
    return (
      <Html>
        <Head></Head>
      </Html>
    )
  }
}
 
export default MyDocument
```

## [Useful Links](#useful-links)

- [Custom Document](/docs/pages/building-your-application/routing/custom-document)

Was this helpful?

supported.

Send
