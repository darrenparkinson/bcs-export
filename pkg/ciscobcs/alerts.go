package ciscobcs

import "context"

func (c *Client) ListSoftwareEOX(ctx context.Context, options ...*ListOptions) (*BCSListResponse[SoftwareEOX], error) {
	urlStr := c.getURLForSlug("/productAlerts/swEox")
	return ListBCSObject[SoftwareEOX](ctx, c, c.CustomerID, urlStr, options...)
}
func (c *Client) ListSoftwareEOXBulletins(ctx context.Context, options ...*ListOptions) (*BCSListResponse[SoftwareEOXBulletin], error) {
	urlStr := c.getURLForSlug("/productAlerts/swEoxBulletins")
	return ListBCSObject[SoftwareEOXBulletin](ctx, c, c.CustomerID, urlStr, options...)
}
func (c *Client) ListHardwareEOX(ctx context.Context, options ...*ListOptions) (*BCSListResponse[HardwareEOX], error) {
	urlStr := c.getURLForSlug("/productAlerts/hwEox")
	return ListBCSObject[HardwareEOX](ctx, c, c.CustomerID, urlStr, options...)
}
func (c *Client) ListHardwareEOXBulletins(ctx context.Context, options ...*ListOptions) (*BCSListResponse[HardwareEOXBulletin], error) {
	urlStr := c.getURLForSlug("/productAlerts/hwEoxBulletins")
	return ListBCSObject[HardwareEOXBulletin](ctx, c, c.CustomerID, urlStr, options...)
}
func (c *Client) ListFieldNotices(ctx context.Context, options ...*ListOptions) (*BCSListResponse[FieldNotice], error) {
	urlStr := c.getURLForSlug("/productAlerts/fieldNotices")
	return ListBCSObject[FieldNotice](ctx, c, c.CustomerID, urlStr, options...)
}
func (c *Client) ListFieldNoticeNBulletins(ctx context.Context, options ...*ListOptions) (*BCSListResponse[FieldNoticeNBulletin], error) {
	urlStr := c.getURLForSlug("/productAlerts/fnBulletins")
	return ListBCSObject[FieldNoticeNBulletin](ctx, c, c.CustomerID, urlStr, options...)
}
func (c *Client) ListSecurityAdvisories(ctx context.Context, options ...*ListOptions) (*BCSListResponse[SecurityAdvisory], error) {
	urlStr := c.getURLForSlug("/productAlerts/psirt")
	return ListBCSObject[SecurityAdvisory](ctx, c, c.CustomerID, urlStr, options...)
}
func (c *Client) ListSecurityAdvisoryBulletins(ctx context.Context, options ...*ListOptions) (*BCSListResponse[SecurityAdvisoryBulletin], error) {
	urlStr := c.getURLForSlug("/productAlerts/psirtBulletins")
	return ListBCSObject[SecurityAdvisoryBulletin](ctx, c, c.CustomerID, urlStr, options...)
}
func (c *Client) ListSoftwareAlerts(ctx context.Context, options ...*ListOptions) (*BCSListResponse[SoftwareAlert], error) {
	urlStr := c.getURLForSlug("/productAlerts/swAlerts")
	return ListBCSObject[SoftwareAlert](ctx, c, c.CustomerID, urlStr, options...)
}

func (c *Client) ListAllSoftwareEOX(ctx context.Context) ([]*SoftwareEOX, error) {
	urlStr := c.getURLForSlug("/productAlerts/swEox")
	return ListAllBCSObjects[SoftwareEOX](ctx, c, c.CustomerID, urlStr)
}
func (c *Client) ListAllSoftwareEOXBulletins(ctx context.Context) ([]*SoftwareEOXBulletin, error) {
	urlStr := c.getURLForSlug("/productAlerts/swEoxBulletins")
	return ListAllBCSObjects[SoftwareEOXBulletin](ctx, c, c.CustomerID, urlStr)
}
func (c *Client) ListAllHardwareEOX(ctx context.Context) ([]*HardwareEOX, error) {
	urlStr := c.getURLForSlug("/productAlerts/hwEox")
	return ListAllBCSObjects[HardwareEOX](ctx, c, c.CustomerID, urlStr)
}
func (c *Client) ListAllHardwareEOXBulletins(ctx context.Context) ([]*HardwareEOXBulletin, error) {
	urlStr := c.getURLForSlug("/productAlerts/hwEoxBulletins")
	return ListAllBCSObjects[HardwareEOXBulletin](ctx, c, c.CustomerID, urlStr)
}
func (c *Client) ListAllFieldNotices(ctx context.Context) ([]*FieldNotice, error) {
	urlStr := c.getURLForSlug("/productAlerts/fieldNotices")
	return ListAllBCSObjects[FieldNotice](ctx, c, c.CustomerID, urlStr)
}
func (c *Client) ListAllFieldNoticeNBulletins(ctx context.Context) ([]*FieldNoticeNBulletin, error) {
	urlStr := c.getURLForSlug("/productAlerts/fnBulletins")
	return ListAllBCSObjects[FieldNoticeNBulletin](ctx, c, c.CustomerID, urlStr)
}
func (c *Client) ListAllSecurityAdvisories(ctx context.Context) ([]*SecurityAdvisory, error) {
	urlStr := c.getURLForSlug("/productAlerts/psirt")
	return ListAllBCSObjects[SecurityAdvisory](ctx, c, c.CustomerID, urlStr)
}
func (c *Client) ListAllSecurityAdvisoryBulletins(ctx context.Context) ([]*SecurityAdvisoryBulletin, error) {
	urlStr := c.getURLForSlug("/productAlerts/psirtBulletins")
	return ListAllBCSObjects[SecurityAdvisoryBulletin](ctx, c, c.CustomerID, urlStr)
}
func (c *Client) ListAllSoftwareAlerts(ctx context.Context) ([]*SoftwareAlert, error) {
	urlStr := c.getURLForSlug("/productAlerts/swAlerts")
	return ListAllBCSObjects[SoftwareAlert](ctx, c, c.CustomerID, urlStr)
}

// SoftwareEOX defines model for softwareEOX.
type SoftwareEOX struct {
	// The current end-of-life milestone as calculated during last NP profile. If more than one milestone falls on the same date, the returned value will be a comma-separated list.
	CurrentEoxMilestone *string `json:"currentEoxMilestone,omitempty"`

	// The date associated with the current end-of-life milestone.
	CurrentEoxMilestoneDate *DateTime `json:"currentEoxMilestoneDate,omitempty"`

	// The unique ID of the NP NetworkElement/Device.
	DeviceId *int `json:"deviceId,omitempty" gorm:"primaryKey;autoIncrement:false"`

	// The Network Element Name of the device. When used as input, it can include % as wildcard.
	DeviceName *string `json:"deviceName,omitempty"`

	// The next end-of-life milestone that is coming up as calculated during last NP profile.  If more than one milestone falls on the same date, the returned value will be a comma-separated list.  If the device is already LDoS it will not have an next milestone.
	NextEoxMilestone *string `json:"nextEoxMilestone,omitempty"`

	// The date associated with the next end-of-life milestone.
	NextEoxMilestoneDate *DateTime `json:"nextEoxMilestoneDate,omitempty"`

	// Internal software end-of-life identifier to allow join with master sw_eox_bulletins API.
	SwEoxId *int `json:"swEoxId,omitempty" gorm:"primaryKey;autoIncrement:false"`

	// The Type of Software running on the NP Network Element.  Common values include IOS, IOS XR, IOS-XE, NX-OS, etc.
	SwType *string `json:"swType,omitempty"`

	// The Software Version of the device.
	SwVersion *string `json:"swVersion,omitempty"`
}

// SoftwareEOXBulletin defines model for SWEOXBulletins.
type SoftwareEOXBulletin struct {
	// The Cisco.com bulletin number for an End-of-Life bulletin or Field Notice.
	BulletinNumber *string `json:"bulletinNumber,omitempty"`

	// The Cisco.com Title/Headline for the bulletin.
	BulletinTitle *string `json:"bulletinTitle,omitempty"`

	// The Cisco.com URL for the bulletin.
	BulletinUrl *string `json:"bulletinUrl,omitempty"`

	// The End-of-Life Announcement (Announced) Date.
	EoLifeAnnouncementDate *DateTime `json:"eoLifeAnnouncementDate,omitempty"`

	// The End-of-Sale (EoSale) Date.
	EoSaleDate *DateTime `json:"eoSaleDate,omitempty"`

	// The End of Vulnerability/Security Support (EoVSS) Date.
	EoSecurityVulSupportDate *DateTime `json:"eoSecurityVulSupportDate,omitempty"`

	// The End of SW Maintenance Releases (EoSWM) Date.
	EoSwMaintenanceReleasesDate *DateTime `json:"eoSwMaintenanceReleasesDate,omitempty"`

	// The Last Date of Support (LDoS).
	LastDateOfSupport *DateTime `json:"lastDateOfSupport,omitempty"`

	// Internal software end-of-life identifier to allow join with master sw_eox_bulletins API.
	SwEoxId *int `json:"swEoxId,omitempty" gorm:"primaryKey;autoIncrement:false"`

	// The maintenance version portion of the software version.  For example, in 12.4(21), it is "21"
	SwMaintenanceVersion *string `json:"swMaintenanceVersion,omitempty"`

	// The major version portion of the software version.
	SwMajorVersion *string `json:"swMajorVersion,omitempty"`

	// The Software Train, typically only applies to IOS.
	SwTrain *string `json:"swTrain,omitempty"`

	// The Type of Software running on the NP Network Element.  Common values include IOS, IOS XR, IOS-XE, NX-OS, etc.
	SwType *string `json:"swType,omitempty"`
}

// HardwareEOX defines model for HWEOX.
type HardwareEOX struct {
	// The current end-of-life milestone as calculated during last NP profile. If more than one milestone falls on the same date, the returned value will be a comma-separated list.
	CurrentEoxMilestone *string `json:"currentEoxMilestone,omitempty"`

	// The date associated with the current end-of-life milestone.
	CurrentEoxMilestoneDate *DateTime `json:"currentEoxMilestoneDate,omitempty"`

	// The unique ID of the NP NetworkElement/Device.
	DeviceId *int `json:"deviceId,omitempty"`

	// The Network Element Name of the device. When used as input, it can include % as wildcard.
	DeviceName *string `json:"deviceName,omitempty"`

	// Internal hardware end-of-life identifier to allow join with master hw_eox_bulletins API.
	HwEoxId *int `json:"hwEoxId,omitempty" gorm:"primaryKey;autoIncrement:false"`

	// The next end-of-life milestone that is coming up as calculated during last NP profile.  If more than one milestone falls on the same date, the returned value will be a comma-separated list.  If the device is already LDoS it will not have an next milestone.
	NextEoxMilestone *string `json:"nextEoxMilestone,omitempty"`

	// The date associated with the next end-of-life milestone.
	NextEoxMilestoneDate *DateTime `json:"nextEoxMilestoneDate,omitempty"`

	// The unique ID in NP for a specific piece of hardware.
	PhysicalElementId *int `json:"physicalElementId,omitempty" gorm:"primaryKey;autoIncrement:false"`

	// The physical type of the hardware.  Valid values are: Chassis, Module, Power Supply, Fan.
	PhysicalType *string `json:"physicalType,omitempty"`

	// The Cisco Product ID (PID) of the hardware.
	ProductId *string `json:"productId,omitempty"`
}

// HardwareEOXBulletin defines model for HWEOXBulletins.
type HardwareEOXBulletin struct {
	// The Cisco.com bulletin number for an End-of-Life bulletin or Field Notice.
	BulletinNumber *string `json:"bulletinNumber,omitempty"`

	// The Cisco.com Title/Headline for the bulletin.
	BulletinTitle *string `json:"bulletinTitle,omitempty"`

	// The Cisco.com URL for the bulletin.
	BulletinUrl *string `json:"bulletinUrl,omitempty"`

	// The End-of-Life Announcement (Announced) Date.
	EoLifeAnnouncementDate *DateTime `json:"eoLifeAnnouncementDate,omitempty"`

	// The End of New Service Attachment Date.
	EoNewServiceAttachDate *DateTime `json:"eoNewServiceAttachDate,omitempty"`

	// The End of Routine Failure Analysis Date (EoRFA) Date.
	EoRoutineFailureAnalysisDate *DateTime `json:"eoRoutineFailureAnalysisDate,omitempty"`

	// The End-of-Sale (EoSale) Date.
	EoSaleDate *DateTime `json:"eoSaleDate,omitempty"`

	// The End of Vulnerability/Security Support (EoVSS) Date.
	EoSecurityVulSupportDate *DateTime `json:"eoSecurityVulSupportDate,omitempty"`

	// The End of Service Contract Renewal (EoSCR) Date.
	EoSoftwareContractRenewalDate *DateTime `json:"eoSoftwareContractRenewalDate,omitempty"`

	// The End of SW Maintenance Releases (EoSWM) Date.
	EoSwMaintenanceReleasesDate *DateTime `json:"eoSwMaintenanceReleasesDate,omitempty"`

	// Internal hardware end-of-life identifier to allow join with master hw_eox_bulletins API.
	HwEoxId *int `json:"hwEoxId,omitempty" gorm:"primaryKey;autoIncrement:false"`

	// The Last Date of Support (LDoS).
	LastDateOfSupport *DateTime `json:"lastDateOfSupport,omitempty"`

	// The Last Ship Date.
	LastShipDate *DateTime `json:"lastShipDate,omitempty"`

	// The Cisco Product ID (PID) of the hardware.
	ProductId *string `json:"productId,omitempty"`
}

// FieldNotice defines model for fieldNotices.
type FieldNotice struct {
	// The unique ID of the NP NetworkElement/Device.
	DeviceId *int `json:"deviceId,omitempty" `

	// Field Notice ID number.
	FieldNoticeId *string `json:"fieldNoticeId,omitempty" gorm:"primaryKey;autoIncrement:false"`

	// The match confidence result from PAS.  Valid values include: Vulnerable, Potentially Vulnerable, Not Vulnerable.
	MatchConfidence *string `json:"matchConfidence,omitempty"`

	// The reason behind the match confidence result from PAS. Explains why you are vulnerable or not vulnerable or what data is missing to cause a potentially vulnerable result.  PAS value is enhanced in NP for readability.
	MatchConfidenceReason *string `json:"matchConfidenceReason,omitempty"`

	// The unique ID in NP for a specific piece of hardware.
	PhysicalElementId *int `json:"physicalElementId,omitempty" gorm:"primaryKey;autoIncrement:false"`
}

// FieldNoticeBulletin defines model for FNBulletins.
type FieldNoticeNBulletin struct {
	// The date when the bulletin was first published to Cisco.com.  Most API calls will allow Regex input for this field.
	BulletinFirstPublished *string `json:"bulletinFirstPublished,omitempty"`

	// The date when the bulletin was last updated on Cisco.com.
	BulletinLastUpdated *DateTime `json:"bulletinLastUpdated,omitempty"`

	// The Bulletin Mapping Caveat gives any explanations why the automation may need additional review by the customer.
	BulletinMappingCaveat *string `json:"bulletinMappingCaveat,omitempty"`

	// The Cisco.com Title/Headline for the bulletin.
	BulletinTitle *string `json:"bulletinTitle,omitempty"`

	// The Cisco.com URL for the bulletin.
	BulletinUrl *string `json:"bulletinUrl,omitempty"`

	// Field Notice ID number.
	FieldNoticeId *string `json:"fieldNoticeId,omitempty" gorm:"primaryKey;autoIncrement:false"`

	// Type of Field Notice as defined from PLATO.  Valid values include: hardware, software, other
	FnType *string `json:"fnType,omitempty"`

	// The description of the problem on a Cisco bulletin.
	ProblemDescription *string `json:"problemDescription,omitempty"`
}

// SecurityAdvisory defines model for securityAdvisories.
type SecurityAdvisory struct {
	// The unique ID of the NP NetworkElement/Device.
	DeviceId *int `json:"deviceId,omitempty" gorm:"primaryKey;autoIncrement:false"`

	// The match confidence result from PAS.  Valid values include: Vulnerable, Potentially Vulnerable, Not Vulnerable.
	MatchConfidence *string `json:"matchConfidence,omitempty"`

	// The reason behind the match confidence result from PAS. Explains why you are vulnerable or not vulnerable or what data is missing to cause a potentially vulnerable result.  PAS value is enhanced in NP for readability.
	MatchConfidenceReason *string `json:"matchConfidenceReason,omitempty"`

	// The internal COLD ID for a PSIRT. This is useful for joining multiple data sources.
	PsirtColdId *int `json:"psirtColdId,omitempty" gorm:"primaryKey;autoIncrement:false"`
}

// SecurityAdvisoryBulletin defines model for PSIRTBulletins.
type SecurityAdvisoryBulletin struct {
	// The date when the bulletin was first published to Cisco.com.  Most API calls will allow Regex input for this field.
	BulletinFirstPublished *string `json:"bulletinFirstPublished,omitempty"`

	// The date when the bulletin was last updated on Cisco.com.
	BulletinLastUpdated *DateTime `json:"bulletinLastUpdated,omitempty"`

	// The Bulletin Mapping Caveat gives any explanations why the automation may need additional review by the customer.
	BulletinMappingCaveat *string `json:"bulletinMappingCaveat,omitempty"`

	// The Summary of a Cisco.com bulletin.
	BulletinSummary *string `json:"bulletinSummary,omitempty"`

	// The Cisco.com Title/Headline for the bulletin.
	BulletinTitle *string `json:"bulletinTitle,omitempty"`

	// The Cisco.com URL for the bulletin.
	BulletinUrl *string `json:"bulletinUrl,omitempty"`

	// The version # of the Cisco.com bulletin.
	BulletinVersion *string `json:"bulletinVersion,omitempty"`

	// Comma-separated list of Cisco Bug IDs.
	CiscoBugIds *string `json:"ciscoBugIds,omitempty"`

	// Common Vulnerabilities and Exposures (CVE) Identifier
	CveId *string `json:"cveId,omitempty"`

	// Common Vulnerability Scoring System (CVSS) Base Score
	CvssBase *string `json:"cvssBase,omitempty"`

	// Common Vulnerability Scoring System (CVSS) Temporal Score
	CvssTemporal *string `json:"cvssTemporal,omitempty"`

	// The Advisory ID of a PSIRT as seen on Cisco.com.
	PsirtAdvisoryId *string `json:"psirtAdvisoryId,omitempty"`

	// The internal COLD ID for a PSIRT.  This is useful for joining multiple data sources.
	PsirtColdId *int `json:"psirtColdId,omitempty" gorm:"primaryKey;autoIncrement:false"`

	// The Security Impact Rating (SIR) for Cisco PSIRTs.
	Sir *string `json:"sir,omitempty"`
}

// SoftwareAlert defines model for softwareAlerts.
type SoftwareAlert struct {
	// The unique ID of the NP NetworkElement/Device.
	DeviceId *int `json:"deviceId,omitempty" gorm:"primaryKey;autoIncrement:false"`

	// The Network Element Name of the device. When used as input, it can include % as wildcard.
	DeviceName *string `json:"deviceName,omitempty"`

	// The Image Name of the software on the Network Element.
	ImageName *string `json:"imageName,omitempty"`

	// The type of Software Alert on the device. Valid values include SA for Software Advisory and DF for Deferral.
	SwAlertType *string `json:"swAlertType,omitempty"`

	// The Cisco.com URL with details for a specific software advisory or deferral.
	SwAlertUrl *string `json:"swAlertUrl,omitempty" gorm:"primaryKey;autoIncrement:false"`

	// The Type of Software running on the NP Network Element.  Common values include IOS, IOS XR, IOS-XE, NX-OS, etc.
	SwType *string `json:"swType,omitempty"`

	// The Software Version of the device.
	SwVersion *string `json:"swVersion,omitempty"`
}
