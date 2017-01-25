package main

import (
	"net/http"
	"fmt"
	"os"
	"strconv"
)

type Player struct {
	PlayerHand   string
	PlayerName   string
	PlayerResult string
}

func (p *Player) setPlayerHand(intHand int) {

	switch intHand {
	case 0:
		p.PlayerHand = "グー"
	case 1:
		p.PlayerHand = "チョキ"
	default:
		p.PlayerHand = "パー"
	}
}


// controller1
func ReadHTML(res http.ResponseWriter, req *http.Request) {
	fmt.Println("ReadHTMLが呼ばれました")
	file, err := os.Open("./resource/index.html")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	buf := make([] byte, 99999)
	for {
		n, err := file.Read(buf)
		if n == 0 {
			break
		}
		if err != nil {
			// Readエラー処理
			fmt.Println(err)
			os.Exit(1)
			break
		}

		//fmt.Print(string(buf[:n]))
		res.Write(buf[:n])
	}
}


// controller2
func sendResponse(res http.ResponseWriter, req *http.Request) {
	fmt.Println("sendResponse")

	user := new(Player)
	cpu := new(Player)

	hand := req.FormValue("hand")

	// modelに渡す
	user.PlayerName = "あなた"

	//　打ち手
	intUserHand, err := strconv.Atoi(hand)
	user.setPlayerHand(intUserHand)

	intCpuHand := DecideCpuHand()
	fmt.Print("cpuの")
	fmt.Print(intCpuHand)
	cpu.setPlayerHand(intCpuHand)

	user.PlayerResult, cpu.PlayerResult = Judge(intUserHand, intCpuHand)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// cpu
	cpu.PlayerName = "AI"

	fmt.Println(cpu.PlayerHand)
	fmt.Println(cpu.PlayerName)

	// 勝敗結果
	user.PlayerResult, cpu.PlayerResult = Judge(intUserHand, intCpuHand)

	resultView(res, req, user, cpu)

}