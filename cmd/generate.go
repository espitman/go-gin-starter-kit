package cmd

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	. "github.com/dave/jennifer/jen"

	"github.com/spf13/cobra"
)

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "generate go gin starter kit",
	Long:  `generate go gin starter kit`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("generate called")
		what := args[0]
		if what == "controller" {
			createController(args[1])
			createRoute(args[1])
		} else if what == "model" {
			createModel(args[1])
		} else if what == "dto" {
			createDto(args[1])
		}
		os.Exit(1)
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)
}

func createController(name string) {
	fmt.Println("createController: " + name)
	f := NewFile(name + "Controller")
	f.Func().Id("Ping").Params(Id("c").Params(Add(Op("*")).Qual("github.com/gin-gonic/gin", "Context"))).Block(
		Id("c.JSON").Params(Qual("net/http", "StatusOK"), Id("gin.H{\"message\": \""+name+"\",}")),
	)
	buf := &bytes.Buffer{}
	if err := f.Render(buf); err != nil {
		panic(err)
	}
	dir := os.Mkdir("./controller/"+name, 0777)
	fmt.Println(dir)
	if err := ioutil.WriteFile("./controller/"+name+"/"+name+".go", buf.Bytes(), 0644); err != nil {
		panic(err)
	}
}

func createRoute(name string) {
	fmt.Println("createRoute: " + name)
	f := NewFile("route")
	f.ImportAlias("jettster/controller/"+name, name+"Controller")
	f.Func().Id("(t *T)").Id(strings.Title(name) + "Routes").Params().Block(
		Id("t.router.GET").Params(Id("\"/"+name+"/ping\""), Qual("jettster/controller/"+name, "Ping")),
	)
	buf := &bytes.Buffer{}
	if err := f.Render(buf); err != nil {
		panic(err)
	}
	if err := ioutil.WriteFile("./route/"+name+".go", buf.Bytes(), 0644); err != nil {
		panic(err)
	}
}

func createModel(name string) {
	fmt.Println("createModel: " + name)
	f := NewFile("model_" + name)
	f.Type().Id(strings.Title(name)).Struct(
		Qual("github.com/Kamva/mgm", "DefaultModel").Id("`bson:\",inline\"`"),
	)
	buf := &bytes.Buffer{}
	if err := f.Render(buf); err != nil {
		panic(err)
	}
	dir := os.Mkdir("./model/"+name, 0777)
	fmt.Println(dir)
	if err := ioutil.WriteFile("./model/"+name+"/"+name+".go", buf.Bytes(), 0644); err != nil {
		panic(err)
	}
}

func createDto(name string) {
	fmt.Println("createDto: " + name)
	dir := os.Mkdir("./dto/"+name, 0777)
	fmt.Println(dir)
	dir = os.Mkdir("./dto/"+name+"/request", 0777)
	fmt.Println(dir)
	dir = os.Mkdir("./dto/"+name+"/response", 0777)
	fmt.Println(dir)
	createRequestDto(name)
	createResponseDto(name)
}

func createRequestDto(name string) {
	f := NewFile("dto_" + name)
	f.Type().Id("Details").Struct(
		Id("ID").String().Id("`uri:\"id\"`"),
	)
	buf := &bytes.Buffer{}
	if err := f.Render(buf); err != nil {
		panic(err)
	}
	if err := ioutil.WriteFile("./dto/"+name+"/request/details.dto.go", buf.Bytes(), 0644); err != nil {
		panic(err)
	}
}

func createResponseDto(name string) {
	f := NewFile("dto_" + name)
	f.Type().Id("Full").Struct(
		Id("ID").Qual("go.mongodb.org/mongo-driver/bson/primitive", "ObjectID").Id("`json:\"_id\"`"),
	)
	buf := &bytes.Buffer{}
	if err := f.Render(buf); err != nil {
		panic(err)
	}
	if err := ioutil.WriteFile("./dto/"+name+"/response/full.dto.go", buf.Bytes(), 0644); err != nil {
		panic(err)
	}
}
