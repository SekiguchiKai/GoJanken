package main

import (
	"net/http"
	"fmt"
	"os"
	"strconv"
)



// controller1
func IndexHandler(res http.ResponseWriter, req *http.Request) {
	http.ServeFile(res, req, "./resource/index.html")
}


// controller2
func ResultHandler(res http.ResponseWriter, req *http.Request) {
	user := new(Player)
	cpu := new(Player)

	hand := req.FormValue("hand")

	user.PlayerName = "あなた"
	cpu.PlayerName = "AI"

	intUserHand, err := strconv.Atoi(hand)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	user.setPlayerHand(intUserHand)

	intCpuHand := DecideCpuHand()
	cpu.setPlayerHand(intCpuHand)

	user.PlayerResult, cpu.PlayerResult = Judge(intUserHand, intCpuHand)

	resultView(res, req, user, cpu)

}