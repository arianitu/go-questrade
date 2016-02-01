package questrade

import (
	"errors"
	"fmt"
	"net/url"
	"time"
)

type Account struct {
	Type              string
	Number            string
	Status            string
	IsPrimary         bool
	IsBilling         bool
	ClientAccountType string
}

type Accounts struct {
	Accounts []Account
	UserId   int64
}

type Position struct {
	Symbol             string
	SymbolId           uint64
	OpenQuantity       float64
	ClosedQuantity     float64
	CurrentMarketValue float64
	CurrentPrice       float64
	AverageEntryPrice  float64
	ClosedPnL          float64
	OpenPnL            float64
	TotalCost          float64
	IsRealTime         bool
	IsUnderReorg       bool
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
	SymbolId                 uint64
	Quantity                 int
	Side                     string
	Price                    float64
	Id                       int
	OrderId                  int
	OrderChainId             int
	ExchangeExecId           string
	Timestamp                time.Time
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

type Activities struct {
	Activities []Activity
}

type Activity struct {
	TradeDate       time.Time
	TransactionDate time.Time
	SettlementDate  time.Time
	Action          string
	Symbol          string
	SymbolId        uint64
	Description     string
	Currency        string
	Quantity        float64
	Price           float64
	GrossAmount     float64
	Commission      float64
	NetAmount       float64
	Type            string
}

var (
	OrderNotFound = errors.New("Order not found")
)

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
	parameters := url.Values{}
	parameters.Add("startTime", startTime.Format(time.RFC3339))
	parameters.Add("endTime", endTime.Format(time.RFC3339))

	u := fmt.Sprintf("accounts/%vp/executions?%v", account.Number, parameters.Encode())
	fmt.Println(u)

	var executions Executions
	err := c.NewRequest("GET", u, nil, &executions)
	if err != nil {
		return nil, err
	}

	return &executions, nil
}

func (c *Client) getOrders(account Account, parameters url.Values) (*Orders, error) {
	u := fmt.Sprintf("accounts/%vp/orders?%v", account.Number, parameters.Encode())
	fmt.Println(u)

	var orders Orders
	err := c.NewRequest("GET", u, nil, &orders)
	if err != nil {
		return nil, err
	}

	return &orders, nil
}

func (c *Client) OrderById(account Account, orderId int, stateFilter string) (*Order, error) {
	parameters := url.Values{}
	parameters.Add("orderId", string(orderId))
	parameters.Add("stateFilter", stateFilter)

	orders, err := c.getOrders(account, parameters)
	if err != nil {
		return nil, err
	}

	if len(orders.Orders) == 0 {
		return nil, OrderNotFound
	}

	return &orders.Orders[0], nil
}

func (c *Client) Orders(account Account, startTime time.Time, endTime time.Time, stateFilter string) (*Orders, error) {
	parameters := url.Values{}
	parameters.Add("startTime", startTime.Format(time.RFC3339))
	parameters.Add("endTime", endTime.Format(time.RFC3339))
	parameters.Add("stateFilter", stateFilter)

	return c.getOrders(account, parameters)
}

func (c *Client) Activities(account Account, startTime time.Time, endTime time.Time) (*Activities, error) {
	parameters := url.Values{}
	parameters.Add("startTime", startTime.Format(time.RFC3339))
	parameters.Add("endTime", endTime.Format(time.RFC3339))

	u := fmt.Sprintf("accounts/%vp/activities?%v", account.Number, parameters.Encode())
	fmt.Println(u)

	var activities Activities
	err := c.NewRequest("GET", u, nil, &activities)
	if err != nil {
		return nil, err
	}

	return &activities, nil

}
