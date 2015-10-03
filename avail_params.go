// Copyright (c) 2015 LunaNode Hosting Inc. All right reserved.
// Use of this source code is governed by the MIT License. See LICENSE file.

package linode

type Datacenter struct {
	// The datacenter ID.
	ID int `json:"DATACENTERID"`

	// String describing datacenter location (e.g. "Newark, NJ, USA").
	Location string `json:"LOCATION"`

	// Short string identifying this datacenter (e.g. "newark").
	Abbreviation string `json:"ABBR"`
}

type Distribution struct {
	// The distribution ID.
	ID int `json:"DISTRIBUTIONID"`

	// 1 if this distribution is 64-bit.
	Is64Bit int `json:"IS64BIT"`

	Label string `json:"LABEL"`
	MinImageSize int `json:"MINIMAGESIZE"`
	CreateTime string `json:"CREATE_DT"`
	RequiresPVOpsKernel int `json:"REQUIRESPVOPSKERNEL"`
}

type Kernel struct {
	// The kernel ID.
	ID int `json:"KERNELID"`

	Label string `json:"LABEL"`
	IsXen int `json:"ISXEN"`
	IsKVM int `json:"ISKVM"`
	IsPVOps int `json:"ISPVOPS"`
}

type Plan struct {
	// The plan ID.
	ID int `json:"PLANID"`

	Label string `json:"LABEL"`
	MonthlyPrice float64 `json:"PRICE"`
	HourlyPrice float64 `json:"HOURLY"`

	// Number of allocated CPU cores.
	Cores int `json:"CORES"`

	// Memory allocation in MB.
	RAM int `json:"RAM"`

	// Bandwidth allocation in GB.
	Bandwidth int `json:"XFER"`

	// Disk allocation in GB.
	Disk int `json:"DISK"`
}
