---
title: "draftMode"
source: "https://nextjs.org/docs/15/app/api-reference/functions/draft-mode"
---

# draftMode

Menu

App Router 15

Using App Router

Features available in /app

Version 15

15.5.6

[API Reference](/docs/15/app/api-reference)[Functions](/docs/15/app/api-reference/functions)draftMode

You are currently viewing documentation for version 15 of Next.js.

# draftMode

`draftMode` is an **async** function allows you to enable and disable [Draft Mode](/docs/15/app/guides/draft-mode), as well as check if Draft Mode is enabled in a [Server Component](/docs/15/app/getting-started/server-and-client-components).

app/page.ts

TypeScript

JavaScriptTypeScript

```
import { draftMode } from 'next/headers'
 
export default async function Page() {
  const { isEnabled } = await draftMode()
}
```

## [Reference](#reference)

The following methods and properties are available:

MethodDescription`isEnabled`A boolean value that indicates if Draft Mode is enabled.`enable()`Enables Draft Mode in a Route Handler by setting a cookie (`__prerender_bypass`).`disable()`Disables Draft Mode in a Route Handler by deleting a cookie.

## [Good to know](#good-to-know)

- `draftMode` is an **asynchronous** function that returns a promise. You must use `async/await` or React's [`use`](https://react.dev/reference/react/use) function.
  
  - In version 14 and earlier, `draftMode` was a synchronous function. To help with backwards compatibility, you can still access it synchronously in Next.js 15, but this behavior will be deprecated in the future.
- A new bypass cookie value will be generated each time you run `next build`. This ensures that the bypass cookie can’t be guessed.
- To test Draft Mode locally over HTTP, your browser will need to allow third-party cookies and local storage access.

## [Examples](#examples)

### [Enabling Draft Mode](#enabling-draft-mode)

To enable Draft Mode, create a new [Route Handler](/docs/15/app/api-reference/file-conventions/route) and call the `enable()` method:

app/draft/route.ts

TypeScript

JavaScriptTypeScript

```
import { draftMode } from 'next/headers'
 
export async function GET(request: Request) {
  const draft = await draftMode()
  draft.enable()
  return new Response('Draft mode is enabled')
}
```

### [Disabling Draft Mode](#disabling-draft-mode)

By default, the Draft Mode session ends when the browser is closed.

To disable Draft Mode manually, call the `disable()` method in your [Route Handler](/docs/15/app/api-reference/file-conventions/route):

app/draft/route.ts

TypeScript

JavaScriptTypeScript

```
import { draftMode } from 'next/headers'
 
export async function GET(request: Request) {
  const draft = await draftMode()
  draft.disable()
  return new Response('Draft mode is disabled')
}
```

Then, send a request to invoke the Route Handler. If calling the route using the [`<Link>` component](/docs/15/app/api-reference/components/link), you must pass `prefetch={false}` to prevent accidentally deleting the cookie on prefetch.

### [Checking if Draft Mode is enabled](#checking-if-draft-mode-is-enabled)

You can check if Draft Mode is enabled in a Server Component with the `isEnabled` property:

app/page.ts

TypeScript

JavaScriptTypeScript

```
import { draftMode } from 'next/headers'
 
export default async function Page() {
  const { isEnabled } = await draftMode()
  return (
    <main>
      <h1>My Blog Post</h1>
      <p>Draft Mode is currently {isEnabled ? 'Enabled' : 'Disabled'}</p>
    </main>
  )
}
```

## [Version History](#version-history)

VersionChanges`v15.0.0-RC``draftMode` is now an async function. A [codemod](/docs/15/app/guides/upgrading/codemods#150) is available.`v13.4.0``draftMode` introduced.

## Next Steps

Learn how to use Draft Mode with this step-by-step guide.

[Next.js Docs  
\
...  
\
Guides  
\
**Draft Mode**  
\
Next.js has draft mode to toggle between static and dynamic pages. You can learn how it works with App Router here.](/docs/15/app/guides/draft-mode)

Was this helpful?

supported.

Send
