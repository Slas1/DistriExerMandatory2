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

func (s *CriticalServer) pop() int32 {
	if s.isEmpty() {
		log.Panic(errors.New("the queue was empty when popping, PANIC"))
		fmt.Println(errors.New("the queue was empty when popping, PANIC"))
	}
	x := s.queue[0]
	s.queue = s.queue[1:]
	return x
}

func (s *CriticalServer) push(new int32) {
	s.queue = append(s.queue, new)
}

func (s *CriticalServer) peek() int32 {
	return s.queue[0]
}

func (s *CriticalServer) isEmpty() bool {
	return len(s.queue) < 1
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
	if (s.isEmpty() && s.clientIDWithAccess == 0) {
		s.clientIDWithAccess = request.SenderID
		fmt.Println("I'm here")
		return &criticalpb.AccessGranted{Message: "You have gained access to Critical information"}, nil
	} else {
		s.push(request.SenderID)
		fmt.Println(len(s.queue))
	}

	for {
		if s.peek() == request.SenderID && s.clientIDWithAccess == 0 {
			break
		}
		time.Sleep(time.Second * 4)
	}

	s.clientIDWithAccess = s.pop()
	return &criticalpb.AccessGranted{Message: "You have gained access to Critical information"}, nil
}

func (s *CriticalServer) RetriveCriticalInformation(ctx context.Context, request *criticalpb.Message) (*criticalpb.Message, error) {
	if request.SenderID == s.clientIDWithAccess {
		return &criticalpb.Message{Message: s.topSecretStuff}, nil
	}
	return nil, errors.New("you can't access the Critical information, you dont have access")
}

func (s *CriticalServer) ReleaseAccessToCritical(ctx context.Context, request *criticalpb.Message) (*criticalpb.AccessReleased, error) {
	if request.SenderID == s.clientIDWithAccess {
		s.clientIDWithAccess = 0
		return &criticalpb.AccessReleased{Message: "You have lost the access to Critical information"}, nil
	} 
	return nil, errors.New("sender ID and ID of client with access don't match")
}

func (s *CriticalServer) ClearFromQueue(ctx context.Context, request *criticalpb.Message) (*criticalpb.Message, error){
}

func main() {
	LOG_FILE := "./ServerLogfile"
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
		fmt.Printf("Failed to listen: %v", err)
	}

	s := newServer()

	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	criticalpb.RegisterCriticalSectionGRPCServer(grpcServer, s)
	grpcServer.Serve(lis)
}
