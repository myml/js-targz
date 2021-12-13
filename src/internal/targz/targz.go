package targz

import (
	"archive/tar"
	"bufio"
	"compress/gzip"
	"io"
)

// TarGz 遍历.tar.gz文件，并返回列表信息
func TarGz(rs io.Reader, callback func(*tar.Header) error) error {
	r, err := gzip.NewReader(bufio.NewReaderSize(rs, 4*1024*1024))
	if err != nil {
		return err
	}
	tr := tar.NewReader(r)
	for {
		header, err := tr.Next()
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}
		err = callback(header)
		if err != nil {
			return err
		}
	}
	return nil
}
