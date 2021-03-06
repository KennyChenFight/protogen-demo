package main

import (
	"io/ioutil"
	"log"
	"os"

	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/pluginpb"
)

func main() {
	data, err := ioutil.ReadFile("./code_generator_request.pb.bin")
	if err != nil {
		log.Fatal("unable to read input: ", err)
	}

	req := &pluginpb.CodeGeneratorRequest{}
	if err = proto.Unmarshal(data, req); err != nil {
		log.Fatal("unable to unmarshal request: ", err)
	}

	options := protogen.Options{}
	plugin, err := options.New(req)
	if err := Test(plugin); err != nil {
		plugin.Error(err)
	}
	resp := plugin.Response()
	out, err := proto.Marshal(resp)
	if err != nil {
		log.Fatal("unable to unmarshal response: ", err)
	}
	if _, err := os.Stdout.Write(out); err != nil {
		log.Fatal("unable to write response: ", err)
	}
}

func Test(gen *protogen.Plugin) error {
	for _, f := range gen.Files {
		if !f.Generate {
			continue
		}
		generateFile(gen, f)
	}
	return nil
}

func generateFile(gen *protogen.Plugin, file *protogen.File) *protogen.GeneratedFile {
	filename := file.GeneratedFilenamePrefix + "_demo.pb.go"
	g := gen.NewGeneratedFile(filename, file.GoImportPath)
	g.P("// Code generated by protoc-gen-demo. DO NOT EDIT.")
	g.P()
	g.P("package ", file.GoPackageName)
	g.P()

	for _, msg := range file.Messages {
		var existGreetingField bool
		for _, field := range msg.Fields {
			if field.GoName == "Greeting" {
				existGreetingField = true
				break
			}
		}
		if existGreetingField {
			g.P("func (x *", msg.GoIdent, ") SayHello() string {")
			g.P("return `", "Greeting!", "`")
			g.P("}")
		}
	}

	return g
}
