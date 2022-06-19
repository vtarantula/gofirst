package util

import (
	"errors"
	"fmt"
	"io"
	"os"
)

func GetFileStat(filename string) (*os.FileInfo, error) {
	finfo, err := os.Stat(filename)
	if err != nil {
		var str_message string = fmt.Sprintf("unable to find file: %s", filename)
		return nil, errors.New(str_message)
	}
	return &finfo, nil
}

func GetFilePerm(filename string) (*os.FileMode, error) {
	finfo, err := GetFileStat(filename)
	if err != nil {
		return nil, err
	}
	fmode := (*finfo).Mode()
	return &fmode, nil
}

func CopyFile(src_filename, dest_filename string) error {
	_, err := GetFileStat(dest_filename)
	if err == nil {
		var str_message string = fmt.Sprintf("file %s already exists.", dest_filename)
		return errors.New(str_message)
	}

	fd_s, err := os.Open(src_filename)
	if err != nil {
		return err
	}
	defer fd_s.Close()

	fperm, err := GetFilePerm(src_filename)
	if err != nil {
		return err
	}

	fd_d, err := os.OpenFile(dest_filename, os.O_CREATE|os.O_APPEND|os.O_RDWR, *fperm)
	if err != nil {
		return err
	}
	defer fd_d.Close()

	buf := make([]byte, DEFAULT_BUFFER_SIZE)
	for {
		read, err := fd_s.Read(buf)
		if err == io.EOF || read == 0 {
			break
		}
		if err != nil {
			return err
		}
		fd_d.Write(buf[:read])
	}
	return nil
}
