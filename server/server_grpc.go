package main

import (
	"context"
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"log"
	"net"
	"strconv"
	"strings"

	"google.golang.org/grpc"
	pb "karthik.com/go-tokenmgmt-grpc/proto"
)

const (
	port = ":50051"
)

type TokenManagementServer struct {
	pb.UnimplementedTokenManagementServer
	token_list *pb.TokenList
}

func (s *TokenManagementServer) CreateNewToken(ctx context.Context, in *pb.NewToken) (*pb.Token, error) {
	log.Printf("Received: %v", in.GetName())
	created_token := &pb.Token{Name: in.GetName(), Domain: in.GetDomain(), Id: in.GetId(), State: in.GetState()}
	s.token_list.Tokens = append(s.token_list.Tokens, created_token)
	return created_token, nil
}

func (s *TokenManagementServer) GetToekns(ctx context.Context, in *pb.Token) (*pb.Token, error) {
	for i := 0; i < len(s.token_list.Tokens); i++ {
		if s.token_list.Tokens[i].Id == in.GetId() {
			if s.token_list.Tokens[i].Domain != "undefined" {
				dom := s.token_list.Tokens[i].Domain
				state := s.token_list.Tokens[i].State
				if strings.Split(strings.Split(state, ";")[1], ":")[1] != "undefined" {
					return s.token_list.Tokens[i], nil
				}
				mid := strings.Split(strings.Split(dom, ";")[1], ":")[1]
				high := strings.Split(strings.Split(dom, ";")[2], ":")[1]
				partVal := strings.Split(strings.Split(state, ";")[0], ":")[1]
				midInt, _ := strconv.Atoi(mid)
				highInt, _ := strconv.Atoi(high)
				_, hashVal := hasher(midInt, highInt)
				parInt, _ := strconv.Atoi(partVal)
				hashInt, _ := strconv.Atoi(hashVal)
				var finalSta string
				if parInt < hashInt {
					finalSta = "PartialValue:" + strconv.Itoa(parInt) + ";FinalValue:" + strconv.Itoa(parInt)
				} else {
					finalSta = "PartialValue:" + strconv.Itoa(parInt) + ";FinalValue:" + strconv.Itoa(hashInt)
				}
				s.token_list.Tokens[i].State = finalSta
				return s.token_list.Tokens[i], nil
			}
			return s.token_list.Tokens[i], nil
		}
	}
	return nil, nil
}

func (s *TokenManagementServer) DropToken(ctx context.Context, in *pb.TokenInfo) (*pb.EmptyToken, error) {
	for i := 0; i < len(s.token_list.Tokens); i++ {
		if s.token_list.Tokens[i].Id == in.GetId() {
			s.token_list.Tokens = append(s.token_list.Tokens[:i], s.token_list.Tokens[i+1:]...)
			break
		}
	}
	return nil, nil
}

func (s *TokenManagementServer) WriteToken(ctx context.Context, in *pb.NewToken) (*pb.Token, error) {
	var index int
	var flag bool
	for i := 0; i < len(s.token_list.Tokens); i++ {
		if s.token_list.Tokens[i].Id == in.GetId() {
			s.token_list.Tokens[i].Name = in.GetName()
			s.token_list.Tokens[i].State = in.GetState()
			s.token_list.Tokens[i].Domain = in.GetDomain()
			index = i
			flag = true
			break
		}
	}
	if flag != true {
		log.Fatal("Token Not Found! Please enter correct ID.")
		return nil, nil
	}
	return s.token_list.Tokens[index], nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterTokenManagementServer(s, &TokenManagementServer{token_list: &pb.TokenList{}})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func hasher(start int, end int) (int, string) {
	hasher := sha256.New()
	var min uint64
	var temp uint64
	var index int
	for i := start; i <= end; i++ {
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
