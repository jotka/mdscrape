---
title: "Google Font Preconnect"
source: "https://nextjs.org/docs/messages/google-font-preconnect"
---

# Google Font Preconnect

Menu

Using App Router

Features available in /app

Latest Version

16.0.5

[Docs](/docs)[Errors](/docs)Google Font Preconnect

# Google Font Preconnect

> **Note**: Next.js automatically adds `<link rel="preconnect" />` after version `12.0.1`.

> Ensure `preconnect` is used with Google Fonts.

## [Why This Error Occurred](#why-this-error-occurred)

A preconnect resource hint was not used with a request to the Google Fonts domain. Adding `preconnect` is recommended to initiate an early connection to the origin.

## [Possible Ways to Fix It](#possible-ways-to-fix-it)

Add `rel="preconnect"` to the Google Font domain `<link>` tag:

pages/\_document.js

```
<link rel="preconnect" href="https://fonts.gstatic.com" />
```

> **Note**: a **separate** link with `dns-prefetch` can be used as a fallback for browsers that don't support `preconnect` although this is not required.

## [Useful Links](#useful-links)

- [Preconnect to required origins](https://web.dev/uses-rel-preconnect/)
- [Preconnect and dns-prefetch](https://web.dev/preconnect-and-dns-prefetch/#resolve-domain-name-early-with-reldns-prefetch)
- [Next.js Font Optimization](/docs/pages/api-reference/components/font)

Was this helpful?

supported.

Send
