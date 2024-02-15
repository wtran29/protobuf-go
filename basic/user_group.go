package basic

import (
	pb "github.com/wtran29/protobuf-go/protogen/basic"
	"golang.org/x/crypto/bcrypt"
)

func BasicUserGroup() *pb.UserGroup {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("alpha"), bcrypt.DefaultCost)
	klaw := pb.User{
		Id:       2,
		Username: "klaw",
		IsActive: true,
		Password: hashedPassword,
		Gender:   pb.Gender_GENDER_MALE,
	}
	pg13 := pb.User{
		Id:       13,
		Username: "pg13",
		IsActive: true,
		Password: hashedPassword,
		Gender:   pb.Gender_GENDER_MALE,
	}
	harden := pb.User{
		Id:       1,
		Username: "uno",
		IsActive: true,
		Password: hashedPassword,
		Gender:   pb.Gender_GENDER_MALE,
	}
	russ := pb.User{
		Id:       0,
		Username: "brodie",
		IsActive: true,
		Password: hashedPassword,
		Gender:   pb.Gender_GENDER_MALE,
	}
	clippers := pb.UserGroup{
		GroupId:   213,
		GroupName: "Los Angeles Clippers",
		Users: []*pb.User{
			&klaw,
			&pg13,
			&harden,
			&russ,
		},
		Description: "Number 1 NBA basketball team in the Western Conference",
	}

	return &clippers
}
