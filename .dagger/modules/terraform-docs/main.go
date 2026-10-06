//
// This module runs terraform-docs as a dagger generator
//

package main

import (
	"dagger/terraform-docs/internal/dagger"
)

type TerraformDocs struct {
	// +private
	BaseImageAddress string
	// +private
	Source *dagger.Directory
	// +private
	Config string
	// +private
	Path string
}

func New(
	// Current workspace auto populated by Dagger
	ws *dagger.Workspace,
	// Base image used to run terraform-docs.
	// +default="quay.io/terraform-docs/terraform-docs"
	baseImage string,
	// Base image version used to run terraform-docs.
	// +default="latest"
	baseImageVersion string,
	// Filepath to terraform-docs [config](https://terraform-docs.io/how-to/configuration-file/) file.
	// +default=".terraform-docs.yaml"
	config string,
	// Module path where the docs will be generated
	// +default="."
	path string,

) *TerraformDocs {
	return &TerraformDocs{
		BaseImageAddress: baseImage + ":" + baseImageVersion,
		Source: ws.Directory("/", dagger.WorkspaceDirectoryOpts{
			Exclude: []string{"**/.git", "**/.dagger", "**/vendor"},
		}),
		Config: config,
		Path:   path,
	}
}

// +generate
//
// Generate runs terraform-docs for the configured Terraform path and returns
// the resulting source changeset.
//
// The workspace must contain the terraform-docs configuration file specified
// by Config. Use `dagger generate terraform-docs -y` to apply the returned changeset to the
// workspace.
func (m *TerraformDocs) Generate() *dagger.Changeset {
	docs := dag.Container().
		From(m.BaseImageAddress).
		WithMountedDirectory("/", m.Source).
		WithWorkdir("/").
		WithExec([]string{
			"terraform-docs",
			"-c",
			m.Config,
			m.Path,
		}).
		Directory("/")

	return docs.Changes(m.Source)
}
