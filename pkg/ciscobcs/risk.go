package ciscobcs

import "context"

func (c *Client) ListRiskMitigationSummaries(ctx context.Context) ([]*RiskMitigationSummary, error) {
	urlStr := c.getURLForSlug("/riskMitigation/summary")
	res, err := ListAllBCSObjects[RiskMitigationSummary](ctx, c, c.CustomerID, urlStr)
	for _, item := range res {
		item.CustomerID = c.CustomerID
	}
	return res, err
}

func (c *Client) ListRiskMitigationDetails(ctx context.Context) ([]*RiskMitigationDetails, error) {
	urlStr := c.getURLForSlug("/riskMitigation/details")
	res, err := ListAllBCSObjects[RiskMitigationDetails](ctx, c, c.CustomerID, urlStr)
	for _, item := range res {
		item.CustomerID = c.CustomerID
	}
	return res, err
}

// RiskMitigationDetails defines model for riskMitigationDetails.
type RiskMitigationDetails struct {
	// Customer ID is not included in the response, but should be added to associate this contract to a specific customer
	CustomerID string `gorm:"primaryKey;autoIncrement:false"`

	// The unique ID of the NP NetworkElement/Device.
	DeviceId *int `json:"deviceId,omitempty" gorm:"primaryKey;autoIncrement:false"`

	// The management IP address of the device.
	DeviceIp *string `json:"deviceIp,omitempty" gorm:"size:255"`

	// The Network Element Name of the device. When used as input, it can include % as wildcard.
	DeviceName *string `json:"deviceName,omitempty" gorm:"size:255"`

	// The Cisco Product Family of the hardware.  Values come from MDF for Chassis.
	ProductFamily *string `json:"productFamily,omitempty" gorm:"size:4000"`

	// The Cisco Product ID (PID) of the hardware.
	ProductId *string `json:"productId,omitempty" gorm:"size:32000"`

	// Severity of risk
	RiskCategory *string `json:"riskCategory,omitempty" gorm:"size:255"`

	// Risk score
	RiskScore *float32 `json:"riskScore,omitempty"`

	// The Type of Software running on the NP Network Element.  Common values include IOS, IOS XR, IOS-XE, NX-OS, etc.
	SwType *string `json:"swType,omitempty" gorm:"size:255"`

	// The Software Version of the device.
	SwVersion *string `json:"swVersion,omitempty" gorm:"size:255"`
}

// RiskMitigationSummary defines model for riskMitigationSummary.
type RiskMitigationSummary struct {
	// Customer ID is not included in the response, but should be added to associate this contract to a specific customer
	CustomerID string `gorm:"primaryKey;autoIncrement:false"`

	// Amount of devices marked as in high risk of crash
	HighRiskDeviceCount *int `json:"highRiskDeviceCount,omitempty"`

	// Amount of devices marked as in low risk of crash
	LowRiskDeviceCount *int `json:"lowRiskDeviceCount,omitempty"`

	// Amount of devices marked as in mediumrisk of crash
	MediumRiskDeviceCount *int `json:"mediumRiskDeviceCount,omitempty"`

	// The Cisco Product Family of the hardware.  Values come from MDF for Chassis.
	ProductFamily *string `json:"productFamily,omitempty" gorm:"primaryKey;autoIncrement:false;size:4000"`
}
