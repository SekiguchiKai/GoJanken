package main

import (
	"net/http"
	"strconv"
	"log"
)

// controller1
func IndexHandler(res http.ResponseWriter, req *http.Request) {
	http.ServeFile(res, req, "./resource/index.html")
}


// controller2
func ResultHandler(res http.ResponseWriter, req *http.Request) {
	user := NewPlayer("あなた")
	cpu := NewPlayer("AI")

	hand := req.FormValue("hand")

	intUserHand, err := strconv.Atoi(hand)
	if err != nil {
		log.Println(err)
	}
	user.setPlayerHand(intUserHand)

	intCpuHand := DecideCpuHand()
	cpu.setPlayerHand(intCpuHand)

	user.PlayerResult, cpu.PlayerResult = Judge(intUserHand, intCpuHand)

	err = resultView(res, req, user, cpu)

	if err != nil {
		log.Println(err)
	}

}