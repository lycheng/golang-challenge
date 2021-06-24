package design

import "fmt"

// SnakesLadders - Stucture that game is played on
//   from https://www.codewars.com/kata/587136ba2eefcb92a9000027/train/go
type SnakesLadders struct {
	gameIsOver bool
	player     int
	location   map[int]int
	transfers  map[int]int
}

// NewGame - Method that starts a new game for the SnakesLadders struct
func (sl *SnakesLadders) NewGame() {
	sl.player = 0
	sl.gameIsOver = false
	sl.location = map[int]int{0: 0, 1: 0}
	sl.transfers = map[int]int{
		2: 38, 7: 14, 8: 31, 15: 26, 21: 42, 28: 84, 36: 44, 51: 67, 71: 91, 78: 98, 87: 94,
		99: 80, 95: 75, 92: 88, 89: 68, 74: 53, 64: 60, 62: 19, 49: 11, 46: 25, 16: 6,
	}
}

// Play - Method that makes turn givem a doce roll from die1 and die2 for the SnakesLadders struct
// Player that dice is for should automatically be determined based on rules
func (sl *SnakesLadders) Play(die1, die2 int) string {
	if sl.gameIsOver {
		return "Game over!"
	}
	defer func() {
		if die1 != die2 {
			sl.player = (sl.player + 1) % 2
		}
	}()
	next := sl.location[sl.player] + (die1 + die2)
	if next > 100 {
		next = 200 - next
	}

	if n := sl.transfers[next]; n > 0 {
		next = n
	}

	sl.location[sl.player] = next
	if next == 100 {
		sl.gameIsOver = true
		return fmt.Sprintf("Player %d Wins!", sl.player+1)
	}
	return fmt.Sprintf("Player %d is on square %d", sl.player+1, next)
}
