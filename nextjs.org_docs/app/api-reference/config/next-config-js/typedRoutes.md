---
title: "typedRoutes"
source: "https://nextjs.org/docs/app/api-reference/config/next-config-js/typedRoutes"
---

# typedRoutes

Menu

Using App Router

Features available in /app

Latest Version

16.0.5

[Configuration](/docs/app/api-reference/config)[next.config.js](/docs/app/api-reference/config/next-config-js)typedRoutes

Copy page

# typedRoutes

> **Note**: This option has been marked as stable, so you should use `typedRoutes` instead of `experimental.typedRoutes`.

Support for [statically typed links](/docs/app/api-reference/config/typescript#statically-typed-links). This feature requires using TypeScript in your project.

next.config.js

```
/** @type {import('next').NextConfig} */
const nextConfig = {
  typedRoutes: true,
}
 
module.exports = nextConfig
```

Was this helpful?

supported.

Send
