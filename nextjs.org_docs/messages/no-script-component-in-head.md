---
title: "No Script Component in Head"
source: "https://nextjs.org/docs/messages/no-script-component-in-head"
---

# No Script Component in Head

Menu

Using App Router

Features available in /app

Latest Version

16.0.5

[Docs](/docs)[Errors](/docs)No Script Component in Head

# No Script Component in Head

> Prevent usage of `next/script` in `next/head` component.

## [Why This Error Occurred](#why-this-error-occurred)

The `next/script` component should not be used in a `next/head` component.

## [Possible Ways to Fix It](#possible-ways-to-fix-it)

Move the `<Script />` component outside of `<Head>` instead.

**Before**

pages/index.js

```
import Script from 'next/script'
import Head from 'next/head'
 
export default function Index() {
  return (
    <Head>
      <title>Next.js</title>
      <Script src="/my-script.js" />
    </Head>
  )
}
```

**After**

pages/index.js

```
import Script from 'next/script'
import Head from 'next/head'
 
export default function Index() {
  return (
    <>
      <Head>
        <title>Next.js</title>
      </Head>
      <Script src="/my-script.js" />
    </>
  )
}
```

## [Useful Links](#useful-links)

- [next/head](/docs/pages/api-reference/components/head)
- [next/script](/docs/pages/guides/scripts)

Was this helpful?

supported.

Send
