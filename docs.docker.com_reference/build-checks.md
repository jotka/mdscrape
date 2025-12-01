---
title: "Build checks"
source: "https://docs.docker.com/reference/build-checks/"
---

# Build checks

Back

Ask AI

Start typing to search or try Ask AI.

[Contact support](https://hub.docker.com/support/contact)

[Reference](https://docs.docker.com/reference/)

- [Get started](/get-started/)
- [Guides](/guides/)
- [Manuals](/manuals/)

# Build checks

Page options

Copy page as Markdown for LLMs

View page as plain text

Ask questions with Docs AI

Claude

Open in Claude

* * *

BuildKit has built-in support for analyzing your build configuration based on a set of pre-defined rules for enforcing Dockerfile and building best practices. Adhering to these rules helps avoid errors and ensures good readability of your Dockerfile.

Checks run as a build invocation, but instead of producing a build output, it performs a series of checks to validate that your build doesn't violate any of the rules. To run a check, use the `--check` flag:

```console
$ docker build --check .
```

To learn more about how to use build checks, see [Checking your build configuration](https://docs.docker.com/build/checks/).

NameDescription[StageNameCasing](./stage-name-casing/)Stage names should be lowercase[FromAsCasing](./from-as-casing/)The 'as' keyword should match the case of the 'from' keyword[NoEmptyContinuation](./no-empty-continuation/)Empty continuation lines will become errors in a future release[ConsistentInstructionCasing](./consistent-instruction-casing/)All commands within the Dockerfile should use the same casing (either upper or lower)[DuplicateStageName](./duplicate-stage-name/)Stage names should be unique[ReservedStageName](./reserved-stage-name/)Reserved words should not be used as stage names[JSONArgsRecommended](./json-args-recommended/)JSON arguments recommended for ENTRYPOINT/CMD to prevent unintended behavior related to OS signals[MaintainerDeprecated](./maintainer-deprecated/)The MAINTAINER instruction is deprecated, use a label instead to define an image author[UndefinedArgInFrom](./undefined-arg-in-from/)FROM command must use declared ARGs[WorkdirRelativePath](./workdir-relative-path/)Relative workdir without an absolute workdir declared within the build can have unexpected results if the base image changes[UndefinedVar](./undefined-var/)Variables should be defined before their use[MultipleInstructionsDisallowed](./multiple-instructions-disallowed/)Multiple instructions of the same type should not be used in the same stage[LegacyKeyValueFormat](./legacy-key-value-format/)Legacy key/value format with whitespace separator should not be used[RedundantTargetPlatform](./redundant-target-platform/)Setting platform to predefined $TARGETPLATFORM in FROM is redundant as this is the default behavior[SecretsUsedInArgOrEnv](./secrets-used-in-arg-or-env/)Sensitive data should not be used in the ARG or ENV commands[InvalidDefaultArgInFrom](./invalid-default-arg-in-from/)Default value for global ARG results in an empty or invalid base image name[FromPlatformFlagConstDisallowed](./from-platform-flag-const-disallowed/)FROM --platform flag should not use a constant value[CopyIgnoredFile (experimental)](./copy-ignored-file/)Attempting to Copy file that is excluded by .dockerignore[InvalidDefinitionDescription (experimental)](./invalid-definition-description/)Comment for build stage or argument should follow the format: \`# \`. If this is not intended to be a description comment, add an empty line or comment between the instruction and the comment.[ExposeProtoCasing](./expose-proto-casing/)Protocol in EXPOSE instruction should be lowercase[ExposeInvalidFormat](./expose-invalid-format/)IP address and host-port mapping should not be used in EXPOSE instruction. This will become an error in a future release

[Request changes](https://github.com/docker/docs/issues/new?template=doc_issue.yml&location=https%3A%2F%2Fdocs.docker.com%2Freference%2Fbuild-checks%2F&labels=status%2Ftriage)
