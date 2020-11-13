package consumer

import (
	"log"
	"os"
	"path/filepath"
	"reflect"
	"strings"
)

type T struct{}

func getFiles(files *[]string) filepath.WalkFunc {
	return func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Fatal(err)
		}
		if path != "./worker/consumer" && path != "worker/consumer/consumer.go" {
			name := strings.ReplaceAll(path, "worker/consumer/", "")
			name = strings.ReplaceAll(name, ".go", "")
			*files = append(*files, name)
		}

		return nil
	}
}

func Start() {
	var files []string
	err := filepath.Walk("./worker/consumer", getFiles(&files))
	if err != nil {
		panic(err)
	}
	var t T
	for _, file := range files {
		go reflect.ValueOf(&t).MethodByName(strings.Title(file) + "Consumer").Call([]reflect.Value{})
	}

}
