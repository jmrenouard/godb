package main

import (
	"os"
)

func TouchFile(name string) error {
	file, err := os.OpenFile(name, os.O_RDONLY|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	return file.Close()
}

func ReadFileContent(name string) ([]byte, error) {
	file, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	fileinfo, err := file.Stat()
	if err != nil {
		return nil, err
	}

	filesize := fileinfo.Size()
	buffer := make([]byte, filesize)

	bytesread, err := file.Read(buffer)
	if err != nil {
		return nil, err
	}

	if int64(bytesread) != filesize {
		return nil, errors.New("ReadFileContent: could not read whole file")
	}

	return buffer, nil
}

func ReadTextFileContent(name string) (string, error) {
	buffer, err := ReadFileContent(name)
	if err != nil {
		return "", err
	}
	return string(buffer), nil
}
