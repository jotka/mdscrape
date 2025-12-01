---
title: "use server"
source: "https://nextjs.org/docs/15/app/api-reference/directives/use-server"
---

# use server

Menu

App Router 15

Using App Router

Features available in /app

Version 15

15.5.6

[API Reference](/docs/15/app/api-reference)[Directives](/docs/15/app/api-reference/directives)use server

You are currently viewing documentation for version 15 of Next.js.

# use server

The `use server` directive designates a function or file to be executed on the **server side**. It can be used at the top of a file to indicate that all functions in the file are server-side, or inline at the top of a function to mark the function as a [Server Function](https://19.react.dev/reference/rsc/server-functions). This is a React feature.

## [Using `use server` at the top of a file](#using-use-server-at-the-top-of-a-file)

The following example shows a file with a `use server` directive at the top. All functions in the file are executed on the server.

app/actions.ts

TypeScript

JavaScriptTypeScript

```
'use server'
import { db } from '@/lib/db' // Your database client
 
export async function createUser(data: { name: string; email: string }) {
  const user = await db.user.create({ data })
  return user
}
```

### [Using Server Functions in a Client Component](#using-server-functions-in-a-client-component)

To use Server Functions in Client Components you need to create your Server Functions in a dedicated file using the `use server` directive at the top of the file. These Server Functions can then be imported into Client and Server Components and executed.

Assuming you have a `fetchUsers` Server Function in `actions.ts`:

app/actions.ts

TypeScript

JavaScriptTypeScript

```
'use server'
import { db } from '@/lib/db' // Your database client
 
export async function fetchUsers() {
  const users = await db.user.findMany()
  return users
}
```

Then you can import the `fetchUsers` Server Function into a Client Component and execute it on the client-side.

app/components/my-button.tsx

TypeScript

JavaScriptTypeScript

```
'use client'
import { fetchUsers } from '../actions'
 
export default function MyButton() {
  return <button onClick={() => fetchUsers()}>Fetch Users</button>
}
```

## [Using `use server` inline](#using-use-server-inline)

In the following example, `use server` is used inline at the top of a function to mark it as a [Server Function](https://19.react.dev/reference/rsc/server-functions):

app/posts/\[id]/page.tsx

TypeScript

JavaScriptTypeScript

```
import { EditPost } from './edit-post'
import { revalidatePath } from 'next/cache'
 
export default async function PostPage({ params }: { params: { id: string } }) {
  const post = await getPost(params.id)
 
  async function updatePost(formData: FormData) {
    'use server'
    await savePost(params.id, formData)
    revalidatePath(`/posts/${params.id}`)
  }
 
  return <EditPost updatePostAction={updatePost} post={post} />
}
```

## [Security considerations](#security-considerations)

When using the `use server` directive, it's important to ensure that all server-side logic is secure and that sensitive data remains protected.

### [Authentication and authorization](#authentication-and-authorization)

Always authenticate and authorize users before performing sensitive server-side operations.

app/actions.ts

TypeScript

JavaScriptTypeScript

```
'use server'
 
import { db } from '@/lib/db' // Your database client
import { authenticate } from '@/lib/auth' // Your authentication library
 
export async function createUser(
  data: { name: string; email: string },
  token: string
) {
  const user = authenticate(token)
  if (!user) {
    throw new Error('Unauthorized')
  }
  const newUser = await db.user.create({ data })
  return newUser
}
```

## [Reference](#reference)

See the [React documentation](https://react.dev/reference/rsc/use-server) for more information on `use server`.

Was this helpful?

supported.

Send
