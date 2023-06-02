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
	"strings"

	"github.com/osbuild/ib/lib/api"
	"github.com/spf13/cobra"
)

var QueueDistribution string
var QueueArchitecture string
var QueueImageType string
var QueuePackages string

// queueCmd represents the compose command
var queueCmd = &cobra.Command{
	Use:   "q",
	Short: "Queue an image build",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(api.NewComposeRequest(QueueDistribution, QueueArchitecture, QueueImageType, "demo", strings.Split(QueuePackages, ",")))
	},
}

func init() {
	composeCmd.AddCommand(queueCmd)

	queueCmd.Flags().StringVarP(&QueueDistribution, "distribution", "d", "", "Request a specific distribution.")
	queueCmd.MarkFlagRequired("distribution")

	queueCmd.Flags().StringVarP(&QueueArchitecture, "architecture", "a", "", "Request a specific architecture.")
	queueCmd.MarkFlagRequired("architecture")

	queueCmd.Flags().StringVarP(&QueueImageType, "image-type", "t", "", "Request a specific image-type.")
	queueCmd.MarkFlagRequired("image-type")

	queueCmd.Flags().StringVarP(&QueuePackages, "packages", "p", "", "Comma-separated list of additional packages to install.")
}
