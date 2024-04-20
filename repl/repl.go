package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/WillBallentine/bark/lexer"
	"github.com/WillBallentine/bark/parser"
	//"github.com/WillBallentine/bark/token"
)

const PROMPT = ">> "
const doggo = `   
   / \__
  (    @\___
  /         O
 /   (_____/
/_____/   U
`

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
		p := parser.New(l)

		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			printParserErrors(out, p.Errors())
			continue
		}

		io.WriteString(out, program.String())
		io.WriteString(out, "\n")
	}
}

func printParserErrors(out io.Writer, errors []string) {
	io.WriteString(out, doggo)
	io.WriteString(out, "Woof! Looks like we got tangled up in our leash there!\n")
	io.WriteString(out, " parser errors:\n")
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
