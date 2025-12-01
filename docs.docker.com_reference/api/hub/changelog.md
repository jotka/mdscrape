---
title: "Docker Hub API changelog"
source: "https://docs.docker.com/reference/api/hub/changelog/"
---

# Docker Hub API changelog

Back

Ask AI

Start typing to search or try Ask AI.

[Contact support](https://hub.docker.com/support/contact)

[Reference](https://docs.docker.com/reference/)

- [Get started](/get-started/)
- [Guides](/guides/)
- [Manuals](/manuals/)

# Docker Hub API changelog

Page options

Copy page as Markdown for LLMs

View page as plain text

Ask questions with Docs AI

Claude

Open in Claude

Table of contents

* * *

Here you can learn about the latest changes, new features, bug fixes, and known issues for Docker Service APIs.

* * *

## [2025-11-21](#2025-11-21)

### [Updates](#updates)

- Add missing `expires_at` fields on [PAT management](/reference/api/hub/latest/#tag/access-tokens) endpoints.

## [2025-09-25](#2025-09-25)

### [Updates](#updates-1)

- Fix [Assign repository group](/reference/api/hub/latest/#tag/repositories/operation/CreateRepositoryGroup) endpoints request/response

* * *

## [2025-09-19](#2025-09-19)

### [New](#new)

- Add [Create repository](/reference/api/hub/latest/#tag/repositories/operation/CreateRepository) endpoints for a given `namespace`.
- Add [Get repository](/reference/api/hub/latest/#tag/repositories/operation/GetRepository) endpoints for a given `namespace`.
- Add [Check repository](/reference/api/hub/latest/#tag/repositories/operation/CheckRepository) endpoints for a given `namespace`.

### [Deprecations](#deprecations)

- [Deprecate POST /v2/repositories](/reference/api/hub/deprecated/#deprecate-legacy-createrepository)
- [Deprecate POST /v2/repositories/{namespace}](/reference/api/hub/deprecated/#deprecate-legacy-createrepository)
- [Deprecate GET /v2/repositories/{namespace}/{repository}](/reference/api/hub/deprecated/#deprecate-legacy-getrepository)
- [Deprecate HEAD /v2/repositories/{namespace}/{repository}](/reference/api/hub/deprecated/#deprecate-legacy-getrepository)

* * *

## [2025-07-29](#2025-07-29)

### [New](#new-1)

- Add [Update repository immutable tags settings](/reference/api/hub/latest/#tag/repositories/operation/UpdateRepositoryImmutableTags) endpoints for a given `namespace` and `repository`.
- Add [Verify repository immutable tags](/reference/api/hub/latest/#tag/repositories/operation/VerifyRepositoryImmutableTags) endpoints for a given `namespace` and `repository`.

* * *

## [2025-06-27](#2025-06-27)

### [New](#new-2)

- Add [List repositories](/reference/api/hub/latest/#tag/repositories/operation/listNamespaceRepositories) endpoints for a given `namespace`.

### [Deprecations](#deprecations-1)

- [Deprecate /v2/repositories/{namespace}](/reference/api/hub/deprecated/#deprecate-legacy-listnamespacerepositories)

* * *

## [2025-03-25](#2025-03-25)

### [New](#new-3)

- Add [APIs](/reference/api/hub/latest/#tag/org-access-tokens) for organization access token (OATs) management.

* * *

## [2025-03-18](#2025-03-18)

### [New](#new-4)

- Add access to [audit logs](/reference/api/hub/latest/#tag/audit-logs) for org access tokens.

[Edit this page](https://github.com/docker/docs/edit/main/content/reference/api/hub/changelog.md)

[Request changes](https://github.com/docker/docs/issues/new?template=doc_issue.yml&location=https%3A%2F%2Fdocs.docker.com%2Freference%2Fapi%2Fhub%2Fchangelog%2F&labels=status%2Ftriage)

Table of contents
