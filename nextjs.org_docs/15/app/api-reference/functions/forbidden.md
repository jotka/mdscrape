---
title: "forbidden"
source: "https://nextjs.org/docs/15/app/api-reference/functions/forbidden"
---

# forbidden

Menu

App Router 15

Using App Router

Features available in /app

Version 15

15.5.6

[API Reference](/docs/15/app/api-reference)[Functions](/docs/15/app/api-reference/functions)forbidden

You are currently viewing documentation for version 15 of Next.js.

# forbidden

This feature is currently experimental and subject to change, it's not recommended for production. Try it out and share your feedback on [GitHub](https://github.com/vercel/next.js/issues).

The `forbidden` function throws an error that renders a Next.js 403 error page. It's useful for handling authorization errors in your application. You can customize the UI using the [`forbidden.js` file](/docs/15/app/api-reference/file-conventions/forbidden).

To start using `forbidden`, enable the experimental [`authInterrupts`](/docs/15/app/api-reference/config/next-config-js/authInterrupts) configuration option in your `next.config.js` file:

next.config.ts

TypeScript

JavaScriptTypeScript

```
import type { NextConfig } from 'next'
 
const nextConfig: NextConfig = {
  experimental: {
    authInterrupts: true,
  },
}
 
export default nextConfig
```

`forbidden` can be invoked in [Server Components](/docs/15/app/getting-started/server-and-client-components), [Server Actions](/docs/15/app/getting-started/updating-data), and [Route Handlers](/docs/15/app/api-reference/file-conventions/route).

app/auth/page.tsx

TypeScript

JavaScriptTypeScript

```
import { verifySession } from '@/app/lib/dal'
import { forbidden } from 'next/navigation'
 
export default async function AdminPage() {
  const session = await verifySession()
 
  // Check if the user has the 'admin' role
  if (session.role !== 'admin') {
    forbidden()
  }
 
  // Render the admin page for authorized users
  return <></>
}
```

## [Good to know](#good-to-know)

- The `forbidden` function cannot be called in the [root layout](/docs/15/app/api-reference/file-conventions/layout#root-layout).

## [Examples](#examples)

### [Role-based route protection](#role-based-route-protection)

You can use `forbidden` to restrict access to certain routes based on user roles. This ensures that users who are authenticated but lack the required permissions cannot access the route.

app/admin/page.tsx

TypeScript

JavaScriptTypeScript

```
import { verifySession } from '@/app/lib/dal'
import { forbidden } from 'next/navigation'
 
export default async function AdminPage() {
  const session = await verifySession()
 
  // Check if the user has the 'admin' role
  if (session.role !== 'admin') {
    forbidden()
  }
 
  // Render the admin page for authorized users
  return (
    <main>
      <h1>Admin Dashboard</h1>
      <p>Welcome, {session.user.name}!</p>
    </main>
  )
}
```

### [Mutations with Server Actions](#mutations-with-server-actions)

When implementing mutations in Server Actions, you can use `forbidden` to only allow users with a specific role to update sensitive data.

app/actions/update-role.ts

TypeScript

JavaScriptTypeScript

```
'use server'
 
import { verifySession } from '@/app/lib/dal'
import { forbidden } from 'next/navigation'
import db from '@/app/lib/db'
 
export async function updateRole(formData: FormData) {
  const session = await verifySession()
 
  // Ensure only admins can update roles
  if (session.role !== 'admin') {
    forbidden()
  }
 
  // Perform the role update for authorized users
  // ...
}
```

## [Version History](#version-history)

VersionChanges`v15.1.0``forbidden` introduced.

## Next Steps

[Next.js Docs  
\
...  
\
File-system conventions  
\
**forbidden.js**  
\
API reference for the forbidden.js special file.](/docs/15/app/api-reference/file-conventions/forbidden)

Was this helpful?

supported.

Send
