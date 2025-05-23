package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/hawkaii/monkey_go/evaluator"
	"github.com/hawkaii/monkey_go/lexer"
	"github.com/hawkaii/monkey_go/object"
	"github.com/hawkaii/monkey_go/parser"
)

const PROMPT = ">> "
const MONKEY_FACE = `
	      .-"      "-.
	     /            \
	    |              |
	    |,  .-.  .-. ,|
	    | )(_o/  \o_)( |
	    |/     /\     \|
	    (_     ^^     _)
	     \__|""""""|__/
	      |  |    |  |
	      \_/      \_/

	`

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	env := object.NewEnvironement()

	for {
		fmt.Print(PROMPT)
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

		evaluated := evaluator.Eval(program, env)
		if evaluated != nil {

			io.WriteString(out, evaluated.Inspect())
			io.WriteString(out, "\n")
		}

	}
}

func printParserErrors(out io.Writer, errors []string) {
	io.WriteString(out, MONKEY_FACE)
	io.WriteString(out, "Whoops! maybe u did some idiot thing, kawaii!\n")
	io.WriteString(out, "	parser errors: \n")

	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
