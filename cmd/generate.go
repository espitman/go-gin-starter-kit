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
