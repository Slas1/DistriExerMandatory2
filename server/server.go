package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"time"

	"criticalpb/criticalpb"

	"google.golang.org/grpc"
)

type CriticalServer struct {
	criticalpb.UnimplementedCriticalSectionGRPCServer
	queue              []int32
	activeClients      []bool
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

func (s *CriticalServer) removeIndex(i int) {
	var part1 = make([]int32, 0, 10)
	var part2 = make([]int32, 0, 10)
	if i > 0 {
		fmt.Printf("%v\n", s.queue)
		part1 = s.queue[0:i]
		fmt.Println(strconv.Itoa(len(s.queue)))
	}
	if i == len(s.queue)-1 {
		s.queue = part1
	} else {
		part2 = s.queue[i+1:]
		s.queue = append(part1, part2...)
	}
}

func newServer() *CriticalServer {
	s := &CriticalServer{
		queue:              make([]int32, 0, 20),
		activeClients: make([]bool, 20),
		topSecretStuff:     "Secret information. shhhh!",
		clientIDWithAccess: 0,
		highestUniqueID:    0,
	}
	return s
}

func (s *CriticalServer) GetIdFromServer(ctx context.Context, request *criticalpb.Message) (*criticalpb.IdResponse, error) {
	s.highestUniqueID += 1
	s.activeClients[s.highestUniqueID] = true
	return &criticalpb.IdResponse{ID: s.highestUniqueID}, nil
}

func (s *CriticalServer) RequestAccessToCritical(ctx context.Context, request *criticalpb.Message) (*criticalpb.AccessGranted, error) {	
	log.Printf("Client with id: " + strconv.Itoa(int(request.SenderID)) + " - Has requested access to Critical")
	fmt.Printf("Client with id: " + strconv.Itoa(int(request.SenderID)) + " - Has requested access to Critical\n")
	if s.isEmpty() && s.clientIDWithAccess == 0 || s.clientIDWithAccess == request.SenderID {
		s.clientIDWithAccess = request.SenderID
		return &criticalpb.AccessGranted{Message: "You have gained access to Critical information"}, nil
	}

	s.push(request.SenderID)
	fmt.Printf("Queue length is : " + strconv.Itoa(len(s.queue)) + "\n")

	for {
		if s.activeClients[request.SenderID] && s.peek() == request.SenderID && s.clientIDWithAccess == 0 {
			break
		}
		time.Sleep(time.Second * 4)
	}
	log.Printf("Client with id: %s - has gained access to Critical", strconv.Itoa(int(request.SenderID)))
	fmt.Printf("Client with id: %s - has gained access to Critical\n", strconv.Itoa(int(request.SenderID)))

	s.clientIDWithAccess = s.pop()
	return &criticalpb.AccessGranted{Message: "You have gained access to Critical information"}, nil
}

func (s *CriticalServer) RetriveCriticalInformation(ctx context.Context, request *criticalpb.Message) (*criticalpb.Message, error) {
	if request.SenderID == s.clientIDWithAccess {
		log.Printf("Client with id: %s - has retrieved Critical Information", strconv.Itoa(int(request.SenderID)))
		fmt.Printf("Client with id: %s - has retrieved Critical Information\n", strconv.Itoa(int(request.SenderID)))
		return &criticalpb.Message{Message: s.topSecretStuff}, nil
	}
	return nil, errors.New("you can't access the Critical information, you dont have access")
}

func (s *CriticalServer) ReleaseAccessToCritical(ctx context.Context, request *criticalpb.Message) (*criticalpb.AccessReleased, error) {
	if request.SenderID == s.clientIDWithAccess {
		s.clientIDWithAccess = 0
		log.Printf("Client with id: %s - has lost access to Critical", strconv.Itoa(int(request.SenderID)))
		fmt.Printf("Client with id: %s - has lost access to Critical\n", strconv.Itoa(int(request.SenderID)))
		return &criticalpb.AccessReleased{Message: "You have lost the access to Critical information"}, nil
	}
	return nil, errors.New("sender ID and ID of client with access don't match")
}

func (s *CriticalServer) ClearFromQueue(ctx context.Context, request *criticalpb.Message) (*criticalpb.Message, error) {
	for i := 0; i < len(s.queue); i++ {
		if s.queue[i] == request.SenderID {
			s.removeIndex(i)
			log.Printf("Client with id: %s - has been removed from queue", strconv.Itoa(int(request.SenderID)))
			fmt.Printf("Client with id: %s - has been removed from queue\n", strconv.Itoa(int(request.SenderID)))
		}
	}
	return &criticalpb.Message{Message: "Sender with id: " + strconv.Itoa(int(request.SenderID)) + " - Was removed from the queue"}, nil
}

func (s *CriticalServer) Leave(ctx context.Context, request *criticalpb.Message) (*criticalpb.Message, error) {
	s.activeClients[request.SenderID] = false
	
	if request.SenderID == s.clientIDWithAccess {
		aR, err := s.ReleaseAccessToCritical(ctx, request)
		return &criticalpb.Message{Message: aR.Message }, err
	} else {
		return s.ClearFromQueue(ctx, request)
	}
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
		fmt.Printf("Failed to listen: %v\n", err)
	}

	s := newServer()

	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	criticalpb.RegisterCriticalSectionGRPCServer(grpcServer, s)
	log.Printf("Server was started.")
	fmt.Printf("Server was started. \n")
	grpcServer.Serve(lis)
}
