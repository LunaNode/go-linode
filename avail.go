package linode

func (client *Client) ListDatacenters() ([]*Datacenter, error) {
	var datacenters []*Datacenter
	err := client.request("avail.datacenters", nil, &datacenters)
	return datacenters, err
}

func (client *Client) ListDistributions() ([]*Distribution, error) {
	var distributions []*Distribution
	err := client.request("avail.distributions", nil, &distributions)
	return distributions, err
}

func (client *Client) ListKernels() ([]*Kernel, error) {
	var kernels []*Kernel
	err := client.request("avail.kernels", nil, &kernels)
	return kernels, err
}

func (client *Client) ListPlans() ([]*Plan, error) {
	var plans []*Plan
	err := client.request("avail.linodeplans", nil, &plans)
	return plans, err
}
