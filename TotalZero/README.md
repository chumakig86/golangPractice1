# Go Console Input Echo Program

## Description

This is a simple Go program that continuously reads integers from standard input and echoes them back to the console. The program terminates when the user inputs the number `0`.

## How It Works

- The program uses a `for` loop to continuously read integers from standard input using `fmt.Scan`.
- If the input is `0`, the loop breaks and the program exits.
- Otherwise, the program prints the input value to the console.

## Example

```bash
$ go run TotalZero.go
5
5
12
12
0
