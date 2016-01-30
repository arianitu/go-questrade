package questrade

import (
	"fmt"
	"os"
	"testing"
	"time"
)

var client *Client

func TestMain(m *testing.M) {
	var err error
	client, err = NewClient(os.Getenv("REFRESH_TOKEN"), false)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	os.Exit(m.Run())
}

func TestTime(t *testing.T) {
	_, err := client.Time()
	if err != nil {
		t.Fatal(err)
	}
}

func TestAccounts(t *testing.T) {
	accounts, err := client.Accounts()
	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("%+v\n", accounts)
}

func TestAccountPositions(t *testing.T) {
	accounts, err := client.Accounts()
	if err != nil {
		t.Fatal(err)
	}
	positions, err := client.Positions(accounts[0])
	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("%+v\n", positions)
}

func TestAccountBalances(t *testing.T) {
	accounts, err := client.Accounts()
	if err != nil {
		t.Fatal(err)
	}
	balances, err := client.Balances(accounts[0])
	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("%+v\n", balances)
}

func TestAccountExecutions(t *testing.T) {
	accounts, err := client.Accounts()
	if err != nil {
		t.Fatal(err)
	}
	executions, err := client.Executions(accounts[0], time.Now().Add(-365*24*time.Hour), time.Now())
	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("%+v\n", executions)
}
