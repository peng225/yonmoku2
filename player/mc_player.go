package player

import(
    "yonmoku2/board"
    "math"
    "math/rand"
    "errors"
    "time"
    "fmt"
)

type McPlayer struct {
    timeInMs int64
}

const(
    C = 1.41421356
    EPSILON = 0.001
)

type node struct {
    children []node
    pos int
    times int
    win int
}

func (mcp *McPlayer) selectBestChild(nd *node) int {
    maxParentWinRate := -1.0
    selectedPos := -1
    for _, child := range nd.children {
        parentWinRate := float64(child.times - child.win) / float64(child.times)
        fmt.Printf("pos, winRate = %v, %v\n", child.pos, parentWinRate)
        if maxParentWinRate < parentWinRate {
            selectedPos = child.pos
            maxParentWinRate = parentWinRate
        }
    }
    fmt.Println("selected pos = ", selectedPos)
    return selectedPos
}

func (mcp *McPlayer) getUcbScore(totalTimes int, node *node) float64 {
    parentWin := node.times - node.win
    return float64(parentWin)/(float64(node.times)+EPSILON) +
            C * math.Sqrt(math.Log2(float64(totalTimes))/(float64(node.times)+EPSILON))
}

func (mcp *McPlayer) selectChild(nd *node) *node {
    maxScore := -1.0
    selectedIndex := -1
    for i, child := range nd.children {
        currentScore := mcp.getUcbScore(nd.times, &child)
        if maxScore < currentScore {
            maxScore = currentScore
            selectedIndex = i
        }
    }
    return &nd.children[selectedIndex]
}

func (mcp *McPlayer) getChildList(nd *node, bd *board.Board) []node {
    children := make([]node, 0)
    for i := 0; i < (*bd).Size(); i++ {
        if (*bd).CanPut(i) {
            child := node{make([]node, 0), i, 0, 0}
            children = append(children, child)
        }
    }
    return children
}

func (mcp *McPlayer) playout(nd *node, bd *board.Board) string {
    nd.times++
    putCount := 0
    for !(*bd).IsEnd() {
        pos := rand.Intn((*bd).Size())
        if (*bd).CanPut(pos) {
            err := (*bd).Put(pos)
            if err != nil {
                panic("Invalid pos")
            }
            putCount++
        }
    }
    winner := (*bd).GetWinner()
    for i := 0; i < putCount; i++ {
        (*bd).Reverse()
    }

    if winner == (*bd).GetTurn() {
        nd.win++
    }
    return winner
}

func (mcp *McPlayer) searchHelper(nd *node, bd *board.Board) string {
    nd.times++
    var winner string
    if (*bd).IsEnd() {
        winner = (*bd).GetWinner()
    } else {
        // If this is the first time to visit the node "bd", get child list.
        if len(nd.children) == 0 {
            children := mcp.getChildList(nd, bd)
            if len(children) == 0 {
                panic("getChildList failed.")
            }
            nd.children = children
        }

        // Move to the selected child node.
        child := mcp.selectChild(nd)
        err := (*bd).Put(child.pos)
        if err != nil {
            panic("Failed to expand node due to the Put failure.")
        }
        if child.times < 3 {
            winner = mcp.playout(child, bd)
        } else {
            // expand
            winner = mcp.searchHelper(child, bd)
        }
        (*bd).Reverse()
    }

    if winner == (*bd).GetTurn() {
        nd.win++
    }
    return winner
}

func (mcp *McPlayer) Search(bd *board.Board) int {
    rand.Seed(time.Now().UnixNano())
    var root node
    root.pos = -1
    startTime := time.Now()
    for time.Since(startTime).Milliseconds() < mcp.timeInMs {
        mcp.searchHelper(&root, bd)
    }
    selectedPos := mcp.selectBestChild(&root)
    fmt.Println("the number of searches: ", root.times)
    return selectedPos
}

func (mcp *McPlayer) Time(timeInMs int64) error {
    if timeInMs <= 0 {
        return errors.New("The time should be positive.")
    }
    mcp.timeInMs = timeInMs
    return nil
}
