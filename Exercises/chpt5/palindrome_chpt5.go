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
	"path/filepath"
	"unicode/utf8"
)

var IsPalindrome func(string) bool // Holds a reference to a function

func init() {
	if len(os.Args) > 1 &&
		(os.Args[1] == "-a" || os.Args[1] == "--ascii") {
		os.Args = append(os.Args[:1], os.Args[2:]...) // Strip out arg.

		fmt.Println(os.Args[1:])

		IsPalindrome = func(s string) bool {

			for i := 0; i < len(s); i++ {
				a := s[(len(s)-1)-i : len(s)-i] // last digit
				b := s[0+i : 1+i]               // first digit

				if a != b {
					fmt.Println("not a palindrome")
					return false
				}
				fmt.Println(a)
				fmt.Println(b)
			}
			fmt.Println("a palindrome")
			return true
		}

	} else {
		//    s := os.Args[1]
		IsPalindrome = func(s string) bool {

			if utf8.RuneCountInString(s) <= 1 {
				fmt.Println(s, "a is a palindrome")
				return true
			}
			for i := 0; i < len(s); i++ {

				first, _ := utf8.DecodeRuneInString(s)
				last, _ := utf8.DecodeLastRuneInString(s)
				if first != last {
					fmt.Println(s, "is not a palindrome")
					return false
				}
			}
			fmt.Println(s, "is a palindrome")
			return true
		}

	}
}

func main() {
	if len(os.Args) == 1 {
		fmt.Printf("usage: %s [-a|--ascii] word1 [word2 [... wordN]]\n",
			filepath.Base(os.Args[0]))
		os.Exit(1)
	}
	words := os.Args[1:]
	for _, word := range words {
		fmt.Printf("%5t %q\n", IsPalindrome(word), word)
	}

}
