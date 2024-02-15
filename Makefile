GO_MODULE := github.com/wtran29/protobuf-go
GOOG_PROTO_PATH = C://vcpkg//installed//x64-windows//include
PROTO_DIRS := ./proto/basic ./proto/job ./proto/software ./proto/jobsearch
PROTO_DIR = proto

.PHONY: tidy
tidy:
	go mod tidy

.PHONY: clean
clean:
	rm -rf ./protogen


.PHONY: protoc
# (un)comment or add folder that contains proto as required
protoc:
	protoc -I${GOOG_PROTO_PATH} -I./proto/basic --go_opt=module=${GO_MODULE} \
	--go_out=. \
	./proto/basic/*.proto \


	protoc --go_opt=module=${GO_MODULE} \
	--go_out=. \
	./proto/job/*.proto \
	./proto/software/*.proto \
	./proto/jobsearch/*.proto \
	
	
	
.PHONY: build
build: clean protoc tidy

.PHONY: run
run:
	go run main.go

