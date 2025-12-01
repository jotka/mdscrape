---
title: "bundlePagesRouterDependencies"
source: "https://nextjs.org/docs/pages/api-reference/config/next-config-js/bundlePagesRouterDependencies"
---

# bundlePagesRouterDependencies

Menu

Using App Router

Features available in /app

Latest Version

16.0.5

[Configuration](/docs/pages/api-reference/config)[next.config.js Options](/docs/pages/api-reference/config/next-config-js)bundlePagesRouterDependencies

Copy page

# bundlePagesRouterDependencies

Enable automatic server-side dependency bundling for Pages Router applications. Matches the automatic dependency bundling in App Router.

next.config.js

```
/** @type {import('next').NextConfig} */
const nextConfig = {
  bundlePagesRouterDependencies: true,
}
 
module.exports = nextConfig
```

Explicitly opt-out certain packages from being bundled using the [`serverExternalPackages`](/docs/pages/api-reference/config/next-config-js/serverExternalPackages) option.

## [Version History](#version-history)

VersionChanges`v15.0.0`Moved from experimental to stable. Renamed from `bundlePagesExternals` to `bundlePagesRouterDependencies`

Was this helpful?

supported.

Send
