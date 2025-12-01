---
title: "forbidden.js"
source: "https://nextjs.org/docs/15/app/api-reference/file-conventions/forbidden"
---

# forbidden.js

Menu

App Router 15

Using App Router

Features available in /app

Version 15

15.5.6

[API Reference](/docs/15/app/api-reference)[File-system conventions](/docs/15/app/api-reference/file-conventions)forbidden.js

You are currently viewing documentation for version 15 of Next.js.

# forbidden.js

This feature is currently experimental and subject to change, it's not recommended for production. Try it out and share your feedback on [GitHub](https://github.com/vercel/next.js/issues).

The **forbidden** file is used to render UI when the [`forbidden`](/docs/15/app/api-reference/functions/forbidden) function is invoked during authentication. Along with allowing you to customize the UI, Next.js will return a `403` status code.

app/forbidden.tsx

TypeScript

JavaScriptTypeScript

```
import Link from 'next/link'
 
export default function Forbidden() {
  return (
    <div>
      <h2>Forbidden</h2>
      <p>You are not authorized to access this resource.</p>
      <Link href="/">Return Home</Link>
    </div>
  )
}
```

## [Reference](#reference)

### [Props](#props)

`forbidden.js` components do not accept any props.

## [Version History](#version-history)

VersionChanges`v15.1.0``forbidden.js` introduced.

## Next Steps

[Next.js Docs  
\
...  
\
Functions  
\
**forbidden**  
\
API Reference for the forbidden function.](/docs/15/app/api-reference/functions/forbidden)

Was this helpful?

supported.

Send
