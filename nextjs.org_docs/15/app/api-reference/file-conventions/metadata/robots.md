---
title: "robots.txt"
source: "https://nextjs.org/docs/15/app/api-reference/file-conventions/metadata/robots"
---

# robots.txt

Menu

App Router 15

Using App Router

Features available in /app

Version 15

15.5.6

[File-system conventions](/docs/15/app/api-reference/file-conventions)[Metadata Files](/docs/15/app/api-reference/file-conventions/metadata)robots.txt

You are currently viewing documentation for version 15 of Next.js.

# robots.txt

Add or generate a `robots.txt` file that matches the [Robots Exclusion Standard](https://en.wikipedia.org/wiki/Robots.txt#Standard) in the **root** of `app` directory to tell search engine crawlers which URLs they can access on your site.

## [Static `robots.txt`](#static-robotstxt)

app/robots.txt

```
User-Agent: *
Allow: /
Disallow: /private/

Sitemap: https://acme.com/sitemap.xml
```

## [Generate a Robots file](#generate-a-robots-file)

Add a `robots.js` or `robots.ts` file that returns a [`Robots` object](#robots-object).

> **Good to know**: `robots.js` is a special Route Handlers that is cached by default unless it uses a [Dynamic API](/docs/15/app/guides/caching#dynamic-apis) or [dynamic config](/docs/15/app/guides/caching#segment-config-options) option.

app/robots.ts

TypeScript

JavaScriptTypeScript

```
import type { MetadataRoute } from 'next'
 
export default function robots(): MetadataRoute.Robots {
  return {
    rules: {
      userAgent: '*',
      allow: '/',
      disallow: '/private/',
    },
    sitemap: 'https://acme.com/sitemap.xml',
  }
}
```

Output:

```
User-Agent: *
Allow: /
Disallow: /private/

Sitemap: https://acme.com/sitemap.xml
```

### [Customizing specific user agents](#customizing-specific-user-agents)

You can customise how individual search engine bots crawl your site by passing an array of user agents to the `rules` property. For example:

app/robots.ts

TypeScript

JavaScriptTypeScript

```
import type { MetadataRoute } from 'next'
 
export default function robots(): MetadataRoute.Robots {
  return {
    rules: [
      {
        userAgent: 'Googlebot',
        allow: ['/'],
        disallow: '/private/',
      },
      {
        userAgent: ['Applebot', 'Bingbot'],
        disallow: ['/'],
      },
    ],
    sitemap: 'https://acme.com/sitemap.xml',
  }
}
```

Output:

```
User-Agent: Googlebot
Allow: /
Disallow: /private/

User-Agent: Applebot
Disallow: /

User-Agent: Bingbot
Disallow: /

Sitemap: https://acme.com/sitemap.xml
```

### [Robots object](#robots-object)

```
type Robots = {
  rules:
    | {
        userAgent?: string | string[]
        allow?: string | string[]
        disallow?: string | string[]
        crawlDelay?: number
      }
    | Array<{
        userAgent: string | string[]
        allow?: string | string[]
        disallow?: string | string[]
        crawlDelay?: number
      }>
  sitemap?: string | string[]
  host?: string
}
```

## [Version History](#version-history)

VersionChanges`v13.3.0``robots` introduced.

Was this helpful?

supported.

Send
