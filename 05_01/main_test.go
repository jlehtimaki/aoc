package main

import (
	"testing"
)

var seatIds = map[string]int{
"FBFBBFFRLR": 357,
"BFFFBBFRRR": 567,
"FFFBBBFRRR": 119,
"BBFFBBFRLL": 820,
}

func TestGetSeatId(t *testing.T) {
	for seatId, _ := range seatIds {
		got := getSeatId(seatId)
		want := seatIds[seatId]
		if got != want {
			t.Errorf("%s --> want: %d - got: %d", seatId, want, got)
		}
	}
}

func testMain(t *testing.T) {

}