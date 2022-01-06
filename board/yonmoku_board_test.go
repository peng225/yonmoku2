package board

import "testing"

func TestGetWinner(t *testing.T) {
    // type args struct {
    //     a int
    //     b int
    // }
    type params struct {
        posList []int
    }
    tests := []struct {
        name string
        // args args
        params params
        want string
    }{
        {
            name: "Empty board",
            // args: args{a: 1, b: 2},
            params: params{[]int{}},
            want: "EMPTY",
        },
        {
            name: "Draw",
            params: params{[]int{0, 1, 1, 2, 2, 3, 3, 5, 7}},
            want: "EMPTY",
        },
        {
            name: "Red fastest",
            params: params{[]int{1, 1, 2, 2, 3, 3, 4}},
            want: "RED",
        },
        {
            name: "Yellow fastest",
            params: params{[]int{0, 1, 1, 2, 2, 3, 3, 4}},
            want: "YELLOW",
        },
        {
            name: "Red vertical",
            params: params{[]int{1, 0, 1, 0, 1, 0, 1}},
            want: "RED",
        },
        {
            name: "Yellow vertical",
            params: params{[]int{1, 7, 2, 7, 3, 7, 1, 7}},
            want: "YELLOW",
        },
        {
            name: "Red left to down",
            params: params{[]int{0, 0, 1, 0, 0,
                                 1, 1, 2, 2,
                                 7, 3}},
            want: "RED",
        },
        {
            name: "Yellow top to left",
            params: params{[]int{5, 5, 5, 5, 5, 5, 4, 5,
                                 4, 4, 4, 4, 3, 4,
                                 3, 3, 3, 3,
                                 1, 2, 2, 2, 0, 2}},
            want: "YELLOW",
        },
    }
    for _, tt := range tests {
        var yb YonmokuBoard
        yb.Init(8)
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
