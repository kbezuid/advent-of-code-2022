package game

import "fmt"

const oppoR = 'A'
const oppoP = 'B'
const oppoS = 'C'

const respR = 'X'
const respP = 'Y'
const respS = 'Z'

const (
	LoseOutcome = 'X'
	DrawOutcome = 'Y'
	WinOutcome  = 'Z'
)

const (
	Rock     = 1
	Paper    = 2
	Scissors = 3
)

const (
	LostScore = 0
	DrawScore = 3
	WonScore  = 6
)

type Game struct {
	TotalScore int
}

type play struct {
	score  int
	option int
}

func newDrawMap() map[byte]byte {
	var m = make(map[byte]byte)

	m[oppoR] = respR
	m[oppoP] = respP
	m[oppoS] = respS

	return m
}

func newWinMap() map[byte]byte {
	var m = make(map[byte]byte)

	m[oppoR] = respP
	m[oppoP] = respS
	m[oppoS] = respR

	return m
}

func newLoseMap() map[byte]byte {
	var m = make(map[byte]byte)

	m[oppoR] = respS
	m[oppoP] = respR
	m[oppoS] = respP

	return m
}

func getInputMap(outcome byte) map[byte]byte {
	switch outcome {
	case DrawOutcome:
		return newDrawMap()
	case WinOutcome:
		return newWinMap()
	case LoseOutcome:
		return newLoseMap()
	default:
		panic("Should not get here")
	}
}

func newRock() play {
	return play{
		score:  Rock,
		option: Rock,
	}
}

func newPaper() play {
	return play{
		score:  Paper,
		option: Paper,
	}
}

func newScissors() play {
	return play{
		score:  Scissors,
		option: Scissors,
	}
}

func getOpponentPlay(opponent byte) play {
	if opponent == oppoR {
		return newRock()
	}

	if opponent == oppoP {
		return newPaper()
	}

	if opponent == oppoS {
		return newScissors()
	}

	fmt.Printf("Panic %c", opponent)
	panic("Should not get here")
}

func getResponsePlay(response byte) play {
	if response == respR {
		return newRock()
	}

	if response == respP {
		return newPaper()
	}

	if response == respS {
		return newScissors()
	}

	panic("Should not get here")
}

func (g *Game) PlayRoundWithOutcome(opponent byte, outcome byte) {
	responseMap := getInputMap(outcome)
	g.PlayRound(opponent, responseMap[opponent])
}

func (g *Game) PlayRound(opponent byte, response byte) {
	oppoPlay := getOpponentPlay(opponent)
	respPlay := getResponsePlay(response)

	fmt.Printf("Play Score %d\n", respPlay.score)
	g.TotalScore += respPlay.score

	if oppoPlay.option == respPlay.option {
		g.TotalScore += DrawScore
		return
	}

	result := LostScore

	if oppoPlay.option == Rock {
		if respPlay.option == Paper {
			result = WonScore
		} else if respPlay.option == Scissors {
			result = LostScore
		}
	}

	if oppoPlay.option == Paper {
		if respPlay.option == Scissors {
			result = WonScore
		} else if respPlay.option == Rock {
			result = LostScore
		}
	}

	if oppoPlay.option == Scissors {
		if respPlay.option == Rock {
			result = WonScore
		} else if respPlay.option == Paper {
			result = LostScore
		}
	}

	g.TotalScore += result
	fmt.Printf("Result Score %d\n", result)
	fmt.Printf("Total %d\n", g.TotalScore)
}

func NewGame() Game {
	return Game{
		TotalScore: 0,
	}
}
