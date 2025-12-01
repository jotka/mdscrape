---
title: "No assign module variable"
source: "https://nextjs.org/docs/messages/no-assign-module-variable"
---

# No assign module variable

Menu

Using App Router

Features available in /app

Latest Version

16.0.5

[Docs](/docs)[Errors](/docs)No assign module variable

# No assign module variable

> Prevent assignment to the `module` variable.

## [Why This Error Occurred](#why-this-error-occurred)

A value is being assigned to the `module` variable. The `module` variable is already used and it is highly likely that assigning to this variable will cause errors.

## [Possible Ways to Fix It](#possible-ways-to-fix-it)

Use a different variable name:

example.js

```
let myModule = {...}
```

Was this helpful?

supported.

Send
