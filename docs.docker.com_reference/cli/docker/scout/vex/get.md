---
title: "docker scout vex get"
source: "https://docs.docker.com/reference/cli/docker/scout/vex/get/"
---

# docker scout vex get

Back

Ask AI

Start typing to search or try Ask AI.

[Contact support](https://hub.docker.com/support/contact)

[Reference](https://docs.docker.com/reference/)

- [Get started](/get-started/)
- [Guides](/guides/)
- [Manuals](/manuals/)

# docker scout vex get

DescriptionGet VEX attestation for imageUsage`docker scout vex get OPTIONS IMAGE`

**Experimental**

**This command is experimental.**

Experimental features are intended for testing and feedback as their functionality or design may change between releases without warning or can be removed entirely in a future release.

## [Description](#description)

The docker scout vex get command gets a VEX attestation for images.

## [Options](#options)

OptionDefaultDescription`--key``https://registry.scout.docker.com/keyring/dhi/latest.pub`Signature key to use for verification`--org`Namespace of the Docker organization`-o, --output`Write the report to a file`--platform`Platform of image to analyze`--ref`Reference to use if the provided tarball contains multiple references.  
Can only be used with archive`--skip-tlog`Skip signature verification against public transaction log`--verify`Verify the signature on the attestation

Table of contents
