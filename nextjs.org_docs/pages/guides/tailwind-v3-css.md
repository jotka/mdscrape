---
title: "Tailwind CSS"
source: "https://nextjs.org/docs/pages/guides/tailwind-v3-css"
---

# Tailwind CSS

Menu

Using App Router

Features available in /app

Latest Version

16.0.5

[Pages Router](/docs/pages)[Guides](/docs/pages/guides)Tailwind CSS

Copy page

# Tailwind CSS

This guide will walk you through how to install [Tailwind CSS v3](https://v3.tailwindcss.com/) in your Next.js application.

> **Good to know:** For the latest Tailwind 4 setup, see the [Tailwind CSS setup instructions](/docs/app/getting-started/css#tailwind-css).

## [Installing Tailwind v3](#installing-tailwind-v3)

Install Tailwind CSS and its peer dependencies, then run the `init` command to generate both `tailwind.config.js` and `postcss.config.js` files:

pnpmnpmyarnbun

Terminal

```
pnpm add -D tailwindcss@^3 postcss autoprefixer
npx tailwindcss init -p
```

## [Configuring Tailwind v3](#configuring-tailwind-v3)

Configure your template paths in your `tailwind.config.js` file:

tailwind.config.js

```
/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    './pages/**/*.{js,ts,jsx,tsx,mdx}',
    './components/**/*.{js,ts,jsx,tsx,mdx}',
    './app/**/*.{js,ts,jsx,tsx,mdx}',
  ],
  theme: {
    extend: {},
  },
  plugins: [],
}
```

Add the Tailwind directives to your global CSS file:

styles/globals.css

```
@tailwind base;
@tailwind components;
@tailwind utilities;
```

Import the CSS file in your `pages/_app.js` file:

pages/\_app.js

```
import '@/styles/globals.css'
 
export default function MyApp({ Component, pageProps }) {
  return <Component {...pageProps} />
}
```

## [Using classes](#using-classes)

After installing Tailwind CSS and adding the global styles, you can use Tailwind's utility classes in your application.

pages/index.tsx

TypeScript

JavaScriptTypeScript

```
export default function Page() {
  return <h1 className="text-3xl font-bold underline">Hello, Next.js!</h1>
}
```

## [Usage with Turbopack](#usage-with-turbopack)

As of Next.js 13.1, Tailwind CSS and PostCSS are supported with [Turbopack](https://turbo.build/pack/docs/features/css#tailwind-css).

Was this helpful?

supported.

Send
