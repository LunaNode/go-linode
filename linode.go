// Copyright (c) 2015 LunaNode Hosting Inc. All right reserved.
// Use of this source code is governed by the MIT License. See LICENSE file.

package linode

import "errors"
import "fmt"

// Creates a new Linode, returning Linode ID on success.
func (client *Client) CreateLinode(datacenterID int, planID int) (int, error) {
	params := map[string]string{
		"DatacenterID": fmt.Sprintf("%d", datacenterID),
		"PlanID": fmt.Sprintf("%d", planID),
	}
	response := new(linodeIDResponse)
	err := client.request("linode.create", params, response)
	return response.LinodeID, err
}

// Deletes the specified Linode.
// If skipChecks is false, the Linode will only be deleted if there are no disk images.
// Otherwise, the Linode is always deleted.
func (client *Client) DeleteLinode(linodeID int, skipChecks bool) error {
	params := map[string]string{
		"LinodeID": fmt.Sprintf("%d", linodeID),
	}
	if skipChecks {
		params["skipChecks"] = "1"
	}
	return client.request("linode.delete", params, nil)
}

func (client *Client) linodeAction(action string, linodeID int) (int, error) {
	params := map[string]string{
		"LinodeID": fmt.Sprintf("%d", linodeID),
	}
	response := new(jobResponse)
	err := client.request("linode." + action, params, response)
	return response.JobID, err
}

// Boots Linode with the last used configuration profile, or the first configuration profile
// if the Linode has not been booted before.
func (client *Client) BootLinode(linodeID int) (int, error) {
	return client.linodeAction("boot", linodeID)
}

func (client *Client) BootLinodeWithConfig(linodeID int, configID int) (int, error) {
	params := map[string]string{
		"LinodeID": fmt.Sprintf("%d", linodeID),
		"ConfigID": fmt.Sprintf("%d", configID),
	}
	response := new(jobResponse)
	err := client.request("linode.boot", params, response)
	return response.JobID, err
}

func (client *Client) RebootLinode(linodeID int) (int, error) {
	return client.linodeAction("reboot", linodeID)
}

func (client *Client) ShutdownLinode(linodeID int) (int, error) {
	return client.linodeAction("shutdown", linodeID)
}

func (client *Client) ListLinodes() ([]*Linode, error) {
	var linodes []*Linode
	err := client.request("linode.list", nil, &linodes)
	if err != nil {
		return nil, err
	}
	for _, linode := range linodes {
		linode.parse()
	}
	return linodes, nil
}

func (client *Client) GetLinode(linodeID int) (*Linode, error) {
	params := map[string]string{
		"LinodeID": fmt.Sprintf("%d", linodeID),
	}
	var linodes []*Linode
	err := client.request("linode.list", params, &linodes)
	if err != nil {
		return nil, err
	} else if len(linodes) != 1 {
		return nil, errors.New("expected one linode in response")
	} else {
		linodes[0].parse()
		return linodes[0], nil
	}
}

func (client *Client) ResizeLinode(linodeID int, planID int) error {
	params := map[string]string{
		"LinodeID": fmt.Sprintf("%d", linodeID),
		"PlanID": fmt.Sprintf("%d", planID),
	}
	return client.request("linode.resize", params, nil)
}

// Clones the Linode to a new instance in the specified datacenter with the specified plan.
// Returns Linode ID on success.
func (client *Client) CloneLinode(linodeID int, datacenterID int, planID int) (int, error) {
	params := map[string]string{
		"LinodeID": fmt.Sprintf("%d", linodeID),
		"DatacenterID": fmt.Sprintf("%d", datacenterID),
		"PlanID": fmt.Sprintf("%d", planID),
	}
	response := new(linodeIDResponse)
	err := client.request("linode.clone", params, response)
	return response.LinodeID, err
}

func (client *Client) UpdateLinode(linodeID int, updateParams map[string]string) error {
	params := map[string]string{
		"LinodeID": fmt.Sprintf("%d", linodeID),
	}
	for k, v := range updateParams {
		params[k] = v
	}
	return client.request("linode.update", params, nil)
}

func (client *Client) RenameLinode (linodeID int, newLabel string) error {
	return client.UpdateLinode(linodeID, map[string]string{"Label": newLabel})
}

func (client *Client) CreateConfig(linodeID int, kernelID int, label string, disks []int, options CreateConfigOptions) (int, error) {
	diskList := ""
	for _, diskID := range disks {
		if diskList != "" {
			diskList += ","
		}
		diskList += fmt.Sprintf("%d", diskID)
	}
	params := map[string]string{
		"LinodeID": fmt.Sprintf("%d", linodeID),
		"KernelID": fmt.Sprintf("%d", kernelID),
		"Label": label,
		"DiskList": diskList,
	}
	if options.Comments != "" {
		params["Comments"] = options.Comments
	}
	if options.RAMLimit != 0 {
		params["RAMLimit"] = fmt.Sprintf("%d", options.RAMLimit)
	}
	if options.VirtMode != "" {
		params["virt_mode"] = options.VirtMode
	}
	if options.RunLevel != "" {
		params["RunLevel"] = options.RunLevel
	}
	if options.Comments != "" {
		params["Comments"] = options.Comments
	}
	if options.RootDeviceNum != 0 {
		params["RootDeviceNum"] = fmt.Sprintf("%d", options.RootDeviceNum)
	}
	if options.RootDeviceCustom != "" {
		// RootDeviceNum must be 0 to use RootDeviceCustom
		params["RootDeviceCustom"] = options.RootDeviceCustom
		if params["RootDeviceNum"] == "" {
			params["RootDeviceNum"] = "0"
		} else if params["RootDeviceNum"] != "0" {
			return 0, errors.New("RootDeviceCustom is set but RootDeviceNum is set non-zero")
		}
	}
	if options.DisableRootDeviceRO {
		params["RootDeviceRO"] = "0"
	}
	if options.DisableDisableUpdateDBHelper {
		params["helper_disableUpdateDB"] = "0"
	}
	if options.DisableDistroHelper {
		params["helper_distro"] = "0"
	}
	if options.DisableDepmodHelper {
		params["helper_depmod"] = "0"
	}
	if options.EnableHelperNetwork {
		params["helper_network"] = "1"
	}
	if options.DisableAutomountDevtmpfs {
		params["devtmpfs_automount"] = "0"
	}
	response := new(configIDResponse)
	err := client.request("linode.config.create", params, response)
	return response.ConfigID, err
}

func (client *Client) DeleteConfig(linodeID int, configID int) error {
	params := map[string]string{
		"LinodeID": fmt.Sprintf("%d", linodeID),
		"ConfigID": fmt.Sprintf("%d", configID),
	}
	return client.request("linode.config.delete", params, nil)
}

func (client *Client) CreateDisk(linodeID int, label string, diskType string, size int, diskOptions CreateDiskOptions) (int, int, error) {
	params := map[string]string{
		"LinodeID": fmt.Sprintf("%d", linodeID),
		"Label": label,
		"Type": diskType,
		"Size": fmt.Sprintf("%d", size),
	}
	if diskOptions.DistributionID != 0 {
		params["FromDistributionID"] = fmt.Sprintf("%d", diskOptions.DistributionID)
	}
	if diskOptions.RootPassword != "" {
		params["rootPass"] = diskOptions.RootPassword
	}
	if diskOptions.RootSSHKey != "" {
		params["rootSSHKey"] = diskOptions.RootSSHKey
	}
	if diskOptions.IsReadOnly {
		params["isReadOnly"] = "1"
	}
	response := new(linodeDiskCreateResponse)
	err := client.request("linode.disk.create", params, response)
	return response.DiskID, response.JobID, err
}

func (client *Client) CreateDiskFromDistribution(linodeID int, label string, distributionID int, size int, rootPass string, rootSSHKey string) (int, int, error) {
	params := map[string]string{
		"LinodeID": fmt.Sprintf("%d", linodeID),
		"Label": label,
		"DistributionID": fmt.Sprintf("%d", distributionID),
		"Size": fmt.Sprintf("%d", size),
		"rootPass": rootPass,
	}
	if rootSSHKey != "" {
		params["rootSSHKey"] = rootSSHKey
	}
	response := new(linodeDiskCreateResponse)
	err := client.request("linode.disk.createfromdistribution", params, response)
	return response.DiskID, response.JobID, err
}

func (client *Client) CreateDiskFromImage(linodeID int, label string, imageID int, size int, rootPass string, rootSSHKey string) (int, int, error) {
	params := map[string]string{
		"LinodeID": fmt.Sprintf("%d", linodeID),
		"Label": label,
		"ImageID": fmt.Sprintf("%d", imageID),
		"Size": fmt.Sprintf("%d", size),
	}
	if rootPass != "" {
		params["rootPass"] = rootPass
	}
	if rootSSHKey != "" {
		params["rootSSHKey"] = rootSSHKey
	}
	response := new(linodeDiskCreateResponse)
	err := client.request("linode.disk.createfromimage", params, response)
	return response.DiskID, response.JobID, err
}

func (client *Client) DeleteDisk(linodeID int, diskID int) (int, error) {
	params := map[string]string{
		"LinodeID": fmt.Sprintf("%d", linodeID),
		"DiskID": fmt.Sprintf("%d", diskID),
	}
	response := new(jobResponse)
	err := client.request("linode.disk.delete", params, response)
	return response.JobID, err
}

func (client *Client) ImagizeDisk(linodeID int, diskID int, label string) (int, int, error) {
	params := map[string]string{
		"LinodeID": fmt.Sprintf("%d", linodeID),
		"DiskID": fmt.Sprintf("%d", diskID),
		"Label": label,
	}
	response := new(imagizeResponse)
	err := client.request("linode.disk.imagize", params, response)
	return response.ImageID, response.JobID, err
}

func (client *Client) ResizeDisk(linodeID int, diskID int, size int) (int, error) {
	params := map[string]string{
		"LinodeID": fmt.Sprintf("%d", linodeID),
		"DiskID": fmt.Sprintf("%d", diskID),
		"size": fmt.Sprintf("%d", size),
	}
	response := new(jobResponse)
	err := client.request("linode.disk.resize", params, response)
	return response.JobID, err
}

func (client *Client) ListIP(linodeID int) ([]*IP, error) {
	params := map[string]string{
		"LinodeID": fmt.Sprintf("%d", linodeID),
	}
	var ips []*IP
	err := client.request("linode.ip.list", params, &ips)
	return ips, err
}

func (client *Client) GetIP(ipID int) (*IP, error) {
	params := map[string]string{
		"IPAddressID": fmt.Sprintf("%d", ipID),
	}
	var ips []*IP
	err := client.request("linode.ip.list", params, &ips)
	if err != nil {
		return nil, err
	} else if len(ips) != 1 {
		return nil, errors.New("expected one IP in response")
	} else {
		return ips[0], nil
	}
}

func (client *Client) SetRDNS(ipID int, hostname string) error {
	params := map[string]string{
		"IPAddressID": fmt.Sprintf("%d", ipID),
		"Hostname": hostname,
	}
	return client.request("linode.ip.setrdns", params, nil)
}

// Moves the specified IP address to the specified Linode.
func (client *Client) MoveIP(ipID int, linodeID int) error {
	params := map[string]string{
		"IPAddressID": fmt.Sprintf("%d", ipID),
		"toLinodeID": fmt.Sprintf("%d", linodeID),
	}
	return client.request("linode.ip.swap", params, nil)
}

// Swaps IP addresses between two different Linodes in the same datacenter.
func (client *Client) SwapIP(ipID1 int, ipID2 int) error {
	params := map[string]string{
		"IPAddressID": fmt.Sprintf("%d", ipID1),
		"withIPAddressID": fmt.Sprintf("%d", ipID2),
	}
	return client.request("linode.ip.swap", params, nil)
}

// Adds an IP address to the specified Linode.
// If private is false, adds a public IP, otherwise adds a private IP.
// Returns the IP address ID on success.
func (client *Client) AddIP(linodeID int, private bool) (int, error) {
	params := map[string]string{
		"LinodeID": fmt.Sprintf("%d", linodeID),
	}
	action := "linode.ip.addpublic"
	if private {
		action = "linode.ip.addprivate"
	}
	response := new(ipIDResponse)
	err := client.request(action, params, response)
	return response.ID, err
}

func (client *Client) JobList(linodeID int, pendingOnly bool) ([]*Job, error) {
	params := map[string]string{
		"LinodeID": fmt.Sprintf("%d", linodeID),
	}
	if pendingOnly {
		params["pendingOnly"] = "1"
	}
	var jobs []*Job
	err := client.request("linode.job.list", params, &jobs)
	return jobs, err
}

func (client *Client) GetJob(linodeID int, jobID int) (*Job, error) {
	params := map[string]string{
		"LinodeID": fmt.Sprintf("%d", linodeID),
		"JobID": fmt.Sprintf("%d", jobID),
	}
	var jobs []*Job
	err := client.request("linode.job.list", params, &jobs)
	if err != nil {
		return nil, err
	} else if len(jobs) != 1 {
		return nil, errors.New("expected one job in response")
	} else {
		return jobs[0], nil
	}
}
