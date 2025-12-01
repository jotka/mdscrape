---
title: "docker image inspect"
source: "https://docs.docker.com/reference/cli/docker/image/inspect/"
---

# docker image inspect

Back

Ask AI

Start typing to search or try Ask AI.

[Contact support](https://hub.docker.com/support/contact)

[Reference](https://docs.docker.com/reference/)

- [Get started](/get-started/)
- [Guides](/guides/)
- [Manuals](/manuals/)

# docker image inspect

DescriptionDisplay detailed information on one or more imagesUsage`docker image inspect [OPTIONS] IMAGE [IMAGE...]`

## [Description](#description)

Display detailed information on one or more images

## [Options](#options)

OptionDefaultDescription`-f, --format`Format output using a custom template:  
'json': Print in JSON format  
'TEMPLATE': Print output using the given Go template.  
Refer to [https://docs.docker.com/go/formatting/](https://docs.docker.com/go/formatting/) for more information about formatting output with templates`--platform`API 1.49+ Inspect a specific platform of the multi-platform image.  
If the image or the server is not multi-platform capable, the command will error out if the platform does not match.  
'os\[/arch\[/variant]]': Explicit platform (eg. linux/amd64)

Table of contents
