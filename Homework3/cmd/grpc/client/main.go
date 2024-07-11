package main

import (
	"awesomeProject/proto"
	"context"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"strconv"
	"time"
)

type Command struct {
	Port    int
	Host    string
	Cmd     string
	Name    string
	NewName string
	Amount  int
}

func (c *Command) Do() error {
	switch c.Cmd {
	case "create":
		if err := c.create(); err != nil {
			return fmt.Errorf("create account failed: %w", err)
		}

		return nil
	case "get":
		if err := c.get(); err != nil {
			return fmt.Errorf("get account failed: %w", err)
		}

		return nil
	case "delete":
		if err := c.delete(); err != nil {
			return fmt.Errorf("delete account failed: %w", err)
		}

		return nil
	case "change":
		if err := c.changeAccount(); err != nil {
			return fmt.Errorf("change account failed: %w", err)
		}

		return nil
	case "path":
		if err := c.pathAccount(); err != nil {
			return fmt.Errorf("path account failed: %w", err)
		}
		return nil

	default:
		return fmt.Errorf("unknown command %s", c.Cmd)
	}
}

func (c *Command) create() error {
	conn, err := grpc.NewClient(c.Host+":"+strconv.Itoa(c.Port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return fmt.Errorf("create grpc client failed: %w", err)
	}
	defer func() {
		_ = conn.Close()

	}()
	con := proto.NewBankClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := con.CreateAccountRequest(ctx, &proto.CreateAccount{Name: c.Name, Balance: int64(c.Amount)})
	if err != nil {
		return fmt.Errorf("create account failed: %w", err)
	}
	if res.Status == "" {
		res.Status = "ok"
	}
	fmt.Println("status:", res.Status)
	return nil
}

func (c *Command) get() error {
	conn, err := grpc.NewClient(c.Host+":"+strconv.Itoa(c.Port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return fmt.Errorf("get grpc client failed: %w", err)
	}
	defer func() {
		_ = conn.Close()

	}()
	con := proto.NewBankClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := con.GetAccountResponse(ctx, &proto.GetName{Name: c.Name})
	if err != nil {
		return fmt.Errorf("create account failed: %w", err)
	}
	if res.Status == "" {
		res.Status = "ok"
	} else {
		fmt.Println("status:", res.Status)
		return nil
	}
	fmt.Println("status:", res.Status)
	fmt.Println("balance:", res.Balance)
	return nil
}

func (c *Command) delete() error {
	conn, err := grpc.NewClient(c.Host+":"+strconv.Itoa(c.Port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return fmt.Errorf("delete grpc client failed: %w", err)
	}
	defer func() {
		_ = conn.Close()

	}()
	con := proto.NewBankClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := con.DeleteAccountRequest(ctx, &proto.DeleteAccount{Name: c.Name})
	if err != nil {
		return fmt.Errorf("delete account failed: %w", err)
	}
	if res.Status == "" {
		res.Status = "ok"
	}
	fmt.Println("status:", res.Status)
	return nil
}

func (c *Command) changeAccount() error {
	conn, err := grpc.NewClient(c.Host+":"+strconv.Itoa(c.Port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return fmt.Errorf("change grpc client failed: %w", err)
	}
	defer func() {
		_ = conn.Close()

	}()
	con := proto.NewBankClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := con.ChangeAccountRequest(ctx, &proto.ChangeAccount{Name: c.Name, NewName: c.NewName})
	if err != nil {
		return fmt.Errorf("change account failed: %w", err)
	}
	if res.Status == "" {
		res.Status = "ok"
	}
	fmt.Println("status:", res.Status)
	return nil
}

func (c *Command) pathAccount() error {
	conn, err := grpc.NewClient(c.Host+":"+strconv.Itoa(c.Port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return fmt.Errorf("path grpc client failed: %w", err)
	}
	defer func() {
		_ = conn.Close()

	}()
	con := proto.NewBankClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := con.PathAccountRequest(ctx, &proto.PathAccount{Name: c.Name, Balance: int64(c.Amount)})
	if err != nil {
		return fmt.Errorf("path account failed: %w", err)
	}
	if res.Status == "" {
		res.Status = "ok"
	}
	fmt.Println("status:", res.Status)
	return nil
}

func main() {
	portVal := flag.Int("port", 4567, "server port")
	hostVal := flag.String("host", "0.0.0.0", "server host")
	cmdVal := flag.String("cmd", "", "command to execute")
	nameVal := flag.String("name", "", "name of account")
	amountVal := flag.Int("amount", 0, "amount of account")
	newName := flag.String("new_name", "", "new name of account")

	flag.Parse()

	cmd := Command{
		Port:    *portVal,
		Host:    *hostVal,
		Cmd:     *cmdVal,
		Name:    *nameVal,
		Amount:  *amountVal,
		NewName: *newName,
	}

	if err := cmd.Do(); err != nil {
		panic(err)
	}
}
