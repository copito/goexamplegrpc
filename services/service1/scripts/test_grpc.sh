#!/bin/bash

# Make sure service is running!
YELLOW="\e[33m"
GREEN="\e[92m"
RED="\e[31m"
STOP="\e[0m"

# List all services
printf "${YELLOW}Listing all services available:\n${STOP}"
printf "${GREEN}rpcurl -plaintext localhost:50051 list\n${STOP}"
grpcurl -plaintext localhost:50051 list
printf "\n"
sleep 4

printf "${YELLOW}Now, lets try to describe the service1.Greeter:\n${STOP}"
printf "${GREEN}grpcurl -plaintext localhost:50051 describe service1.Greeter\n${STOP}"
grpcurl -plaintext localhost:50051 describe service1.Greeter
printf "\n"


sleep 4

printf "${YELLOW}Lets also decribe the request message - .service1.HelloRequest:\n${STOP}"
printf "${GREEN}grpcurl -plaintext localhost:50051 describe .service1.HelloRequest\n${STOP}"
grpcurl -plaintext localhost:50051 describe .service1.HelloRequest
printf "\n"


sleep 4

printf "${YELLOW}Lets make our first request to the service1.Greeter/SayHello:\n${STOP}"
printf "${GREEN}grpcurl -d '{\"name\": \"Copito\", \"age\": 100}' -plaintext localhost:50051 service1.Greeter/SayHello\n${STOP}"
grpcurl -d '{"name": "Copito", "age": 100}' -plaintext localhost:50051 service1.Greeter/SayHello
printf "\n"

