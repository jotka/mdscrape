---
title: "No Cache Detected"
source: "https://nextjs.org/docs/messages/no-cache"
---

# No Cache Detected

Menu

Using App Router

Features available in /app

Latest Version

16.0.5

[Docs](/docs)[Errors](/docs)No Cache Detected

# No Cache Detected

## [Why This Error Occurred](#why-this-error-occurred)

A Next.js build was triggered in a continuous integration environment, but no cache was detected.

This results in slower builds and can hurt Next.js' filesystem cache of client-side bundles across builds.

## [Possible Ways to Fix It](#possible-ways-to-fix-it)

> **Note**: If this is a new project, or being built for the first time in your CI, you can ignore this error. However, you'll want to make sure it doesn't continue to happen and fix it if it does!

Follow the instructions in [CI Build Caching](/docs/pages/guides/ci-build-caching) to ensure your Next.js cache is persisted between builds.

Was this helpful?

supported.

Send
