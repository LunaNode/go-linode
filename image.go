// Copyright (c) 2015 LunaNode Hosting Inc. All right reserved.
// Use of this source code is governed by the MIT License. See LICENSE file.

package linode

import "errors"
import "fmt"

// Returns a list of images on the user account.
// Restricts to pending images if pendingOnly is true.
func (client *Client) ListImages(pendingOnly bool) ([]*Image, error) {
	params := make(map[string]string)
	if pendingOnly {
		params["pending"] = "1"
	}
	var images []*Image
	err := client.request("image.list", params, &images)
	return images, err
}

func (client *Client) GetImage(imageID int) (*Image, error) {
	params := map[string]string{
		"ImageID": fmt.Sprintf("%d", imageID),
	}
	var images []*Image
	err := client.request("image.list", params, &images)
	if err != nil {
		return nil, err
	} else if len(images) != 1 {
		return nil, errors.New("expected one image in response")
	} else {
		return images[0], nil
	}
}

func (client *Client) DeleteImage(imageID int) error {
	params := map[string]string{
		"ImageID": fmt.Sprintf("%d", imageID),
	}
	return client.request("image.delete", params, nil)
}
