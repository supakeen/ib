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

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"sort"
)

func NewComposeRequest(distribution string, architecture string, imageType string, name string, packages []string) string {
	EnsureToken()

	var composeRequest ComposeRequest
	var imageRequest ImageRequest
	var uploadRequest UploadRequest
	var customizationRequest Customizations

	composeRequest.Distribution = distribution
	composeRequest.Name = name
	composeRequest.Customizations = &customizationRequest

	imageRequest.Architecture = architecture
	imageRequest.ImageType = imageType

	uploadRequest.Type = "aws.s3"

	imageRequest.UploadRequest = uploadRequest

	composeRequest.ImageRequests = []ImageRequest{imageRequest}

	data, err := json.Marshal(composeRequest)

	if err != nil {
		log.Fatal(err)
	}

	url := "https://console.redhat.com/api/image-builder/v1/compose"
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(data))

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+AccessToken)

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		log.Fatal(err)
	}

	var composeResponse ComposeResponse

	err = json.Unmarshal(body, &composeResponse)

	if err != nil {
		log.Fatal(err)
	}

	return composeResponse.ID
}

func NewComposeStatusRequest(composeID string) string {
	EnsureToken()

	url := "https://console.redhat.com/api/image-builder/v1/composes/" + composeID
	req, err := http.NewRequest(http.MethodGet, url, nil)

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", "Bearer "+AccessToken)

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		log.Fatal(err)
	}

	var imageStatus ComposeStatus

	err = json.Unmarshal(body, &imageStatus)

	if err != nil {
		log.Fatal(err)
	}

	return imageStatus.Status.Status
}

func NewComposeDownloadRequest(composeID string) string {
	EnsureToken()

	url := "https://console.redhat.com/api/image-builder/v1/composes/" + composeID
	req, err := http.NewRequest(http.MethodGet, url, nil)

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", "Bearer "+AccessToken)

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		log.Fatal(err)
	}

	var imageStatus ComposeStatus

	err = json.Unmarshal(body, &imageStatus)

	if err != nil {
		log.Fatal(err)
	}

	return imageStatus.Status.UploadStatus.Options.URL
}

func NewDistributionsRequest() Distributions {
	EnsureToken()

	url := "https://console.redhat.com/api/image-builder/v1/distributions"
	req, err := http.NewRequest(http.MethodGet, url, nil)

	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", "Bearer "+AccessToken)

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		log.Fatal(err)
	}

	var distributions Distributions

	err = json.Unmarshal(body, &distributions)

	if err != nil {
		log.Fatal(err)
	}

	sort.Slice(distributions, func(i, j int) bool {
		return distributions[i].Name < distributions[j].Name
	})

	return distributions
}

func NewArchitecturesRequest(distribution string) Architectures {
	EnsureToken()

	url := "https://console.redhat.com/api/image-builder/v1/architectures/" + distribution
	req, err := http.NewRequest(http.MethodGet, url, nil)

	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", "Bearer "+AccessToken)

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		log.Fatal(err)
	}

	var architectures Architectures

	err = json.Unmarshal(body, &architectures)

	if err != nil {
		log.Fatal(err)
	}

	sort.Slice(architectures, func(i, j int) bool {
		return architectures[i].Name < architectures[j].Name
	})

	return architectures
}

func NewComposeListRequest() []ComposesResponseItem {
	EnsureToken()

	url := "https://console.redhat.com/api/image-builder/v1/composes"
	req, err := http.NewRequest(http.MethodGet, url, nil)

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", "Bearer "+AccessToken)

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		log.Fatal(err)
	}

	var listResponse ComposesResponse

	err = json.Unmarshal(body, &listResponse)

	if err != nil {
		log.Fatal(err)
	}

	return listResponse.Data
}
