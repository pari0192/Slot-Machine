# Slot-Machine

It is a fast , interactive 3x3 terminal slot machine betting game written in Go. Players start with $200 , place custom bets , and test their luck against weighted random reels.
## Descritption 
Slot -Machine is a command line game built in Go . It simulates a 3 by 3 slot machine where symbols have different weights and multipliers.The game valudates bets , matches lines , and lets you crash out when done.

## Preview
- starting point

![Slot Machine Start Point](./screenshots/Screenshot%203.png)
- GamePlay
![Slot Machine gameplay](./screenshots/Screenshot%202.png)

## Getting Started 
### Dependencies

- Go 1.2.1 or higher
- Any OS with a terminal (Windows , macOS , Linux)

### Installing
1. Clone the repository:
``` git clone https://github.com/pari0192/Slot-Machine.git```
2. Navigate into the project folder:
```cd Slot-Machine```


### Run the game from the project root folder

## features

- weighted symbol pool
- Dynamic Multipliers
- Line Matching
- Smart bet validation
- Clean Cash Out

## Folder Strucute

```
main.go
spin.go
utils.go
go.mod
screenshots/
    Screenshot 1.png
    Screenshot 2.png
    Screenshot 3.png
README.md
```


## How to run

```go run .``` in root terminal of the project

## GamePlay
- start with **$200** balance
- Each round , enter your bet within your balance
- The machine spins a **3X3 grid** of symbols: `A` , `B` , `C` , `D`
- win by matching all 3 symbols in a row 
- each symbol has a different multiplier:
    - `A` -> 20x
    - `B` -> 10x
    - `C` -> 5x
    - `D` -> 2x
- winning = bet x multiplier per winning line
- Enter `0` as your bet to cash out and end the game

## Help
If you get a `command not found: go` error , make sure Go is installed :
- Download from https://go.dev/dl/
## Release

[![Release](https://img.shields.io/github/v/release/pari0192/Slot-Machine)](https://github.com/pari0192/Slot-Machine/releases/latest)

## License
MIT License © 2026 Pari. Free to use , modify and distribute.