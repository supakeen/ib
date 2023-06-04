/*
Copyright © 2023 Simon de Vlieger <supakeen@redhat.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/

/*
Functions and structs to perform requests to the Image Builder API.
*/

package api

type Package string
type Empty struct{}

type CustomRepository struct{}
type PayloadRepository struct{}

type Filesystem struct{}
type Subscription struct{}

type User struct {
	Name   string `json:name`
	SSHKey string `json:ssh_key`
}

type Customizations struct {
	Packages            []Package           `json:"packages,omitempty"`
	CustomRepositories  []CustomRepository  `json:"custom_repositories,omitempty"`
	PayloadRepositories []PayloadRepository `json:"payload_repositories,omitempty"`
	SubscriptionDetails *Subscription       `json:"subscription,omitempty"`
	Users               []User              `json:"users,omitempty"`
}

type UploadRequest struct {
	Type    string `json:"type"`
	Options Empty  `json:"options"`
}

type ImageRequest struct {
	Architecture  string        `json:"architecture"`
	ImageType     string        `json:"image_type"`
	UploadRequest UploadRequest `json:"upload_request"`
}

type ComposeRequest struct {
	Distribution   string          `json:"distribution"`
	Name           string          `json:"image_name"`
	Customizations *Customizations `json:"customizations,omitempty"`
	ImageRequests  []ImageRequest  `json:"image_requests"`
}
