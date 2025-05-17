# Monkey Interpreter in Go

A Go implementation of the Monkey programming language, built by following Thorsten Ball’s book [Writing An Interpreter In Go](https://interpreterbook.com/).

## Overview

This project is an educational interpreter for the Monkey language, developed step-by-step as described in Thorsten Ball’s book. The goal is to learn about programming language design and interpreter implementation in Go.

Currently, the project covers the material from chapters 1–3 of the book:
- A fully functional lexer that tokenizes Monkey source code.
- A simple REPL (Read-Eval-Print Loop) for interactive exploration.
- Recognition of basic language constructs: identifiers, integer literals, operators, delimiters, and keywords.

## Features

- **Lexer:** Tokenizes input into Monkey language tokens.
- **REPL:** Interactive shell to type Monkey statements and see their tokens.
- **Basic Token Support:** Identifiers, numbers, operators, delimiters, and keywords.
- **Error Handling:** Reports unknown or illegal tokens.

## Getting Started

### Prerequisites

- [Go](https://golang.org/dl/) (version 1.18+ recommended)

### Building

Clone the repository and build the interpreter:
```sh
git clone https://github.com/hawkaii/monkey_go.git
cd monkey_go
go build -o monkey
```

### Running

Start the interactive REPL:
```sh
./monkey
```
Type Monkey language code and see the output!

## Project Structure

- `lexer/` — The Monkey lexer implementation.
- `token/` — Definitions for Monkey language tokens.
- `repl/` — The interactive Read-Eval-Print Loop.
- `main.go` — Entry point for the interpreter.

## Example

```text
>> let five = 5;
[LET] [IDENT] [ASSIGN] [INT] [SEMICOLON]
```

## References

- [Writing An Interpreter In Go](https://interpreterbook.com/) by Thorsten Ball

## Roadmap

- [x] Chapters 1–3: Lexer, tokens, and REPL
- [x] Chapters 4: Parser, AST, and evaluation
- [ ] Extending the Interpreter

## License

### MIT License

Copyright (c) 2025 hawkaii

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.

---
Happy hacking! 🚀
