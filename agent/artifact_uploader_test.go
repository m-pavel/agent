package agent

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"
	"testing"

	"github.com/buildkite/agent/api"
	"github.com/stretchr/testify/assert"
)

func TestCollectArtifactsForUploading(t *testing.T) {
	wd, _ := os.Getwd()
	rootDir := filepath.Join(wd, "..")
	os.Chdir(rootDir)

	uploader := ArtifactUploader{
		Paths: strings.Join([]string{
			filepath.FromSlash(`test/fixtures/artifacts/**/*.jpg`),
			filepath.Join(rootDir, filepath.FromSlash(`test/fixtures/artifacts/**/*.gif`)),
		}, ArtifactPathDelimiter),
	}

	artifacts, err := uploader.Collect()
	if err != nil {
		t.Fatal(err)
	}

	var testCases = []struct {
		Search   string
		Path     string
		GlobPath string
		FileSize int
		Sha1Sum  string
	}{
		{
			Search:   `Mr Freeze.jpg`,
			Path:     `test/fixtures/artifacts/Mr Freeze.jpg`,
			GlobPath: `test/fixtures/artifacts/**/*.jpg`,
			FileSize: 362371,
			Sha1Sum:  "f5bc7bc9f5f9c3e543dde0eb44876c6f9acbfb6b",
		},
		{
			Search:   `Commando.jpg`,
			Path:     `test/fixtures/artifacts/folder/Commando.jpg`,
			GlobPath: `test/fixtures/artifacts/**/*.jpg`,
			FileSize: 113000,
			Sha1Sum:  "811d7cb0317582e22ebfeb929d601cdabea4b3c0",
		},
		{
			Search:   `The Terminator.jpg`,
			Path:     `test/fixtures/artifacts/this is a folder with a space/The Terminator.jpg`,
			GlobPath: `test/fixtures/artifacts/**/*.jpg`,
			FileSize: 47301,
			Sha1Sum:  "ed76566ede9cb6edc975fcadca429665aad8785a",
		},
		{
			Search:   `Smile.gif`,
			Path:     `test/fixtures/artifacts/gifs/Smile.gif`,
			GlobPath: path.Join(filepath.ToSlash(rootDir), `test/fixtures/artifacts/**/*.gif`),
			FileSize: 2038453,
			Sha1Sum:  "bd4caf2e01e59777744ac1d52deafa01c2cb9bfd",
		},
	}

	for idx, testCase := range testCases {
		t.Run(fmt.Sprintf("Artifact #%d %s", idx+1, testCase.Search), func(t *testing.T) {
			var a *api.Artifact

			// find the artifact in the returned set
			for _, candidate := range artifacts {
				if filepath.Base(candidate.Path) == testCase.Search {
					a = candidate
				}
			}

			if a == nil {
				t.Fatalf("Failed to find an artifact for %q", testCase.Search)
			}

			assert.Equal(t, a.Path, filepath.FromSlash(testCase.Path))
			assert.Equal(t, a.AbsolutePath, filepath.Join(rootDir, filepath.FromSlash(testCase.Path)))
			assert.Equal(t, a.GlobPath, filepath.FromSlash(testCase.GlobPath))
			assert.Equal(t, int(a.FileSize), testCase.FileSize)
			assert.Equal(t, a.Sha1Sum, testCase.Sha1Sum)
		})
	}

	// // Need to trim the first character because it's path doesn't contain
	// // the root, which in this case is / or x:\\
	// a = findArtifact(artifacts, "Smile.gif")
	// assert.NotNil(t, a)
	// gifPath := filepath.Join(root, filepath.Join("test/fixtures/artifacts/gifs/Smile.gif"))
	// if runtime.GOOS == "windows" {
	// 	assert.Equal(t, a.Path, gifPath[3:])
	// } else {
	// 	assert.Equal(t, a.Path, gifPath[1:])
	// }
	// assert.Equal(t, a.AbsolutePath, filepath.Join(root, "test/fixtures/artifacts/gifs/Smile.gif"))
	// assert.Equal(t, a.GlobPath, filepath.Join(root, "test/fixtures/artifacts/**/*.gif"))
	// assert.Equal(t, int(a.FileSize), 2038453)
	// assert.Equal(t, a.Sha1Sum, "bd4caf2e01e59777744ac1d52deafa01c2cb9bfd")
}
