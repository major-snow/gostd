package tar

import (
	"archive/tar"
	"gostd/tools"
	"io"
	"os"
	"testing"
)

// 打包就是将指定的文件目录打包成一个文件，便于存放下载等

// 打包的基本流程：创建打包对象，写入打包的文件信息，写入文件
func TestTarWrite(t *testing.T) {
	// 创建一个文件作为writer
	fw, err := os.Create("basic.tar")
	tools.Err(err)
	defer fw.Close()

	// 创建打包写入对象
	tw := tar.NewWriter(fw)
	defer tw.Close()

	// 读取需要打包的文件
	fr, err := os.Open("file1")
	tools.Err(err)
	defer fr.Close()

	// 写入待打包的文件信息
	fi, err := fr.Stat()
	tools.Err(err)
	th, err := tar.FileInfoHeader(fi, "")
	tools.Err(err)
	err = tw.WriteHeader(th)
	tools.Err(err)

	// 写入文件
	_, err = io.Copy(tw, fr)
	tools.Err(err)
}

// 接包的基本流程：读取包，遍历处理，根据相应类型处理
func TestTarRead(t *testing.T) {
	// 读取包，创建打包读取对象
	fr, err := os.Open("basic.tar")
	tools.Err(err)
	tr := tar.NewReader(fr)

	// 遍历处理
	for head, err := tr.Next(); err != io.EOF; head, err = tr.Next() {
		tools.Err(err)

		// 根据head信息处理
		fi := head.FileInfo()
		fw, err := os.Create(fi.Name())
		defer fw.Close()
		tools.Err(err)

		_, err = io.Copy(fw, tr)
		tools.Err(err)
	}
}
