package questrade

type OrderLeg struct {
}

type Order struct {
	Id                       int
	Symbol                   string
	SymbolId                 int
	TotalQuantity            int
	OpenQuantity             int
	FilledQuantity           int
	CanceledQuantity         int
	Side                     string
	OrderType                string
	LimitPrice               int
	StopPrice                float64
	IsAllOrNone              bool
	IsAnonymous              bool
	IcebergQuantity          int
	MinQuantity              int
	AvgExecPrice             float64
	LastExecPrice            float64
	Source                   string
	TimeInForce              string
	GtdDate                  string
	State                    string
	ClientReasonStr          string
	ChainId                  int
	CreationTime             string
	UpdateTime               string
	Notes                    string
	PrimaryRoute             string
	SecondaryRoute           string
	OrderRoute               string
	VenueHoldingOrder        string
	CommissionCharged        float64
	ExchangeOrderId          string
	IsSignificantShareholder bool
	IsInsider                bool
	IsLimitOffsetInDollar    bool
	UserId                   int
	PlacementCommission      float64
	Legs                     []OrderLeg
	StrategyType             string
	TriggerStopPrice         float64
	OrderGroupId             int
	OrderClass               string
}

type Orders struct {
	Orders []Order
}
