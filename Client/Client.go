package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	pb "github.com/MieMilvang/DISYS_gRPC_Eample/Proto"
	"google.golang.org/grpc"
)

const(
	address = "localhost: 8080"
)

var (
	clientid                    int64
)

func main(){
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewServiceClient(conn)
	if clientid == 0{
		fmt.Printf("please put in 'clientID' and then the wanted id")
		reader, err := bufio.NewReader(os.Stdin).ReadString('\n')
		input := strings.TrimSuffix(reader, "\n")
		if err != nil {
			log.Fatalf("failed: %v", err )
		}
		if strings.EqualFold(input, "clientID"){
			if strings.EqualFold(input, "clientID") {
				fmt.Println("what ID do you want?")
			i, err := strconv.ParseInt(input, 10, 32)
			checkErr(err)
			if i >= 0 {
				clientid = i
			}else {
				fmt.Printf("pleasse use a posetive integer as ID value")
				
			}

		}
	}

	for (clientid > 0) {
		// to ensure "enter" has been hit before publishing - skud ud til Mie
		// remove newline windows format "\r\n"
		//reader, err := bufio.NewReader(os.Stdin).ReadString('\n')
		//input := strings.TrimSuffix(reader, "\r\n")
		//inputMac := strings.TrimSuffix(reader, "\n");

			
			fmt.Println("how much do you want to increment?")
			_reader, err := bufio.NewReader(os.Stdin).ReadString('\n')
			_input := strings.TrimSuffix(_reader, "\r\n")
			if err != nil {
				log.Fatalf("failed: %v", err )
			}
			In, err := strconv.ParseInt(_input, 10, 32)
			if In > 0 {
				in32 := int32(In)
				c.Increment(context.Background(), &pb.IncrementValue{Value: in32, Clientid: int32(clientid)} )
				}else {
					fmt.Printf("pleasse use a posetive integer as increment value")
					break;
				}

		}
	}
}

func checkErr(err error) {
	if err != nil {
		log.Fatalln(err.Error())
	}
}

