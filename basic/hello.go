package basic

import (
	"log"

	pb "github.com/wtran29/protobuf-go/protogen/basic"
)

func BasicHello() {
	h := pb.Hello{
		Name: "Jack Black",
	}

	log.Println(&h)
}
