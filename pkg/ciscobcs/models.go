package ciscobcs

import (
	"database/sql/driver"
	"encoding/json"
	"time"
)

// BCSListResponse is used for all currently known list types with a generic Items field to capture the various results.
type BCSListResponse[T any] struct {
	Items   []*T `json:"items"`
	Page    int  `json:"page"`
	Pages   int  `json:"pages"`
	PerPage int  `json:"perPage"`
	Total   int  `json:"total"`
}

// ListOptions specifies the optional parameters to various List methods that
// support offset pagination and other features.
type ListOptions struct {
	// Query filter in JSON to search for specific fields. <br/>Example: {"physicaltype":"Module"}
	Filter *string `url:"filter,omitempty"`
	// Mask filter to lower the amount of fields returned, comma separated values, nested field between brackets. Example: items{physicaltype,currenteoxmilestone,currenteoxmilestonedate,productid}}
	Mask *string `url:"mask,omitempty"`
	// Page Number
	Page *int `url:"page,omitempty"`
	// Results Per Page
	PerPage *int `url:"perPage,omitempty"`
}

// DateFormat represents the date only format provided in the Cisco results.
const DateFormat = "2006-01-02"

// Date represente a date provided in the Cisco results.
type Date struct {
	time.Time
}

// MarshalJSON will marshal the date format provided in the Cisco results.
func (d Date) MarshalJSON() ([]byte, error) {
	return json.Marshal(d.Time.Format(DateFormat))
}

// UnmarshalJSON will unmarshal the date format provided in the Cisco results.
func (d *Date) UnmarshalJSON(data []byte) error {
	var dateStr string
	err := json.Unmarshal(data, &dateStr)
	if err != nil {
		return err
	}
	parsed, err := time.Parse(DateFormat, dateStr)
	if err != nil {
		return err
	}
	d.Time = parsed
	return nil
}

// String will return the date representation in the format provided by Cisco.
func (d Date) String() string {
	return d.Time.Format(DateFormat)
}

// DateTimeMinusTimezoneFormat represents the datetime field provided in the Cisco results which
// they have chosen not to provide any timezone information for.
const DateTimeMinusTimezoneFormat = "2006-01-02T15:04:05"

// DateTime represents the datetime field in the Cisco results which has no timezone information.
type DateTime struct {
	time.Time
}

// MarshalJSON will marshal the datetime format provided in the Cisco results.
func (d DateTime) MarshalJSON() ([]byte, error) {
	return json.Marshal(d.Time.Format(DateTimeMinusTimezoneFormat))
}

// UnmarshalJSON will unmarshal the datetime format provided in the Cisco results.
func (d *DateTime) UnmarshalJSON(data []byte) error {
	var dateStr string
	err := json.Unmarshal(data, &dateStr)
	if err != nil {
		return err
	}
	parsed, err := time.Parse(DateTimeMinusTimezoneFormat, dateStr)
	if err != nil {
		return err
	}
	d.Time = parsed
	return nil
}

// String will return the datetime representation in the format provided by Cisco.
func (d DateTime) String() string {
	return d.Time.Format(DateTimeMinusTimezoneFormat)
}

// Value is used by gorm to get a time.Time back instead of a Cisco DateTime.
func (d *DateTime) Value() (driver.Value, error) {
	if d == nil {
		return nil, nil
	}
	return d.Time, nil
}
