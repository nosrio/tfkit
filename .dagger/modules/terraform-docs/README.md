# terraform-docs

This Dagger module runs
[terraform-docs](https://terraform-docs.io/) against a Terraform directory and
returns the generated changes as a Dagger changeset.

## Requirements

- [Dagger CLI](https://docs.dagger.io/) with a working Dagger Engine
- A Terraform project in the Dagger workspace
- A `terraform-docs` configuration file available in that workspace

## Generate documentation

List available generators and run the module from the root of the Dagger
workspace:

```sh
dagger generate terraform-docs -l
dagger generate terraform-docs -y
```

`-l` lists available generators.  
`-y` applies the generated changeset to the workspace. Review the resulting diff before committing it.

The generator defaults are:

| Setting | Description | Default |
| --- | --- | --- | 
| `--base-image` | Container image | `quay.io/terraform-docs/terraform-docs` |
| `--base-image-version` | Container image version | `latest` |
| `--config` | Configuration file | `.terraform-docs.yaml` |
| `--path` | Terraform directory | `.` |

The configuration file controls the output format and destination. For example,
to inject generated documentation into a README, configure it like this:

```yaml
# .terraform-docs.yaml
formatter: markdown

output:
  file: README.md
```

The `config` and `path` inputs can be changed when invoking the generator. The
configuration file and Terraform directory should be relative to the workspace
root.
