---
title: "docker trust key load"
source: "https://docs.docker.com/reference/cli/docker/trust/key/load/"
---

# docker trust key load

Back

Ask AI

Start typing to search or try Ask AI.

[Contact support](https://hub.docker.com/support/contact)

[Reference](https://docs.docker.com/reference/)

- [Get started](/get-started/)
- [Guides](/guides/)
- [Manuals](/manuals/)

# docker trust key load

DescriptionLoad a private key file for signingUsage`docker trust key load [OPTIONS] KEYFILE`

## [Description](#description)

`docker trust key load` adds private keys to the local Docker trust keystore.

To add a signer to a repository use `docker trust signer add`.

## [Options](#options)

OptionDefaultDescription`--name``signer`Name for the loaded key

## [Examples](#examples)

### [Load a single private key](#load-a-single-private-key)

For a private key `alice.pem` with permissions `-rw-------`

```console
$ docker trust key load alice.pem

Loading key from "alice.pem"...
Enter passphrase for new signer key with ID f8097df:
Repeat passphrase for new signer key with ID f8097df:
Successfully imported key from alice.pem
```

To specify a name use the `--name` flag:

```console
$ docker trust key load --name alice-key alice.pem

Loading key from "alice.pem"...
Enter passphrase for new alice-key key with ID f8097df:
Repeat passphrase for new alice-key key with ID f8097df:
Successfully imported key from alice.pem
```

Table of contents
