package main

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

var s1min []int
var s2min []int

func main() {
  file, err := os.ReadFile("input")
  if err != nil {
    log.Fatal(err)
  }
  s :=  strings.Fields(string(file)) 
  s1, s2 := SeperateSlice(s)

  var score int
  for _, e := range s1 {
    s := SimilarList(e,s2)
    score += s
  }

  for len(s1) > 0 { 
    s1 = GetMin(s1, &s1min)
    s2 = GetMin(s2, &s2min)
  }


  fmt.Println(score) 
  length := TotalDistance(s1min, s2min)
  fmt.Println(length)
  
}
func SimilarList(num string, s []string) int {
  var score int 
  for _, e := range s {
    if num == e  {
      score += 1
      fmt.Println("score", num, e)
    }
  } 

  x, err := strconv.Atoi(num)
  if err != nil {
    log.Fatal(err)
  }
  return score * x

}

func SeperateSlice(s []string) ([]string, []string) {
  var s1 []string
  var s2 []string

  for i,e := range s {
    if i % 2 == 0 {
      s1 = append(s1, e)
    } else {
      s2 = append(s2, e)
    }
  }

  return s1, s2
}


func GetMin(s []string, min *[]int) []string {
  Min := 100000
  MinIndex := 0


  for i, e := range s {
    num, err := strconv.Atoi(e)
    
    if err != nil {
      log.Fatal("failed to convert int",err)
    }

    if num < Min {
      Min = num
      MinIndex = i
    }
  
  }
  
  *min = append(*min, Min)
  s = slices.Delete(s, MinIndex, MinIndex + 1)
  return s
}

func TotalDistance(s1, s2 []int) int {
  var length int 
  for i := range s1 {
    if s1[i] < s2[i] {
      length += s2[i] - s1[i]
    } else  {
      length += s1[i] - s2[i]
    }
  }
  return length
}
