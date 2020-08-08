protoc --proto_path ../../../ -I=./proto --go_out=plugins=grpc:./proto proto/gobuildmaster.proto 

mv proto/github.com/brotherlogic/gobuildmaster/proto/* ./proto
