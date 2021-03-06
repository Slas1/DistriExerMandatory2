# Mandatory Exercise 2 - Distributed Mutual Exclusion

## How to use (Target: TA)

Limitations: Maximum 20 clients.

1. Start server by typing: "go run server.go" in a terminal while in the server folder

1. Start client by typing: "go run client.go" in a terminal while in the client folder. A client vil have a linear behavior, as seen in its "behavior" function. RequestAccess -> wait for access -> Retrive Information 5 times -> Release Access. Every 3'rd client will do a random rpc function, to display that only the Client with access have access.

## Description of Submited LogFile (before TA runs)
In the log file we created we created 2 normal clients. Then a troublelsome 3'rd client, that tried to retrive information without having access and died. We created a 4'th client and removed it while in was in queue. We then removed 2 while it had access, and the next in the queue (client 1) got access.

## Description:

You have to implement distributed mutual exclusion between nodes in your distributed system. 

You can choose to implement any of the algorithms, that were discussed in lecture 7.

## System Requirements:

R1: Any node can at any time decide it wants access to the Critical Section

R2: Only one node at the same time is allowed to enter the Critical Section 

R2: Every node that requests access to the Critical Section, will get access to the Critical Section (at some point in time)

## Technical Requirements:

1. Use Golang to implement the service's nodes

1. Use gRPC for message passing between nodes
 
1. Your nodes need to find each other.  For service discovery, you can choose one of the following options
 
  - supply a file with  ip addresses/ports of other nodes

  - enter ip adress/ports trough command line

  - use the Serf package for service discovery

4. Demonstrate that the system can be started with at least 3 nodes

5. Demonstrate using logs,  that a node gets access to the Critical Section

## Hand-in requirements:

1. Hand in a single report in a pdf file

1. Provide a link to a Git repo with your source code in the report

1. Include system logs, that document the requirements are met, in the appendix of your report


## Grading notes

Partial implementations may be accepted, if the students can reason what they should have done in the report.
