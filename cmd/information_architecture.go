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
	"log"
)

var Distribution string

// architectureCmd represents the architecture command
var architectureCmd = &cobra.Command{
	Use:   "a",
	Short: "List available architectures for a distribution",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if len(Distribution) == 0 {
			log.Fatal("missing d!")
		}

		for _, architecture := range api.NewArchitecturesRequest(Distribution) {
			fmt.Printf("%s -- %s\n", architecture.Name, architecture.ImageTypes)
		}
	},
}

func init() {
	informationCmd.AddCommand(architectureCmd)

	architectureCmd.Flags().StringVarP(&Distribution, "distribution", "d", "", "Request a specific distribution.")
	architectureCmd.MarkFlagRequired("distribution")
}
