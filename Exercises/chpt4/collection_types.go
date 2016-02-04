
package main

import (
	"fmt"
  "sort"
	"strings"
//	"math"
)

func main (){

     //Exercise 1.
     var numbers sort.IntSlice = []int{9,2,12,1,16,12,10, 12, 12, 9, 8, 9, 0, 11, 2, 1}
     numbers.Sort()  //numbers are sorted and can be used as is.
     result := UniqueInts(numbers)
     fmt.Println("Exercise 1\n", result)

     //Exercise 2.
     irregularMatrix := [][]int{{1,2,3,4},{5,6,7,8},{9,10,11},{12,13,14,15},{16,17,18,19,20}}
     slice := Flatten(irregularMatrix)

		 fmt.Println("Exercise 2")
     fmt.Printf("1x%d: %v\n", len(slice), slice)

     //Exercise 3.
     numbers_exer3 := []int{155,2,12,1,16,12,222,12,12,11,8,9,5,11,2,20}

     var column_count = 2
     var fraction int
		 var fraction_float float64

     fraction_float = float64(len(numbers_exer3)) / float64(column_count)
     fraction = len(numbers_exer3) / column_count

          if fraction_float != float64(int64(fraction_float)) {

          numbers_exer3 := Make2D_helper(fraction_float, numbers_exer3, column_count)
          numbers_exer3_result := Make2D(numbers_exer3, column_count, fraction + 1)
          fmt.Println("Exercise 3",column_count,numbers_exer3_result[column_count:])

        } else {
             numbers_exer3_result := Make2D(numbers_exer3, column_count, fraction)
             fmt.Println("Exercise 3\n",column_count,numbers_exer3_result[column_count:])
            }

	  //Exercise 4.

						iniData := []string {
						"#filter substitution",
						"[App]",
						";",
						"; This field specifies your organization's name.  This field is recommended,",
						"; but optional.",
						"Vendor=MozillaTest",
						";",
						"; This field specifies your application's name.  This field is required.",
						"Name=Simple",
						";",
						"; This field specifies your application's version.  This field is required.",
						"Version=0.1",
						";",
						"; This field specifies your application's build ID (timestamp).  This field is",
						"; required.",
						"BuildID=20070625",
						";",
						"; This field specifies a compact copyright notice for your application.  This",
						"; field is optional.",
						"Copyright=Copyright (c) 2004 Mozilla.org",
						";",
						"; This ID is just an example.  Every XUL app ought to have it's own unique ID.",
						"; You can use the microsoft 'guidgen' or 'uuidgen' tools, or go on",
						"; irc.mozilla.org and /msg botbot uuid.  This field is optional.",
						"ID={3aea3f07-ffe3-4060-bb03-bff3a5365e90}",
						"",
						"[Gecko]",
						";",
						"; This field is required.  It specifies the minimum Gecko version that this",
						"; application requires.",
						"MinVersion=@MOZILLA_VERSION_U@",
						";",
						"; This field is optional.  It specifies the maximum Gecko version that this",
						"; application requires.  It should be specified if your application uses",
						"; unfrozen interfaces.",
						"MaxVersion=@MOZILLA_VERSION_U@",
						"",
						"[Shell]",
						";",
						"; This field specifies the location of your application's main icon with file",
						"; extension excluded.  NOTE: Unix style file separators are required.  This",
						"; field is optional.",
						"Icon=chrome/icons/default/simple",
						}

				fmt.Println("Exercise 4,5")

				//Exercise 5

			//	PrintIni()

      map_data :=  ParseIni(iniData)
			PrintIni(map_data)
     }


	func PrintIni(ini map[string]map[string]string) {
	    groups := make([]string, 0, len(ini))
	    for group := range ini {
	        groups = append(groups, group)
	    }
	    sort.Strings(groups)
	    for i, group := range groups {
	        fmt.Printf("[%s]\n", group)
	        keys := make([]string, 0, len(ini[group]))
	        for key := range ini[group] {
	            keys = append(keys, key)
	        }
	        sort.Strings(keys)
	        for _, key := range keys {
	            fmt.Printf("%s=%s\n", key, ini[group][key])
	        }
	        if i+1 < len(groups) {
	            fmt.Println()
	        }
	    }
	}


	func ParseIni(iniData []string) (map[string]map[string]string) {

			result_final := make(map[string]map[string]string)

			var iniDataString string
    	var bracket_string string
    	var equals string
			const separator = "="
			var value string

        for _, line:= range iniData {
					iniDataString = line

				for _, line := range strings.Split(iniDataString, "\n") {
			  		line = strings.TrimSpace(line)
						if strings.HasPrefix(line, "[") && line != ";"  {
							bracket_string = line  //[App], [Gecko]..
							result_final[bracket_string] = make(map[string]string)
						}

					if j := strings.Index(line, separator); j > -1 && line != ";" {
  			 		equals = line[:j+len(separator)-1] //-1 additional position over from "="
	  				value = line[j+len(separator):]
            result_final[bracket_string][equals] = value
						}
        }
   	}

			/*	for z, x := range result_final {
			//	  if x["Vendor"] == "MozillaTest" {
					  		fmt.Println(z,"\n",x)
				//						}
        //	fmt.Println(result_final)
				}   */
      return result_final
}

func Make2D_helper(fraction_float float64, numbers_exer3 []int, column_count int) ([]int) {
      for {
        numbers_exer3 = append(numbers_exer3, 0) //temporary for prototype
				fraction_float = float64(len(numbers_exer3)) / float64(column_count)

				if fraction_float != float64(int64(fraction_float)) {
              return numbers_exer3
            }
       }
      return numbers_exer3
}


func Make2D(numbers_exer3 []int, column_count int, fraction int) [][]int {
     result := make([][]int, column_count)
     var count = []int{}
     value := 0
     var column_count_int = column_count
     count = numbers_exer3[value:len(numbers_exer3)] // set count = numbers_exer3[:3]
  		  for _, _ = range count[:fraction] {
            count = numbers_exer3[value:column_count]// set count = numbers_exer3[:3]
            result = append(result, count)
            value += column_count_int
            column_count += column_count_int
     }
    return result
}


func Flatten( irregularMatrix [][]int) []int {
     var result []int
     var matrix_x_result []int
     for _, x:= range irregularMatrix{
          matrix_x_result = x
     for _, x := range matrix_x_result {
         result = append(result, x)
     }
  }
      return result
}

func UniqueInts(numbers []int) []int {
     var result []int
     var previous int

for _, x := range numbers {
    if previous != x  {
       result = append(result, x)
        }
      previous = x
    }
return result
  }
