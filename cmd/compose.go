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
	"fmt"
	"log"
	"strings"
	"time"

	"io"
	"net/http"
	"os"

	"github.com/supakeen/ib/lib/api"
	"github.com/spf13/cobra"
)

var ComposeDistribution string
var ComposeArchitecture string
var ComposeImageType string
var ComposeOutputFile string
var ComposePackages string

// composeCmd represents the compose command
var composeCmd = &cobra.Command{
	Use:   "c",
	Short: "Build an image",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		composeID := api.NewComposeRequest(ComposeDistribution, ComposeArchitecture, ComposeImageType, "demo", strings.Split(ComposePackages, ","))

		fmt.Printf("queued %s\n", composeID)

		fmt.Printf("waiting ")

		for api.NewComposeStatusRequest(composeID) != "success" {
			fmt.Print(".")
			time.Sleep(10 * time.Second)
		}

		fmt.Println()
		fmt.Printf("downloading %s\n", composeID)

		url := api.NewComposeDownloadRequest(composeID)

		out, err := os.Create(ComposeOutputFile)
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
	rootCmd.AddCommand(composeCmd)

	composeCmd.Flags().StringVarP(&ComposeDistribution, "distribution", "d", "", "Request a specific distribution.")
	composeCmd.MarkFlagRequired("distribution")

	composeCmd.Flags().StringVarP(&ComposeArchitecture, "architecture", "a", "", "Request a specific architecture.")
	composeCmd.MarkFlagRequired("architecture")

	composeCmd.Flags().StringVarP(&ComposeImageType, "image-type", "t", "", "Request a specific image-type.")
	composeCmd.MarkFlagRequired("image-type")

	composeCmd.Flags().StringVarP(&ComposeOutputFile, "output-file", "o", "", "File to write to.")
	composeCmd.MarkFlagRequired("output-file")

	composeCmd.Flags().StringVarP(&ComposePackages, "packages", "p", "", "Comma-separated list of additional packages to install.")

}
