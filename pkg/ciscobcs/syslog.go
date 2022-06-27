package ciscobcs

import "context"

func (c *Client) ListAllDailySyslog(ctx context.Context) ([]*SyslogDaily, error) {
	urlStr := c.getURLForSlug("/syslog/daily")
	res, err := ListAllBCSObjects[SyslogDaily](ctx, c, c.CustomerID, urlStr)
	for _, item := range res {
		item.CustomerID = c.CustomerID
	}
	return res, err
}

// SyslogDailyModel defines model for syslogDailyModel.
type SyslogDaily struct {
	// Customer ID is not included in the response, but should be added to associate this device to a specific customer
	CustomerID string `gorm:"primary_key;autoIncrement:false"`

	// Date of insertion
	Date *Date `json:"date,omitempty"`

	Description *string `json:"description,omitempty"`

	Devices *map[string]interface{} `json:"devices,omitempty"`

	Icseverity *string `json:"icseverity,omitempty"`

	// Lifecycle      *map[string]interface{} `json:"lifecycle,omitempty"`

	Msgtype *string `json:"msgtype,omitempty"`

	Recommendation *string `json:"recommendation,omitempty"`

	Severity *int `json:"severity,omitempty"`

	Timestamp *DateTime `json:"timestamp,omitempty"`

	Total *int `json:"total,omitempty"`
}

type SyslogDevice struct {
	Device *string `json:"device,omitempty"`

	DeviceId *int `json:"deviceId,omitempty"`

	Events *int `json:"events,omitempty"`
}
