package player

import(
    "yonmoku2/board"
    "math"
    "math/rand"
    "errors"
    "time"
    "fmt"
    "sort"
    "sync"
    "sync/atomic"
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
    times int32
    win int32
    mu sync.Mutex
}


func (mcp *McPlayer) selectBestChild(nd *node) int {
    // Sort the child nodes according to the parent's win rate.
    compareChildren := func(i, j int) bool {
        iParentWinRate := float64(nd.children[i].times - nd.children[i].win) / float64(nd.children[i].times)
        jParentWinRate := float64(nd.children[j].times - nd.children[j].win) / float64(nd.children[j].times)
        return iParentWinRate > jParentWinRate
    }
    sort.Slice(nd.children, compareChildren)

    for _, child := range nd.children {
        parentWinRate := float64(child.times - child.win) / float64(child.times)
        fmt.Printf("pos, winRate = %v, %v\n", child.pos, parentWinRate)
    }

    selectedPos := nd.children[0].pos
    fmt.Println("selected pos = ", selectedPos)
    return selectedPos
}

func (mcp *McPlayer) getUcbScore(totalTimes int32, node *node) float64 {
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
            child := node{children: make([]node, 0),
                          pos: i,
                          times: 0,
                          win: 0}
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
    atomic.AddInt32(&nd.times, 1)
    var winner string
    if (*bd).IsEnd() {
        winner = (*bd).GetWinner()
    } else {
        // If this is the first time to visit the node "bd", get child list.
        nd.mu.Lock()
        if len(nd.children) == 0 {
            children := mcp.getChildList(nd, bd)
            if len(children) == 0 {
                panic("getChildList failed.")
            }
            nd.children = children
        }
        nd.mu.Unlock()

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
        atomic.AddInt32(&nd.win, 1)
    }
    return winner
}

func (mcp *McPlayer) Search(bd *board.Board) int {
    rand.Seed(time.Now().UnixNano())
    var root node
    root.pos = -1
    startTime := time.Now()
    const UNIT_SEARCH_TIMES = 500
    wg := &sync.WaitGroup{}
    for time.Since(startTime).Milliseconds() < mcp.timeInMs*98/100 {
        wg.Add(1)
        go func() {
            tmpBd := (*bd).Copy()
            for i := 0;
                i < UNIT_SEARCH_TIMES &&
                    time.Since(startTime).Milliseconds() < mcp.timeInMs*98/100;
                i++ {
                mcp.searchHelper(&root, &tmpBd)
            }
            wg.Done()
        }()
    }
    wg.Wait()
    selectedPos := mcp.selectBestChild(&root)
    fmt.Println("the number of searches: ", root.times)
    fmt.Println("elapsed time[ms]: ", time.Since(startTime).Milliseconds())
    return selectedPos
}

func (mcp *McPlayer) Time(timeInMs int64) error {
    if timeInMs <= 0 {
        return errors.New("The time should be positive.")
    }
    mcp.timeInMs = timeInMs
    return nil
}
