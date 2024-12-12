/*
Copyright © 2024 49pctber

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
	"io"
	"net/http"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

func getInput(year, day int) error {

	if year < 2015 {
		return fmt.Errorf("invalid year")
	}

	if day > 25 || day < 1 {
		return fmt.Errorf("invalid day")
	}

	url := fmt.Sprintf("https://adventofcode.com/%d/day/%d/input", year, day)

	session, exists := os.LookupEnv("ADVENT_OF_CODE_SESSION")
	if !exists {
		return fmt.Errorf("ADVENT_OF_CODE_SESSION environment variable does not exist")
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	cookie := &http.Cookie{
		Name:  "session",
		Value: session,
	}
	req.AddCookie(cookie)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	fname := fmt.Sprintf("%d_%d.txt", year, day)
	err = os.WriteFile(fname, body, 0666)
	if err != nil {
		return err
	}

	fmt.Println("saved", fname)

	return nil
}

var getinputCmd = &cobra.Command{
	Use:   "getday <year> <day>",
	Short: "Get input for a given day",
	Long:  `Get input for a given day`,
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {

		year, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Printf("error parsing year: %v\n", err)
			os.Exit(1)
		}

		day, err := strconv.Atoi(args[1])
		if err != nil {
			fmt.Printf("error parsing day: %v\n", err)
			os.Exit(1)
		}

		err = getInput(year, day)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println("Success!")
	},
}

func init() {
	rootCmd.AddCommand(getinputCmd)
}
