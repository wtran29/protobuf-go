package main

import (
	"fmt"
	"log"
	"time"

	"github.com/wtran29/protobuf-go/jobsearch"
)

type logWriter struct {
}

func (writer logWriter) Write(bytes []byte) (int, error) {
	return fmt.Print(time.Now().Format("15:04:05") + " " + string(bytes))
}

func main() {
	log.SetFlags(0)
	log.SetOutput(new(logWriter))

	// basic.BasicHello()
	// u := basic.CreateUser()
	// basic.ToJSON(u)
	// basic.FromJSON()

	// ug := basic.BasicUserGroup()
	// basic.ToJSON(ug)

	jobsearch.JobSearchSoftware()
	jobsearch.JobSearchApplication()
}
