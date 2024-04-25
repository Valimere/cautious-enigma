# cautious-enigma
Simplified Tetris calc

## Build process ssh
```shell
git pull git@github.com:Valimere/cautious-enigma.git
go build -o tetris
```
## Usage
```shell
Usage of ./tetris:
  -d	Debug Mode will print interpreted input file and grid after each placement
  -i string
    	Input File path (default "input.txt")

# Example run
./tetris -i input.txt
Case 1 Resulting Height:4
...
Case 10 Resulting Height:4

# Example debug run
./tetris -i input.txt -d
...
. . . . . . . . . #
. . . . . . # # # #
# # # # # # # # # #
. # # # . . # # # #
. . . # # . # # # #
---------------------------
Clearing row: 3
. . . . . . . . . #
. . . . . . # # # #
. # # # . . # # # #
. . . # # . # # # #
---------------------------
Grid during height calculation:
. . . . . . . . . #
. . . . . . # # # #
. # # # . . # # # #
. . . # # . # # # #
---------------------------
Case 10 Resulting Height:4
```

ChatGPT was used to help fill out some of the boilerplate, including expanding unit tests, and walking the grid (top down, bottom up, etc)

Built with go version: 1.21.1