package main

import (
  "bufio"
  "os"
  "fmt"
  "strings"
  "os/exec"
)

func execInput(input string) error {
  input = strings.TrimSuffix(input, "\n")

  args := strings.Split(input, " ")

  cmd := exec.Command(args[0], args[1:]...)

  cmd.Stderr = os.Stderr
  cmd.Stdout = os.Stdout

  return cmd.Run()
}

func main() {
  reader := bufio.NewReader(os.Stdin)

  for {
    fmt.Print("(bananas) > ")
    input, err := reader.ReadString('\n')

    if err != nil {
      fmt.Fprintln(os.Stderr, err)
    }

    err = execInput(input)

    if err != nil {
      fmt.Fprintln(os.Stderr, err)
    }
  }
}
