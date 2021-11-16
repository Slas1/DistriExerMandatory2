package main

import (
	"context"
	"criticalpb/criticalpb"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/thecodeteam/goodbye"
	"google.golang.org/grpc"
)

var tcpServer = flag.String("server", ":8080", "TCP server")
var id int32

func getIdFromServer(client criticalpb.CriticalSectionGRPCClient, request criticalpb.Message) {
	response, err := client.GetIdFromServer(context.Background(), &request)
	if err != nil {
		log.Fatalf("Client with id: %s - Error when calling getIdFromServer: %s", strconv.Itoa(int(id)), err)
		fmt.Printf("Client with id: %s - Error when calling getIdFromServer: %s\n", strconv.Itoa(int(id)), err)
	}
	id = response.ID
	log.Printf("Client with id: %s - Got ID from server", strconv.Itoa(int(id)))
	fmt.Printf("Client with id: %s - Got ID from server\n", strconv.Itoa(int(id)))
}

func requestAccessToCritical(client criticalpb.CriticalSectionGRPCClient, request criticalpb.Message) {
	response, err := client.RequestAccessToCritical(context.Background(), &request)
	if err != nil {
		log.Fatalf("Client with id: %s - Error when calling requestAccesToCritical: %s", strconv.Itoa(int(id)), err)
		fmt.Printf("Client with id: %s - Error when calling requestAccesToCritical: %s\n", strconv.Itoa(int(id)), err)
	}

	log.Printf("Client with id: %s - %v\n", strconv.Itoa(int(id)), response.Message)
	fmt.Printf("Client with id: %s - %v\n", strconv.Itoa(int(id)), response.Message)
}

func retriveCriticalInformation(client criticalpb.CriticalSectionGRPCClient, request criticalpb.Message) {
	response, err := client.RetriveCriticalInformation(context.Background(), &request)
	if err != nil {
		log.Fatalf("Client with id: %s - Error when calling retriveCriticalInformation: %s", strconv.Itoa(int(id)), err)
		fmt.Printf("Client with id: %s - Error when calling retriveCriticalInformation: %s\n", strconv.Itoa(int(id)), err)
	}

	log.Printf("Client with id: %s - %v\n", strconv.Itoa(int(id)), response.Message)
	fmt.Printf("Client with id: %s - %v\n", strconv.Itoa(int(id)), response.Message)
}

func releaseAccessToCritical(client criticalpb.CriticalSectionGRPCClient, request criticalpb.Message) {
	response, err := client.ReleaseAccessToCritical(context.Background(), &request)
	if err != nil {
		log.Fatalf("Client with id: %s - Error when calling retriveCriticalInformation: %s", strconv.Itoa(int(id)), err)
		fmt.Printf("Client with id: %s - Error when calling retriveCriticalInformation: %s\n", strconv.Itoa(int(id)), err)
	}
	log.Printf("Client with id: %s - %v\n", strconv.Itoa(int(id)), response.Message)
	fmt.Printf("Client with id: %s - %v\n", strconv.Itoa(int(id)), response.Message)
}

func clearFromQueue(client criticalpb.CriticalSectionGRPCClient, request criticalpb.Message) {
	response, err := client.ClearFromQueue(context.Background(), &request)
	if err != nil {
		log.Fatalf("Client with id: %s - Error when calling ClearFromQueue: %s", strconv.Itoa(int(id)), err)
		fmt.Printf("Client with id: %s - Error when calling ClearFromQueue: %s\n", strconv.Itoa(int(id)), err)
	}
	log.Printf("Client with id: %s - %v\n", strconv.Itoa(int(id)), response.Message)
	fmt.Printf("Client with id: %s - %v\n", strconv.Itoa(int(id)), response.Message)
}

func leave(client criticalpb.CriticalSectionGRPCClient, request criticalpb.Message) {
	response, err := client.Leave(context.Background(), &request)
	if err != nil {
		log.Fatalf("Client with id: %s - Error when calling Leave: %s", strconv.Itoa(int(id)), err)
		fmt.Printf("Client with id: %s - Error when calling Leave: %s\n", strconv.Itoa(int(id)), err)
	}
	log.Printf("Client with id: %s - %v\n", strconv.Itoa(int(id)), response.Message)
	fmt.Printf("Client with id: %s - %v\n", strconv.Itoa(int(id)), response.Message)
}

func behavior(ctx context.Context, client criticalpb.CriticalSectionGRPCClient) {
	if int(id) % 3 == 0 {
		var prefix string = "Client with id: " + strconv.Itoa(int(id)) + " - "
		for {
			var random = rand.Intn(2)
			
			switch random {
			case 0:
				requestAccessToCritical(client, criticalpb.Message{Message: prefix + "Give me access you filthy casual", SenderID: id})
			case 1:
				retriveCriticalInformation(client, criticalpb.Message{Message: prefix + "Give me critical information", SenderID: id})
			case 2:
				releaseAccessToCritical(client, criticalpb.Message{Message: prefix + "I'm done with you peasant, release my access to Critical", SenderID: id})
			}
		}
	}else {
		for {
			var randomTime = rand.Intn(5) + 1
			time.Sleep(time.Second * time.Duration(randomTime))
			var prefix string = "Client with id: " + strconv.Itoa(int(id)) + " - "
			requestAccessToCritical(client, criticalpb.Message{Message: prefix + "Give me access you filthy casual", SenderID: id})
			for i := 0; i < 5; i++ {
				retriveCriticalInformation(client, criticalpb.Message{Message: prefix + "Give me critical information", SenderID: id})
				time.Sleep(time.Second * 2)
			}
			releaseAccessToCritical(client, criticalpb.Message{Message: prefix + "I'm done with you peasant, release my access to Critical", SenderID: id})
		}
	}
	
}

func main() {
	LOG_FILE := "./ClientLogfile"
	logFile, err := os.OpenFile(LOG_FILE, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Panic(err)
	}

	// Set log out put and enjoy :)
	log.SetOutput(logFile)
	log.SetFlags(log.Lmicroseconds)

	flag.Parse()

	var opts []grpc.DialOption
	opts = append(opts, grpc.WithBlock(), grpc.WithInsecure())

	conn, err := grpc.Dial(*tcpServer, opts...)
	if err != nil {
		log.Fatalf("Fail to dial(connect): %v", err)
	}

	ctx := context.Background()
	client := criticalpb.NewCriticalSectionGRPCClient(conn)

	var message string = "Client with id: " + strconv.Itoa(int(id)) + " - "

	defer goodbye.Exit(ctx, -1)
	goodbye.Notify(ctx)
	goodbye.RegisterWithPriority(func(ctx context.Context, sig os.Signal) {
		leave(client, criticalpb.Message{Message: message + "I'm leaving, remove me.", SenderID: id })
	}, 1)
	goodbye.RegisterWithPriority(func(ctx context.Context, sig os.Signal) { logFile.Close() }, 4)
	goodbye.RegisterWithPriority(func(ctx context.Context, sig os.Signal) { conn.Close() }, 5)

	getIdFromServer(client, criticalpb.Message{Message: message + "Please give me a unique ID", SenderID: id})

	randomJoiner(ctx, client)

}
