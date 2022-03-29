package main

import (
	"context"
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"google.golang.org/grpc"
	pb "karthik.com/go-tokenmgmt-grpc/proto"
)

const (
	address = "localhost:50051"
)

func main() {
	args := os.Args[1:]
	if args[0] == "-create" {
		fmt.Println("create")
		fmt.Println(args)
		conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
		if err != nil {
			log.Fatalf("did not connect: %v", err)
		}
		defer conn.Close()
		c := pb.NewTokenManagementClient(conn)

		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		r, err := c.CreateNewToken(ctx, &pb.NewToken{Name: "undefined", Domain: "undefined", State: "undefined", Id: args[2]})
		if err != nil {
			log.Fatalf("could not create token: %v", err)
		}
		log.Printf(`Token Details:
		NAME: %s
		AGE: %s
		State: %s
		ID: %s`, r.GetName(), r.GetDomain(), r.GetState(), r.GetId())
	}
	if args[0] == "-read" {
		fmt.Println("read")

		fmt.Println(args)
		conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
		if err != nil {
			log.Fatalf("did not connect: %v", err)
		}
		defer conn.Close()
		c := pb.NewTokenManagementClient(conn)

		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		r, err := c.GetToekns(ctx, &pb.Token{Id: args[2]})
		fmt.Println(r)
	}
	if args[0] == "-write" {
		fmt.Println(args)
		fmt.Println("write")
		fmt.Println(args)
		conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
		if err != nil {
			log.Fatalf("did not connect: %v", err)
		}
		defer conn.Close()
		c := pb.NewTokenManagementClient(conn)

		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		low := args[6]
		mid := args[8]
		high := args[10]
		dom := "low:" + low + ";mid:" + mid + ";high:" + high
		lowInt, _ := strconv.Atoi(args[6])
		midInt, _ := strconv.Atoi(args[8])
		_, partVal := hasher(lowInt, midInt)
		state := "partialValue:" + partVal + ";finalValue:undefined"
		r, err := c.WriteToken(ctx, &pb.NewToken{Name: args[4], Domain: dom, State: state, Id: args[2]})
		if err != nil {
			log.Fatalf("could not create token: %v", err)
		}
		fmt.Println("writen token: ", r)
	}
	if args[0] == "-drop" {
		fmt.Println(args)
		conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
		if err != nil {
			log.Fatalf("did not connect: %v", err)
		}
		defer conn.Close()
		c := pb.NewTokenManagementClient(conn)

		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		r, err := c.DropToken(ctx, &pb.TokenInfo{Id: args[1]})
		fmt.Println("Token has been dropped: ", r)
	}

}

func hasher(start int, end int) (int, string) {
	hasher := sha256.New()
	var min uint64
	var temp uint64
	var index int
	for i := start; i < end; i++ {
		hasher.Write([]byte(fmt.Sprintf("%s %d", "karthik", i)))
		temp = binary.BigEndian.Uint64(hasher.Sum(nil))
		if i == start {
			min = temp
			index = 0
		}
		if i != start && min > temp {
			min = temp
			index = i
		}
	}

	return index, strconv.Itoa(int(min))
}
