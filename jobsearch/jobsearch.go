package jobsearch

import (
	"log"

	pb "github.com/wtran29/protobuf-go/basic"
	"github.com/wtran29/protobuf-go/protogen/job"
	"github.com/wtran29/protobuf-go/protogen/jobsearch"
	"github.com/wtran29/protobuf-go/protogen/software"
)

func JobSearchSoftware() {
	js := jobsearch.JobSoftware{
		JobSoftwareId: 777,
		Application: &software.Application{
			Version:   "1.0.0",
			Name:      "Protosoft",
			Platforms: []string{"Windows", "Linux", "MacOS"},
		},
	}

	jsonBytes := pb.ToJSON(&js)
	log.Printf("Software : %v", string(jsonBytes))
}

func JobSearchApplication() {
	jc := jobsearch.JobApplication{
		JobCandidateId: 420,
		Application: &job.Application{
			AppId:   888,
			AppName: "The Chosen One",
			Phone:   "999-999-9999",
			Email:   "chosen@gmail.com",
		},
	}
	jsonBytes := pb.ToJSON(&jc)
	log.Printf("Candidate : %v", string(jsonBytes))
}
