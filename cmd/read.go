/*
Copyright © 2023 Fabrizio Gortani <fabrizio.gortani@pm.me>

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
	"os"

	"github.com/CasvalDOT/akichat-cli/presenters"
	core "github.com/CasvalDOT/akichat-core"

	"github.com/spf13/cobra"
)

var readCmd = &cobra.Command{
	Use:   "read",
	Short: "Read the messages",
	Long:  "Command for read the current messages in the public chat",
	Run: func(cmd *cobra.Command, args []string) {
		c := core.NewChat(core.ChatTypeHentakihabara)
		messages, err := c.ReadMessages("0")
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		p := presenters.NewMessagesPresenter(messages)
		p.Print()
	},
}

func init() {
	rootCmd.AddCommand(readCmd)
}
