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

	"github.com/spf13/cobra"
	"github.com/supakeen/ib/lib/api"
)

var QueueDistribution string
var QueueArchitecture string
var QueueImageType string
var QueuePackages string
var QueueUsers string

// queueCmd represents the compose command
var queueCmd = &cobra.Command{
	Use:   "q",
	Short: "Queue an image build",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		var users []api.User

		for _, userFromCli := range strings.Split(QueueUsers, ",") {
			part := strings.Split(userFromCli, ":")

			if len(part) != 2 {
				log.Fatal("invalid user")
			}

			users = append(users, api.User{
				Name:   part[0],
				SSHKey: part[1],
			})
		}

		fmt.Println(api.NewComposeRequest(QueueDistribution, QueueArchitecture, QueueImageType, "demo", strings.Split(QueuePackages, ","), users))
	},
}

func init() {
	composeCmd.AddCommand(queueCmd)

	queueCmd.Flags().StringVarP(&QueueDistribution, "distribution", "d", "", "Request a specific distribution for example `centos-9`.")
	queueCmd.MarkFlagRequired("distribution")

	queueCmd.Flags().StringVarP(&QueueArchitecture, "architecture", "a", "", "Request a specific architecture for example `x86_64`.")
	queueCmd.MarkFlagRequired("architecture")

	queueCmd.Flags().StringVarP(&QueueImageType, "image-type", "t", "", "Request a specific image-type for example `guest-image`.")
	queueCmd.MarkFlagRequired("image-type")

	queueCmd.Flags().StringVarP(&QueuePackages, "packages", "p", "", "Additional packages to install in `nginx,tmux` format.")

	queueCmd.Flags().StringVarP(&QueueUsers, "users", "u", "", "Additional users to add in `user:ssh-key,user:ssh-key` format.")
}
