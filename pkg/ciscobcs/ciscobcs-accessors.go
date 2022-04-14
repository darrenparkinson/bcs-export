// Copyright 2021 The ciscobcs AUTHORS. All rights reserved.
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.
// Code generated by gen-accessors; DO NOT EDIT.
package ciscobcs

// GetChassisName returns the ChassisName field if it's non-nil, zero value otherwise.
func (a *Asset) GetChassisName() string {
	if a == nil || a.ChassisName == nil {
		return ""
	}
	return *a.ChassisName
}

// GetDeviceId returns the DeviceId field if it's non-nil, zero value otherwise.
func (a *Asset) GetDeviceId() int {
	if a == nil || a.DeviceId == nil {
		return 0
	}
	return *a.DeviceId
}

// GetDeviceName returns the DeviceName field if it's non-nil, zero value otherwise.
func (a *Asset) GetDeviceName() string {
	if a == nil || a.DeviceName == nil {
		return ""
	}
	return *a.DeviceName
}

// GetHwRev returns the HwRev field if it's non-nil, zero value otherwise.
func (a *Asset) GetHwRev() string {
	if a == nil || a.HwRev == nil {
		return ""
	}
	return *a.HwRev
}

// GetInstalledFlash returns the InstalledFlash field if it's non-nil, zero value otherwise.
func (a *Asset) GetInstalledFlash() int {
	if a == nil || a.InstalledFlash == nil {
		return 0
	}
	return *a.InstalledFlash
}

// GetInstalledMemory returns the InstalledMemory field if it's non-nil, zero value otherwise.
func (a *Asset) GetInstalledMemory() int {
	if a == nil || a.InstalledMemory == nil {
		return 0
	}
	return *a.InstalledMemory
}

// GetPcb returns the Pcb field if it's non-nil, zero value otherwise.
func (a *Asset) GetPcb() string {
	if a == nil || a.Pcb == nil {
		return ""
	}
	return *a.Pcb
}

// GetPcbRev returns the PcbRev field if it's non-nil, zero value otherwise.
func (a *Asset) GetPcbRev() string {
	if a == nil || a.PcbRev == nil {
		return ""
	}
	return *a.PcbRev
}

// GetPhysicalElementId returns the PhysicalElementId field if it's non-nil, zero value otherwise.
func (a *Asset) GetPhysicalElementId() int {
	if a == nil || a.PhysicalElementId == nil {
		return 0
	}
	return *a.PhysicalElementId
}

// GetPhysicalSubtype returns the PhysicalSubtype field if it's non-nil, zero value otherwise.
func (a *Asset) GetPhysicalSubtype() string {
	if a == nil || a.PhysicalSubtype == nil {
		return ""
	}
	return *a.PhysicalSubtype
}

// GetPhysicalType returns the PhysicalType field if it's non-nil, zero value otherwise.
func (a *Asset) GetPhysicalType() string {
	if a == nil || a.PhysicalType == nil {
		return ""
	}
	return *a.PhysicalType
}

// GetProductFamily returns the ProductFamily field if it's non-nil, zero value otherwise.
func (a *Asset) GetProductFamily() string {
	if a == nil || a.ProductFamily == nil {
		return ""
	}
	return *a.ProductFamily
}

// GetProductId returns the ProductId field if it's non-nil, zero value otherwise.
func (a *Asset) GetProductId() string {
	if a == nil || a.ProductId == nil {
		return ""
	}
	return *a.ProductId
}

// GetProductType returns the ProductType field if it's non-nil, zero value otherwise.
func (a *Asset) GetProductType() string {
	if a == nil || a.ProductType == nil {
		return ""
	}
	return *a.ProductType
}

// GetSerialNumber returns the SerialNumber field if it's non-nil, zero value otherwise.
func (a *Asset) GetSerialNumber() string {
	if a == nil || a.SerialNumber == nil {
		return ""
	}
	return *a.SerialNumber
}

// GetSerialNumberStatus returns the SerialNumberStatus field if it's non-nil, zero value otherwise.
func (a *Asset) GetSerialNumberStatus() string {
	if a == nil || a.SerialNumberStatus == nil {
		return ""
	}
	return *a.SerialNumberStatus
}

// GetSlot returns the Slot field if it's non-nil, zero value otherwise.
func (a *Asset) GetSlot() string {
	if a == nil || a.Slot == nil {
		return ""
	}
	return *a.Slot
}

// GetSwVersion returns the SwVersion field if it's non-nil, zero value otherwise.
func (a *Asset) GetSwVersion() string {
	if a == nil || a.SwVersion == nil {
		return ""
	}
	return *a.SwVersion
}

// GetTan returns the Tan field if it's non-nil, zero value otherwise.
func (a *Asset) GetTan() string {
	if a == nil || a.Tan == nil {
		return ""
	}
	return *a.Tan
}

// GetTanRev returns the TanRev field if it's non-nil, zero value otherwise.
func (a *Asset) GetTanRev() string {
	if a == nil || a.TanRev == nil {
		return ""
	}
	return *a.TanRev
}

// GetBpNuggetId returns the BpNuggetId field if it's non-nil, zero value otherwise.
func (c *CBPDetail) GetBpNuggetId() int {
	if c == nil || c.BpNuggetId == nil {
		return 0
	}
	return *c.BpNuggetId
}

// GetBpRuleId returns the BpRuleId field if it's non-nil, zero value otherwise.
func (c *CBPDetail) GetBpRuleId() int {
	if c == nil || c.BpRuleId == nil {
		return 0
	}
	return *c.BpRuleId
}

// GetConfigSource returns the ConfigSource field if it's non-nil, zero value otherwise.
func (c *CBPDetail) GetConfigSource() string {
	if c == nil || c.ConfigSource == nil {
		return ""
	}
	return *c.ConfigSource
}

// GetDeviceId returns the DeviceId field if it's non-nil, zero value otherwise.
func (c *CBPDetail) GetDeviceId() int {
	if c == nil || c.DeviceId == nil {
		return 0
	}
	return *c.DeviceId
}

// GetBpCaveat returns the BpCaveat field if it's non-nil, zero value otherwise.
func (c *CBPRule) GetBpCaveat() string {
	if c == nil || c.BpCaveat == nil {
		return ""
	}
	return *c.BpCaveat
}

// GetBpCorrectiveAction returns the BpCorrectiveAction field if it's non-nil, zero value otherwise.
func (c *CBPRule) GetBpCorrectiveAction() string {
	if c == nil || c.BpCorrectiveAction == nil {
		return ""
	}
	return *c.BpCorrectiveAction
}

// GetBpDescription returns the BpDescription field if it's non-nil, zero value otherwise.
func (c *CBPRule) GetBpDescription() string {
	if c == nil || c.BpDescription == nil {
		return ""
	}
	return *c.BpDescription
}

// GetBpNuggetId returns the BpNuggetId field if it's non-nil, zero value otherwise.
func (c *CBPRule) GetBpNuggetId() int {
	if c == nil || c.BpNuggetId == nil {
		return 0
	}
	return *c.BpNuggetId
}

// GetBpPrimaryTechnology returns the BpPrimaryTechnology field if it's non-nil, zero value otherwise.
func (c *CBPRule) GetBpPrimaryTechnology() string {
	if c == nil || c.BpPrimaryTechnology == nil {
		return ""
	}
	return *c.BpPrimaryTechnology
}

// GetBpRecommendation returns the BpRecommendation field if it's non-nil, zero value otherwise.
func (c *CBPRule) GetBpRecommendation() string {
	if c == nil || c.BpRecommendation == nil {
		return ""
	}
	return *c.BpRecommendation
}

// GetBpRisk returns the BpRisk field if it's non-nil, zero value otherwise.
func (c *CBPRule) GetBpRisk() string {
	if c == nil || c.BpRisk == nil {
		return ""
	}
	return *c.BpRisk
}

// GetBpRuleId returns the BpRuleId field if it's non-nil, zero value otherwise.
func (c *CBPRule) GetBpRuleId() int {
	if c == nil || c.BpRuleId == nil {
		return 0
	}
	return *c.BpRuleId
}

// GetBpSecondaryTechnology returns the BpSecondaryTechnology field if it's non-nil, zero value otherwise.
func (c *CBPRule) GetBpSecondaryTechnology() string {
	if c == nil || c.BpSecondaryTechnology == nil {
		return ""
	}
	return *c.BpSecondaryTechnology
}

// GetBpTitle returns the BpTitle field if it's non-nil, zero value otherwise.
func (c *CBPRule) GetBpTitle() string {
	if c == nil || c.BpTitle == nil {
		return ""
	}
	return *c.BpTitle
}

// GetCreateDate returns the CreateDate field.
func (c *CBPRule) GetCreateDate() *DateTime {
	if c == nil {
		return nil
	}
	return c.CreateDate
}

// GetSwType returns the SwType field if it's non-nil, zero value otherwise.
func (c *CBPRule) GetSwType() string {
	if c == nil || c.SwType == nil {
		return ""
	}
	return *c.SwType
}

// GetUpdateDate returns the UpdateDate field.
func (c *CBPRule) GetUpdateDate() *DateTime {
	if c == nil {
		return nil
	}
	return c.UpdateDate
}

// GetBpRuleId returns the BpRuleId field if it's non-nil, zero value otherwise.
func (c *CBPRulesReference) GetBpRuleId() int {
	if c == nil || c.BpRuleId == nil {
		return 0
	}
	return *c.BpRuleId
}

// GetBpUrl returns the BpUrl field if it's non-nil, zero value otherwise.
func (c *CBPRulesReference) GetBpUrl() string {
	if c == nil || c.BpUrl == nil {
		return ""
	}
	return *c.BpUrl
}

// GetBpUrlTitle returns the BpUrlTitle field if it's non-nil, zero value otherwise.
func (c *CBPRulesReference) GetBpUrlTitle() string {
	if c == nil || c.BpUrlTitle == nil {
		return ""
	}
	return *c.BpUrlTitle
}

// GetBpNuggetId returns the BpNuggetId field if it's non-nil, zero value otherwise.
func (c *CBPSummary) GetBpNuggetId() int {
	if c == nil || c.BpNuggetId == nil {
		return 0
	}
	return *c.BpNuggetId
}

// GetBpPrimaryTechnology returns the BpPrimaryTechnology field if it's non-nil, zero value otherwise.
func (c *CBPSummary) GetBpPrimaryTechnology() string {
	if c == nil || c.BpPrimaryTechnology == nil {
		return ""
	}
	return *c.BpPrimaryTechnology
}

// GetBpRisk returns the BpRisk field if it's non-nil, zero value otherwise.
func (c *CBPSummary) GetBpRisk() string {
	if c == nil || c.BpRisk == nil {
		return ""
	}
	return *c.BpRisk
}

// GetBpRuleId returns the BpRuleId field if it's non-nil, zero value otherwise.
func (c *CBPSummary) GetBpRuleId() int {
	if c == nil || c.BpRuleId == nil {
		return 0
	}
	return *c.BpRuleId
}

// GetBpSecondaryTechnology returns the BpSecondaryTechnology field if it's non-nil, zero value otherwise.
func (c *CBPSummary) GetBpSecondaryTechnology() string {
	if c == nil || c.BpSecondaryTechnology == nil {
		return ""
	}
	return *c.BpSecondaryTechnology
}

// GetBpTitle returns the BpTitle field if it's non-nil, zero value otherwise.
func (c *CBPSummary) GetBpTitle() string {
	if c == nil || c.BpTitle == nil {
		return ""
	}
	return *c.BpTitle
}

// GetSwType returns the SwType field if it's non-nil, zero value otherwise.
func (c *CBPSummary) GetSwType() string {
	if c == nil || c.SwType == nil {
		return ""
	}
	return *c.SwType
}

// GetTotalDevices returns the TotalDevices field if it's non-nil, zero value otherwise.
func (c *CBPSummary) GetTotalDevices() int {
	if c == nil || c.TotalDevices == nil {
		return 0
	}
	return *c.TotalDevices
}

// GetCollector returns the Collector field if it's non-nil, zero value otherwise.
func (d *Device) GetCollector() string {
	if d == nil || d.Collector == nil {
		return ""
	}
	return *d.Collector
}

// GetConfigRegister returns the ConfigRegister field if it's non-nil, zero value otherwise.
func (d *Device) GetConfigRegister() string {
	if d == nil || d.ConfigRegister == nil {
		return ""
	}
	return *d.ConfigRegister
}

// GetConfigStatus returns the ConfigStatus field if it's non-nil, zero value otherwise.
func (d *Device) GetConfigStatus() string {
	if d == nil || d.ConfigStatus == nil {
		return ""
	}
	return *d.ConfigStatus
}

// GetConfigTime returns the ConfigTime field.
func (d *Device) GetConfigTime() *DateTime {
	if d == nil {
		return nil
	}
	return d.ConfigTime
}

// GetCreateDate returns the CreateDate field.
func (d *Device) GetCreateDate() *DateTime {
	if d == nil {
		return nil
	}
	return d.CreateDate
}

// GetDeviceId returns the DeviceId field if it's non-nil, zero value otherwise.
func (d *Device) GetDeviceId() int {
	if d == nil || d.DeviceId == nil {
		return 0
	}
	return *d.DeviceId
}

// GetDeviceIp returns the DeviceIp field if it's non-nil, zero value otherwise.
func (d *Device) GetDeviceIp() string {
	if d == nil || d.DeviceIp == nil {
		return ""
	}
	return *d.DeviceIp
}

// GetDeviceName returns the DeviceName field if it's non-nil, zero value otherwise.
func (d *Device) GetDeviceName() string {
	if d == nil || d.DeviceName == nil {
		return ""
	}
	return *d.DeviceName
}

// GetDeviceStatus returns the DeviceStatus field if it's non-nil, zero value otherwise.
func (d *Device) GetDeviceStatus() string {
	if d == nil || d.DeviceStatus == nil {
		return ""
	}
	return *d.DeviceStatus
}

// GetDeviceSysName returns the DeviceSysName field if it's non-nil, zero value otherwise.
func (d *Device) GetDeviceSysName() string {
	if d == nil || d.DeviceSysName == nil {
		return ""
	}
	return *d.DeviceSysName
}

// GetDeviceType returns the DeviceType field if it's non-nil, zero value otherwise.
func (d *Device) GetDeviceType() string {
	if d == nil || d.DeviceType == nil {
		return ""
	}
	return *d.DeviceType
}

// GetFeatureSetdesc returns the FeatureSetdesc field if it's non-nil, zero value otherwise.
func (d *Device) GetFeatureSetdesc() string {
	if d == nil || d.FeatureSetdesc == nil {
		return ""
	}
	return *d.FeatureSetdesc
}

// GetImageName returns the ImageName field if it's non-nil, zero value otherwise.
func (d *Device) GetImageName() string {
	if d == nil || d.ImageName == nil {
		return ""
	}
	return *d.ImageName
}

// GetInSeedFile returns the InSeedFile field if it's non-nil, zero value otherwise.
func (d *Device) GetInSeedFile() bool {
	if d == nil || d.InSeedFile == nil {
		return false
	}
	return *d.InSeedFile
}

// GetInventoryStatus returns the InventoryStatus field if it's non-nil, zero value otherwise.
func (d *Device) GetInventoryStatus() string {
	if d == nil || d.InventoryStatus == nil {
		return ""
	}
	return *d.InventoryStatus
}

// GetInventoryTime returns the InventoryTime field.
func (d *Device) GetInventoryTime() *DateTime {
	if d == nil {
		return nil
	}
	return d.InventoryTime
}

// GetIpAddress returns the IpAddress field if it's non-nil, zero value otherwise.
func (d *Device) GetIpAddress() string {
	if d == nil || d.IpAddress == nil {
		return ""
	}
	return *d.IpAddress
}

// GetLastReset returns the LastReset field.
func (d *Device) GetLastReset() *DateTime {
	if d == nil {
		return nil
	}
	return d.LastReset
}

// GetProductFamily returns the ProductFamily field if it's non-nil, zero value otherwise.
func (d *Device) GetProductFamily() string {
	if d == nil || d.ProductFamily == nil {
		return ""
	}
	return *d.ProductFamily
}

// GetProductId returns the ProductId field if it's non-nil, zero value otherwise.
func (d *Device) GetProductId() string {
	if d == nil || d.ProductId == nil {
		return ""
	}
	return *d.ProductId
}

// GetProductType returns the ProductType field if it's non-nil, zero value otherwise.
func (d *Device) GetProductType() string {
	if d == nil || d.ProductType == nil {
		return ""
	}
	return *d.ProductType
}

// GetResetReason returns the ResetReason field if it's non-nil, zero value otherwise.
func (d *Device) GetResetReason() string {
	if d == nil || d.ResetReason == nil {
		return ""
	}
	return *d.ResetReason
}

// GetSwType returns the SwType field if it's non-nil, zero value otherwise.
func (d *Device) GetSwType() string {
	if d == nil || d.SwType == nil {
		return ""
	}
	return *d.SwType
}

// GetSwVersion returns the SwVersion field if it's non-nil, zero value otherwise.
func (d *Device) GetSwVersion() string {
	if d == nil || d.SwVersion == nil {
		return ""
	}
	return *d.SwVersion
}

// GetSysContact returns the SysContact field if it's non-nil, zero value otherwise.
func (d *Device) GetSysContact() string {
	if d == nil || d.SysContact == nil {
		return ""
	}
	return *d.SysContact
}

// GetSysDescription returns the SysDescription field if it's non-nil, zero value otherwise.
func (d *Device) GetSysDescription() string {
	if d == nil || d.SysDescription == nil {
		return ""
	}
	return *d.SysDescription
}

// GetSysLocation returns the SysLocation field if it's non-nil, zero value otherwise.
func (d *Device) GetSysLocation() string {
	if d == nil || d.SysLocation == nil {
		return ""
	}
	return *d.SysLocation
}

// GetSysObjectId returns the SysObjectId field if it's non-nil, zero value otherwise.
func (d *Device) GetSysObjectId() string {
	if d == nil || d.SysObjectId == nil {
		return ""
	}
	return *d.SysObjectId
}

// GetUserField1 returns the UserField1 field if it's non-nil, zero value otherwise.
func (d *Device) GetUserField1() string {
	if d == nil || d.UserField1 == nil {
		return ""
	}
	return *d.UserField1
}

// GetUserField2 returns the UserField2 field if it's non-nil, zero value otherwise.
func (d *Device) GetUserField2() string {
	if d == nil || d.UserField2 == nil {
		return ""
	}
	return *d.UserField2
}

// GetUserField3 returns the UserField3 field if it's non-nil, zero value otherwise.
func (d *Device) GetUserField3() string {
	if d == nil || d.UserField3 == nil {
		return ""
	}
	return *d.UserField3
}

// GetUserField4 returns the UserField4 field if it's non-nil, zero value otherwise.
func (d *Device) GetUserField4() string {
	if d == nil || d.UserField4 == nil {
		return ""
	}
	return *d.UserField4
}

// GetDeviceId returns the DeviceId field if it's non-nil, zero value otherwise.
func (f *FieldNotice) GetDeviceId() int {
	if f == nil || f.DeviceId == nil {
		return 0
	}
	return *f.DeviceId
}

// GetFieldNoticeId returns the FieldNoticeId field if it's non-nil, zero value otherwise.
func (f *FieldNotice) GetFieldNoticeId() string {
	if f == nil || f.FieldNoticeId == nil {
		return ""
	}
	return *f.FieldNoticeId
}

// GetMatchConfidence returns the MatchConfidence field if it's non-nil, zero value otherwise.
func (f *FieldNotice) GetMatchConfidence() string {
	if f == nil || f.MatchConfidence == nil {
		return ""
	}
	return *f.MatchConfidence
}

// GetMatchConfidenceReason returns the MatchConfidenceReason field if it's non-nil, zero value otherwise.
func (f *FieldNotice) GetMatchConfidenceReason() string {
	if f == nil || f.MatchConfidenceReason == nil {
		return ""
	}
	return *f.MatchConfidenceReason
}

// GetPhysicalElementId returns the PhysicalElementId field if it's non-nil, zero value otherwise.
func (f *FieldNotice) GetPhysicalElementId() int {
	if f == nil || f.PhysicalElementId == nil {
		return 0
	}
	return *f.PhysicalElementId
}

// GetBulletinFirstPublished returns the BulletinFirstPublished field if it's non-nil, zero value otherwise.
func (f *FieldNoticeNBulletin) GetBulletinFirstPublished() string {
	if f == nil || f.BulletinFirstPublished == nil {
		return ""
	}
	return *f.BulletinFirstPublished
}

// GetBulletinLastUpdated returns the BulletinLastUpdated field.
func (f *FieldNoticeNBulletin) GetBulletinLastUpdated() *DateTime {
	if f == nil {
		return nil
	}
	return f.BulletinLastUpdated
}

// GetBulletinMappingCaveat returns the BulletinMappingCaveat field if it's non-nil, zero value otherwise.
func (f *FieldNoticeNBulletin) GetBulletinMappingCaveat() string {
	if f == nil || f.BulletinMappingCaveat == nil {
		return ""
	}
	return *f.BulletinMappingCaveat
}

// GetBulletinTitle returns the BulletinTitle field if it's non-nil, zero value otherwise.
func (f *FieldNoticeNBulletin) GetBulletinTitle() string {
	if f == nil || f.BulletinTitle == nil {
		return ""
	}
	return *f.BulletinTitle
}

// GetBulletinUrl returns the BulletinUrl field if it's non-nil, zero value otherwise.
func (f *FieldNoticeNBulletin) GetBulletinUrl() string {
	if f == nil || f.BulletinUrl == nil {
		return ""
	}
	return *f.BulletinUrl
}

// GetFieldNoticeId returns the FieldNoticeId field if it's non-nil, zero value otherwise.
func (f *FieldNoticeNBulletin) GetFieldNoticeId() string {
	if f == nil || f.FieldNoticeId == nil {
		return ""
	}
	return *f.FieldNoticeId
}

// GetFnType returns the FnType field if it's non-nil, zero value otherwise.
func (f *FieldNoticeNBulletin) GetFnType() string {
	if f == nil || f.FnType == nil {
		return ""
	}
	return *f.FnType
}

// GetProblemDescription returns the ProblemDescription field if it's non-nil, zero value otherwise.
func (f *FieldNoticeNBulletin) GetProblemDescription() string {
	if f == nil || f.ProblemDescription == nil {
		return ""
	}
	return *f.ProblemDescription
}

// GetCurrentEoxMilestone returns the CurrentEoxMilestone field if it's non-nil, zero value otherwise.
func (h *HardwareEOX) GetCurrentEoxMilestone() string {
	if h == nil || h.CurrentEoxMilestone == nil {
		return ""
	}
	return *h.CurrentEoxMilestone
}

// GetCurrentEoxMilestoneDate returns the CurrentEoxMilestoneDate field.
func (h *HardwareEOX) GetCurrentEoxMilestoneDate() *DateTime {
	if h == nil {
		return nil
	}
	return h.CurrentEoxMilestoneDate
}

// GetDeviceId returns the DeviceId field if it's non-nil, zero value otherwise.
func (h *HardwareEOX) GetDeviceId() int {
	if h == nil || h.DeviceId == nil {
		return 0
	}
	return *h.DeviceId
}

// GetDeviceName returns the DeviceName field if it's non-nil, zero value otherwise.
func (h *HardwareEOX) GetDeviceName() string {
	if h == nil || h.DeviceName == nil {
		return ""
	}
	return *h.DeviceName
}

// GetHwEoxId returns the HwEoxId field if it's non-nil, zero value otherwise.
func (h *HardwareEOX) GetHwEoxId() int {
	if h == nil || h.HwEoxId == nil {
		return 0
	}
	return *h.HwEoxId
}

// GetNextEoxMilestone returns the NextEoxMilestone field if it's non-nil, zero value otherwise.
func (h *HardwareEOX) GetNextEoxMilestone() string {
	if h == nil || h.NextEoxMilestone == nil {
		return ""
	}
	return *h.NextEoxMilestone
}

// GetNextEoxMilestoneDate returns the NextEoxMilestoneDate field.
func (h *HardwareEOX) GetNextEoxMilestoneDate() *DateTime {
	if h == nil {
		return nil
	}
	return h.NextEoxMilestoneDate
}

// GetPhysicalElementId returns the PhysicalElementId field if it's non-nil, zero value otherwise.
func (h *HardwareEOX) GetPhysicalElementId() int {
	if h == nil || h.PhysicalElementId == nil {
		return 0
	}
	return *h.PhysicalElementId
}

// GetPhysicalType returns the PhysicalType field if it's non-nil, zero value otherwise.
func (h *HardwareEOX) GetPhysicalType() string {
	if h == nil || h.PhysicalType == nil {
		return ""
	}
	return *h.PhysicalType
}

// GetProductId returns the ProductId field if it's non-nil, zero value otherwise.
func (h *HardwareEOX) GetProductId() string {
	if h == nil || h.ProductId == nil {
		return ""
	}
	return *h.ProductId
}

// GetBulletinNumber returns the BulletinNumber field if it's non-nil, zero value otherwise.
func (h *HardwareEOXBulletin) GetBulletinNumber() string {
	if h == nil || h.BulletinNumber == nil {
		return ""
	}
	return *h.BulletinNumber
}

// GetBulletinTitle returns the BulletinTitle field if it's non-nil, zero value otherwise.
func (h *HardwareEOXBulletin) GetBulletinTitle() string {
	if h == nil || h.BulletinTitle == nil {
		return ""
	}
	return *h.BulletinTitle
}

// GetBulletinUrl returns the BulletinUrl field if it's non-nil, zero value otherwise.
func (h *HardwareEOXBulletin) GetBulletinUrl() string {
	if h == nil || h.BulletinUrl == nil {
		return ""
	}
	return *h.BulletinUrl
}

// GetEoLifeAnnouncementDate returns the EoLifeAnnouncementDate field.
func (h *HardwareEOXBulletin) GetEoLifeAnnouncementDate() *DateTime {
	if h == nil {
		return nil
	}
	return h.EoLifeAnnouncementDate
}

// GetEoNewServiceAttachDate returns the EoNewServiceAttachDate field.
func (h *HardwareEOXBulletin) GetEoNewServiceAttachDate() *DateTime {
	if h == nil {
		return nil
	}
	return h.EoNewServiceAttachDate
}

// GetEoRoutineFailureAnalysisDate returns the EoRoutineFailureAnalysisDate field.
func (h *HardwareEOXBulletin) GetEoRoutineFailureAnalysisDate() *DateTime {
	if h == nil {
		return nil
	}
	return h.EoRoutineFailureAnalysisDate
}

// GetEoSaleDate returns the EoSaleDate field.
func (h *HardwareEOXBulletin) GetEoSaleDate() *DateTime {
	if h == nil {
		return nil
	}
	return h.EoSaleDate
}

// GetEoSecurityVulSupportDate returns the EoSecurityVulSupportDate field.
func (h *HardwareEOXBulletin) GetEoSecurityVulSupportDate() *DateTime {
	if h == nil {
		return nil
	}
	return h.EoSecurityVulSupportDate
}

// GetEoSoftwareContractRenewalDate returns the EoSoftwareContractRenewalDate field.
func (h *HardwareEOXBulletin) GetEoSoftwareContractRenewalDate() *DateTime {
	if h == nil {
		return nil
	}
	return h.EoSoftwareContractRenewalDate
}

// GetEoSwMaintenanceReleasesDate returns the EoSwMaintenanceReleasesDate field.
func (h *HardwareEOXBulletin) GetEoSwMaintenanceReleasesDate() *DateTime {
	if h == nil {
		return nil
	}
	return h.EoSwMaintenanceReleasesDate
}

// GetHwEoxId returns the HwEoxId field if it's non-nil, zero value otherwise.
func (h *HardwareEOXBulletin) GetHwEoxId() int {
	if h == nil || h.HwEoxId == nil {
		return 0
	}
	return *h.HwEoxId
}

// GetLastDateOfSupport returns the LastDateOfSupport field.
func (h *HardwareEOXBulletin) GetLastDateOfSupport() *DateTime {
	if h == nil {
		return nil
	}
	return h.LastDateOfSupport
}

// GetLastShipDate returns the LastShipDate field.
func (h *HardwareEOXBulletin) GetLastShipDate() *DateTime {
	if h == nil {
		return nil
	}
	return h.LastShipDate
}

// GetProductId returns the ProductId field if it's non-nil, zero value otherwise.
func (h *HardwareEOXBulletin) GetProductId() string {
	if h == nil || h.ProductId == nil {
		return ""
	}
	return *h.ProductId
}

// GetFilter returns the Filter field if it's non-nil, zero value otherwise.
func (l *ListOptions) GetFilter() string {
	if l == nil || l.Filter == nil {
		return ""
	}
	return *l.Filter
}

// GetMask returns the Mask field if it's non-nil, zero value otherwise.
func (l *ListOptions) GetMask() string {
	if l == nil || l.Mask == nil {
		return ""
	}
	return *l.Mask
}

// GetPage returns the Page field if it's non-nil, zero value otherwise.
func (l *ListOptions) GetPage() int {
	if l == nil || l.Page == nil {
		return 0
	}
	return *l.Page
}

// GetPerPage returns the PerPage field if it's non-nil, zero value otherwise.
func (l *ListOptions) GetPerPage() int {
	if l == nil || l.PerPage == nil {
		return 0
	}
	return *l.PerPage
}

// GetDeviceId returns the DeviceId field if it's non-nil, zero value otherwise.
func (s *SecurityAdvisory) GetDeviceId() int {
	if s == nil || s.DeviceId == nil {
		return 0
	}
	return *s.DeviceId
}

// GetMatchConfidence returns the MatchConfidence field if it's non-nil, zero value otherwise.
func (s *SecurityAdvisory) GetMatchConfidence() string {
	if s == nil || s.MatchConfidence == nil {
		return ""
	}
	return *s.MatchConfidence
}

// GetMatchConfidenceReason returns the MatchConfidenceReason field if it's non-nil, zero value otherwise.
func (s *SecurityAdvisory) GetMatchConfidenceReason() string {
	if s == nil || s.MatchConfidenceReason == nil {
		return ""
	}
	return *s.MatchConfidenceReason
}

// GetPsirtColdId returns the PsirtColdId field if it's non-nil, zero value otherwise.
func (s *SecurityAdvisory) GetPsirtColdId() int {
	if s == nil || s.PsirtColdId == nil {
		return 0
	}
	return *s.PsirtColdId
}

// GetBulletinFirstPublished returns the BulletinFirstPublished field if it's non-nil, zero value otherwise.
func (s *SecurityAdvisoryBulletin) GetBulletinFirstPublished() string {
	if s == nil || s.BulletinFirstPublished == nil {
		return ""
	}
	return *s.BulletinFirstPublished
}

// GetBulletinLastUpdated returns the BulletinLastUpdated field.
func (s *SecurityAdvisoryBulletin) GetBulletinLastUpdated() *DateTime {
	if s == nil {
		return nil
	}
	return s.BulletinLastUpdated
}

// GetBulletinMappingCaveat returns the BulletinMappingCaveat field if it's non-nil, zero value otherwise.
func (s *SecurityAdvisoryBulletin) GetBulletinMappingCaveat() string {
	if s == nil || s.BulletinMappingCaveat == nil {
		return ""
	}
	return *s.BulletinMappingCaveat
}

// GetBulletinSummary returns the BulletinSummary field if it's non-nil, zero value otherwise.
func (s *SecurityAdvisoryBulletin) GetBulletinSummary() string {
	if s == nil || s.BulletinSummary == nil {
		return ""
	}
	return *s.BulletinSummary
}

// GetBulletinTitle returns the BulletinTitle field if it's non-nil, zero value otherwise.
func (s *SecurityAdvisoryBulletin) GetBulletinTitle() string {
	if s == nil || s.BulletinTitle == nil {
		return ""
	}
	return *s.BulletinTitle
}

// GetBulletinUrl returns the BulletinUrl field if it's non-nil, zero value otherwise.
func (s *SecurityAdvisoryBulletin) GetBulletinUrl() string {
	if s == nil || s.BulletinUrl == nil {
		return ""
	}
	return *s.BulletinUrl
}

// GetBulletinVersion returns the BulletinVersion field if it's non-nil, zero value otherwise.
func (s *SecurityAdvisoryBulletin) GetBulletinVersion() string {
	if s == nil || s.BulletinVersion == nil {
		return ""
	}
	return *s.BulletinVersion
}

// GetCiscoBugIds returns the CiscoBugIds field if it's non-nil, zero value otherwise.
func (s *SecurityAdvisoryBulletin) GetCiscoBugIds() string {
	if s == nil || s.CiscoBugIds == nil {
		return ""
	}
	return *s.CiscoBugIds
}

// GetCveId returns the CveId field if it's non-nil, zero value otherwise.
func (s *SecurityAdvisoryBulletin) GetCveId() string {
	if s == nil || s.CveId == nil {
		return ""
	}
	return *s.CveId
}

// GetCvssBase returns the CvssBase field if it's non-nil, zero value otherwise.
func (s *SecurityAdvisoryBulletin) GetCvssBase() string {
	if s == nil || s.CvssBase == nil {
		return ""
	}
	return *s.CvssBase
}

// GetCvssTemporal returns the CvssTemporal field if it's non-nil, zero value otherwise.
func (s *SecurityAdvisoryBulletin) GetCvssTemporal() string {
	if s == nil || s.CvssTemporal == nil {
		return ""
	}
	return *s.CvssTemporal
}

// GetPsirtAdvisoryId returns the PsirtAdvisoryId field if it's non-nil, zero value otherwise.
func (s *SecurityAdvisoryBulletin) GetPsirtAdvisoryId() string {
	if s == nil || s.PsirtAdvisoryId == nil {
		return ""
	}
	return *s.PsirtAdvisoryId
}

// GetPsirtColdId returns the PsirtColdId field if it's non-nil, zero value otherwise.
func (s *SecurityAdvisoryBulletin) GetPsirtColdId() int {
	if s == nil || s.PsirtColdId == nil {
		return 0
	}
	return *s.PsirtColdId
}

// GetSir returns the Sir field if it's non-nil, zero value otherwise.
func (s *SecurityAdvisoryBulletin) GetSir() string {
	if s == nil || s.Sir == nil {
		return ""
	}
	return *s.Sir
}

// GetDeviceId returns the DeviceId field if it's non-nil, zero value otherwise.
func (s *SoftwareAlert) GetDeviceId() int {
	if s == nil || s.DeviceId == nil {
		return 0
	}
	return *s.DeviceId
}

// GetDeviceName returns the DeviceName field if it's non-nil, zero value otherwise.
func (s *SoftwareAlert) GetDeviceName() string {
	if s == nil || s.DeviceName == nil {
		return ""
	}
	return *s.DeviceName
}

// GetImageName returns the ImageName field if it's non-nil, zero value otherwise.
func (s *SoftwareAlert) GetImageName() string {
	if s == nil || s.ImageName == nil {
		return ""
	}
	return *s.ImageName
}

// GetSwAlertType returns the SwAlertType field if it's non-nil, zero value otherwise.
func (s *SoftwareAlert) GetSwAlertType() string {
	if s == nil || s.SwAlertType == nil {
		return ""
	}
	return *s.SwAlertType
}

// GetSwAlertUrl returns the SwAlertUrl field if it's non-nil, zero value otherwise.
func (s *SoftwareAlert) GetSwAlertUrl() string {
	if s == nil || s.SwAlertUrl == nil {
		return ""
	}
	return *s.SwAlertUrl
}

// GetSwType returns the SwType field if it's non-nil, zero value otherwise.
func (s *SoftwareAlert) GetSwType() string {
	if s == nil || s.SwType == nil {
		return ""
	}
	return *s.SwType
}

// GetSwVersion returns the SwVersion field if it's non-nil, zero value otherwise.
func (s *SoftwareAlert) GetSwVersion() string {
	if s == nil || s.SwVersion == nil {
		return ""
	}
	return *s.SwVersion
}

// GetCurrentEoxMilestone returns the CurrentEoxMilestone field if it's non-nil, zero value otherwise.
func (s *SoftwareEOX) GetCurrentEoxMilestone() string {
	if s == nil || s.CurrentEoxMilestone == nil {
		return ""
	}
	return *s.CurrentEoxMilestone
}

// GetCurrentEoxMilestoneDate returns the CurrentEoxMilestoneDate field.
func (s *SoftwareEOX) GetCurrentEoxMilestoneDate() *DateTime {
	if s == nil {
		return nil
	}
	return s.CurrentEoxMilestoneDate
}

// GetDeviceId returns the DeviceId field if it's non-nil, zero value otherwise.
func (s *SoftwareEOX) GetDeviceId() int {
	if s == nil || s.DeviceId == nil {
		return 0
	}
	return *s.DeviceId
}

// GetDeviceName returns the DeviceName field if it's non-nil, zero value otherwise.
func (s *SoftwareEOX) GetDeviceName() string {
	if s == nil || s.DeviceName == nil {
		return ""
	}
	return *s.DeviceName
}

// GetNextEoxMilestone returns the NextEoxMilestone field if it's non-nil, zero value otherwise.
func (s *SoftwareEOX) GetNextEoxMilestone() string {
	if s == nil || s.NextEoxMilestone == nil {
		return ""
	}
	return *s.NextEoxMilestone
}

// GetNextEoxMilestoneDate returns the NextEoxMilestoneDate field.
func (s *SoftwareEOX) GetNextEoxMilestoneDate() *DateTime {
	if s == nil {
		return nil
	}
	return s.NextEoxMilestoneDate
}

// GetSwEoxId returns the SwEoxId field if it's non-nil, zero value otherwise.
func (s *SoftwareEOX) GetSwEoxId() int {
	if s == nil || s.SwEoxId == nil {
		return 0
	}
	return *s.SwEoxId
}

// GetSwType returns the SwType field if it's non-nil, zero value otherwise.
func (s *SoftwareEOX) GetSwType() string {
	if s == nil || s.SwType == nil {
		return ""
	}
	return *s.SwType
}

// GetSwVersion returns the SwVersion field if it's non-nil, zero value otherwise.
func (s *SoftwareEOX) GetSwVersion() string {
	if s == nil || s.SwVersion == nil {
		return ""
	}
	return *s.SwVersion
}

// GetBulletinNumber returns the BulletinNumber field if it's non-nil, zero value otherwise.
func (s *SoftwareEOXBulletin) GetBulletinNumber() string {
	if s == nil || s.BulletinNumber == nil {
		return ""
	}
	return *s.BulletinNumber
}

// GetBulletinTitle returns the BulletinTitle field if it's non-nil, zero value otherwise.
func (s *SoftwareEOXBulletin) GetBulletinTitle() string {
	if s == nil || s.BulletinTitle == nil {
		return ""
	}
	return *s.BulletinTitle
}

// GetBulletinUrl returns the BulletinUrl field if it's non-nil, zero value otherwise.
func (s *SoftwareEOXBulletin) GetBulletinUrl() string {
	if s == nil || s.BulletinUrl == nil {
		return ""
	}
	return *s.BulletinUrl
}

// GetEoLifeAnnouncementDate returns the EoLifeAnnouncementDate field.
func (s *SoftwareEOXBulletin) GetEoLifeAnnouncementDate() *DateTime {
	if s == nil {
		return nil
	}
	return s.EoLifeAnnouncementDate
}

// GetEoSaleDate returns the EoSaleDate field.
func (s *SoftwareEOXBulletin) GetEoSaleDate() *DateTime {
	if s == nil {
		return nil
	}
	return s.EoSaleDate
}

// GetEoSecurityVulSupportDate returns the EoSecurityVulSupportDate field.
func (s *SoftwareEOXBulletin) GetEoSecurityVulSupportDate() *DateTime {
	if s == nil {
		return nil
	}
	return s.EoSecurityVulSupportDate
}

// GetEoSwMaintenanceReleasesDate returns the EoSwMaintenanceReleasesDate field.
func (s *SoftwareEOXBulletin) GetEoSwMaintenanceReleasesDate() *DateTime {
	if s == nil {
		return nil
	}
	return s.EoSwMaintenanceReleasesDate
}

// GetLastDateOfSupport returns the LastDateOfSupport field.
func (s *SoftwareEOXBulletin) GetLastDateOfSupport() *DateTime {
	if s == nil {
		return nil
	}
	return s.LastDateOfSupport
}

// GetSwEoxId returns the SwEoxId field if it's non-nil, zero value otherwise.
func (s *SoftwareEOXBulletin) GetSwEoxId() int {
	if s == nil || s.SwEoxId == nil {
		return 0
	}
	return *s.SwEoxId
}

// GetSwMaintenanceVersion returns the SwMaintenanceVersion field if it's non-nil, zero value otherwise.
func (s *SoftwareEOXBulletin) GetSwMaintenanceVersion() string {
	if s == nil || s.SwMaintenanceVersion == nil {
		return ""
	}
	return *s.SwMaintenanceVersion
}

// GetSwMajorVersion returns the SwMajorVersion field if it's non-nil, zero value otherwise.
func (s *SoftwareEOXBulletin) GetSwMajorVersion() string {
	if s == nil || s.SwMajorVersion == nil {
		return ""
	}
	return *s.SwMajorVersion
}

// GetSwTrain returns the SwTrain field if it's non-nil, zero value otherwise.
func (s *SoftwareEOXBulletin) GetSwTrain() string {
	if s == nil || s.SwTrain == nil {
		return ""
	}
	return *s.SwTrain
}

// GetSwType returns the SwType field if it's non-nil, zero value otherwise.
func (s *SoftwareEOXBulletin) GetSwType() string {
	if s == nil || s.SwType == nil {
		return ""
	}
	return *s.SwType
}