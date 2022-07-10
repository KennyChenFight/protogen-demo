package main

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/pluginpb"
)

func main() {
	data, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal("unable to read input: ", err)
	}

	req := &pluginpb.CodeGeneratorRequest{}
	if err = proto.Unmarshal(data, req); err != nil {
		log.Fatal("unable to unmarshal request: ", err)
	}

	path := req.GetParameter()
	if path == "" {
		log.Fatal(`please execute the plugin with the output path to properly write the output file: --debug_out="{PATH}:{PATH}"`)
	}

	err = os.MkdirAll(path, 0755)
	if err != nil {
		log.Fatal("unable to create output dir: ", err)
	}

	err = ioutil.WriteFile(filepath.Join(path, "code_generator_request.pb.bin"), data, 0644)
	if err != nil {
		log.Fatal("unable to write request to disk: ", err)
	}
}
