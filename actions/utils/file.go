package utils

import (
	"errors"
	"fmt"
	"io/fs"
	"os"

	"github.com/colinso/lego/actions/config"
)

type FileLocation int64

const (
	Model FileLocation = iota
	Handler
	Root
	Wire
	Config
	API
	Logic
)

var dirPaths = map[FileLocation]string{
	Root:    ".",
	Model:   "internal/models",
	Handler: "internal/handlers",
	Wire:    "internal/wire",
	Config:  "internal/config",
	API:     "internal/api",
	Logic:   "internal/logic",
}

func CreateFileForType(fLocation FileLocation, fname string) *os.File {
	pathPrefix := config.GetConfig().ProjectPath
	dirPath := fmt.Sprintf("%v/%v/%v", pathPrefix, config.GetConfig().Name, dirPaths[fLocation])
	CreatePathIfNotExists(dirPath)

	fpath := fmt.Sprintf("%v/%v.go", dirPath, fname)
	var f *os.File
	f, err := os.Create(fpath)
	if err != nil {
		panic(err)
	}
	return f
}

func CreatePathIfNotExists(path string) {
	if err := os.MkdirAll(path, os.ModePerm); err != nil && !errors.Is(err, fs.ErrExist) {
		panic(err)
	}
}
