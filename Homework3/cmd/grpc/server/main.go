package main

import (
	"awesomeProject/proto"
	"context"
	"fmt"
	"sync"

	"google.golang.org/grpc"
	"net"
)

type server struct {
	proto.UnimplementedBankServer
}

type Handler struct {
	accounts map[string]*proto.Account
	guard    *sync.RWMutex
}

var h Handler

func (s *server) DeleteAccountRequest(ctx context.Context, req *proto.DeleteAccount) (*proto.Status, error) {
	h.guard.Lock()
	if _, ok := h.accounts[req.Name]; !ok {
		h.guard.Unlock()
		return &proto.Status{Status: "account does not exist"}, nil
	}
	delete(h.accounts, req.Name)
	h.guard.Unlock()
	return &proto.Status{Status: ""}, nil
}

func (s *server) PathAccountRequest(ctx context.Context, req *proto.PathAccount) (*proto.Status, error) {
	h.guard.Lock()
	if _, ok := h.accounts[req.Name]; !ok {
		h.guard.Unlock()
		return &proto.Status{Status: "account does not exist"}, nil
	}
	h.accounts[req.Name] = &proto.Account{Name: req.Name, Balance: req.Balance}
	h.guard.Unlock()
	return &proto.Status{Status: ""}, nil
}

func (s *server) ChangeAccountRequest(ctx context.Context, req *proto.ChangeAccount) (*proto.Status, error) {
	h.guard.Lock()
	account, ok := h.accounts[req.Name]
	if !ok {
		h.guard.Unlock()
		return &proto.Status{Status: "account does not exist"}, nil
	}
	h.accounts[req.NewName] = &proto.Account{Name: req.Name, Balance: account.Balance}
	if req.Name != req.NewName {
		delete(h.accounts, req.Name)
	}
	h.guard.Unlock()
	return &proto.Status{Status: ""}, nil
}

func (s *server) CreateAccountRequest(ctx context.Context, req *proto.CreateAccount) (*proto.Status, error) {
	h.guard.Lock()
	h.accounts[req.Name] = &proto.Account{Name: req.Name, Balance: req.Balance}
	h.guard.Unlock()
	return &proto.Status{Status: ""}, nil
}

func (s *server) GetAccountResponse(ctx context.Context, req *proto.GetName) (*proto.Balance, error) {
	h.guard.Lock()
	account, ok := h.accounts[req.Name]
	if !ok {
		h.guard.Unlock()
		return &proto.Balance{Balance: 0, Status: "account does not exist"}, nil
	}
	h.guard.Unlock()
	return &proto.Balance{Balance: account.Balance, Status: ""}, nil
}

func main() {
	h.guard = new(sync.RWMutex)
	h.accounts = make(map[string]*proto.Account)
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 4567))
	if err != nil {
		panic(err)
	}
	s := grpc.NewServer()
	proto.RegisterBankServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		panic(err)
	}

}
