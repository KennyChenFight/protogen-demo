package main

import (
	"github.com/KennyChenFight/protogen-demo/protoc-gen-demo2/module"
	pgs "github.com/lyft/protoc-gen-star"
	pgsgo "github.com/lyft/protoc-gen-star/lang/go"
	"google.golang.org/protobuf/types/pluginpb"
)

func main() {
	optionalFeature := uint64(pluginpb.CodeGeneratorResponse_FEATURE_PROTO3_OPTIONAL)
	pgs.Init(
		pgs.DebugMode(),
		pgs.SupportedFeatures(&optionalFeature),
	).RegisterModule(
		module.NewJSONifyModule(),
	).RegisterPostProcessor(
		pgsgo.GoFmt(),
	).Render()
}
