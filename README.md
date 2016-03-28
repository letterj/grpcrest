
## Install dependencies

* go get -u github.com/gengo/grpc-gateway/protoc-gen-grpc-gateway
* go get -u github.com/gengo/grpc-gateway/protoc-gen-swagger

Maybe?
* go get -u github.com/golang/protobuf/protoc-gen-go



proto file changes

```
syntax = "proto3";
package example;
+
+import "google/api/annotations.proto";
+
message StringMessage {
	  string value = 1;
}

service YourService {
	-  rpc Echo(StringMessage) returns (StringMessage) {}
	+  rpc Echo(StringMessage) returns (StringMessage) {
		+    option (google.api.http) = {
			+      post: "/v1/example/echo"
			+      body: "*"
			+    };
			+  }
}
```

### Generate pb.proto file

```
protoc -I/usr/local/include -I. \
 -I$GOPATH/src \
  -I$GOPATH/src/github.com/gengo/grpc-gateway/third_party/googleapis \
   --go_out=Mgoogle/api/annotations.proto=github.com/gengo/grpc-gateway/third_party/googleapis/google/api,plugins=grpc:. grpcreset.proto
```

### Generate the reverse proxy
```
protoc -I/usr/local/include -I. \
 -I$GOPATH/src \
  -I$GOPATH/src/github.com/gengo/grpc-gateway/third_party/googleapis \
   --grpc-gateway_out=logtostderr=true:. grpcreset.proto
```

#!/bin/bash 
set -x
protoc --gofast_out=plugins=grpc:. *.proto
go install .
