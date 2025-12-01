---
title: "No Title in Document Head"
source: "https://nextjs.org/docs/messages/no-title-in-document-head"
---

# No Title in Document Head

Menu

Using App Router

Features available in /app

Latest Version

16.0.5

[Docs](/docs)[Errors](/docs)No Title in Document Head

# No Title in Document Head

> Prevent usage of `<title>` with `Head` component from `next/document`.

## [Why This Error Occurred](#why-this-error-occurred)

A `<title>` element was defined within the `Head` component imported from `next/document`, which should only be used for any `<head>` code that is common for all pages. Title tags should be defined at the page-level using `next/head` instead.

## [Possible Ways to Fix It](#possible-ways-to-fix-it)

Within a page or component, import and use `next/head` to define a page title:

pages/index.js

```
import Head from 'next/head'
 
export function Home() {
  return (
    <div>
      <Head>
        <title>My page title</title>
      </Head>
    </div>
  )
}
```

## [Useful Links](#useful-links)

- [next/head](/docs/pages/api-reference/components/head)
- [Custom Document](/docs/pages/building-your-application/routing/custom-document)

Was this helpful?

supported.

Send
