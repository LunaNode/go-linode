// Copyright (c) 2015 LunaNode Hosting Inc. All right reserved.
// Use of this source code is governed by the MIT License. See LICENSE file.

package linode

import "fmt"

type LinodeIDResponse struct {
	LinodeID int `json:"LinodeID"`
}

type JobResponse struct {
	JobID int `json:"JobID"`
}

type Linode struct {
	ID int `json:"LINODEID"`
	Status int `json:"STATUS"`
	StatusString string
	Label string `json:"LABEL"`
	CreateTime string `json:"CREATE_DT"`
	Plan int `json:"PLANID"`
	DistributionVendor string `json:"DISTRIBUTIONVENDOR"`
	IsXen int `json:"ISXEN"`
	IsKVM int `json:"ISKVM"`
	TotalBandwidth int `json:"TOTALXFER"`
	TotalRAM int `json:"TOTALRAM"`
	TotalHD int `json:"TOTALHD"`
	DatacenterID int `json:"DATACENTERID"`
	AlertBandwidth int `json:"ALERT_BWQUOTA_ENABLED"`
	AlertBandwidthThreshold int `json:"ALERT_BWQUOTA_THRESHOLD"`
	AlertDiskIO int `json:"ALERT_DISKIO_ENABLED"`
	AlertDiskIOThreshold int `json:"ALERT_DISKIO_THRESHOLD"`
	AlertBandwidthOut int `json:"ALERT_BWOUT_ENABLED"`
	AlertBandwidthOutThreshold int `json:"ALERT_BWOUT_THRESHOLD"`
	AlertBandwidthIn int `json:"ALERT_BWIN_ENABLED"`
	AlertBandwidthInThreshold int `json:"ALERT_BWIN_THRESHOLD"`
	AlertCPU int `json:"ALERT_CPU_ENABLED"`
	AlertCPUThreshold int `json:"ALERT_CPU_THRESHOLD"`
	BackupEnabled int `json:"BACKUPSENABLED"`
	BackupWindow int `json:"BACKUPWINDOW"`
	BackupWeeklyDay int `json:"BACKUPWEEKLYDAY"`
	Watchdog int `json:"WATCHDOG"`
}

func (linode *Linode) Parse() {
	if linode.Status == -1 {
		linode.StatusString = "Being Created"
	} else if linode.Status == 0 {
		linode.StatusString = "Brand New"
	} else if linode.Status == 1 {
		linode.StatusString = "Running"
	} else if linode.Status == 2 {
		linode.StatusString = "Powered Off"
	} else {
		linode.StatusString = fmt.Sprintf("Unknown (%d)", linode.Status)
	}
}

type ConfigIDResponse struct {
	ConfigID int `json:"ConfigID"`
}

type CreateConfigOptions struct {
	Comments string
	RAMLimit int
	VirtMode string
	RunLevel string
	RootDeviceNum int
	RootDeviceCustom string
	DisableRootDeviceRO bool
	DisableDisableUpdateDBHelper bool
	DisableDistroHelper bool
	DisableDepmodHelper bool
	EnableHelperNetwork bool
	DisableAutomountDevtmpfs bool
}

type CreateDiskOptions struct {
	DistributionID int
	RootPassword string
	RootSSHKey string
	IsReadOnly bool
}

type LinodeDiskCreateResponse struct {
	JobID int `json:"JobID"`
	DiskID int `json:"DiskID"`
}

type ImagizeResponse struct {
	JobID int `json:"JobID"`
	ImageID int `json:"ImageID"`
}

type IPIDResponse struct {
	ID int `json:"IPAddressID"`
}

type IP struct {
	ID int `json:"IPADDRESSID"`
	LinodeID int `json:"LINODEID"`
	IsPublic int `json:"ISPUBLIC"`
	Address string `json:"IPADDRESS"`
	Hostname string `json:"RDNS_NAME"`
}

type Job struct {
	ID int `json:"JOBID"`
	Action string `json:"ACTION"`
	Label string `json:"LABEL"`
	LinodeID int `json:"LINODEID"`
	Duration int `json:"DURATION"`
	Success int `json:"HOST_SUCCESS"`
	Message string `json:"HOST_MESSAGE"`
	EnteredTime string `json:"ENTERED_DT"`
	StartTime string `json:"HOST_START_DT"`
	FinishTime string `json:"HOST_FINISH_DT"`
}
