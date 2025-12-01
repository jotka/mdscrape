---
title: "No Duplicate Head"
source: "https://nextjs.org/docs/messages/no-duplicate-head"
---

# No Duplicate Head

Menu

Using App Router

Features available in /app

Latest Version

16.0.5

[Docs](/docs)[Errors](/docs)No Duplicate Head

# No Duplicate Head

> Prevent duplicate usage of `<Head>` in `pages/_document.js`.

## [Why This Error Occurred](#why-this-error-occurred)

More than a single instance of the `<Head />` component was used in a single custom document. This can cause unexpected behavior in your application.

## [Possible Ways to Fix It](#possible-ways-to-fix-it)

Only use a single `<Head />` component in your custom document in `pages/_document.js`.

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
        <Head />
        <body>
          <Main />
          <NextScript />
        </body>
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
