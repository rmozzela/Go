// Copyright © 2011-12 Qtrac Ltd.
//
// This program or package and any associated files are licensed under the
// Apache License, Version 2.0 (the "License"); you may not use these files
// except in compliance with the License. You can get a copy of the License
// at: http://www.apache.org/licenses/LICENSE-2.0.
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"fmt"
	"os"
	//    "path/filepath"
	//    "unicode/utf8"
	"bytes"
)

func CommonPrefix(s string) {

	test := "/home/user/prefix"
	result := make([]rune, len(test))

	var b bytes.Buffer // A Buffer needs no initialization.
	n := 0
	i := 0
	var equal_not_equal string
	// compare utf value; write to utf value to buffer ==
	// convert utf value to string value; print out string value

	//test2 = s[0]

	for _, sdata := range test {
		//  utf8.DcodeRuneInString(s[n:])
		result[0] = sdata
		//fmt.Println("sdata",result[0])
		n++

		for i == n-1 {
			if result[0] == rune(s[i]) {
				b.WriteRune(result[0])
				equal_not_equal = "=="

			} else {
				equal_not_equal = "!="
				break
			}
			i++
		}
	}
	if equal_not_equal == "!=" {
		fmt.Printf("Exited on:--->")
		b.WriteTo(os.Stdout)
		fmt.Println("  ", test, equal_not_equal, s)

		///  fmt.Println()
	} else {
		fmt.Printf("Prefix is Equal:")
		fmt.Println(test, equal_not_equal, s)
	}
	//b.WriteTo(os.Stdout)
	//return test

}

func main() {

	iniData := []string{
		"/home/user/prefix",
		"/home/#/prefix",
		"/user/timestamp",
		"/home/user/prefix/dolphin",
		"/home/user/dolphin",
		"/home/test/prefix",
	}

	for _, files := range iniData {
		words := files
		//  fmt.Println(CommonPrefix(words))
		CommonPrefix(words)
	}

}
