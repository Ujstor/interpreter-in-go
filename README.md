# Go Interpreter

This repository contains the source code for a custom interpreter written in Go. The interpreter is designed to parse and evaluate the Monkey programming language. C-like language with features such as variable bindings, arithmetic expressions, built-in functions, first-class and higher-order functions, closures and complex data structures like strings, arrays, and hash.

This project is an implementation of a programming language interpreter, as presented in Thorsten Ball's book: ["Writing An Interpreter In Go"](https://interpreterbook.com/).

## Features

- **Lexer and Parser**: Custom-built lexer and parser to handle the syntax of Monkey.
- **Abstract Syntax Tree (AST)**: Building and evaluating AST for interpreting the Monkey language.
- **Evaluator**: A core component that interprets the AST.
- **Data Structures**: Support for integers, booleans, strings, arrays, and hashes.
- **Functions**: First-class and higher-order functions with closures.
- **REPL**: A Read-Eval-Print Loop for interactive execution.

## Getting Started

### Prerequisites

- Go programming language (version 1.14 or higher recommended)
- Basic understanding of compilers and interpreters

### Installation

1. Clone the repository:
   ```sh
   git clone https://github.com/Ujstor/interpreter-in-go.git
   ```

2. Navigate to the project directory:
   ```sh
   cd go-interpreter
   ```

### Running the Interpreter

To launch the REPL:

```sh
go run main.go
```

## Usage Examples

### Variable Bindings

```monkey
let age = 30;
let name = "Monkey";
let result = 10 * (20 / 2);
```

### Working with Arrays and Hashes

```monkey
let myArray = [1, 2, 3, 4, 5];
let userInfo = {"name": "Jane", "age": 25};
```

### Functions

```monkey
let add = fn(a, b) { return a + b; };
add(10, 20);

let fibonacci = fn(x) {
    if (x == 0) {
        0
    } else {
        if (x == 1) {
            1
        } else {
            fibonacci(x - 1) + fibonacci(x - 2);
        }
    }
};
fibonacci(10);
```

### Higher-Order Functions

```monkey
let twice = fn(f, x) { return f(f(x)); };
let addTwo = fn(x) { return x + 2; };
twice(addTwo, 2); // => 6
```