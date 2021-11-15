package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net"
	"os"
	"time"

	"criticalpb/criticalpb"

	"google.golang.org/grpc"
)

type CriticalServer struct {
	criticalpb.UnimplementedCriticalSectionGRPCServer
	queue              []int32
	clientIDWithAccess int32
	topSecretStuff     string
	highestUniqueID    int32
}

func pop(server CriticalServer) int32 {
	if !isEmpty(server) {
		log.Panic(errors.New("the queue was empty when popping, PANIC"))
	}
	x := server.queue[0]
	server.queue = server.queue[1:]
	return x
}

func push(server CriticalServer, new int32) {
	server.queue = append(server.queue, new)
}

func peek(server CriticalServer) int32 {
	return server.queue[0]
}

func isEmpty(server CriticalServer) bool {
	return len(server.queue) < 1
}

func newServer() *CriticalServer {
	s := &CriticalServer{
		queue:              make([]int32, 0, 20),
		topSecretStuff:     "Secret information. shhhh!",
		clientIDWithAccess: 0,
		highestUniqueID:    0,
	}
	return s
}

func (s *CriticalServer) GetIdFromServer(ctx context.Context, request *criticalpb.Message) (*criticalpb.IdResponse, error) {
	s.highestUniqueID += 1
	return &criticalpb.IdResponse{ID: s.highestUniqueID}, nil
}

func (s *CriticalServer) RequestAccessToCritical(ctx context.Context, request *criticalpb.Message) (*criticalpb.AccessGranted, error) {
	if isEmpty(*s) && s.clientIDWithAccess == 0 {
		s.clientIDWithAccess = request.SenderID
		fmt.Println("im am here")
		return &criticalpb.AccessGranted{Message: "You have gained access to Critical information"}, nil
	} else {
		push(*s, request.SenderID)
	}

	for {
		if peek(*s) == request.SenderID && s.clientIDWithAccess == 0 {
			break
		}
		time.Sleep(time.Second * 4)
	}

	s.clientIDWithAccess = pop(*s)
	return &criticalpb.AccessGranted{Message: "You have gained access to Critical information"}, nil
}

func (s *CriticalServer) RetriveCriticalInformation(ctx context.Context, request *criticalpb.Message) (*criticalpb.Message, error) {
	if request.SenderID == s.clientIDWithAccess {
		return &criticalpb.Message{Message: s.topSecretStuff}, nil
	}
	return nil, errors.New("you can access the Critical information, you dont have access")
}

func (s *CriticalServer) ReleaseAccessToCritical(ctx context.Context, request *criticalpb.Message) (*criticalpb.AccessReleased, error) {
	s.clientIDWithAccess = 0
	return &criticalpb.AccessReleased{Message: "You have lost the access to Critical information"}, nil
}

func main() {
	LOG_FILE := "./logfile"
	logFile, err := os.OpenFile(LOG_FILE, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Panic(err)
	}
	defer logFile.Close()

	// Set log out put and enjoy :)
	log.SetOutput(logFile)
	log.SetFlags(log.Lmicroseconds)

	lis, err := net.Listen("tcp", "localhost:8080")

	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := newServer()

	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	criticalpb.RegisterCriticalSectionGRPCServer(grpcServer, s)
	grpcServer.Serve(lis)
}
