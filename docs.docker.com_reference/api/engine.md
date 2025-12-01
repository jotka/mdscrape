---
title: "Docker Engine API"
source: "https://docs.docker.com/reference/api/engine/"
---

# Docker Engine API

Back

Ask AI

Start typing to search or try Ask AI.

[Contact support](https://hub.docker.com/support/contact)

[Reference](https://docs.docker.com/reference/)

- [Get started](/get-started/)
- [Guides](/guides/)
- [Manuals](/manuals/)

# Docker Engine API

Page options

Copy page as Markdown for LLMs

View page as plain text

Ask questions with Docs AI

Claude

Open in Claude

Table of contents

* * *

Docker provides an API for interacting with the Docker daemon (called the Docker Engine API), as well as SDKs for Go and Python. The SDKs allow you to efficiently build and scale Docker apps and solutions. If Go or Python don't work for you, you can use the Docker Engine API directly.

For information about Docker Engine SDKs, see [Develop with Docker Engine SDKs](https://docs.docker.com/reference/api/engine/sdk/).

The Docker Engine API is a RESTful API accessed by an HTTP client such as `wget` or `curl`, or the HTTP library which is part of most modern programming languages.

## [View the API reference](#view-the-api-reference)

You can [view the reference for the latest version of the API](https://docs.docker.com/reference/api/engine/version/v1.52/) or [choose a specific version](/reference/api/engine/#api-version-matrix).

## [Versioned API and SDK](#versioned-api-and-sdk)

The version of the Docker Engine API you should use depends upon the version of your Docker daemon and Docker client.

A given version of the Docker Engine SDK supports a specific version of the Docker Engine API, as well as all earlier versions. If breaking changes occur, they are documented prominently.

> Note
> 
> The Docker daemon and client don't necessarily need to be the same version at all times. However, keep the following in mind.
> 
> - If the daemon is newer than the client, the client doesn't know about new features or deprecated API endpoints in the daemon.
> - If the client is newer than the daemon, the client can request API endpoints that the daemon doesn't know about.

A new version of the API is released when new features are added. The Docker API is backward-compatible, so you don't need to update code that uses the API unless you need to take advantage of new features.

To see the highest version of the API your Docker daemon and client support, use `docker version`:

```console
$ docker version
Client: Docker Engine - Community
 Version:           29.0.4
 API version:       1.52
 Go version:        go1.25.4
 Git commit:        3247a5a
 Built:             Mon Nov 24 21:59:50 2025
 OS/Arch:           linux/arm64
 Context:           default

Server: Docker Engine - Community
 Engine:
  Version:          29.0.4
  API version:      1.52 (minimum version 1.44)
  Go version:       go1.25.4
  Git commit:       4612690
  Built:            Mon Nov 24 21:59:50 2025
  OS/Arch:          linux/arm64
  ...
```

You can specify the API version to use in any of the following ways:

- When using the SDK, use the latest version. At a minimum, use the version that incorporates the API version with the features you need.
- When using `curl` directly, specify the version as the first part of the URL. For instance, if the endpoint is `/containers/` you can use `/v1.52/containers/`.
- To force the Docker CLI or the Docker Engine SDKs to use an older version of the API than the version reported by `docker version`, set the environment variable `DOCKER_API_VERSION` to the correct version. This works on Linux, Windows, or macOS clients.
  
  ```console
  $ DOCKER_API_VERSION=1.51
  ```
  
  While the environment variable is set, that version of the API is used, even if the Docker daemon supports a newer version. This environment variable disables API version negotiation, so you should only use it if you must use a specific version of the API, or for debugging purposes.
- The Docker Go SDK allows you to enable API version negotiation, automatically selects an API version that's supported by both the client and the Docker Engine that's in use.
- For the SDKs, you can also specify the API version programmatically as a parameter to the `client` object. See the [Go constructor](https://pkg.go.dev/github.com/docker/docker/client#NewClientWithOpts) or the [Python SDK documentation for `client`](https://docker-py.readthedocs.io/en/stable/client.html).

### [API version matrix](#api-version-matrix)

Docker versionMaximum API versionChange log29.0[1.52](/reference/api/engine/version/v1.52/)[changes](/reference/api/engine/version-history/#v152-api-changes)28.3[1.51](/reference/api/engine/version/v1.51/)[changes](/reference/api/engine/version-history/#v151-api-changes)28.2[1.50](/reference/api/engine/version/v1.50/)[changes](/reference/api/engine/version-history/#v150-api-changes)28.1[1.49](/reference/api/engine/version/v1.49/)[changes](/reference/api/engine/version-history/#v149-api-changes)28.0[1.48](/reference/api/engine/version/v1.48/)[changes](/reference/api/engine/version-history/#v148-api-changes)27.5[1.47](/reference/api/engine/version/v1.47/)[changes](/reference/api/engine/version-history/#v147-api-changes)27.4[1.47](/reference/api/engine/version/v1.47/)[changes](/reference/api/engine/version-history/#v147-api-changes)27.3[1.47](/reference/api/engine/version/v1.47/)[changes](/reference/api/engine/version-history/#v147-api-changes)27.2[1.47](/reference/api/engine/version/v1.47/)[changes](/reference/api/engine/version-history/#v147-api-changes)27.1[1.46](/reference/api/engine/version/v1.46/)[changes](/reference/api/engine/version-history/#v146-api-changes)27.0[1.46](/reference/api/engine/version/v1.46/)[changes](/reference/api/engine/version-history/#v146-api-changes)26.1[1.45](/reference/api/engine/version/v1.45/)[changes](/reference/api/engine/version-history/#v145-api-changes)26.0[1.45](/reference/api/engine/version/v1.45/)[changes](/reference/api/engine/version-history/#v145-api-changes)25.0[1.44](/reference/api/engine/version/v1.44/)[changes](/reference/api/engine/version-history/#v144-api-changes)24.01.43[changes](/reference/api/engine/version-history/#v143-api-changes)23.01.42[changes](/reference/api/engine/version-history/#v142-api-changes)20.101.41[changes](/reference/api/engine/version-history/#v141-api-changes)19.031.40[changes](/reference/api/engine/version-history/#v140-api-changes)18.091.39[changes](/reference/api/engine/version-history/#v139-api-changes)18.061.38[changes](/reference/api/engine/version-history/#v138-api-changes)18.051.37[changes](/reference/api/engine/version-history/#v137-api-changes)18.041.37[changes](/reference/api/engine/version-history/#v137-api-changes)18.031.37[changes](/reference/api/engine/version-history/#v137-api-changes)18.021.36[changes](/reference/api/engine/version-history/#v136-api-changes)17.121.35[changes](/reference/api/engine/version-history/#v135-api-changes)17.111.34[changes](/reference/api/engine/version-history/#v134-api-changes)17.101.33[changes](/reference/api/engine/version-history/#v133-api-changes)17.091.32[changes](/reference/api/engine/version-history/#v132-api-changes)17.071.31[changes](/reference/api/engine/version-history/#v131-api-changes)17.061.30[changes](/reference/api/engine/version-history/#v130-api-changes)17.051.29[changes](/reference/api/engine/version-history/#v129-api-changes)17.041.28[changes](/reference/api/engine/version-history/#v128-api-changes)17.03.11.27[changes](/reference/api/engine/version-history/#v127-api-changes)17.031.26[changes](/reference/api/engine/version-history/#v126-api-changes)1.13.11.26[changes](/reference/api/engine/version-history/#v126-api-changes)1.131.25[changes](/reference/api/engine/version-history/#v125-api-changes)1.121.24[changes](/reference/api/engine/version-history/#v124-api-changes)

### [Deprecated API versions](#deprecated-api-versions)

API versions before v1.44 are deprecated. You can find archived documentation for deprecated versions of the API in the code repository on GitHub:

- [Documentation for API versions 1.24–1.43](https://github.com/moby/moby/tree/28.x/docs/api).
- [Documentation for API versions 1.18–1.23](https://github.com/moby/moby/tree/v25.0.0/docs/api).
- [Documentation for API versions 1.17 and before](https://github.com/moby/moby/tree/v1.9.1/docs/reference/api).

[Edit this page](https://github.com/docker/docs/edit/main/content/reference/api/engine/_index.md)

[Request changes](https://github.com/docker/docs/issues/new?template=doc_issue.yml&location=https%3A%2F%2Fdocs.docker.com%2Freference%2Fapi%2Fengine%2F&labels=status%2Ftriage)

Table of contents
