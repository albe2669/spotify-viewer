// Module for testing the go backend code of spotify-viewer

package main

import (
	"context"
	"dagger/spotify-viewer/internal/dagger"
)

type SpotifyViewer struct{}

// Build a ready-to-use development environment
func (m *SpotifyViewer) BuildEnv(source *dagger.Directory) *dagger.Container {
	// create a Dagger cache volume for dependencies
	goCache := dag.CacheVolume("go-modules")
	return dag.Go().WithSource(source).WithModuleCache(goCache).Container().WithExec([]string{"go", "generate", "./ent"})
}

// Runs the `spotify-viewer` module tests
func (m *SpotifyViewer) Test(ctx context.Context, source *dagger.Directory) (string, error) {
	output, err := m.BuildEnv(source). // call the test runner
						WithExec([]string{"go", "test", "./lib/..."}).Stdout(ctx)
	if err != nil {
		return "", err
	}
	println(output)

	return "", nil
}

// Build the application container
// To Export please call it with ... export --path=./dist from the source folder
func (m *SpotifyViewer) Build(source *dagger.Directory) *dagger.Directory {
	return m.BuildEnv(source).
		WithExec([]string{"go", "build", "-C", "cmd", "-o", "../dist/backend"}).
		Directory("./dist")
}
