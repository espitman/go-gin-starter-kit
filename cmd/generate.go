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
		} else if what == "consumer" {
			createConsumer(args[1])
		} else if what == "cron" {
			createCron(args[1])
		}
		os.Exit(1)
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)
}

func createController(name string) {
	createDto(name)
	fmt.Println("createController: " + name)
	f := NewFile(name + "Controller")
	f.Comment("Single " + name)
	f.Comment("@Summary Single " + name)
	f.Comment("@ID single-" + name)
	f.Comment("@Accept  json")
	f.Comment("@Param id path string true \"url params\"")
	f.Comment("@Success 200 {object} dto_" + name + "_response.Full")
	f.Comment("@Failure 400 {object} dto_error.Error")
	f.Comment("@Router /" + name + "/{id} [get]")

	f.Func().Id("Ping").Params(Id("c").Params(Add(Op("*")).Qual("github.com/gin-gonic/gin", "Context"))).Block(
		Qual("jettster/utils", "FormatResponse").Params(Id("c"), Id("gin.H{\"ID\": \""+name+"\",}")),
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
		Id("t.router.GET").Params(Id("\"/"+name+"/:id\""), Qual("jettster/controller/"+name, "Ping")),
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
	f := NewFile("dto_" + name + "_request")
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
	f := NewFile("dto_" + name + "_response")
	f.Type().Id("Full_Payload").Struct(
		Id("ID").Id("string").Id("`json:\"ID\"`"),
	)
	f.Type().Id("Full").Struct(
		Id("Message").Id("string").Id("`json:\"message\"`"),
		Id("Status").Id("int").Id("`json:\"status\"`"),
		Id("Payload").Id("Full_Payload").Id("`json:\"payload\"`"),
	)
	buf := &bytes.Buffer{}
	if err := f.Render(buf); err != nil {
		panic(err)
	}
	if err := ioutil.WriteFile("./dto/"+name+"/response/full.dto.go", buf.Bytes(), 0644); err != nil {
		panic(err)
	}
}

func createConsumer(name string) {
	f := NewFile("consumer")
	f.Func().Params(Id("t *T")).Id(strings.Title(name)+"Consumer").Params().Block(
		Id("msgs, _ :=").Qual("jettster/provider/rabbitmq", "Consume").Call(Id("\"ginTestExchange\""), Id("\"ginTestQueue\"")),
		Id("forever :=").Make(Chan().Bool()),
		Id("go func").Params().Block(
			Id("for msg := range msgs").Block(
				Id("body := ").Id("string(msg.Body)"),
				Qual("fmt", "Println").Call(Id("\""+name+"Consumer::: "+"\"").Op("+").Id("\"message received:\" ").Op("+").Id(" body")),
				//	_ = msg.Ack(true)
				Id("_ =").Id("msg.Ack(true)"),
			),
		).Call(),
		Id("	<-forever"),
	)

	buf := &bytes.Buffer{}
	if err := f.Render(buf); err != nil {
		panic(err)
	}
	if err := ioutil.WriteFile("./worker/consumer/"+name+".go", buf.Bytes(), 0644); err != nil {
		panic(err)
	}
}

func createCron(name string) {
	f := NewFile("cron")

	f.Func().Id(strings.Title(name) + "Task").Params().Block(
		Qual("fmt", "Println").Call(Id("\"I am running " + strings.Title(name) + "Task.\"")),
	)
	// func (t *T) TestRun() {
	// 	s1 := gocron.NewScheduler(time.UTC)
	// 	_, _ = s1.Every(1).Seconds().Do(TestTask)
	// 	s1.StartAsync()
	// }
	f.Func().Id("(t *T)").Id(strings.Title(name)+"Run").Params().Block(
		Id("s1:=").Qual("github.com/go-co-op/gocron", "NewScheduler").Params(Qual("time", "UTC")),
		Id("_, _ =").Id("s1.Every(1).Seconds().Do("+strings.Title(name)+"Task)"),
		Id("s1.StartAsync()"),
	)

	buf := &bytes.Buffer{}
	if err := f.Render(buf); err != nil {
		panic(err)
	}
	if err := ioutil.WriteFile("./worker/cron/"+name+".go", buf.Bytes(), 0644); err != nil {
		panic(err)
	}
}
