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

	"github.com/spf13/cobra"
	"github.com/supakeen/ib/lib/api"
)

var ComposeDistribution string
var ComposeArchitecture string
var ComposeImageType string
var ComposeOutputFile string
var ComposePackages string
var ComposeUsers string

// composeCmd represents the compose command
var composeCmd = &cobra.Command{
	Use:   "c",
	Short: "Build a (customized) image",
	Long:  `Build a (customized) image. Wait for the build. Download the result.`,
	Run: func(cmd *cobra.Command, args []string) {
		var users []api.User

		for _, userFromCli := range strings.Split(ComposeUsers, ",") {
			part := strings.Split(userFromCli, ":")

			if len(part) != 2 {
				log.Fatal("invalid user")
			}

			users = append(users, api.User{
				Name:   part[0],
				SSHKey: part[1],
			})
		}

		composeID := api.NewComposeRequest(ComposeDistribution, ComposeArchitecture, ComposeImageType, "demo", strings.Split(ComposePackages, ","), users)

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

	composeCmd.Flags().StringVarP(&ComposeDistribution, "distribution", "d", "", "Request a specific distribution for example `centos-9`.")
	composeCmd.MarkFlagRequired("distribution")

	composeCmd.Flags().StringVarP(&ComposeArchitecture, "architecture", "a", "", "Request a specific architecture for example `x86_64`.")
	composeCmd.MarkFlagRequired("architecture")

	composeCmd.Flags().StringVarP(&ComposeImageType, "image-type", "t", "", "Request a specific image-type for example `guest-image`.")
	composeCmd.MarkFlagRequired("image-type")

	composeCmd.Flags().StringVarP(&ComposeOutputFile, "output-file", "o", "", "File to write to for example `image.qcow`.")
	composeCmd.MarkFlagRequired("output-file")

	composeCmd.Flags().StringVarP(&ComposePackages, "packages", "p", "", "Additional packages to install in `nginx,tmux` format.")

	composeCmd.Flags().StringVarP(&ComposeUsers, "users", "u", "", "Additional users to add in `user:ssh-key,user:ssh-key` format.")

}
