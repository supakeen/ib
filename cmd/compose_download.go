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
package cmd

import (
	"log"

	"io"
	"net/http"
	"os"

	"github.com/supakeen/ib/lib/api"
	"github.com/spf13/cobra"
)

var DownloadComposeID string
var DownloadOutputFile string

// downloadCmd represents the download command
var downloadCmd = &cobra.Command{
	Use:   "d",
	Short: "Download a compose by identifier.",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		url := api.NewComposeDownloadRequest(DownloadComposeID)

		out, err := os.Create(DownloadOutputFile)
		defer out.Close()

		if err != nil {
			log.Fatal(err)
		}

		res, err := http.Get(url)
		defer res.Body.Close()

		_, err = io.Copy(out, res.Body)

		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	composeCmd.AddCommand(downloadCmd)

	downloadCmd.Flags().StringVarP(&DownloadComposeID, "compose-id", "i", "", "Request a specific compose id.")
	downloadCmd.MarkFlagRequired("compose-id")

	downloadCmd.Flags().StringVarP(&DownloadOutputFile, "output-file", "o", "", "File to write to.")
	downloadCmd.MarkFlagRequired("output-file")

}
