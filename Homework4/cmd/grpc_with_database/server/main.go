package main

import (
	"awesomeProject/proto1"
	"context"
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/jackc/pgx/v5/stdlib"
	"google.golang.org/grpc"
	"net"
)

type server struct {
	proto1.UnimplementedBankServer
}

var connectionString = "host=0.0.0.0 port=5432 dbname=postgres user=postgres password=mysecretpassword sslmode=disable"

func (s *server) DeleteAccountRequest(ctx context.Context, req *proto1.DeleteAccount) (*proto1.Status, error) {
	db, err := sql.Open("pgx", connectionString)
	if err != nil {
		return &proto1.Status{}, err
	}

	defer db.Close()

	if err := db.Ping(); err != nil {
	}
	rows, err := db.QueryContext(ctx, "SELECT name, balance FROM accounts where name = $1", req.Name)
	if err != nil {
		return &proto1.Status{}, err
	}
	defer rows.Close()
	var accounts []*proto1.Account
	for rows.Next() {
		var account proto1.Account
		if err := rows.Scan(&account.Name, &account.Balance); err != nil {
		}
		accounts = append(accounts, &account)
	}
	if len(accounts) == 0 {
		return &proto1.Status{Status: "account does not exist"}, nil
	}
	if len(accounts) > 1 {
		return &proto1.Status{}, errors.New("multiple accounts found")
	}
	db.Exec("DELETE FROM accounts WHERE name = $1", req.Name)
	return &proto1.Status{Status: "ok"}, nil
}

func (s *server) PathAccountRequest(ctx context.Context, req *proto1.PathAccount) (*proto1.Status, error) {
	db, err := sql.Open("pgx", connectionString)
	if err != nil {
		return &proto1.Status{}, err
	}

	defer db.Close()

	if err := db.Ping(); err != nil {
	}
	rows, err := db.QueryContext(ctx, "SELECT name, balance FROM accounts where name = $1", req.Name)
	if err != nil {
		return &proto1.Status{}, err
	}
	defer rows.Close()
	var accounts []*proto1.Account
	for rows.Next() {
		var account proto1.Account
		if err := rows.Scan(&account.Name, &account.Balance); err != nil {
		}
		accounts = append(accounts, &account)
	}
	if len(accounts) == 0 {
		return &proto1.Status{Status: "account does not exist"}, nil
	}
	if len(accounts) > 1 {
		return &proto1.Status{}, errors.New("multiple accounts found")
	}
	if _, err := db.Exec("UPDATE accounts SET balance = $1 WHERE name = $2", req.Balance, req.Name); err != nil {
		return &proto1.Status{}, err
	}
	return &proto1.Status{}, nil
}

func (s *server) ChangeAccountRequest(ctx context.Context, req *proto1.ChangeAccount) (*proto1.Status, error) {
	db, err := sql.Open("pgx", connectionString)
	if err != nil {
		return &proto1.Status{}, err
	}

	defer db.Close()

	if err := db.Ping(); err != nil {
	}
	rows, err := db.QueryContext(ctx, "SELECT name, balance FROM accounts where name = $1", req.Name)
	if err != nil {
		return &proto1.Status{}, err
	}
	defer rows.Close()
	var accounts []*proto1.Account
	for rows.Next() {
		var account proto1.Account
		if err := rows.Scan(&account.Name, &account.Balance); err != nil {
		}
		accounts = append(accounts, &account)
	}
	if len(accounts) == 0 {
		return &proto1.Status{Status: "account does not exist"}, nil
	}
	if len(accounts) > 1 {
		return &proto1.Status{}, errors.New("multiple accounts found")
	}
	if _, err := db.Exec("UPDATE accounts SET name = $1 WHERE name = $2", req.NewName, req.Name); err != nil {
		return &proto1.Status{}, err
	}
	return &proto1.Status{}, nil
}

func (s *server) CreateAccountRequest(ctx context.Context, req *proto1.CreateAccount) (*proto1.Status, error) {
	db, err := sql.Open("pgx", connectionString)
	if err != nil {
		return &proto1.Status{Status: "err"}, err
	}

	defer db.Close()

	if err := db.Ping(); err != nil {
		return &proto1.Status{Status: "err"}, err
	}
	rows, err := db.QueryContext(ctx, "SELECT name, balance FROM accounts where name = $1", req.Name)
	if err != nil {
		return &proto1.Status{Status: "err"}, err
	}
	defer rows.Close()
	var accounts []*proto1.Account
	for rows.Next() {
		var account proto1.Account
		if err := rows.Scan(&account.Name, &account.Balance); err != nil {
			return &proto1.Status{Status: "err"}, err
		}
		accounts = append(accounts, &account)
	}
	if len(accounts) > 0 {
		return &proto1.Status{Status: "account already exists"}, err
	}
	if _, err := db.ExecContext(ctx, "INSERT INTO accounts(name, balance) VALUES ($1, $2)", req.Name, req.Balance); err != nil {
		return &proto1.Status{Status: "err"}, err
	}
	return &proto1.Status{Status: "err"}, nil
}

func (s *server) GetAccountResponse(ctx context.Context, req *proto1.GetName) (*proto1.Balance, error) {
	db, err := sql.Open("pgx", connectionString)
	if err != nil {
		return &proto1.Balance{}, err
	}

	defer db.Close()

	if err := db.Ping(); err != nil {
		return &proto1.Balance{Status: "err"}, err
	}
	rows, err := db.QueryContext(ctx, "SELECT name, balance FROM accounts where name = $1", req.Name)
	if err != nil {
		return &proto1.Balance{}, err
	}

	defer rows.Close()
	var accounts []*proto1.Account
	for rows.Next() {
		var account proto1.Account
		if err := rows.Scan(&account.Name, &account.Balance); err != nil {
		}
		accounts = append(accounts, &account)
	}
	if len(accounts) == 0 {
		return &proto1.Balance{Status: "account does not exist"}, nil
	}
	if len(accounts) > 1 {
		return &proto1.Balance{Status: "multiple accounts found"}, errors.New("multiple accounts found")
	}
	return &proto1.Balance{Status: "", Balance: accounts[0].Balance}, nil
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 4567))
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	proto1.RegisterBankServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		panic(err)
	}

}
