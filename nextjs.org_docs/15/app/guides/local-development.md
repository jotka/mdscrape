---
title: "How to optimize your local development environment"
source: "https://nextjs.org/docs/15/app/guides/local-development"
---

# How to optimize your local development environment

Menu

App Router 15

Using App Router

Features available in /app

Version 15

15.5.6

[App Router](/docs/15/app)[Guides](/docs/15/app/guides)Development Environment

You are currently viewing documentation for version 15 of Next.js.

# How to optimize your local development environment

Next.js is designed to provide a great developer experience. As your application grows, you might notice slower compilation times during local development. This guide will help you identify and fix common compile-time performance issues.

## [Local dev vs. production](#local-dev-vs-production)

The development process with `next dev` is different than `next build` and `next start`.

`next dev` compiles routes in your application as you open or navigate to them. This enables you to start the dev server without waiting for every route in your application to compile, which is both faster and uses less memory. Running a production build applies other optimizations, like minifying files and creating content hashes, which are not needed for local development.

## [Improving local dev performance](#improving-local-dev-performance)

### [1. Check your computer's antivirus](#1-check-your-computers-antivirus)

Antivirus software can slow down file access.

Try adding your project folder to the antivirus exclusion list. While this is more common on Windows machines, we recommend this for any system with an antivirus tool installed.

### [2. Update Next.js and enable Turbopack](#2-update-nextjs-and-enable-turbopack)

Make sure you're using the latest version of Next.js. Each new version often includes performance improvements.

Turbopack is a new bundler integrated into Next.js that can improve local performance.

```
npm install next@latest
npm run dev --turbopack
```

[Learn more about Turbopack](/blog/turbopack-for-development-stable). See our [upgrade guides](/docs/15/app/guides/upgrading) and codemods for more information.

### [3. Check your imports](#3-check-your-imports)

The way you import code can greatly affect compilation and bundling time. Learn more about [optimizing package bundling](/docs/15/app/guides/package-bundling) and explore tools like [Dependency Cruiser](https://github.com/sverweij/dependency-cruiser) or [Madge](https://github.com/pahen/madge).

#### [Icon libraries](#icon-libraries)

Libraries like `@material-ui/icons`, `@phosphor-icons/react`, or `react-icons` can import thousands of icons, even if you only use a few. Try to import only the icons you need:

```
// Instead of this:
import { TriangleIcon } from '@phosphor-icons/react'
 
// Do this:
import { TriangleIcon } from '@phosphor-icons/react/dist/csr/Triangle'
```

You can often find what import pattern to use in the documentation for the icon library you're using. This example follows [`@phosphor-icons/react`](https://www.npmjs.com/package/@phosphor-icons/react#import-performance-optimization) recommendation.

Libraries like `react-icons` includes many different icon sets. Choose one set and stick with that set.

For example, if your application uses `react-icons` and imports all of these:

- `pi` (Phosphor Icons)
- `md` (Material Design Icons)
- `tb` (tabler-icons)
- `cg` (cssgg)

Combined they will be tens of thousands of modules that the compiler has to handle, even if you only use a single import from each.

#### [Barrel files](#barrel-files)

"Barrel files" are files that export many items from other files. They can slow down builds because the compiler has to parse them to find if there are side-effects in the module scope by using the import.

Try to import directly from specific files when possible. [Learn more about barrel files](https://vercel.com/blog/how-we-optimized-package-imports-in-next-js) and the built-in optimizations in Next.js.

#### [Optimize package imports](#optimize-package-imports)

Next.js can automatically optimize imports for certain packages. If you are using packages that utilize barrel files, add them to your `next.config.js`:

```
module.exports = {
  experimental: {
    optimizePackageImports: ['package-name'],
  },
}
```

Turbopack automatically analyzes imports and optimizes them. It does not require this configuration.

### [4. Check your Tailwind CSS setup](#4-check-your-tailwind-css-setup)

If you're using Tailwind CSS, make sure it's set up correctly.

A common mistake is configuring your `content` array in a way which includes `node_modules` or other large directories of files that should not be scanned.

Tailwind CSS version 3.4.8 or newer will warn you about settings that might slow down your build.

1. In your `tailwind.config.js`, be specific about which files to scan:
   
   ```
   module.exports = {
     content: [
       './src/**/*.{js,ts,jsx,tsx}', // Good
       // This might be too broad
       // It will match `packages/**/node_modules` too
       // '../../packages/**/*.{js,ts,jsx,tsx}',
     ],
   }
   ```
2. Avoid scanning unnecessary files:
   
   ```
   module.exports = {
     content: [
       // Better - only scans the 'src' folder
       '../../packages/ui/src/**/*.{js,ts,jsx,tsx}',
     ],
   }
   ```

### [5. Check custom webpack settings](#5-check-custom-webpack-settings)

If you've added custom webpack settings, they might be slowing down compilation.

Consider if you really need them for local development. You can optionally only include certain tools for production builds, or explore moving to Turbopack and using [loaders](/docs/15/app/api-reference/config/next-config-js/turbopack#supported-loaders).

### [6. Optimize memory usage](#6-optimize-memory-usage)

If your app is very large, it might need more memory.

[Learn more about optimizing memory usage](/docs/15/app/guides/memory-usage).

### [7. Server Components and data fetching](#7-server-components-and-data-fetching)

Changes to Server Components cause the entire page to re-render locally in order to show the new changes, which includes fetching new data for the component.

The experimental `serverComponentsHmrCache` option allows you to cache `fetch` responses in Server Components across Hot Module Replacement (HMR) refreshes in local development. This results in faster responses and reduced costs for billed API calls.

[Learn more about the experimental option](/docs/15/app/api-reference/config/next-config-js/serverComponentsHmrCache).

### [8. Consider local development over Docker](#8-consider-local-development-over-docker)

If you're using Docker for development on Mac or Windows, you may experience significantly slower performance compared to running Next.js locally.

Docker's filesystem access on Mac and Windows can cause Hot Module Replacement (HMR) to take seconds or even minutes, while the same application runs with fast HMR when developed locally.

This performance difference is due to how Docker handles filesystem operations outside of Linux environments. For the best development experience:

- Use local development (`npm run dev` or `pnpm dev`) instead of Docker during development
- Reserve Docker for production deployments and testing production builds
- If you must use Docker for development, consider using Docker on a Linux machine or VM

[Learn more about Docker deployment](/docs/15/app/getting-started/deploying#docker) for production use.

## [Tools for finding problems](#tools-for-finding-problems)

### [Detailed fetch logging](#detailed-fetch-logging)

Use the `logging.fetches` option in your `next.config.js` file, to see more detailed information about what's happening during development:

```
module.exports = {
  logging: {
    fetches: {
      fullUrl: true,
    },
  },
}
```

[Learn more about fetch logging](/docs/15/app/api-reference/config/next-config-js/logging).

## [Turbopack tracing](#turbopack-tracing)

Turbopack tracing is a tool that helps you understand the performance of your application during local development. It provides detailed information about the time taken for each module to compile and how they are related.

1. Make sure you have the latest version of Next.js installed.
2. Generate a Turbopack trace file:
   
   ```
   NEXT_TURBOPACK_TRACING=1 npm run dev
   ```
3. Navigate around your application or make edits to files to reproduce the problem.
4. Stop the Next.js development server.
5. A file called `trace-turbopack` will be available in the `.next` folder.
6. You can interpret the file using `npx next internal trace [path-to-file]`:
   
   ```
   npx next internal trace .next/trace-turbopack
   ```
   
   On versions where `trace` is not available, the command was named `turbo-trace-server`:
   
   ```
   npx next internal turbo-trace-server .next/trace-turbopack
   ```
7. Once the trace server is running you can view the trace at [https://trace.nextjs.org/](https://trace.nextjs.org/).
8. By default the trace viewer will aggregate timings, in order to see each individual time you can switch from "Aggregated in order" to "Spans in order" at the top right of the viewer.

## [Still having problems?](#still-having-problems)

Share the trace file generated in the [Turbopack Tracing](#turbopack-tracing) section and share it on [GitHub Discussions](https://github.com/vercel/next.js/discussions) or [Discord](https://nextjs.org/discord).

Was this helpful?

supported.

Send
