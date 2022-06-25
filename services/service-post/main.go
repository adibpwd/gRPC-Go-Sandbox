package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"noval-agung-grpc-protobuf/common/config"
	"noval-agung-grpc-protobuf/common/model"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
)

var localStorage *model.PostListByUser

func init() {
	localStorage = new(model.PostListByUser)
	localStorage.List = make(map[string]*model.PostList)
}

type PostServer struct{}

func (PostServer) Add(ctx context.Context, param *model.PostAndUserId) (*empty.Empty, error) {
	userId := param.UserId
	post := param.Post

	if _, ok := localStorage.List[userId]; !ok {
		localStorage.List[userId] = new(model.PostList)
		localStorage.List[userId].List = make([]*model.Post, 0)
	}

	localStorage.List[userId].List = append(localStorage.List[userId].List, post)
	fmt.Println("Adding post", post.String(), "for user", userId)
	return new(empty.Empty), nil
}

func (PostServer) List(ctx context.Context, param *model.PostByUserId) (*model.PostList, error) {
	userId := param.UserId

	return localStorage.List[userId], nil
}

func main() {
	srv := grpc.NewServer()
	var postsSrv PostServer
	model.RegisterPostsServer(srv, postsSrv)
	portConfig := config.SERVICE_POST_PORT
	log.Println("Starting RPC server at", portConfig)

	l, err := net.Listen("tcp", portConfig)
	if err != nil {
		log.Fatalf("could not listen to %s: %v", portConfig, err)
	}

	log.Fatal(srv.Serve(l))
}
