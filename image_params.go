package linode

type Image struct {
	ID int `json:"IMAGEID"`
	Label string `json:"LABEL"`
	Description string `json:"DESCRIPTION"`
	Creator string `json:"CREATOR"`
	CreateTime string `json:"CREATE_DT"`
	FilesystemType string `json:"FS_TYPE"`
	IsPublic int `json:"ISPUBLIC"`
	MinSize int64 `json:"MINSIZE"`
	Status string `json:"STATUS"`
	Type string `json:"TYPE"`
}
