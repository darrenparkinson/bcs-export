package ciscobcs

import (
	"context"
)

func (c *Client) ListAllDevices(ctx context.Context) ([]*Device, error) {
	urlStr := c.getURLForSlug("/inventory/devices")
	res, err := ListAllBCSObjects[Device](ctx, c, c.CustomerID, urlStr)
	for _, item := range res {
		item.CustomerID = c.CustomerID
	}
	return res, err
}

func (c *Client) ListDevices(ctx context.Context, options ...*ListOptions) (*BCSListResponse[Device], error) {
	urlStr := c.getURLForSlug("/inventory/devices")
	res, err := ListBCSObject[Device](ctx, c, c.CustomerID, urlStr, options...)
	for _, item := range res.Items {
		item.CustomerID = c.CustomerID
	}
	return res, err
}

func (c *Client) ListAllAssets(ctx context.Context) ([]*Asset, error) {
	urlStr := c.getURLForSlug("/inventory/assets")
	res, err := ListAllBCSObjects[Asset](ctx, c, c.CustomerID, urlStr)
	for _, item := range res {
		item.CustomerID = c.CustomerID
	}
	return res, err
}

func (c *Client) ListAssets(ctx context.Context, options ...*ListOptions) (*BCSListResponse[Asset], error) {
	urlStr := c.getURLForSlug("/inventory/assets")
	res, err := ListBCSObject[Asset](ctx, c, c.CustomerID, urlStr, options...)
	for _, item := range res.Items {
		item.CustomerID = c.CustomerID
	}
	return res, err
}

// Device defines model for Device.
type Device struct {
	// Customer ID is not included in the response, but should be added to associate this device to a specific customer
	CustomerID string `gorm:"primaryKey;autoIncrement:false"`

	// The collector identifier, which can be either a 4 character collectorid or the applianceid.
	Collector *string `json:"collector,omitempty" gorm:"size:255"`

	// The Configuration register of the device.
	ConfigRegister *string `json:"configRegister,omitempty" gorm:"size:255"`

	// The status of Configuration collection.  Completed means the config was successfully collected.  NotAvailable means the config was not collected.  NotSupported means the device does not support collection of an ASCii config via CLI.
	ConfigStatus *string `json:"configStatus,omitempty" gorm:"size:255"`

	// The time when the collector last successfully collected the configuration from the device.
	ConfigTime *DateTime `json:"configTime,omitempty"`

	// The date the record or rule was created in NP database.  For devices, a new record is created whenever a unique name+sysobjectid combination is seen in the collector.
	CreateDate *DateTime `json:"createDate,omitempty"`

	// The unique ID of the NP NetworkElement/Device.
	DeviceId *int `json:"deviceId,omitempty" gorm:"primaryKey;autoIncrement:false"`

	// The management IP address of the device.
	DeviceIp *string `json:"deviceIp,omitempty" gorm:"size:255"`

	// The Network Element Name of the device.  When used as input, it can include % as wildcard.
	DeviceName *string `json:"deviceName,omitempty" gorm:"size:255"`

	// The status of the device as reported by the collector.  Usually will be either ACTIVE or DEVICE NOT REACHABLE.
	DeviceStatus *string `json:"deviceStatus,omitempty" gorm:"size:255"`

	// The hostname (SNMP sysName) of the device.  It will be fully-qualified name, if domain name is set on the device.
	DeviceSysName *string `json:"deviceSysName,omitempty" gorm:"size:255"`

	// The type of the device.  Values include Managed Chassis, Managed Multi-Chassis, SDR, Contexts, and IOS-XR Admin
	DeviceType *string `json:"deviceType,omitempty" gorm:"size:255"`

	// The name of the software feature set running on the device.  This data is primarily available for IOS.
	FeatureSetdesc *string `json:"featureSetdesc,omitempty" gorm:"size:4000"`

	// The Image Name of the software on the Network Element.
	ImageName *string `json:"imageName,omitempty" gorm:"size:255"`

	// Indicates whether the device is in a collector seedfile (true) or has been logically created by NP (false).  This is important for some KPI measurements to be accurate.
	InSeedFile *bool `json:"inSeedFile,omitempty"`

	// The status of Inventory collection.  Completed means some SNMP inventory was successfully collected.  NotAvailable means SNMP inventory was not collected.  NotSupported means the device was not in CSPC to be collected.
	InventoryStatus *string `json:"inventoryStatus,omitempty" gorm:"size:255"`

	// The time when the collector last successfully collected inventory from the device.
	InventoryTime *DateTime `json:"inventoryTime,omitempty"`

	// An IPv4 Address.
	IpAddress *string `json:"ipAddress,omitempty" gorm:"size:255"`

	// The date timestamp of the last reset of the device as reported by the show version command.
	LastReset *DateTime `json:"lastReset,omitempty"`

	// The Cisco Product Family of the hardware.  Values come from MDF for Chassis.
	ProductFamily *string `json:"productFamily,omitempty" gorm:"size:4000"`

	// The Cisco Product ID (PID) of the hardware.
	ProductId *string `json:"productId,omitempty" gorm:"size:32000"`

	// The Cisco Product Type in COLD of the hardware.  Values usually come from MDF.
	ProductType *string `json:"productType,omitempty" gorm:"size:255"`

	// The reason for the last system reset as reported in the show version output.
	ResetReason *string `json:"resetReason,omitempty" gorm:"size:255"`

	// The Type of Software running on the NP Network Element.  Common values include IOS, IOS XR, IOS-XE, NX-OS, etc.
	SwType *string `json:"swType,omitempty" gorm:"size:255"`

	// The Software Version of the device.
	SwVersion *string `json:"swVersion,omitempty" gorm:"size:255"`

	// The SNMP sysContact of the device which is populated in most devices using a configuration command.
	SysContact *string `json:"sysContact,omitempty" gorm:"size:4000"`

	// The SNMP system description from the device.
	SysDescription *string `json:"sysDescription,omitempty" gorm:"size:32000"`

	// The SNMP sysLocation of the device which is populated in most devices using a configuration command.
	SysLocation *string `json:"sysLocation,omitempty" gorm:"size:255"`

	// The SNMP sysObjectID of the device.
	SysObjectId *string `json:"sysObjectId,omitempty" gorm:"size:255"`

	// The user field1 value populated in the collector seedfile.
	UserField1 *string `json:"userField1,omitempty" gorm:"size:4000"`

	// The user field2 value populated in the collector seedfile.
	UserField2 *string `json:"userField2,omitempty" gorm:"size:4000"`

	// The user field3 value populated in the collector seedfile.
	UserField3 *string `json:"userField3,omitempty" gorm:"size:4000"`

	// The user field4 value populated in the collector seedfile.
	UserField4 *string `json:"userField4,omitempty" gorm:"size:4000"`
}

// Asset defines model for Asset.
type Asset struct {
	// Customer ID is not included in the response, but should be added to associate this device to a specific customer
	CustomerID string `gorm:"primaryKey;autoIncrement:false"`

	// The name of the chassis.  This is useful to reference child hardware to its parent chassis in a multi-chassis set-up.
	ChassisName *string `json:"chassisName,omitempty"`

	// The unique ID of the NP NetworkElement/Device.
	DeviceId *int `json:"deviceId,omitempty"`

	// The Network Element Name of the device. When used as input, it can include % as wildcard.
	DeviceName *string `json:"deviceName,omitempty"`

	// The hardware revision.
	HwRev *string `json:"hwRev,omitempty"`

	// The amount of installed flash in the chassis (in megabytes).
	InstalledFlash *int `json:"installedFlash,omitempty"`

	// The amount of installed memory in the chassis (in megabytes).
	InstalledMemory *int `json:"installedMemory,omitempty"`

	// The printed circuit board (PCB) number of the hardware.
	Pcb *string `json:"pcb,omitempty"`

	// The printed circuit board (PCB) number revision of the hardware.
	PcbRev *string `json:"pcbRev,omitempty"`

	// The unique ID in NP for a specific piece of hardware.
	PhysicalElementId *int `json:"physicalElementId,omitempty" gorm:"primaryKey;autoIncrement:false"`

	// For chassis physicalType, this field will indicate which are IP-PHONE, LWAP, or UCSB.
	PhysicalSubtype *string `json:"physicalSubtype,omitempty"`

	// The physical type of the hardware.  Valid values are: Chassis, Module, Power Supply, Fan.
	PhysicalType *string `json:"physicalType,omitempty"`

	// The Cisco Product Family of the hardware.  Values come from MDF for Chassis.
	ProductFamily *string `json:"productFamily,omitempty"`

	// The Cisco Product ID (PID) of the hardware.
	ProductId *string `json:"productId,omitempty" gorm:"size:32000"`

	// The Cisco Product Type in COLD of the hardware.  Values usually come from MDF.
	ProductType *string `json:"productType,omitempty"`

	// The serial number of the hardware.
	SerialNumber *string `json:"serialNumber,omitempty"`

	// The validation status of the serial number of the hardware.  VALID means the SN was found in Cisco MFG or Contract DB. INVALID means the SN was not found in either of those DBs. UNKNOWN means the SN validation has been completed. N/A means the SN is null, so validation is not applicable.
	SerialNumberStatus *string `json:"serialNumberStatus,omitempty"`

	// The slot where a hardware component is located in a chassis.
	Slot *string `json:"slot,omitempty"`

	// The Software Version of the device.
	SwVersion *string `json:"swVersion,omitempty"`

	// The Top Assembly Number (TAN) of the hardware.
	Tan *string `json:"tan,omitempty"`

	// The Top Assembly Number (TAN) Revision of the hardware.
	TanRev *string `json:"tanRev,omitempty"`
}
