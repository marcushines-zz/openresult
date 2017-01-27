# openresult
Open Test Result Format

# Build Instructions
Build the proto and move to workspace
```
cd $WORKSPACE
protoc -I . github.com/marcushines/openresult/proto/result.proto  --go_out=plugins=grpc:.
protoc -I . github.com/marcushines/openresult/proto/testbed.proto  --go_out=plugins=grpc:.
mv github.com/marcushines/openresult/proto/result.pb.go github.com/marcushines/openresult
```

# Run Tests
```
go test github.com/marcushines/openresult/testing
```
