# 安装
- proto 安装

-proto-gen-go
-proto-gen-grpc文件

```
 go get google.golang.org/protobuf/cmd/protoc-gen-go && google.golang.org/grpc/cmd/protoc-gen-go-grpc
 
 修改环境变量
 export PATH="$PATH:$(go env GOPATH)/bin"
```

# 生成文件

- 单纯文件
```
protoc --go_out=. --go_opt=paths=source_relative  --go-grpc_out=. --go-grpc_opt=paths=source_relative helloworld.proto 

```

- grpc-gateway

```
protoc -I/usr/local/include -I. -I/home/cc/go/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.9.0/third_party/googleapis --go-grpc_out=. --go-grpc_opt=paths=source_relative helloworld.proto

protoc -I/usr/local/include -I. -I/home/cc/go/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.9.0/third_party/googleapis --grpc-gateway_out=logtostderr=true:. helloworld.proto

```

# 参考文档
https://developer.aliyun.com/article/770106
