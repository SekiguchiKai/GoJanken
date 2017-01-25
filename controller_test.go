package main

import "testing"

func TestSetPlayerHand(t *testing.T) {

	gu := "グー"
	tyoki := "チョキ"
	pa := "パー"

	user := new(Player)
	user.setPlayerHand(0)
	testSetPlayerHandHelper(t, user, gu)
	user.setPlayerHand(1)
	testSetPlayerHandHelper(t, user, tyoki)
	user.setPlayerHand(2)
	testSetPlayerHandHelper(t, user, pa)
}

func testSetPlayerHandHelper(t *testing.T, user *Player, expected string) {
	userHand := user.PlayerHand

	if userHand != expected {
		t.Fatal(expected + "で期待していたsetPlayerHandでintからstringへの変換が正常に行われていません")
	}
}


