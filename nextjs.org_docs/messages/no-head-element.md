---
title: "No Head Element"
source: "https://nextjs.org/docs/messages/no-head-element"
---

# No Head Element

Menu

Using App Router

Features available in /app

Latest Version

16.0.5

[Docs](/docs)[Errors](/docs)No Head Element

# No Head Element

> Prevent usage of `<head>` element.

## [Why This Error Occurred](#why-this-error-occurred)

A `<head>` element was used to include page-level metadata, but this can cause unexpected behavior in a Next.js application. Use Next.js' built-in `next/head` component instead.

## [Possible Ways to Fix It](#possible-ways-to-fix-it)

Import and use the `<Head />` component:

pages/index.js

```
import Head from 'next/head'
 
function Index() {
  return (
    <>
      <Head>
        <title>My page title</title>
        <meta name="viewport" content="initial-scale=1.0, width=device-width" />
      </Head>
    </>
  )
}
 
export default Index
```

## [Useful Links](#useful-links)

- [next/head](/docs/pages/api-reference/components/head)

Was this helpful?

supported.

Send
