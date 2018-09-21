// Copyright © 2018 Jason Lu <luhonghai@gmail.com>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package cmd

import (
	"fmt"
	"strconv"

	"github.com/luhonghai/wsdl-example/pkg/calculator"
	"github.com/spf13/cobra"
)

// calculateCmd represents the calculate command
var calculateCmd = &cobra.Command{
	Use:   "calculate",
	Short: "Simple number calculation command",
	Long: `Available command add, subtract, devine and multiply. For example:
				- wsdl-example calculate add 1 5
				- wsdl-example calculate subtract 5 2
				- wsdl-example calculate subtract 5 2
				- wsdl-example calculate devine 4 2
				- wsdl-example calculate multiply 2 7
		`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 3 {
			var err error
			service := calculator.NewCalculatorSoap("", false, nil)
			intA, err := strconv.ParseInt(args[1], 10, 32)
			intB, err := strconv.ParseInt(args[2], 10, 32)
			var result int32
			switch args[0] {
			case "add":
				resp, rErr := service.Add(&calculator.Add{IntA: int32(intA), IntB: int32(intB)})
				err = rErr
				result = resp.AddResult
				break
			case "subtract":
				resp, rErr := service.Subtract(&calculator.Subtract{IntA: int32(intA), IntB: int32(intB)})
				err = rErr
				result = resp.SubtractResult
				break
			case "devine":
				resp, rErr := service.Divide(&calculator.Divide{IntA: int32(intA), IntB: int32(intB)})
				err = rErr
				result = resp.DivideResult
				break
			case "multiply":
				resp, rErr := service.Multiply(&calculator.Multiply{IntA: int32(intA), IntB: int32(intB)})
				err = rErr
				result = resp.MultiplyResult
				break
			default:
				fmt.Println("Invalid method " + args[0])
				return
			}
			if err != nil {
				fmt.Println("Error", err)
			} else {
				fmt.Println("Result: ", result)
			}
		} else {
			fmt.Println("Require 3 args")
		}
	},
}

func init() {
	rootCmd.AddCommand(calculateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// calculateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// calculateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
