---
title: "generateSitemaps"
source: "https://nextjs.org/docs/app/api-reference/functions/generate-sitemaps"
---

# generateSitemaps

Menu

Using App Router

Features available in /app

Latest Version

16.0.5

[API Reference](/docs/app/api-reference)[Functions](/docs/app/api-reference/functions)generateSitemaps

Copy page

# generateSitemaps

You can use the `generateSitemaps` function to generate multiple sitemaps for your application.

## [Returns](#returns)

The `generateSitemaps` returns an array of objects with an `id` property.

## [URLs](#urls)

Your generated sitemaps will be available at `/.../sitemap/[id].xml`. For example, `/product/sitemap/1.xml`.

## [Example](#example)

For example, to split a sitemap using `generateSitemaps`, return an array of objects with the sitemap `id`. Then, use the `id` to generate the unique sitemaps.

app/product/sitemap.ts

TypeScript

JavaScriptTypeScript

```
import type { MetadataRoute } from 'next'
import { BASE_URL } from '@/app/lib/constants'
 
export async function generateSitemaps() {
  // Fetch the total number of products and calculate the number of sitemaps needed
  return [{ id: 0 }, { id: 1 }, { id: 2 }, { id: 3 }]
}
 
export default async function sitemap(props: {
  id: Promise<number>
}): Promise<MetadataRoute.Sitemap> {
  const id = await props.id
  // Google's limit is 50,000 URLs per sitemap
  const start = id * 50000
  const end = start + 50000
  const products = await getProducts(
    `SELECT id, date FROM products WHERE id BETWEEN ${start} AND ${end}`
  )
  return products.map((product) => ({
    url: `${BASE_URL}/product/${product.id}`,
    lastModified: product.date,
  }))
}
```

## [Version History](#version-history)

VersionChanges`v16.0.0`The `id` values returned from `generateSitemaps` are now passed as a promise to the sitemap function.`v15.0.0``generateSitemaps` now generates consistent URLs between development and production`v13.3.2``generateSitemaps` introduced. In development, you can view the generated sitemap on `/.../sitemap.xml/[id]`. For example, `/product/sitemap.xml/1`.

## Next Steps

Learn how to create sitemaps for your Next.js application.

[**sitemap.xml**  
\
API Reference for the sitemap.xml file.](/docs/app/api-reference/file-conventions/metadata/sitemap)

Was this helpful?

supported.

Send
