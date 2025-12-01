---
title: "cookies"
source: "https://nextjs.org/docs/15/app/api-reference/functions/cookies"
---

# cookies

Menu

App Router 15

Using App Router

Features available in /app

Version 15

15.5.6

[API Reference](/docs/15/app/api-reference)[Functions](/docs/15/app/api-reference/functions)cookies

You are currently viewing documentation for version 15 of Next.js.

# cookies

`cookies` is an **async** function that allows you to read the HTTP incoming request cookies in [Server Components](/docs/15/app/getting-started/server-and-client-components), and read/write outgoing request cookies in [Server Actions](/docs/15/app/getting-started/updating-data) or [Route Handlers](/docs/15/app/api-reference/file-conventions/route).

app/page.tsx

TypeScript

JavaScriptTypeScript

```
import { cookies } from 'next/headers'
 
export default async function Page() {
  const cookieStore = await cookies()
  const theme = cookieStore.get('theme')
  return '...'
}
```

## [Reference](#reference)

### [Methods](#methods)

The following methods are available:

MethodReturn TypeDescription`get('name')`ObjectAccepts a cookie name and returns an object with the name and value.`getAll()`Array of objectsReturns a list of all the cookies with a matching name.`has('name')`BooleanAccepts a cookie name and returns a boolean based on if the cookie exists.`set(name, value, options)`-Accepts a cookie name, value, and options and sets the outgoing request cookie.`delete(name)`-Accepts a cookie name and deletes the cookie.`clear()`-Deletes all cookies.`toString()`StringReturns a string representation of the cookies.

### [Options](#options)

When setting a cookie, the following properties from the `options` object are supported:

OptionTypeDescription`name`StringSpecifies the name of the cookie.`value`StringSpecifies the value to be stored in the cookie.`expires`DateDefines the exact date when the cookie will expire.`maxAge`NumberSets the cookie’s lifespan in seconds.`domain`StringSpecifies the domain where the cookie is available.`path`String, default: `'/'`Limits the cookie's scope to a specific path within the domain.`secure`BooleanEnsures the cookie is sent only over HTTPS connections for added security.`httpOnly`BooleanRestricts the cookie to HTTP requests, preventing client-side access.`sameSite`Boolean, `'lax'`, `'strict'`, `'none'`Controls the cookie's cross-site request behavior.`priority`String (`"low"`, `"medium"`, `"high"`)Specifies the cookie's priority`partitioned`BooleanIndicates whether the cookie is [partitioned](https://github.com/privacycg/CHIPS).

The only option with a default value is `path`.

To learn more about these options, see the [MDN docs](https://developer.mozilla.org/en-US/docs/Web/HTTP/Cookies).

## [Good to know](#good-to-know)

- `cookies` is an **asynchronous** function that returns a promise. You must use `async/await` or React's [`use`](https://react.dev/reference/react/use) function to access cookies.
  
  - In version 14 and earlier, `cookies` was a synchronous function. To help with backwards compatibility, you can still access it synchronously in Next.js 15, but this behavior will be deprecated in the future.
- `cookies` is a [Dynamic API](/docs/15/app/getting-started/partial-prerendering#dynamic-rendering) whose returned values cannot be known ahead of time. Using it in a layout or page will opt a route into [dynamic rendering](/docs/15/app/getting-started/partial-prerendering#dynamic-rendering).
- The `.delete` method can only be called:
  
  - In a [Server Action](/docs/15/app/getting-started/updating-data) or [Route Handler](/docs/15/app/api-reference/file-conventions/route).
  - If it belongs to the same domain from which `.set` is called. For wildcard domains, the specific subdomain must be an exact match. Additionally, the code must be executed on the same protocol (HTTP or HTTPS) as the cookie you want to delete.
- HTTP does not allow setting cookies after streaming starts, so you must use `.set` in a [Server Action](/docs/15/app/getting-started/updating-data) or [Route Handler](/docs/15/app/api-reference/file-conventions/route).

## [Understanding Cookie Behavior in Server Components](#understanding-cookie-behavior-in-server-components)

When working with cookies in Server Components, it's important to understand that cookies are fundamentally a client-side storage mechanism:

- **Reading cookies** works in Server Components because you're accessing the cookie data that the client's browser sends to the server in the HTTP request headers.
- **Setting cookies** cannot be done directly in a Server Component, even when using a Route Handler or Server Action. This is because cookies are actually stored by the browser, not the server.

The server can only send instructions (via `Set-Cookie` headers) to tell the browser to store cookies - the actual storage happens on the client side. This is why cookie operations that modify state (`.set`, `.delete`, `.clear`) must be performed in a Route Handler or Server Action where the response headers can be properly set.

## [Understanding Cookie Behavior in Server Actions](#understanding-cookie-behavior-in-server-actions)

After you set or delete a cookie in a Server Action, Next.js re-renders the current page and its layouts on the server so the UI reflects the new cookie value. See the [Caching guide](/docs/15/app/guides/caching#cookies).

The UI is not unmounted, but effects that depend on data coming from the server will re-run.

To refresh cached data too, call [`revalidatePath`](/docs/15/app/api-reference/functions/revalidatePath) or [`revalidateTag`](/docs/15/app/api-reference/functions/revalidateTag) inside the action.

## [Examples](#examples)

### [Getting a cookie](#getting-a-cookie)

You can use the `(await cookies()).get('name')` method to get a single cookie:

app/page.tsx

TypeScript

JavaScriptTypeScript

```
import { cookies } from 'next/headers'
 
export default async function Page() {
  const cookieStore = await cookies()
  const theme = cookieStore.get('theme')
  return '...'
}
```

### [Getting all cookies](#getting-all-cookies)

You can use the `(await cookies()).getAll()` method to get all cookies with a matching name. If `name` is unspecified, it returns all the available cookies.

app/page.tsx

TypeScript

JavaScriptTypeScript

```
import { cookies } from 'next/headers'
 
export default async function Page() {
  const cookieStore = await cookies()
  return cookieStore.getAll().map((cookie) => (
    <div key={cookie.name}>
      <p>Name: {cookie.name}</p>
      <p>Value: {cookie.value}</p>
    </div>
  ))
}
```

### [Setting a cookie](#setting-a-cookie)

You can use the `(await cookies()).set(name, value, options)` method in a [Server Action](/docs/15/app/getting-started/updating-data) or [Route Handler](/docs/15/app/api-reference/file-conventions/route) to set a cookie. The [`options` object](#options) is optional.

app/actions.ts

TypeScript

JavaScriptTypeScript

```
'use server'
 
import { cookies } from 'next/headers'
 
export async function create(data) {
  const cookieStore = await cookies()
 
  cookieStore.set('name', 'lee')
  // or
  cookieStore.set('name', 'lee', { secure: true })
  // or
  cookieStore.set({
    name: 'name',
    value: 'lee',
    httpOnly: true,
    path: '/',
  })
}
```

### [Checking if a cookie exists](#checking-if-a-cookie-exists)

You can use the `(await cookies()).has(name)` method to check if a cookie exists:

app/page.ts

TypeScript

JavaScriptTypeScript

```
import { cookies } from 'next/headers'
 
export default async function Page() {
  const cookieStore = await cookies()
  const hasCookie = cookieStore.has('theme')
  return '...'
}
```

### [Deleting cookies](#deleting-cookies)

There are three ways you can delete a cookie.

Using the `delete()` method:

app/actions.ts

TypeScript

JavaScriptTypeScript

```
'use server'
 
import { cookies } from 'next/headers'
 
export async function delete(data) {
  (await cookies()).delete('name')
}
```

Setting a new cookie with the same name and an empty value:

app/actions.ts

TypeScript

JavaScriptTypeScript

```
'use server'
 
import { cookies } from 'next/headers'
 
export async function delete(data) {
  (await cookies()).set('name', '')
}
```

Setting the `maxAge` to 0 will immediately expire a cookie. `maxAge` accepts a value in seconds.

app/actions.ts

TypeScript

JavaScriptTypeScript

```
'use server'
 
import { cookies } from 'next/headers'
 
export async function delete(data) {
  (await cookies()).set('name', 'value', { maxAge: 0 })
}
```

## [Version History](#version-history)

VersionChanges`v15.0.0-RC``cookies` is now an async function. A [codemod](/docs/15/app/guides/upgrading/codemods#150) is available.`v13.0.0``cookies` introduced.

Was this helpful?

supported.

Send
