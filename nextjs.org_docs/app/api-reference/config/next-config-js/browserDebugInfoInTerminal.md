---
title: "browserDebugInfoInTerminal"
source: "https://nextjs.org/docs/app/api-reference/config/next-config-js/browserDebugInfoInTerminal"
---

# browserDebugInfoInTerminal

Menu

Using App Router

Features available in /app

Latest Version

16.0.5

[Configuration](/docs/app/api-reference/config)[next.config.js](/docs/app/api-reference/config/next-config-js)browserDebugInfoInTerminal

Copy page

# browserDebugInfoInTerminal

This feature is currently experimental and subject to change, it's not recommended for production. Try it out and share your feedback on [GitHub](https://github.com/vercel/next.js/issues).

The `experimental.browserDebugInfoInTerminal` option forwards console output and runtime errors originating in the browser to the dev server terminal.

This option is disabled by default. When enabled it only works in development mode.

## [Usage](#usage)

Enable forwarding:

next.config.ts

TypeScript

JavaScriptTypeScript

```
import type { NextConfig } from 'next'
 
const nextConfig: NextConfig = {
  experimental: {
    browserDebugInfoInTerminal: true,
  },
}
 
export default nextConfig
```

### [Serialization limits](#serialization-limits)

Deeply nested objects/arrays are truncated using **sensible defaults**. You can tweak these limits:

- **depthLimit**: (optional) Limit stringification depth for nested objects/arrays. Default: 5
- **edgeLimit**: (optional) Max number of properties or elements to include per object or array. Default: 100

next.config.ts

TypeScript

JavaScriptTypeScript

```
import type { NextConfig } from 'next'
 
const nextConfig: NextConfig = {
  experimental: {
    browserDebugInfoInTerminal: {
      depthLimit: 5,
      edgeLimit: 100,
    },
  },
}
 
export default nextConfig
```

### [Source location](#source-location)

Source locations are included by default when this feature is enabled.

app/page.tsx

```
'use client'
 
export default function Home() {
  return (
    <button
      type="button"
      onClick={() => {
        console.log('Hello World')
      }}
    >
      Click me
    </button>
  )
}
```

Clicking the button prints this message to the terminal.

Terminal

```
[browser] Hello World (app/page.tsx:8:17)
```

To suppress them, set `showSourceLocation: false`.

- **showSourceLocation**: Include source location info when available.

next.config.ts

TypeScript

JavaScriptTypeScript

```
import type { NextConfig } from 'next'
 
const nextConfig: NextConfig = {
  experimental: {
    browserDebugInfoInTerminal: {
      showSourceLocation: false,
    },
  },
}
 
export default nextConfig
```

VersionChanges`v15.4.0`experimental `browserDebugInfoInTerminal` introduced

Was this helpful?

supported.

Send
