package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
  file, err := os.ReadFile("input")
  if err != nil {
    log.Fatal(err)
  }

  s := bufio.NewScanner(bytes.NewReader(file))
  
  var safereports []bool

  for i := 0; s.Scan(); i++ {
    
    report := strings.Fields(s.Text())
    
    var reportNums []int
    
    for _, e := range report {
      num, err := strconv.Atoi(e) 
      if err != nil {
        log.Fatal(err)
      }
      reportNums = append(reportNums, num)
    } 
    
    safelen := len(safereports)
   
    for num := 0; num != len(reportNums); num++  {
      var x []int
      for i, e := range reportNums{
        if i == num {
          continue
        } else {
          x = append(x, e)    
        }
      }
      safe := safeCalc(x, safereports)
      if safe {
        safereports = append(safereports, true)
        break
      } 

    }
  
    if safelen == len(safereports) {
      safereports = append(safereports, false)
    }

    fmt.Println(reportNums)
    fmt.Println(safereports)
    var t int
    for _, e := range safereports {
      if e {
        t += 1
      }
    }
    fmt.Println(t)

  }
}


func safeCalc(reportNums []int, safereports []bool) bool{

  if reportNums[0] < reportNums[1] {
  } else {
    slices.Reverse(reportNums)
  }
  
  s := slices.IsSorted(reportNums)
  safelen := len(safereports)
    
  if s {
    for i := range reportNums {
      if i > 0 {
        
        if reportNums[i] == 0 {
          continue
        }

        if reportNums[i] == reportNums[i - 1] {
          safereports = append(safereports, false)
          break
        }

        if reportNums[i] - reportNums[i - 1] > 3 {
          safereports = append(safereports, false)
          break
        }

      }
    }
  } else {
    safereports = append(safereports, false)
  }

  if safelen == len(safereports) {
    safereports = append(safereports, true)
  }

  return safereports[len(safereports) - 1]
}
