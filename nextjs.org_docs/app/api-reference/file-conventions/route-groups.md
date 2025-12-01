---
title: "Route Groups"
source: "https://nextjs.org/docs/app/api-reference/file-conventions/route-groups"
---

# Route Groups

Menu

Using App Router

Features available in /app

Latest Version

16.0.5

[API Reference](/docs/app/api-reference)[File-system conventions](/docs/app/api-reference/file-conventions)Route Groups

Copy page

# Route Groups

Route Groups are a folder convention that let you organize routes by category or team.

## [Convention](#convention)

A route group can be created by wrapping a folder's name in parenthesis: `(folderName)`.

This convention indicates the folder is for organizational purposes and should **not be included** in the route's URL path.

![An example folder structure using route groups](/_next/image?url=https%3A%2F%2Fh8DxKfmAPhn8O0p3.public.blob.vercel-storage.com%2Fdocs%2Flight%2Fproject-organization-route-groups.png&w=3840&q=75)![An example folder structure using route groups](/_next/image?url=https%3A%2F%2Fh8DxKfmAPhn8O0p3.public.blob.vercel-storage.com%2Fdocs%2Fdark%2Fproject-organization-route-groups.png&w=3840&q=75)

## [Use cases](#use-cases)

- Organizing routes by team, concern, or feature.
- Defining multiple [root layouts](/docs/app/api-reference/file-conventions/layout#root-layout).
- Opting specific route segments into sharing a layout, while keeping others out.

## [Caveats](#caveats)

- **Full page load**: If you navigate between routes that use different root layouts, it'll trigger a full page reload. For example, navigating from `/cart` that uses `app/(shop)/layout.js` to `/blog` that uses `app/(marketing)/layout.js`. This **only** applies to multiple root layouts.
- **Conflicting paths**: Routes in different groups should not resolve to the same URL path. For example, `(marketing)/about/page.js` and `(shop)/about/page.js` would both resolve to `/about` and cause an error.
- **Top-level root layout**: If you use multiple root layouts without a top-level `layout.js` file, make sure your home route (/) is defined within one of the route groups, e.g. app/(marketing)/page.js.

Was this helpful?

supported.

Send
