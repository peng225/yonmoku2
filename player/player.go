package player

import "yonmoku2/board"

type Player interface {
    Search(board *board.Board) int
    Time(timeInMs int64) error
}

