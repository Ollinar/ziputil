package ziputil

import (
	"archive/zip"
	"io"
	"os"
)

// AddFilestoZipWriter add the files passed in to the dest.
// files should be a map of paths. path of the source file as key and path inside the zip as value.
// This wont close the dest.
func AddFilesToZipWriter(dest *zip.Writer, compressionMethod uint16, files map[string]string) error {
	for srcPath, zipPath := range files {
		inf, err := os.Stat(srcPath)
		if err != nil {
			return err
		}
		zInf, err := zip.FileInfoHeader(inf)
		if err != nil {
			return err
		}
		zInf.Name = zipPath
		zInf.Method = compressionMethod
		zw, err := dest.CreateHeader(zInf)
		if err != nil {
			return err
		}
		f, err := os.Open(srcPath)
		if err != nil {
			return err
		}
		_, err = io.Copy(zw, f)
		if err != nil {
			f.Close()
			return err
		}
		f.Close()
	}

	return nil
}

// AddFilesToWriter creates a *zip.Writer with w and write the files in there like AddFilesToZipWriter does.
// files should be a map of paths. path of the source file as key and path inside the zip as value.
func AddFilesToWriter(w io.Writer, compressionMethod uint16, files map[string]string) error {
	dest := zip.NewWriter(w)
	defer dest.Close()
	return AddFilesToZipWriter(dest, compressionMethod, files)

}
