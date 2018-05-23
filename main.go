package main

import (
	"flag"
	"path/filepath"
	"os"
	"./image-converter"
)

func main() {
	path := flag.String("path", ".", "TargetDir")
	from := flag.String("from","jpeg", "from")
	to := flag.String("to", "png", "to")

	flag.Parse()

	err := convert(*path, *from, *to)
	if err != nil {
		panic(err)
	}
}

func convert(path string, from string, to string) error {
	return filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		p := image_converter.Filepath(path)

		if filepath.Ext(path) == "."+from {
			err = image_converter.Writer(p)
			return err
		}
		return nil

	})
}
