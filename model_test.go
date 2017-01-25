package main

import (
	"testing"
)

func TestDecideCpuHand(t *testing.T) {

	// スライス(Listのようなもの)
	s := make([]int, 100)

	// 100回回す
	for i := 0; i < 100; i++ {
		cpuHand := DecideCpuHand()
		s = append(s, cpuHand)
	}

	b0, b1, b2 := testDecideCpuHandHelper(s)

	if !b0 {
		t.Fatal("擬似乱数によって0が生み出されていな")
	} else if !b1 {
		t.Fatal("擬似乱数によって1が生み出されていない")
	} else if !b2 {
		t.Fatal("擬似乱数によって2が生み出されていない")
	}

}

func testDecideCpuHandHelper(s []int) (bool, bool, bool) {

	b0 := false
	b1 := false
	b2 := false

	for _, v := range s {
		if v == 0 {
			b0 = true
		} else if v == 1 {
			b1 = true
		} else if v == 2 {
			b2 = true
		}
	}

	return b0, b1, b2
}

func TestJudge(t *testing.T) {
	gu := 0
	tyoki := 1
	pa := 2
	win := "勝ち"
	lose := "負け"
	draw := "引き分け"

	testJudgeHelper(t, gu, tyoki, win, lose)
	testJudgeHelper(t, tyoki, pa, win, lose)
	testJudgeHelper(t, pa, gu, win, lose)

	testJudgeHelper(t, tyoki, gu, lose, win)
	testJudgeHelper(t, pa, tyoki, lose, win)
	testJudgeHelper(t, gu, pa, lose, win)

	testJudgeHelper(t, gu, gu, draw, draw)
	testJudgeHelper(t, tyoki, tyoki, draw, draw)
	testJudgeHelper(t, pa, pa, draw, draw)

}

func testJudgeHelper(t *testing.T, hand1, hand2 int, expected1, expected2 string) {
	actual1, actual2 := Judge(hand1, hand2)
	if (actual1 != expected1) {
		t.Fatal("hand1に対する勝敗診断が誤っています")
	} else if actual2 != expected2 {
		t.Fatal("hand2に対する勝敗診断が誤っています")
	}

}
