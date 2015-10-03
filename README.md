go-linode
=========

A Go library for interacting with the Linode API. [View the GoDoc](https://godoc.org/github.com/LunaNode/go-linode).

Usage
-----

Import with:
```
import github.com/LunaNode/go-linode
```

And use as linode:
```
apiKey := "myKey"
client := linode.NewClient(apiKey)

linodeID, err := client.CreateLinode(6, 1) // NJ, 1024 MB
if err != nil {
	panic(err)
}
// create a 24 GB Ubuntu 14.04 disk with root password and no SSH key
diskID, _, err := client.CreateDiskFromDistribution(linodeID, "go-linode disk", 124, 24 * 1024, "mypassword", "")
if err != nil {
	panic(err)
}
// create configuration with the latest 64-bit kernel
kernelID := 138
kernels, err := client.ListKernels()
if err != nil {
	panic(err)
}
for _, kernel := range kernels {
	if strings.Contains(kernel.Label, "Latest 64 bit") {
		kernelID = kernel.ID
		break
	}
}
configID, err := client.CreateConfig(linodeID, kernelID, "mylinode", []int{diskID}, linode.CreateConfigOptions{})
if err != nil {
	panic(err)
}
client.BootLinodeWithConfig(linodeID, configID)
```

Examples
--------

See the [Lobster module for Linode API](https://github.com/LunaNode/lobster/tree/master/vmlinode) for an example
use of this library.
