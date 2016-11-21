package main

import (
	"fmt"

	"github.com/cl0ne/go-checkers/checkers"
	"github.com/fatih/color"
)

var bg = color.New(color.BgBlack, color.FgHiWhite).SprintFunc()
var empty = color.New(color.FgBlack, color.BgHiWhite).SprintFunc()

var checker = map[bool]map[bool]string{
	false: {false: "▪ ", true: "◈ "},
	true:  {false: "▫ ", true: "◇ "},
	// another style:
	// false: {false: "⛂ ", true: "⛃ "},
	// true:  {false: "⛀ ", true: "⛁ "},
}

func drawBoard(b *checkers.Board) {
	var header = "a b c d e f g h i j k l m n o p q r s t u v w x y z "
	fmt.Print()
	fmt.Println(" ", header[0:b.Size()*2])
	for r := b.Size(); r > 0; r-- {
		fmt.Print(r, " ")
		for c := 0; c < b.Size(); c++ {
			if b.IsWhiteSquare(checkers.Point{c, r - 1}) {
				fmt.Print(bg("  "))
			} else {
				ch := b.GetChecker(c, r-1)
				if ch == nil {
					fmt.Print(empty("  "))
				} else {
					fmt.Print(empty(checker[ch.IsWhite()][ch.IsQueen()]))
				}
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

type PlayerInput struct {
	name string
}

func (i PlayerInput) SelectChecker(availableCheckers []*checkers.Checker) int {
	var header = "abcdefghijklmnopqrstuvwxyz"
	for {
		fmt.Printf("%s, take a checker:\n", i.name)
		for i := range availableCheckers {
			var posX = availableCheckers[i].Position().X
			var posY = availableCheckers[i].Position().Y
			if availableCheckers[i].IsWhite() {
				if availableCheckers[i].IsQueen() {
					fmt.Printf("%d. White queen on %c%d\n", i+1, header[posX], posY+1)
				} else {
					fmt.Printf("%d. White on %c%d\n", i+1, header[posX], posY+1)
				}
			} else {
				if availableCheckers[i].IsQueen() {
					fmt.Printf("%d. %s queen on %c%d\n", i+1, empty("Black"), header[posX], posY+1)
				} else {
					fmt.Printf("%d. %s on %c%d\n", i+1, empty("Black"), header[posX], posY+1)
				}
			}
		}
		var checkerNumber int
		_, ok := fmt.Scanln(&checkerNumber)
		if ok == nil && (checkerNumber-1) >= 0 && (checkerNumber-1) < len(availableCheckers) {
			return checkerNumber - 1
		} else {
			if ok != nil {
				var discard string
				fmt.Scanln(&discard)
			}
			fmt.Printf("Nope, dude. This variant doesn't exist! Pick another\n")
			continue
		}
	}
}

func (i PlayerInput) SelectTargetPos(availableMoves []checkers.Move) int {
	var header = "abcdefghijklmnopqrstuvwxyz"
	for {
		fmt.Printf("%s, where would you go?\n", i.name)
		for i := range availableMoves {
			var posX = availableMoves[i].Target.X
			var posY = availableMoves[i].Target.Y
			fmt.Printf("%d. To %c%d\n", i+1, header[posX], posY+1)
		}
		var targetNumber int
		_, ok := fmt.Scanln(&targetNumber)
		if ok == nil && (targetNumber-1) >= 0 && (targetNumber-1) < len(availableMoves) {
			return targetNumber - 1
		} else {
			if ok != nil {
				var discard string
				fmt.Scanln(&discard)
			}
			fmt.Printf("Nope, dude. This variant doesn't exist! Pick another\n")
			continue
		}
	}
}

func main() {
	inputs := []checkers.PlayerPoller{
		&PlayerInput{"White"},
		&PlayerInput{"Black"},
	}
	g := checkers.NewGame(inputs[0], inputs[1])
	g.Start()

	drawBoard(g.GetBoard())
	for !g.IsFinished() {
		if g.IsWhitesMove() {
			fmt.Println("White moves:")
		} else {
			fmt.Println(empty("Black"), "moves:")
		}
		g.Update()
		drawBoard(g.GetBoard())
		fmt.Println()
	}
	if g.IsBlackWin() {
		fmt.Println(empty("Black"), "wins")
	} else {
		fmt.Println("White wins")
	}
}
