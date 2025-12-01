---
title: "Project Structure and Organization"
source: "https://nextjs.org/docs/pages/getting-started/project-structure"
---

# Project Structure and Organization

Menu

Using App Router

Features available in /app

Latest Version

16.0.5

[Pages Router](/docs/pages)[Getting Started](/docs/pages/getting-started)Project Structure

Copy page

# Project Structure and Organization

This page provides an overview of **all** the folder and file conventions in Next.js, and recommendations for organizing your project.

## [Folder and file conventions](#folder-and-file-conventions)

### [Top-level folders](#top-level-folders)

Top-level folders are used to organize your application's code and static assets.

![Route segments to path segments](/_next/image?url=https%3A%2F%2Fh8DxKfmAPhn8O0p3.public.blob.vercel-storage.com%2Fdocs%2Flight%2Ftop-level-folders.png&w=3840&q=75)![Route segments to path segments](/_next/image?url=https%3A%2F%2Fh8DxKfmAPhn8O0p3.public.blob.vercel-storage.com%2Fdocs%2Fdark%2Ftop-level-folders.png&w=3840&q=75)

[`app`](/docs/app)App Router[`pages`](/docs/pages/building-your-application/routing)Pages Router[`public`](/docs/app/api-reference/file-conventions/public-folder)Static assets to be served[`src`](/docs/app/api-reference/file-conventions/src-folder)Optional application source folder

### [Top-level files](#top-level-files)

Top-level files are used to configure your application, manage dependencies, run proxy, integrate monitoring tools, and define environment variables.

**Next.js**[`next.config.js`](/docs/app/api-reference/config/next-config-js)Configuration file for Next.js[`package.json`](/docs/app/getting-started/installation#manual-installation)Project dependencies and scripts[`instrumentation.ts`](/docs/app/guides/instrumentation)OpenTelemetry and Instrumentation file[`proxy.ts`](/docs/app/api-reference/file-conventions/proxy)Next.js request proxy[`.env`](/docs/app/guides/environment-variables)Environment variables[`.env.local`](/docs/app/guides/environment-variables)Local environment variables[`.env.production`](/docs/app/guides/environment-variables)Production environment variables[`.env.development`](/docs/app/guides/environment-variables)Development environment variables[`eslint.config.mjs`](/docs/app/api-reference/config/eslint)Configuration file for ESLint`.gitignore`Git files and folders to ignore`next-env.d.ts`TypeScript declaration file for Next.js`tsconfig.json`Configuration file for TypeScript`jsconfig.json`Configuration file for JavaScript

### [File conventions](#file-conventions)

[`_app`](/docs/pages/building-your-application/routing/custom-app)`.js` `.jsx` `.tsx`Custom App[`_document`](/docs/pages/building-your-application/routing/custom-document)`.js` `.jsx` `.tsx`Custom Document[`_error`](/docs/pages/building-your-application/routing/custom-error#more-advanced-error-page-customizing)`.js` `.jsx` `.tsx`Custom Error Page[`404`](/docs/pages/building-your-application/routing/custom-error#404-page)`.js` `.jsx` `.tsx`404 Error Page[`500`](/docs/pages/building-your-application/routing/custom-error#500-page)`.js` `.jsx` `.tsx`500 Error Page

### [Routes](#routes)

**Folder convention**[`index`](/docs/pages/building-your-application/routing/pages-and-layouts#index-routes)`.js` `.jsx` `.tsx`Home page[`folder/index`](/docs/pages/building-your-application/routing/pages-and-layouts#index-routes)`.js` `.jsx` `.tsx`Nested page**File convention**[`index`](/docs/pages/building-your-application/routing/pages-and-layouts#index-routes)`.js` `.jsx` `.tsx`Home page[`file`](/docs/pages/building-your-application/routing/pages-and-layouts)`.js` `.jsx` `.tsx`Nested page

### [Dynamic routes](#dynamic-routes-1)

**Folder convention**[`[folder]/index`](/docs/pages/building-your-application/routing/dynamic-routes)`.js` `.jsx` `.tsx`Dynamic route segment[`[...folder]/index`](/docs/pages/building-your-application/routing/dynamic-routes#catch-all-segments)`.js` `.jsx` `.tsx`Catch-all route segment[`[[...folder]]/index`](/docs/pages/building-your-application/routing/dynamic-routes#optional-catch-all-segments)`.js` `.jsx` `.tsx`Optional catch-all route segment**File convention**[`[file]`](/docs/pages/building-your-application/routing/dynamic-routes)`.js` `.jsx` `.tsx`Dynamic route segment[`[...file]`](/docs/pages/building-your-application/routing/dynamic-routes#catch-all-segments)`.js` `.jsx` `.tsx`Catch-all route segment[`[[...file]]`](/docs/pages/building-your-application/routing/dynamic-routes#optional-catch-all-segments)`.js` `.jsx` `.tsx`Optional catch-all route segment

Was this helpful?

supported.

Send
