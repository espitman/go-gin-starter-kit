package cron

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
		if path != "./worker/cron" && path != "worker/cron/cron.go" {
			name := strings.ReplaceAll(path, "worker/cron/", "")
			name = strings.ReplaceAll(name, ".go", "")
			*files = append(*files, name)
		}

		return nil
	}
}

func Run() {
	var files []string
	err := filepath.Walk("./worker/cron", getFiles(&files))
	if err != nil {
		panic(err)
	}
	var t T
	for _, file := range files {
		reflect.ValueOf(&t).MethodByName(strings.Title(file) + "Run").Call([]reflect.Value{})
	}
}
