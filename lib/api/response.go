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

type Architectures []ArchitectureItem

type ArchitectureItem struct {
	Name       string   `json:"arch"`
	ImageTypes []string `json:"image_types"`
}

type Distributions []DistributionItem

type DistributionItem struct {
	Description string `json:"description"`
	Name        string `json:"name"`
}

type ComposeResponse struct {
	ID string `json:"id"`
}

type ComposesResponse struct {
	Data  []ComposesResponseItem `json:"data"`
	Links struct {
		First string `json:"first"`
		Last  string `json:"last"`
	} `json:"links"`
}

type ComposesResponseItem struct {
	CreatedAt string `json:"created_at"`
	ID        string `json:"id"`
	Name      string `json:"image_name"`
}

type ImageStatus struct {
	Status       string `json:"status"`
	UploadStatus struct {
		Options struct {
			URL string `json:"url"`
		} `json:"options"`
	} `json:"upload_status"`
}

type ComposeStatus struct {
	Status  ImageStatus    `json:"image_status"`
	Request ComposeRequest `json:"request"`
}
