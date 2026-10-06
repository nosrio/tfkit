# tfkit

`tfkit` is a Dagger workspace for running CI checks against Terraform projects.
Checks are organized as Dagger modules under `.dagger/modules`.

## Requirements

- [Dagger CLI](https://docs.dagger.io/) with a working Dagger Engine

## Available modules

- [`terraform-docs`](.dagger/modules/terraform-docs/README.md): generates
  Terraform documentation using
  [terraform-docs](https://terraform-docs.io/).

## Usage

```sh
dagger generate -l
```

List available Dagger generators with `dagger generate -l`. For module-specific
instructions, see that module's README.

## Development

Module source code and dependency manifests live under `.dagger/modules`.
Dagger workspace configuration is in [`dagger.toml`](dagger.toml), and pinned
dependency references are recorded in [`dagger.lock`](dagger.lock).

Commit changes to `dagger.lock` when Dagger module dependencies are updated.
