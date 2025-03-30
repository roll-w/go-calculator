# Go Calculator

This is a simple calculator written in Go. It can perform basic arithmetic 
operations like addition, subtraction, multiplication, and division, also
it can calculate the square root or logarithm of a number.

## Build

To build the calculator, you need to have Go installed on your machine.
Then, you can run the following command:

```bash
go build -o calculator
```

## Usage

To use the calculator, you can run the following command:

```bash
./calculator <expression>
```

Where `<expression>` is a string that represents an arithmetic operation, like
`2 + 2` or `sqrt(4)`. The calculator will evaluate the expression and print
the result.

Or you can run the following command:

```bash
./calculator
```

And then you can enter the expression interactively.

## Examples

```bash
$ ./calculator 2+2
4

$ ./calculator (1+2)*sqrt(4)-log(1)+3!*2^2
14
```

## License

```text
MIT License

Copyright (c) 2025 RollW

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
```
