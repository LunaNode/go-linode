// Copyright (c) 2015 LunaNode Hosting Inc. All right reserved.
// Use of this source code is governed by the MIT License. See LICENSE file.

package linode

type Image struct {
	// The image ID.
	ID int `json:"IMAGEID"`

	Label string `json:"LABEL"`
	Description string `json:"DESCRIPTION"`
	Creator string `json:"CREATOR"`
	CreateTime string `json:"CREATE_DT"`

	// Filesystem type, e.g. "ext4".
	FilesystemType string `json:"FS_TYPE"`

	IsPublic int `json:"ISPUBLIC"`

	// Minimum size in MB for disks created from this image.
	MinSize int64 `json:"MINSIZE"`

	// String status identifier, set to "available" if available.
	Status string `json:"STATUS"`

	Type string `json:"TYPE"`
}
