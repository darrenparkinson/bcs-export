package ciscobcs

import (
	"context"
)

func (c *Client) ListAllCBPDetailsPerDeviceID(ctx context.Context) ([]*CBPDetail, error) {
	urlStr := c.getURLForSlug("/cbp/details")
	res, err := ListAllBCSObjects[CBPDetail](ctx, c, c.CustomerID, urlStr)
	for _, item := range res {
		item.CustomerID = c.CustomerID
	}
	return res, err
}

func (c *Client) ListCBPDetailsPerDeviceID(ctx context.Context, options ...*ListOptions) (*BCSListResponse[CBPDetail], error) {
	urlStr := c.getURLForSlug("/cbp/details")
	res, err := ListBCSObject[CBPDetail](ctx, c, c.CustomerID, urlStr, options...)
	for _, item := range res.Items {
		item.CustomerID = c.CustomerID
	}
	return res, err
}

func (c *Client) ListAllCBPRules(ctx context.Context) ([]*CBPRule, error) {
	urlStr := c.getURLForSlug("/cbp/rules")
	res, err := ListAllBCSObjects[CBPRule](ctx, c, c.CustomerID, urlStr)
	return res, err
}

func (c *Client) ListCBPRules(ctx context.Context, options ...*ListOptions) (*BCSListResponse[CBPRule], error) {
	urlStr := c.getURLForSlug("/cbp/rules")
	res, err := ListBCSObject[CBPRule](ctx, c, c.CustomerID, urlStr, options...)
	return res, err
}

// CBPDetails defines model for CBPDetails.
type CBPDetail struct {
	// // Customer ID is not included in the response, but should be added to associate this device to a specific customer
	CustomerID string `gorm:"primary_key;autoIncrement:false"`

	// The internal Config BP Nugget ID from the AI-X system. Used to cross-reference with other API results.
	BpNuggetId *int `json:"bpNuggetId,omitempty"`

	// The internal Config BP Rule ID from the AI-X system. Used to cross-reference with other API results.
	BpRuleId *int `json:"bpRuleId,omitempty" gorm:"primaryKey;autoIncrement:false"`

	// The source of a Configuration file.  The primary source will be "STANDARD". But in some devices, it might be "CONTEXT" or "ADMIN".
	ConfigSource *string `json:"configSource,omitempty"`

	// The ID of the NP NetworkElement/Device.
	DeviceId *int `json:"deviceId,omitempty" gorm:"primaryKey;autoIncrement:false"`
}

// CBPRules defines model for CBPRules.
type CBPRule struct {
	// The Caveat associated with a Config BP Rule.
	BpCaveat *string `json:"bpCaveat,omitempty" gorm:"size:4000"`

	// The Corrective Action associated with a Config BP Rule.
	BpCorrectiveAction *string `json:"bpCorrectiveAction,omitempty" gorm:"size:4000"`

	// The Description associated with a Config BP Rule.
	BpDescription *string `json:"bpDescription,omitempty" gorm:"size:4000"`

	// The internal Config BP Nugget ID from the AI-X system. Used to cross-reference with other API results.
	BpNuggetId *int `json:"bpNuggetId,omitempty"`

	// The Primary Technology associated with a Config BP Rule.
	BpPrimaryTechnology *string `json:"bpPrimaryTechnology,omitempty"`

	// The Recommendation associated with a Config BP Rule.
	BpRecommendation *string `json:"bpRecommendation,omitempty" gorm:"size:4000"`

	// The Risk associated with a Config BP Rule. Valid values include: High, Medium, Low.
	BpRisk *string `json:"bpRisk,omitempty"`

	// The internal Config BP Rule ID from the AI-X system. Used to cross-reference with other API results.
	BpRuleId *int `json:"bpRuleId,omitempty" gorm:"primaryKey;autoIncrement:false"`

	// The Secondary Technologies associated with a Config BP Rule. Can be multiple values separated with commas.
	BpSecondaryTechnology *string `json:"bpSecondaryTechnology,omitempty"`

	// The Config BP exception headline / title.
	BpTitle *string `json:"bpTitle,omitempty"`

	// The date the record orrule was created in NP database.  For devices, a new record is created whenever a unique name+sysobjectid combination is seen in the collector.
	CreateDate *DateTime `json:"createDate,omitempty"`

	// The Type of Software running on the NP Network Element.  Common values include IOS, IOS XR, IOS-XE, NX-OS, etc.
	SwType *string `json:"swType,omitempty"`

	// The timestamp when the data or rule was last updated.
	UpdateDate *DateTime `json:"updateDate,omitempty"`
}

// CBPRulesReferences defines model for CBPRulesReferences.
type CBPRulesReference struct {
	// The internal Config BP Rule ID from the AI-X system. Used to cross-reference with other API results.
	BpRuleId *int `json:"bpRuleId,omitempty"`

	// The Config BP reference URL.  Any rule can have 1 or more reference URLs.
	BpUrl *string `json:"bpUrl,omitempty"`

	// The Config BP reference URL Title associated with bpUrl.
	BpUrlTitle *string `json:"bpUrlTitle,omitempty"`
}

// CBPSummary defines model for CBPSummary.
type CBPSummary struct {
	// The internal Config BP Nugget ID from the AI-X system. Used to cross-reference with other API results.
	BpNuggetId *int `json:"bpNuggetId,omitempty"`

	// The Primary Technology associated with a Config BP Rule.
	BpPrimaryTechnology *string `json:"bpPrimaryTechnology,omitempty"`

	// The Risk associated with a Config BP Rule. Valid values include: High, Medium, Low.
	BpRisk *string `json:"bpRisk,omitempty"`

	// The internal Config BP Rule ID from the AI-X system. Used to cross-reference with other API results.
	BpRuleId *int `json:"bpRuleId,omitempty"`

	// The Secondary Technologies associated with a Config BP Rule. Can be multiple values separated with commas.
	BpSecondaryTechnology *string `json:"bpSecondaryTechnology,omitempty"`

	// The Config BP exception headline / title.
	BpTitle *string `json:"bpTitle,omitempty"`

	// The Type of Software running on the NP Network Element.  Common values include IOS, IOS XR, IOS-XE, NX-OS, etc.
	SwType *string `json:"swType,omitempty"`

	// The number of unique matching devices in summary APIs.  For Custom Config and BP, this is the number of devices with exceptions.  For Feature, this is the number of devices with the feature match.  For tracks, this is the total number of devices in the track.
	TotalDevices *int `json:"totalDevices,omitempty"`
}
