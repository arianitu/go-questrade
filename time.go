package questrade

import (
	"time"
)

type Time struct {
	Time string
}

func (c *Client) Time() (time.Time, error) {
	var timeResp Time
	err := c.NewRequest("GET", "time", nil, &timeResp)
	if err != nil {
		return time.Time{}, err
	}

	parsedTime, err := time.Parse(time.RFC3339, timeResp.Time)
	if err != nil {
		return time.Time{}, err
	}

	return parsedTime, nil
}
