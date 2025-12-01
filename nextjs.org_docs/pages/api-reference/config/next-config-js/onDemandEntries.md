---
title: "onDemandEntries"
source: "https://nextjs.org/docs/pages/api-reference/config/next-config-js/onDemandEntries"
---

# onDemandEntries

Menu

Using App Router

Features available in /app

Latest Version

16.0.5

[Configuration](/docs/pages/api-reference/config)[next.config.js Options](/docs/pages/api-reference/config/next-config-js)onDemandEntries

Copy page

# onDemandEntries

Next.js exposes some options that give you some control over how the server will dispose or keep in memory built pages in development.

To change the defaults, open `next.config.js` and add the `onDemandEntries` config:

next.config.js

```
module.exports = {
  onDemandEntries: {
    // period (in ms) where the server will keep pages in the buffer
    maxInactiveAge: 25 * 1000,
    // number of pages that should be kept simultaneously without being disposed
    pagesBufferLength: 2,
  },
}
```

Was this helpful?

supported.

Send
