---
title: "Inline script id"
source: "https://nextjs.org/docs/messages/inline-script-id"
---

# Inline script id

Menu

Using App Router

Features available in /app

Latest Version

16.0.5

[Docs](/docs)[Errors](/docs)Inline script id

# Inline script id

> Enforce `id` attribute on `next/script` components with inline content.

## [Why This Error Occurred](#why-this-error-occurred)

`next/script` components with inline content require an `id` attribute to be defined to track and optimize the script.

## [Possible Ways to Fix It](#possible-ways-to-fix-it)

Add an `id` attribute to the `next/script` component.

pages/\_app.js

```
import Script from 'next/script'
 
export default function App({ Component, pageProps }) {
  return (
    <>
      <Script id="my-script">{`console.log('Hello world!');`}</Script>
      <Component {...pageProps} />
    </>
  )
}
```

## [Useful links](#useful-links)

- [Docs for Next.js Script component](/docs/pages/guides/scripts)

Was this helpful?

supported.

Send
