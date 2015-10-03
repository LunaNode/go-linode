package linode

type Datacenter struct {
	ID int `json:"DATACENTERID"`
	Location string `json:"LOCATION"`
	Abbreviation string `json:"ABBR"`
}

type Distribution struct {
	ID int `json:"DISTRIBUTIONID"`
	Is64Bit int `json:"IS64BIT"`
	Label string `json:"LABEL"`
	MinImageSize int `json:"MINIMAGESIZE"`
	CreateTime string `json:"CREATE_DT"`
	RequiresPVPVOpsKernel int `json:"REQUIRESPVOPSKERNEL"`
}

type Kernel struct {
	ID int `json:"KERNELID"`
	Label string `json:"LABEL"`
	IsXen int `json:"ISXEN"`
	IsKVM int `json:"ISKVM"`
	IsPVOps int `json:"ISPVOPS"`
}

type Plan struct {
	ID int `json:"PLANID"`
	Label string `json:"LABEL"`
	MonthlyPrice float64 `json:"PRICE"`
	HourlyPrice float64 `json:"HOURLY"`
	Cores int `json:"CORES"`
	RAM int `json:"RAM"`
	Bandwidth int `json:"XFER"`
	Disk int `json:"DISK"`
}
