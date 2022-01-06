package board

import "testing"

func TestGetWinner(t *testing.T) {
    // type args struct {
    //     a int
    //     b int
    // }
    const DEFAULT_SIZE = 8
    type params struct {
        posList []int
        size int
    }
    tests := []struct {
        name string
        // args args
        params params
        want string
    }{
        // size == DEFAULT_SIZE
        {
            name: "Empty board",
            // args: args{a: 1, b: 2},
            params: params{[]int{}, DEFAULT_SIZE},
            want: "EMPTY",
        },
        {
            name: "Draw",
            params: params{[]int{0, 1, 1, 2, 2, 3, 3, 5, 7}, DEFAULT_SIZE},
            want: "EMPTY",
        },
        {
            name: "Red fastest",
            params: params{[]int{1, 1, 2, 2, 3, 3, 4}, DEFAULT_SIZE},
            want: "RED",
        },
        {
            name: "Yellow fastest",
            params: params{[]int{0, 1, 1, 2, 2, 3, 3, 4}, DEFAULT_SIZE},
            want: "YELLOW",
        },
        {
            name: "Red vertical",
            params: params{[]int{1, 0, 1, 0, 1, 0, 1}, DEFAULT_SIZE},
            want: "RED",
        },
        {
            name: "Yellow vertical",
            params: params{[]int{1, 7, 2, 7, 3, 7, 1, 7}, DEFAULT_SIZE},
            want: "YELLOW",
        },
        {
            name: "Red left to bottom",
            params: params{[]int{0, 0, 1, 0, 0,
                                 1, 1, 2, 2,
                                 7, 3}, DEFAULT_SIZE},
            want: "RED",
        },
        {
            name: "Yellow top to left",
            params: params{[]int{5, 5, 5, 5, 5, 5, 4, 5,
                                 4, 4, 4, 4, 3, 4,
                                 3, 3, 3, 3,
                                 1, 2, 2, 2, 0, 2}, DEFAULT_SIZE},
            want: "YELLOW",
        },
        // size == 4
        {
            name: "Empty board (size == 4)",
            params: params{[]int{}, 4},
            want: "EMPTY",
        },
        {
            name: "Draw (size == 4)",
            params: params{[]int{0, 1, 1, 2, 2, 3, 3, 1, 1}, 4},
            want: "EMPTY",
        },
        {
            name: "Red fastest (size == 4)",
            params: params{[]int{1, 1, 2, 2, 3, 3, 0}, 4},
            want: "RED",
        },
        {
            name: "Yellow fastest (size == 4)",
            params: params{[]int{0, 0, 1, 1, 2, 2, 2, 3, 2, 3}, 4},
            want: "YELLOW",
        },
        {
            name: "Red vertical (size == 4)",
            params: params{[]int{1, 0, 1, 0, 1, 0, 1}, 4},
            want: "RED",
        },
        {
            name: "Yellow vertical (size == 4)",
            params: params{[]int{1, 3, 2, 3, 1, 3, 2, 3}, 4},
            want: "YELLOW",
        },
        {
            name: "Red diagonal (left-top to right-bottom) (size == 4)",
            params: params{[]int{0, 0, 1, 0, 0,
                                 1, 1, 2, 2,
                                 2, 3}, 4},
            want: "RED",
        },
        {
            name: "Yellow diagonal (right-top to left-bottom) (size == 4)",
            params: params{[]int{3, 3, 3, 3, 2, 2, 1, 2,
                                 2, 1, 1, 0}, 4},
            want: "YELLOW",
        },
        // size == 10
        {
            name: "Empty board (size == 10)",
            params: params{[]int{}, 10},
            want: "EMPTY",
        },
        {
            name: "Draw (size == 10)",
            params: params{[]int{0, 1, 1, 2, 2, 3, 3, 5, 9}, 10},
            want: "EMPTY",
        },
        {
            name: "Red fastest (size == 10)",
            params: params{[]int{6, 6, 7, 7, 8, 8, 9}, 10},
            want: "RED",
        },
        {
            name: "Yellow fastest (size == 10)",
            params: params{[]int{5, 6, 6, 7, 7, 8, 8, 9}, 10},
            want: "YELLOW",
        },
        {
            name: "Red vertical (size == 10)",
            params: params{[]int{8, 0, 8, 0, 8, 0, 8}, 10},
            want: "RED",
        },
        {
            name: "Yellow vertical (size == 10)",
            params: params{[]int{1, 9, 2, 9, 3, 9, 8, 9}, 10},
            want: "YELLOW",
        },
        {
            name: "Red left to bottom (size == 10)",
            params: params{[]int{0, 0, 1, 0, 0,
                                 1, 1, 2, 2,
                                 9, 3}, 10},
            want: "RED",
        },
        {
            name: "Yellow top to right (size == 10)",
            params: params{[]int{6, 6, 6, 6, 6, 6, 6, 6, 7, 6,
                                 7, 7, 7, 7, 7, 7, 8, 7,
                                 8, 8, 8, 8, 8, 8,
                                 1, 9, 9, 9, 9, 9, 0, 9}, 10},
            want: "YELLOW",
        },
    }
    for _, tt := range tests {
        var yb YonmokuBoard
        yb.Init(tt.params.size)
        for _, pos := range(tt.params.posList) {
            yb.Put(pos)
        }
        t.Run(tt.name, func(t *testing.T) {
            if got := yb.GetWinner(); got != tt.want {
                t.Errorf("GetWinner() = %v, want %v", got, tt.want)
            }
        })
    }
}
