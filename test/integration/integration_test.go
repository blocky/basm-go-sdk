package integration_test

import (
	"io/fs"
	"os"
	"path/filepath"
	"testing"

	"github.com/rogpeppe/go-internal/testscript"
	"github.com/stretchr/testify/require"
)

func Test(t *testing.T) {
	testscript.Run(t, testscript.Params{
		Dir: "script",
		Setup: func(env *testscript.Env) error {
			// Copy the wasm binary into the testscript work directory
			src := filepath.Join("testdata", "x.wasm")
			dest := filepath.Join(env.WorkDir, "x.wasm")
			wasmBinary, err := os.ReadFile(src)
			require.NoError(t, err)
			return os.WriteFile(dest, wasmBinary, fs.ModePerm)
		},
	})
}
