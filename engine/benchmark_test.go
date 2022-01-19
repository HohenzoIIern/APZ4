package engine

import (
  "fmt"
  "strings"
  "testing"
)

var constInput = "cat " //base function
var command Command

func BenchmarkCount(b *testing.B) {
  repeatNum := 1
  for i := 0; i < 20; i++ {
    repeatNum = 2 * repeatNum
    input := constInput
    input += strings.Repeat("abracadabrasomethingelse", repeatNum) + " " + strings.Repeat("abracadabrasomethingelse", repeatNum)

    b.Run(fmt.Sprintf("len=%d", 24*repeatNum), func(b *testing.B) {
      command = Parse(input)
    })
  }
}