package zip

import (
	"archive/zip"
	"gostd/tools"
	"io"
	"os"
	"testing"
)

func TestZipWrite(t *testing.T) {
	fw, err := os.Create("basic.zip")
	tools.Err(err)
	defer fw.Close()

	zw := zip.NewWriter(fw)
	defer zw.Close()

	fr, err := os.Open("readme.md")
	tools.Err(err)
	defer fr.Close()

	fi, err := fr.Stat()
	tools.Err(err)
	th, err := zip.FileInfoHeader(fi)
	tools.Err(err)
	w, err := zw.CreateHeader(th)
	tools.Err(err)

	_, err = io.Copy(w, fr)
	tools.Err(err)
}

func TestZipRead(t *testing.T) {
	zr, err := zip.OpenReader("basic.zip")
	defer zr.Close()
	tools.Err(err)

	for _, file := range zr.File {
		fr, err := file.Open()
		tools.Err(err)

		fw, err := os.OpenFile(file.Name, os.O_CREATE|os.O_RDWR|os.O_TRUNC, file.Mode())
		tools.Err(err)

		_, err = io.Copy(fw, fr)
		tools.Err(err)
	}
}
