# Knight Move Validator

## Description
This Go program checks whether a knight's move on an 8×8 chessboard is valid. A knight moves in an "L" shape, either:
- **Two squares horizontally and one square vertically**, or
- **Two squares vertically and one square horizontally**.

The program reads four integers representing the knight's current position and the target position and determines whether the move is legal.

## Usage

### Input
The program takes four natural numbers as input:
- `X1, Y1` – The coordinates of the knight's current position.
- `X2, Y2` – The coordinates of the target position.

It is guaranteed that the input values are within the bounds of an 8×8 chessboard.

### Output
The program prints:
- `true` if the move is valid.
- `false` if the move is invalid.

## Example

### Input
```
1 
1 
3 
2
```
### Output
```
true
```

### Input
```
1 
1 
2 
2
```
### Output
```
false
```

## How to Run
Ensure you have Go installed, then compile and run the program:

```sh
go run KnightMove.go
```

Provide the input when prompted and check the output for move validation.

