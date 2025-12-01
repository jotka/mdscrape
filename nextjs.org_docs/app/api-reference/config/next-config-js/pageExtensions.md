---
title: "pageExtensions"
source: "https://nextjs.org/docs/app/api-reference/config/next-config-js/pageExtensions"
---

# pageExtensions

Menu

Using App Router

Features available in /app

Latest Version

16.0.5

[Configuration](/docs/app/api-reference/config)[next.config.js](/docs/app/api-reference/config/next-config-js)pageExtensions

Copy page

# pageExtensions

By default, Next.js accepts files with the following extensions: `.tsx`, `.ts`, `.jsx`, `.js`. This can be modified to allow other extensions like markdown (`.md`, `.mdx`).

next.config.js

```
const withMDX = require('@next/mdx')()
 
/** @type {import('next').NextConfig} */
const nextConfig = {
  pageExtensions: ['js', 'jsx', 'ts', 'tsx', 'md', 'mdx'],
}
 
module.exports = withMDX(nextConfig)
```

Was this helpful?

supported.

Send
