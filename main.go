package main

import(
    "yonmoku2/board"
    "yonmoku2/player"
    "fmt"
    "bufio"
    "strconv"
    "strings"
    "os"
)

const LINE_SIZE = 1000

var rdr = bufio.NewReaderSize(os.Stdin, LINE_SIZE)
func readLine() string {
    buf := make([]byte, 0, LINE_SIZE)
    for {
        l, p, e := rdr.ReadLine()
        if e != nil {
            panic(e)
        }
        buf = append(buf, l...)
        if !p {
            break
        }
    }
    return string(buf)
}

func main() {
    var bd board.Board
    bd = &board.YonmokuBoard{}
    bd.Init(8)
    bd.Display()

    var pl player.Player
    pl = &player.McPlayer{}
    pl.Time(200)

mainLoop:
    for {
        fmt.Print("command> ")
        input := readLine()
        if input == "" {
            continue
        }
        tokens :=  strings.Split(input, " ")

        command := tokens[0]

        switch(command) {
            case "q", "quit":
                break mainLoop
            case "i", "init":
                if len(tokens) < 2 || tokens[1] == "" {
                    fmt.Println("usage: i(init) size")
                    break
                }
                size, err := strconv.Atoi(tokens[1])
                if err != nil {
                    fmt.Println(err)
                    break
                }
                bd.Init(size)
                bd.Display()
            case "d", "display":
                bd.Display()
            case "p", "put":
                if len(tokens) < 2 || tokens[1] == "" {
                    fmt.Println("usage: p(put) pos")
                    break
                }
                if bd.IsEnd() {
                    fmt.Println("The game is over.")
                    break
                }
                pos, err := strconv.Atoi(tokens[1])
                if err != nil {
                    fmt.Println(err)
                    break
                }
                err = bd.Put(pos)
                if err != nil {
                    fmt.Println(err)
                    break
                }
                bd.Display()
                if bd.IsEnd() {
                    fmt.Println("Winner: ", bd.GetWinner())
                }
            case "r", "reverse":
                bd.Reverse()
                bd.Display()
            case "s", "search":
                if bd.IsEnd() {
                    fmt.Println("The game is over.")
                    break
                }
                pos := pl.Search(&bd)
                err := bd.Put(pos)
                if err != nil {
                    panic("Search selected an invalid pos")
                }
                bd.Display()
                if bd.IsEnd() {
                    fmt.Println("Winner: ", bd.GetWinner())
                }
            default:
                fmt.Println("Invalid command")
        }
    }
}

