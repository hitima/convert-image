package image_converter

import (
	"os"
	"path/filepath"
	"image"
	"image/png"
	_"image/jpeg" // decode はこのインポートないとこける。40分はまった
	"fmt"
)

type Path struct {
	path string
}

func Filepath(path string) *Path {
	return &Path{path: path}
}

func (s *Path) GetPath() string {
	return s.path
}

func (s *Path) ChangeFileExtension(ext string) string {
	limit := len(s.path) - len(filepath.Ext(s.path))
	return s.path[:limit]+"."+ext
}

func Writer (p *Path) error {
	from, err := os.Open(p.GetPath())
	if err != nil {
		fmt.Println("error " + err.Error())
	}
	defer from.Close()

	to, err := os.Create(p.ChangeFileExtension("png"))
	defer from.Close()
	if err != nil {
		fmt.Println("error " + err.Error())
	}
	defer to.Close()

	img, _, err := image.Decode(from)
	if err != nil {
		fmt.Println("error " + err.Error())
	}
	println("Convert success: "+p.GetPath()+" to "+p.ChangeFileExtension("png"))
	return png.Encode(to, img)
}
