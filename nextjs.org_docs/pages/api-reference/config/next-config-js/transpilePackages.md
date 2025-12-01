---
title: "transpilePackages"
source: "https://nextjs.org/docs/pages/api-reference/config/next-config-js/transpilePackages"
---

# transpilePackages

Menu

Using App Router

Features available in /app

Latest Version

16.0.5

[Configuration](/docs/pages/api-reference/config)[next.config.js Options](/docs/pages/api-reference/config/next-config-js)transpilePackages

Copy page

# transpilePackages

Next.js can automatically transpile and bundle dependencies from local packages (like monorepos) or from external dependencies (`node_modules`). This replaces the `next-transpile-modules` package.

next.config.js

```
/** @type {import('next').NextConfig} */
const nextConfig = {
  transpilePackages: ['package-name'],
}
 
module.exports = nextConfig
```

## [Version History](#version-history)

VersionChanges`v13.0.0``transpilePackages` added.

Was this helpful?

supported.

Send
