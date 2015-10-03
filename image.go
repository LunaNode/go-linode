package linode

import "errors"
import "fmt"

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
