package board

import(
    "fmt"
    "errors"
)

type STATE int

const (
    EMPTY = iota
    RED = iota
    YELLOW = iota
)

type YonmokuBoard struct {
    size int
    turn STATE
    board []STATE
    history []int
}

func (yb *YonmokuBoard) Size() int {
    return yb.size
}

func (yb *YonmokuBoard) Init(size int) {
    yb.size = size
    yb.turn = RED
    yb.board = make([]STATE, yb.size*yb.size)
    yb.history = make([]int, 0)

    for i := range yb.board {
        yb.board[i] = EMPTY
    }
}

func (yb *YonmokuBoard) Put(pos int) error {
    if pos < 0 || yb.size <= pos {
        return errors.New("The pos is out of range.")
    }
    if yb.board[pos] != EMPTY {
        return errors.New("The pos is already full.")
    }

    isPut := false
    for i := yb.size-1; i >= 0; i-- {
        if yb.board[i*yb.size + pos] == EMPTY {
            yb.board[i*yb.size + pos] = yb.turn
            isPut = true
            break
        }
    }
    if(!isPut) {
        panic("Logic error")
    }
    yb.history = append(yb.history, pos)
    yb.changeTurn()
    return nil
}

func (yb *YonmokuBoard) Reverse() {
    if len(yb.history) == 0 {
        return
    }
    pos := yb.history[len(yb.history)-1]
    yb.history = yb.history[:len(yb.history)-1]
    for i := 0; i < yb.size; i++ {
        if yb.board[i*yb.size + pos] != EMPTY {
            yb.board[i*yb.size + pos] = EMPTY
            yb.changeTurn()
            break
        }
    }
}

func (yb *YonmokuBoard) Display() {
    for i:= 0; i < yb.size; i++ {
        fmt.Print(" ", i)
    }
    fmt.Println("")
    for i:= 0; i < yb.size; i++ {
        fmt.Print("|")
        for j:= 0; j < yb.size; j++ {
            switch(yb.board[yb.size*i + j]) {
                case EMPTY:
                    fmt.Print(" ")
                case RED:
                    fmt.Print("r")
                case YELLOW:
                    fmt.Print("y")
            }
            fmt.Print("|")
        }
        fmt.Println("")
    }
    yb.printTurn()
    fmt.Println()
}

func (yb *YonmokuBoard) changeTurn() {
    switch(yb.turn) {
        case RED:
            yb.turn = YELLOW
        case YELLOW:
            yb.turn = RED
        default:
            panic("Invalid turn")
    }
}

func (yb *YonmokuBoard) printTurn() {
    fmt.Print("turn: ")
    switch(yb.turn) {
        case RED:
            fmt.Println("RED")
        case YELLOW:
            fmt.Println("YELLOW")
        default:
            panic("Invalid turn")
    }
}


func (yb *YonmokuBoard) IsEnd() bool {
    if yb.GetWinner() != "EMPTY" {
        return true
    }

    for i := 0; i < yb.size; i++ {
        if yb.board[i] == EMPTY {
            return false
        }
    }
    return true
}


func stateToString(st STATE) string {

    switch(st) {
        case EMPTY:
            return "EMPTY"
        case RED:
            return "RED"
        case YELLOW:
            return "YELLOW"
        default:
            panic("Invalid state")
    }
}

func (yb *YonmokuBoard) CanPut(pos int) bool {
    if pos < 0 || yb.size <= pos {
        return false
    }
    if yb.board[pos] != EMPTY {
        return false
    }
    return true
}

func (yb *YonmokuBoard) GetWinner() string {
    return stateToString(yb.getWinner())
}

func (yb *YonmokuBoard) getWinner() STATE {
    // Check horizontal line
    for row := 0; row < yb.size; row++ {
        count := 0
        var prevState STATE = EMPTY
        for col := 0; col < yb.size; col++ {
            currentState := yb.board[row*yb.size + col]
            if currentState == EMPTY {
                count = 0
            } else if prevState == EMPTY || currentState == prevState {
                count++
            } else {
                count = 1
            }
            if count == 4 {
                return currentState
            }
            prevState = currentState
        }
    }

    // Check vertical line
    for col := 0; col < yb.size; col++ {
        count := 0
        var prevState STATE = EMPTY
        for row := 0; row < yb.size; row++ {
            currentState := yb.board[row*yb.size + col]
            if currentState == EMPTY {
                count = 0
            } else if prevState == EMPTY || currentState == prevState {
                count++
            } else {
                count = 1
            }
            if count == 4 {
                return currentState
            }
            prevState = currentState
        }
    }

    // Check top-left to bottom-right line
    for row := 0; row < yb.size-3; row++ {
        count := 0
        var prevState STATE = EMPTY
        for r, c := row, 0; r < yb.size && c < yb.size; r, c = r+1, c+1 {
            currentState := yb.board[r*yb.size + c]
            if currentState == EMPTY {
                count = 0
            } else if prevState == EMPTY || currentState == prevState {
                count++
            } else {
                count = 1
            }
            if count == 4 {
                return currentState
            }
            prevState = currentState
        }
    }
    for col := 1; col < yb.size-3; col++ {
        count := 0
        var prevState STATE = EMPTY
        for r, c := 0, col; r < yb.size && c < yb.size; r, c = r+1, c+1 {
            currentState := yb.board[r*yb.size + c]
            if currentState == EMPTY {
                count = 0
            } else if prevState == EMPTY || currentState == prevState {
                count++
            } else {
                count = 1
            }
            if count == 4 {
                return currentState
            }
            prevState = currentState
        }
    }

    // Check top-right to bottom-left line
    for row := 0; row < yb.size-3; row++ {
        count := 0
        var prevState STATE = EMPTY
        for r, c := row, yb.size-1; r < yb.size && 0 <= c; r, c = r+1, c-1 {
            currentState := yb.board[r*yb.size + c]
            if currentState == EMPTY {
                count = 0
            } else if prevState == EMPTY || currentState == prevState {
                count++
            } else {
                count = 1
            }
            if count == 4 {
                return currentState
            }
            prevState = currentState
        }
    }
    for col := yb.size-2; 3 <= col; col-- {
        count := 0
        var prevState STATE = EMPTY
        for r, c := 0, col; r < yb.size && c < yb.size; r, c = r+1, c-1 {
            currentState := yb.board[r*yb.size + c]
            if currentState == EMPTY {
                count = 0
            } else if prevState == EMPTY || currentState == prevState {
                count++
            } else {
                count = 1
            }
            if count == 4 {
                return currentState
            }
            prevState = currentState
        }
    }

    return EMPTY
}

func (yb *YonmokuBoard) GetTurn() string {
    switch(yb.turn) {
        case RED:
            return "RED"
        case YELLOW:
            return "YELLOW"
        default:
            panic("Invalid turn")
    }
}


func (yb *YonmokuBoard) Copy() Board {
    var copiedYb YonmokuBoard = *yb
    copiedYb.board = make([]STATE, len(yb.board))
    copiedYb.history = make([]int, len(yb.history))
    copy(copiedYb.board, yb.board)
    copy(copiedYb.history, yb.history)
    return &copiedYb
}

