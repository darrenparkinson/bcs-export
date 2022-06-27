package ciscobcs

import "context"

func (c *Client) ListAllContractInformationPerSerialNumber(ctx context.Context) ([]*SerialNumberDetails, error) {
	urlStr := c.getURLForSlug("/contract/serials")
	res, err := ListAllBCSObjects[SerialNumberDetails](ctx, c, c.CustomerID, urlStr)
	for _, item := range res {
		item.CustomerID = c.CustomerID
		if item.BasePidList != nil && len(*item.BasePidList) > 0 {
			bpl := item.GetBasePidList()
			item.FirstBasePid = bpl[0].GetBasePid()
		}
		if item.OrderablePidList != nil && len(*item.OrderablePidList) > 0 {
			opl := *item.OrderablePidList
			item.FirstOrderablePid = opl[0].GetOrderablePid()
			item.FirstOrderablePidDescription = opl[0].GetItemDescription()
			item.FirstOrderablePidPillarCode = opl[0].GetPillarCode()
			item.FirstOrderablePidPosition = opl[0].GetItemPosition()
			item.FirstOrderablePidType = opl[0].GetItemType()
		}
	}
	return res, err
}

// SerialNumberDetails defines model for serialNumberDetails.
type SerialNumberDetails struct {
	// Customer ID is not included in the response, but should be added to associate this contract to a specific customer
	CustomerID string `gorm:"primaryKey;autoIncrement:false"`

	// FirstBasePID is not included in the response -- we will add it if there is at least one base PID in the Base PID List
	FirstBasePid string

	BasePidList *[]ContractBasePID `json:"basePidList,omitempty" gorm:"-"`

	// Address field for the contract install site.
	ContractSiteAddress1 *string `json:"contractSiteAddress1,omitempty"`

	// City field for the contract install site; for example,
	ContractSiteCity *string `json:"contractSiteCity,omitempty"`

	// Country field for the contract install site
	ContractSiteCountry *string `json:"contractSiteCountry,omitempty"`

	// Customer name associated to the contract install site.
	ContractSiteCustomerName *string `json:"contractSiteCustomerName,omitempty"`

	// State field for the contract install site
	ContractSiteStateProvince *string `json:"contractSiteStateProvince,omitempty"`

	// End date of the covered product line in the following format: YYYY-MM-DD
	CoveredProductLineEndDate *Date `json:"coveredProductLineEndDate,omitempty"`

	// Number of the record in the results. Appears to be null in most cases.
	Id *int `json:"id,omitempty"`

	// Indicates whether the specified serial number is covered by a service contract; one of the following values: YES or NO. If the serial number is covered by a service contract, the value is Yes.
	IsCovered *string `json:"isCovered,omitempty"`

	OrderablePidList *[]OrderablePid `json:"orderablePidList,omitempty" gorm:"-"`

	// FirstOrderablePid is not included in the response. We should set it separately.
	FirstOrderablePid            string
	FirstOrderablePidType        string
	FirstOrderablePidDescription string
	FirstOrderablePidPosition    string
	FirstOrderablePidPillarCode  string

	// Parent serial number. The value of parent_sr_no will be the same as the value for sr_no if the item is a MAJOR item.
	ParentSrNo *string `json:"parentSrNo,omitempty"`

	// Service contract number
	ServiceContractNumber *string `json:"serviceContractNumber,omitempty"`

	// Description of the service type
	ServiceLineDescr *string `json:"serviceLineDescr,omitempty"`

	// Serial number of the device.
	SrNo *string `json:"srNo,omitempty" gorm:"primaryKey;autoIncrement:false"`

	// End date of the warranty for the specified serial number in the following format: YYYY-MM-DD
	WarrantyEndDate *Date `json:"warrantyEndDate,omitempty"`

	// Warranty service type
	WarrantyType *string `json:"warrantyType,omitempty"`

	// Link to the description of the warranty type.
	WarrantyTypeDescription *string `json:"warrantyTypeDescription,omitempty"`
}

// OrderablePid defines model for orderablePid.
type OrderablePid struct {
	// Orderable product description for the specified serial number
	ItemDescription *string `json:"itemDescription,omitempty"`

	// Orderable product position for the specified serial number
	ItemPosition *string `json:"itemPosition,omitempty"`

	// Orderable product type for the specified serial number
	ItemType *string `json:"itemType,omitempty"`

	// Orderable product identifiers for the specified serial number
	OrderablePid *string `json:"orderablePid,omitempty"`

	// Orderable product identifiers for the specified serial number
	PillarCode *string `json:"pillarCode,omitempty"`
}

// ContractBasePID defines model for contractBasePID.
type ContractBasePID struct {
	// Base or manufacturing product identifiers related to the specified serial number.
	BasePid *string `json:"basePid,omitempty"`
}
