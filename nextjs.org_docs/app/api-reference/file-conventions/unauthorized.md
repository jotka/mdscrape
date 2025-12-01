---
title: "unauthorized.js"
source: "https://nextjs.org/docs/app/api-reference/file-conventions/unauthorized"
---

# unauthorized.js

Menu

Using App Router

Features available in /app

Latest Version

16.0.5

[API Reference](/docs/app/api-reference)[File-system conventions](/docs/app/api-reference/file-conventions)unauthorized.js

Copy page

# unauthorized.js

This feature is currently experimental and subject to change, it's not recommended for production. Try it out and share your feedback on [GitHub](https://github.com/vercel/next.js/issues).

The **unauthorized** file is used to render UI when the [`unauthorized`](/docs/app/api-reference/functions/unauthorized) function is invoked during authentication. Along with allowing you to customize the UI, Next.js will return a `401` status code.

app/unauthorized.tsx

TypeScript

JavaScriptTypeScript

```
import Login from '@/app/components/Login'
 
export default function Unauthorized() {
  return (
    <main>
      <h1>401 - Unauthorized</h1>
      <p>Please log in to access this page.</p>
      <Login />
    </main>
  )
}
```

## [Reference](#reference)

### [Props](#props)

`unauthorized.js` components do not accept any props.

## [Examples](#examples)

### [Displaying login UI to unauthenticated users](#displaying-login-ui-to-unauthenticated-users)

You can use [`unauthorized`](/docs/app/api-reference/functions/unauthorized) function to render the `unauthorized.js` file with a login UI.

app/dashboard/page.tsx

TypeScript

JavaScriptTypeScript

```
import { verifySession } from '@/app/lib/dal'
import { unauthorized } from 'next/navigation'
 
export default async function DashboardPage() {
  const session = await verifySession()
 
  if (!session) {
    unauthorized()
  }
 
  return <div>Dashboard</div>
}
```

app/unauthorized.tsx

TypeScript

JavaScriptTypeScript

```
import Login from '@/app/components/Login'
 
export default function UnauthorizedPage() {
  return (
    <main>
      <h1>401 - Unauthorized</h1>
      <p>Please log in to access this page.</p>
      <Login />
    </main>
  )
}
```

## [Version History](#version-history)

VersionChanges`v15.1.0``unauthorized.js` introduced.

## Next Steps

[**unauthorized**  
\
API Reference for the unauthorized function.](/docs/app/api-reference/functions/unauthorized)

Was this helpful?

supported.

Send
