# protogen-demo

show how to develop your own proto codegen plugin, you can also read [my blog post](https://blog.kennycoder.io/2022/07/09/Protobuf-%E5%A6%82%E4%BD%95%E9%96%8B%E7%99%BC-codegen-plugin/) !

## protoc-gen-demo

use **google.golang.org/protobuf/compiler/protogen** lib to demo how to develop proto codegen plugin

```bash
go build -o ./protoc-gen-demo/protoc-gen-demo.bin ./protoc-gen-demo
protoc \
--plugin=protoc-gen-demo=./protoc-gen-demo/protoc-gen-demo.bin \
--go_out=paths=source_relative:. \
--demo_out=foo=bar,paths=source_relative:. \
proto/pb.proto
```

## protoc-gen-demo2

use **github.com/lyft/protoc-gen-star** lib to demo how to develop proto codegen plugin

```bash
go build -o ./protoc-gen-demo/protoc-gen-demo.bin ./protoc-gen-demo
protoc \
--plugin=protoc-gen-demo2=./protoc-gen-demo2/protoc-gen-demo2.bin \
--go_out=paths=source_relative:. \
--demo2_out=foo=bar,paths=source_relative:. \
proto/pb.proto
```

### protoc-gen-debug

show how to debug your own plugin

1. ```bash
    protoc \                                                                                 
    --plugin=protoc-gen-debug=./protoc-gen-debug/protoc-gen-debug.bin \
    --go_out=paths=source_relative:. \
    --debug_out=.:. \                    
    proto/pb.proto
    ```
2. and run `go run main.go` in protogen-demo folder`
