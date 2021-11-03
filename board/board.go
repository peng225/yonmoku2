package board

type Board interface {
    Init(size int)
    Size() int
    Put(pos int) error
    CanPut(pos int) bool
    Reverse()
    Display()
    IsEnd() bool
    GetWinner() string
    GetTurn() string
}

