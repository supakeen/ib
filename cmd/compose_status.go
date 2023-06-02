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

	"github.com/supakeen/ib/lib/api"
	"github.com/spf13/cobra"
)

var StatusComposeID string

// statusCmd represents the status command
var statusCmd = &cobra.Command{
	Use:   "s",
	Short: "Get the status of an image build by its identifier",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(api.NewComposeStatusRequest(StatusComposeID))
	},
}

func init() {
	composeCmd.AddCommand(statusCmd)

	statusCmd.Flags().StringVarP(&StatusComposeID, "compose-id", "i", "", "Request a specific compose id.")
	statusCmd.MarkFlagRequired("compose-id")
}
