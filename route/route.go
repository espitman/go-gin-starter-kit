package route

import (
	"log"
	"os"
	"path/filepath"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin"
)

type T struct {
	router *gin.Engine
}

func getFiles(files *[]string) filepath.WalkFunc {
	return func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Fatal(err)
		}
		if path != "./route" && path != "route/route.go" {
			name := strings.ReplaceAll(path, "route/", "")
			name = strings.ReplaceAll(name, ".go", "")
			*files = append(*files, name)
		}

		return nil
	}
}

func GetRoutes(router *gin.Engine) {
	var files []string
	err := filepath.Walk("./route", getFiles(&files))
	if err != nil {
		panic(err)
	}
	var t T
	t.router = router
	for _, file := range files {
		reflect.ValueOf(&t).MethodByName(strings.Title(file) + "Routes").Call([]reflect.Value{})
	}
}
