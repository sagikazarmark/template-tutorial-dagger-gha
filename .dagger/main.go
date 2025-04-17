package main

import (
	"context"
	"tutorial-dagger-gha/.dagger/internal/dagger"

	"github.com/sourcegraph/conc/pool"
	_ "github.com/sourcegraph/conc/pool"
)

type Tutorial struct {
	// Project source directory
	//
	// +private
	Source *dagger.Directory
}

func New(
	// Project source directory.
	//
	// +defaultPath="/"
	// +ignore=[".devenv", ".direnv", ".github"]
	source *dagger.Directory,
) *Tutorial {
	return &Tutorial{
		Source: source,
	}
}

// Build the application and run all checks (tests, linters).
func (m *Tutorial) Check(ctx context.Context) error {
	p := pool.New().WithErrors().WithContext(ctx)

	p.Go(func(ctx context.Context) error {
		_, err := m.Build("").Sync(ctx)

		return err
	})
	p.Go(func(ctx context.Context) error {
		_, err := m.Test().Sync(ctx)

		return err
	})
	p.Go(func(ctx context.Context) error {
		_, err := m.Lint().Sync(ctx)

		return err
	})

	return p.Wait()
}

// Build the application.
func (m *Tutorial) Build(
	// Target platform in "[os]/[platform]/[version]" format (e.g., "darwin/arm64/v7", "windows/amd64", "linux/arm64").
	//
	// +optional
	platform dagger.Platform,
) *dagger.File {
	var opts dagger.GoBuildOpts

	if platform != "" {
		opts.Platform = platform
	}

	return dag.Go().Build(m.Source, opts)
}

// Run tests.
func (m *Tutorial) Test() *dagger.Container {
	return dag.Go().WithSource(m.Source).Exec([]string{"go", "test", "-race", "-v", "-shuffle=on", "./..."})
}

// Run linters.
func (m *Tutorial) Lint() *dagger.Container {
	return dag.GolangciLint().Run(m.Source)
}
