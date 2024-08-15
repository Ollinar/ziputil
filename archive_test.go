package ziputil_test

import (
	"archive/zip"
	"os"
	"testing"

	"github.com/Ollinar/ziputil"
	"github.com/stretchr/testify/require"
)

func TestAddFilesToZipWriter(t *testing.T) {
	f, err := os.Create("testResult/AddFilesToZipWriter.zip")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	zw := zip.NewWriter(f)
	defer zw.Close()
	err = ziputil.AddFilesToZipWriter(zw, zip.Deflate, map[string]string{
		"testSample/testz/1.txt":      "testz/1.txt",
		"testSample/testz/dir1/2.txt": "testz/dir1/2.txt",
	})
	require.NoError(t, err)
}
