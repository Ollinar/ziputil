package ziputil

import (
	"archive/zip"
	"io"
	"os"
	"path/filepath"
	"time"
)

// ExtractFromPath will extract all of the contents from srcFile to destDir preserving modified times.
func ExtractFromPath(srcFile string, destDir string) error {
	srcZ, err := zip.OpenReader(srcFile)
	if err != nil {
		return err
	}
	defer srcZ.Close()
	err = os.MkdirAll(destDir, os.ModeAppend)
	if err != nil {
		return err
	}
	dirModTimeMap := make(map[string]time.Time, len(srcZ.File))
	for _, zf := range srcZ.File {
		fPath := filepath.Join(destDir, zf.Name)
		if zf.FileInfo().IsDir() {
			err = os.MkdirAll(fPath, os.ModeAppend)
			if err != nil {
				return err
			}
			dirModTimeMap[fPath] = zf.Modified
			continue
		} else {
			err = os.MkdirAll(filepath.Dir(fPath), os.ModeAppend)
			if err != nil {
				return err
			}
		}

		srcF, err := zf.Open()
		if err != nil {
			return err
		}

		f, err := os.Create(fPath)
		if err != nil {
			srcF.Close()
			return err
		}
		_, err = io.Copy(f, srcF)
		if err != nil {
			srcF.Close()
			f.Close()
			return err
		}
		err = os.Chtimes(f.Name(), time.Time{}, zf.Modified)
		if err != nil {
			srcF.Close()
			f.Close()
			return err
		}
		srcF.Close()
		f.Close()
	}

	// Preserve the mod times of all dir.
	// Doing this after all files are extracted becasue creating a file inside a dir automatically changes the modtime
	for fName, tm := range dirModTimeMap {
		err = os.Chtimes(fName, time.Time{}, tm)
		if err != nil {
			return err
		}
	}

	return nil
}
