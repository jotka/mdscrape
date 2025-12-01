---
title: "Turbopack FileSystem Caching"
source: "https://nextjs.org/docs/app/api-reference/config/next-config-js/turbopackFileSystemCache"
---

# Turbopack FileSystem Caching

Menu

Using App Router

Features available in /app

Latest Version

16.0.5

[Configuration](/docs/app/api-reference/config)[next.config.js](/docs/app/api-reference/config/next-config-js)turbopackFileSystemCache

Copy page

# Turbopack FileSystem Caching

## [Usage](#usage)

Turbopack FileSystem Cache enables Turbopack to reduce work across `next dev` or `next build` commands. When enabled, Turbopack will save and restore data to the `.next` folder between builds, which can greatly speed up subsequent builds and dev sessions.

> **Good to know:** The FileSystem Cache feature is Beta and is still under development. Users adopting should expect some stability issues. We recommend first adopting it for development.

next.config.ts

TypeScript

JavaScriptTypeScript

```
import type { NextConfig } from 'next'
 
const nextConfig: NextConfig = {
  experimental: {
    // Enable filesystem caching for `next dev`
    turbopackFileSystemCacheForDev: true,
    // Enable filesystem caching for `next build`
    turbopackFileSystemCacheForBuild: true,
  },
}
 
export default nextConfig
```

## [Version Changes](#version-changes)

VersionChanges`v16.0.0`Beta release with separate flags for build and dev`v15.5.0`Persistent caching released as experimental on canary releases

Was this helpful?

supported.

Send
