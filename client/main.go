package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"noval-agung-grpc-protobuf/common/config"
	"noval-agung-grpc-protobuf/common/model"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// func serviceGarage() model.GaragesClient {
// 	port := config.SERVICE_GARAGE_PORT
// 	conn, err := grpc.Dial(port, grpc.WithTransportCredentials(insecure.NewCredentials()))
// 	if err != nil {
// 		log.Fatal("could not connect to", port, err)
// 	}

// 	return model.NewGaragesClient(conn)
// }

func servicePost() model.PostsClient {
	port := config.SERVICE_POST_PORT
	conn, err := grpc.Dial(port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("could not connect to", port, err)
	}

	return model.NewPostsClient(conn)
}

func serviceUser() model.UsersClient {
	port := config.SERVICE_USER_PORT
	conn, err := grpc.Dial(port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("could not connect to", port, err)
	}

	return model.NewUsersClient(conn)
}

func main() {
	user1 := model.User{
		Id:       "n001",
		Name:     "Noval Agung",
		Password: "kw8d hl12/3m,a",
		Gender:   model.UserGender(model.UserGender_value["MALE"]),
	}
	user2 := model.User{
		Id:       "n002",
		Name:     "Muslikhul adib",
		Password: "bebaskanlah",
		Gender:   model.UserGender(model.UserGender_value["MALE"]),
	}

	// garage1 := model.Garage{
	// 	Id:   "q001",
	// 	Name: "Quel'thalas",
	// 	Coordinate: &model.GarageCoordinate{
	// 		Latitude:  45.123123123,
	// 		Longitude: 54.1231313123,
	// 	},
	// }
	// garage2 := model.Garage{
	// 	Id:   "q002",
	// 	Name: "mantape pol",
	// 	Coordinate: &model.GarageCoordinate{
	// 		Latitude:  15.123123123,
	// 		Longitude: 14.1231313123,
	// 	},
	// }

	user := serviceUser()

	fmt.Println("\n", "===========> user test")

	// register user1
	user.Register(context.Background(), &user1)

	// register user2
	user.Register(context.Background(), &user2)

	res1, err := user.List(context.Background(), new(empty.Empty))
	if err != nil {
		log.Fatal("error bos", err.Error())
	}

	res1String, _ := json.Marshal(res1.List)
	fmt.Printf("res1String = %v \n", string(res1String))

	// garage := serviceGarage()

	// fmt.Println("\n", "===========> user garage")

	// garage.Add(context.Background(), &model.GarageAndUserId{
	// 	UserId: user1.Id,
	// 	Garage: &garage1,
	// })
	// garage.Add(context.Background(), &model.GarageAndUserId{
	// 	UserId: user2.Id,
	// 	Garage: &garage2,
	// })

	// res2, err := garage.List(context.Background(), &model.GarageUserId{UserId: user1.Id})
	// if err != nil {
	// 	log.Fatal("error bos", err.Error())
	// }

	// res2String, _ := json.Marshal(res2)
	// fmt.Printf("res2String = %v \n", string(res2String))

	post := servicePost()

	fmt.Println("\n", "user post")

	post.Add(context.Background(), &model.PostAndUserId{
		UserId: user1.Id,
		Post: &model.Post{
			Title:   "Test title",
			Content: "Ini content ya guys",
			Id:      "12",
		},
	})

	res3, err := post.List(context.Background(), &model.PostByUserId{UserId: user1.Id})
	if err != nil {
		log.Fatal("error euy", err.Error())
	}

	res3String, _ := json.Marshal(res3)
	fmt.Printf("res3String = %v \n", string(res3String))
}
