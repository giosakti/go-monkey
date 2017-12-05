// repl/repl.go

package repl

import (
  "bufio"
  "fmt"
  "io"
  "github.com/giosakti/go-monkey/lexer"
  "github.com/giosakti/go-monkey/token"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
  scanner := bufio.NewScanner(in)

  for {
    fmt.Printf(PROMPT)
    scanned := scanner.Scan()
    if !scanned {
      return
    }

    line := scanner.Text()
    l := lexer.New(line)

    for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
      fmt.Printf("%+v\n", tok)
    }
  }
}
