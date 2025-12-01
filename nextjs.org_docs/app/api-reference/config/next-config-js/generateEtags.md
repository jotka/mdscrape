---
title: "generateEtags"
source: "https://nextjs.org/docs/app/api-reference/config/next-config-js/generateEtags"
---

# generateEtags

Menu

Using App Router

Features available in /app

Latest Version

16.0.5

[Configuration](/docs/app/api-reference/config)[next.config.js](/docs/app/api-reference/config/next-config-js)generateEtags

Copy page

# generateEtags

Next.js will generate [etags](https://en.wikipedia.org/wiki/HTTP_ETag) for every page by default. You may want to disable etag generation for HTML pages depending on your cache strategy.

Open `next.config.js` and disable the `generateEtags` option:

next.config.js

```
module.exports = {
  generateEtags: false,
}
```

Was this helpful?

supported.

Send
