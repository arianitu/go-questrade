package questrade

import (
	"fmt"
	"net/url"
	"time"
)

type Account struct {
	Type              string
	Number            string
	Status            string
	isPrimary         bool
	isBilling         bool
	ClientAccountType string
}

type Accounts struct {
	Accounts []Account
	UserId   int64
}

type Position struct {
	Symbol             string
	SymbolId           int
	OpenQuantity       float64
	ClosedQuantity     float64
	CurrentMarketValue float64
	CurrentPrice       float64
	AverageEntryPrice  float64
	ClosedPnL          float64
	OpenPnL            float64
	TotalCost          float64
	IsRealTime         bool
	isUnderReorg       bool
}

type Positions struct {
	Positions []Position
}

type Balance struct {
	Currency          string
	Cash              float64
	MarketValue       float64
	TotalEquity       float64
	BuyingPower       float64
	MaintenanceExcess float64
	IsRealTime        bool
}

type Balances struct {
	PerCurrencyBalances    []Balance
	CombinedBalances       []Balance
	SodPerCurrencyBalances []Balance
	SodCombinedBalances    []Balance
}

type Execution struct {
	Symbol                   string
	SymbolId                 int
	Quantity                 int
	Side                     string
	Price                    float64
	Id                       int
	OrderId                  int
	OrderChainId             int
	ExchangeExecId           string
	Timestamp                string
	Notes                    string
	Venue                    string
	TotalCost                float64
	OrderPlacementCommission float64
	Commission               float64
	ExecutionFee             float64
	SecFee                   float64
	CanadianExecutionFee     int
	ParentId                 int
}

type Executions struct {
	Executions []Execution
}

func (c *Client) Accounts() ([]Account, error) {
	var accounts Accounts
	err := c.NewRequest("GET", "accounts", nil, &accounts)
	if err != nil {
		return nil, err
	}

	return accounts.Accounts, nil
}

func (c *Client) Positions(account Account) (*Positions, error) {
	u := fmt.Sprintf("accounts/%v/positions", account.Number)

	var positions Positions
	err := c.NewRequest("GET", u, nil, &positions)
	if err != nil {
		return nil, err
	}

	return &positions, nil
}

func (c *Client) Balances(account Account) (*Balances, error) {
	u := fmt.Sprintf("accounts/%v/balances", account.Number)

	var balances Balances
	err := c.NewRequest("GET", u, nil, &balances)
	if err != nil {
		return nil, err
	}

	return &balances, nil
}

func (c *Client) Executions(account Account, startTime time.Time, endTime time.Time) (*Executions, error) {
	v := url.Values{}
	v.Add("startTime", startTime.Format(time.RFC3339))
	v.Add("endTime", endTime.Format(time.RFC3339))

	u := fmt.Sprintf("accounts/%v/executions?%v", account.Number, v.Encode())
	fmt.Println(u)

	var executions Executions
	err := c.NewRequest("GET", u, nil, &executions)
	if err != nil {
		return nil, err
	}

	return &executions, nil
}
