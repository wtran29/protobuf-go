package basic

import (
	"log"

	pb "github.com/wtran29/protobuf-go/protogen/basic"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

func CreateUser() *pb.User {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)
	u := pb.User{
		Id:       42,
		Username: "Bentley Benjamins",
		IsActive: true,
		Password: hashedPassword,
		Emails:   []string{"bb@corgi.com", "benben@gmail.com"},
		Gender:   pb.Gender_GENDER_MALE,
		Address: &pb.Address{
			Street:  "123 Jump St",
			City:    "Buena Park",
			Country: "USA",
			Coordinate: &pb.Address_Coordinate{
				Longitude: -37.3847,
				Latitude:  34.3429,
			},
		},
	}
	// log.Println(&u)
	return &u

}

func ToJSON(m proto.Message) string {
	option := protojson.MarshalOptions{
		// Multiline:     true,
		UseProtoNames: true,
	}
	jsonBytes, _ := option.Marshal(m)
	jsonString := string(jsonBytes)
	log.Println(jsonString)
	return jsonString
}

func FromJSON() {
	json := `{
        "id": 1,
        "username": "Bret",
		"is_active": true,
		"password": "JDJhJDEwJE0yTE5SQ01SYklyOEpWTllEM3BkWmVFUEtuV25zTWUvZzgxMUdtTG90djZVOW1lQUJ1N0Fh",
        "emails": ["Sincere@april.biz"],
        "gender": "GENDER_MALE"
        }`
	// option := protojson.UnmarshalOptions{
	// 	AllowPartial:   true,
	// 	DiscardUnknown: true,
	// }
	var u pb.User
	err := protojson.Unmarshal([]byte(json), &u)
	if err != nil {
		return
	}
	log.Println(&u)

}
