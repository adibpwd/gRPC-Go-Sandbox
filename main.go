package main

import (
	"bytes"
	"fmt"
	"noval-agung-grpc-protobuf/common/model"
	"os"
	"strings"

	"github.com/golang/protobuf/jsonpb"
)

func main() {
	var user1 = &model.User{
		Id:       "1",
		Name:     "Adib",
		Password: "sembarang",
		Gender:   model.UserGender_Male,
	}

	var userList = &model.UserList{
		List: []*model.User{
			user1,
		},
	}

	var garage1 = &model.Garage{
		Id:   "g001",
		Name: "Kalimdor",
		Coordinate: &model.GarageCoordinate{
			Latitude:  23.2212847,
			Longitude: 32.2394234,
		},
	}

	var garageList = &model.GarageList{
		List: []*model.Garage{
			garage1,
		},
	}

	var garageListByUser = &model.GarageListByUser{
		List: map[string]*model.GarageList{
			user1.Id: garageList,
		},
	}

	var buf bytes.Buffer
	err1 := (&jsonpb.Marshaler{}).Marshal(&buf, garageList)
	if err1 != nil {
		fmt.Println(err1.Error())
		os.Exit(0)
	}

	jsonString := buf.String()

	buf2 := strings.NewReader(jsonString)
	protoObject := new(model.GarageList)

	err2 := (&jsonpb.Unmarshaler{}).Unmarshal(buf2, protoObject)
	if err2 != nil {
		fmt.Println(err2.Error())
		os.Exit(0)
	}

	fmt.Printf("# ==== As String\n       %v \n", protoObject.String())
	fmt.Println("-----------------------------------------------")

	fmt.Printf("Hasi ke json string %v \n", jsonString)
	fmt.Println("-----------------------------------------------")

	fmt.Printf("ini adalah userList = %v", userList)
	fmt.Println("-----------------------------------------------")
	fmt.Printf("ini adalah garage1 = %v", garage1)
	fmt.Println("-----------------------------------------------")
	fmt.Printf("ini adalah listgarage = %v", garageListByUser)

}
