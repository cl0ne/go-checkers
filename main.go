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

// TODO: implement fancy input
type PlayerInput struct {
	name string
}

func (i PlayerInput) SelectChecker(availableCheckers []*checkers.Checker) int {
	fmt.Println(i.name, ", take a checker:")
	for i := range availableCheckers {
		fmt.Println(i+1, availableCheckers[i].String())
	}
	var checkerNumber int
	fmt.Scanln(&checkerNumber)
	return checkerNumber - 1
}

func (i PlayerInput) SelectTargetPos(availableMoves []checkers.Move) int {
	fmt.Println(i.name, ", where would you go?")
	for i := range availableMoves {
		fmt.Println(i+1, availableMoves[i].Target.String())
	}
	var targetNumber int
	fmt.Scanln(&targetNumber)
	return targetNumber - 1
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
