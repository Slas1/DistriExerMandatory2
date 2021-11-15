package main

import (
	"context"
	"criticalpb/criticalpb"
	"flag"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/thecodeteam/goodbye"
	"google.golang.org/grpc"
)

var tcpServer = flag.String("server", ":8080", "TCP server")
var accessStatus = false
var id int32

func getIdFromServer(client criticalpb.CriticalSectionGRPCClient, request criticalpb.Message) {
	response, err := client.GetIdFromServer(context.Background(), &request)
	if err != nil {
		log.Fatalf("Error when calling getIdFromServer: %s", err)
	}
	id = response.ID
	log.Printf("Got ID from server")
}

func requestAccessToCritical(client criticalpb.CriticalSectionGRPCClient, request criticalpb.Message) {
	response, err := client.RequestAccessToCritical(context.Background(), &request)
	if err != nil {
		log.Fatalf("Error when calling requestAccesToCritical: %s", err)
	}

	accessStatus = true
	log.Printf(response.Message)
}

func retriveCriticalInformation(client criticalpb.CriticalSectionGRPCClient, request criticalpb.Message) {
	response, err := client.RetriveCriticalInformation(context.Background(), &request)
	if err != nil {
		log.Fatalf("Error when calling retriveCriticalInformation: %s", err)
	}

	log.Printf(response.Message)
}

func releaseAccessToCritical(client criticalpb.CriticalSectionGRPCClient, request criticalpb.Message) {
	response, err := client.ReleaseAccessToCritical(context.Background(), &request)
	if err != nil {
		log.Fatalf("Error when calling retriveCriticalInformation: %s", err)
	}
	accessStatus = false
	log.Printf(response.Message)
}

func randomJoiner(ctx context.Context, client criticalpb.CriticalSectionGRPCClient) {
	for {
		//var randomTime = rand.Intn(180-30) + 30
		var randomTime = rand.Intn(5) + 1
		time.Sleep(time.Second * time.Duration(randomTime))
		requestAccessToCritical(client, criticalpb.Message{Message: "Give me access you filthy casual", SenderID: id})
		for i := 0; i < 5; i++ {
			retriveCriticalInformation(client, criticalpb.Message{Message: "Give me critical information", SenderID: id})
			time.Sleep(time.Second * 5)
		}
		releaseAccessToCritical(client, criticalpb.Message{Message: "I'm done with you peasant, release my access to Critical", SenderID: id})
	}
}

func main() {
	LOG_FILE := "./logfile"
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

	defer goodbye.Exit(ctx, -1)
	goodbye.Notify(ctx)
	goodbye.RegisterWithPriority(func(ctx context.Context, sig os.Signal) {
		if accessStatus {
			releaseAccessToCritical(client, criticalpb.Message{Message: "Release my acccess from CriticalSection", SenderID: id})
		}
	}, 1)
	goodbye.RegisterWithPriority(func(ctx context.Context, sig os.Signal) { logFile.Close() }, 4)
	goodbye.RegisterWithPriority(func(ctx context.Context, sig os.Signal) { conn.Close() }, 5)

	getIdFromServer(client, criticalpb.Message{Message: "Please give me and unique ID", SenderID: id})

	randomJoiner(ctx, client)

	//requestAccessToCritical(client, criticalpb.Message{Message: "Give me access you filthy casual", SenderID: id})

	//time.Sleep(time.Second * 10)

}
