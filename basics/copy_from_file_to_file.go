package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
)

func OpenFile(filename string, fileType string) (*os.File, error) {
	if filename != "" {
		var file *os.File
		var err error

		if fileType == "source" {
			file, err = os.OpenFile(filename, os.O_RDONLY, 0777)
		} else if fileType == "destination" {
			file, err = os.OpenFile(filename, os.O_CREATE|os.O_TRUNC, 0777)
		} else {
			return nil, nil
		}

		if err != nil {
			if os.IsNotExist(err) {
				return nil, errors.New(fmt.Sprintf("no such %s file exists", fileType))
			} else {
				return nil, errors.New("some error with source file")
			}
		} else {
			return file, nil
		}

	} else {
		return nil, errors.New("no source file")
	}
}

func Copy(from string, to string, offset int, limit int) {
	source, err := OpenFile(from, "source")

	if err != nil {
		fmt.Println(err)
	} else {
		dest, err := OpenFile(to, "destination")

		if err != nil {
			fmt.Println(err)
		}

		ret, err := source.Seek(int64(offset), io.SeekStart)

		if err != nil {
			fmt.Println(err)
		} else {
			written, err := io.CopyN(dest, source, int64(limit))

			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Printf("Copied %d bytes with offset %d.", written, ret)
			}
		}
	}
}

func main() {
	var from string
	var to string
	var offset int
	var limit int

	flag.StringVar(&from, "from", "", "File to copy from.")
	flag.StringVar(&to, "to", "", "File to copy to.")
	flag.IntVar(&offset, "offset", 0, "Offset bytes to copy from.")
	flag.IntVar(&limit, "limit", 0, "Limit bytes to copy upto.")

	flag.Parse()

	Copy(from, to, offset, limit)
}
