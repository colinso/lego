package utils

import (
	"errors"
	"fmt"
	"io/fs"
	"os"

	"github.com/colinso/lego/config"
)

type FileExtension string

const (
	Go   FileExtension = ".go"
	Yaml FileExtension = ".yaml"
	None FileExtension = ""
)

type FileLocation int64

const (
	Model FileLocation = iota
	Handler
	Root
	Cmd
	Wire
	Config
	API
	Service
	DockerCompose
	Dockerfile
)

var dirPaths = map[FileLocation]string{
	Root:    ".",
	Cmd:     "cmd",
	Model:   "internal/models",
	Handler: "internal/handlers",
	Wire:    "internal/wire",
	Config:  "internal/config",
	API:     "internal/api",
	Service: "internal/services",
}

func CreateFileForType(fLocation FileLocation, fname string, ext FileExtension) *os.File {
	pathPrefix := config.GetConfig().ProjectPath
	dirPath := fmt.Sprintf("%v/%v", pathPrefix, dirPaths[fLocation])
	CreatePathIfNotExists(dirPath)

	fpath := fmt.Sprintf("%v/%v%v", dirPath, fname, ext)
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
