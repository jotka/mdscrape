---
title: "No Sync Scripts"
source: "https://nextjs.org/docs/messages/no-sync-scripts"
---

# No Sync Scripts

Menu

Using App Router

Features available in /app

Latest Version

16.0.5

[Docs](/docs)[Errors](/docs)No Sync Scripts

# No Sync Scripts

> Prevent synchronous scripts.

## [Why This Error Occurred](#why-this-error-occurred)

A synchronous script was used which can impact your webpage performance.

## [Possible Ways to Fix It](#possible-ways-to-fix-it)

### [Script component (recommended)](#script-component-recommended)

pages/index.js

```
import Script from 'next/script'
 
function Home() {
  return (
    <div class="container">
      <Script src="https://third-party-script.js"></Script>
      <div>Home Page</div>
    </div>
  )
}
 
export default Home
```

### [Use `async` or `defer`](#use-async-or-defer)

```
<script src="https://third-party-script.js" async />
<script src="https://third-party-script.js" defer />
```

## [Useful Links](#useful-links)

- [Efficiently load third-party JavaScript](https://web.dev/efficiently-load-third-party-javascript/)

Was this helpful?

supported.

Send
