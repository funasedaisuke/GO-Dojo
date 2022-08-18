package game

import (
	"bufio"
	"fmt"
	"io"
	"main/word"
	"os"
	"time"
)

type Game struct {
	OutStream, ErrStream io.Writer
}

func (g *Game) Run() int {
	//GAMEを始める
	st_ch := start_game(os.Stdout)
	<-st_ch
	//答えを入力する
	//ch :=inputWords()
	ch := input(os.Stdout)

	//問題を入力
	words, err := word.GetWords()
	if err != nil {
		return 0
	}

	answerCount := 0
	correctAnswerCount := 0
	timerch := time.NewTimer(time.Second * 30).C
		fmt.Println("GameStart")
Outer:
	for _,wd := range words {

		fmt.Printf("type:%s\n",wd)
	Inner:
		for {
			select {
			case <-timerch:
				break Outer
			case input :=<-ch:
				if wd == input {
					answerCount += 1
					correctAnswerCount += 1
					fmt.Println("正解")
				break Inner

				} else {
					answerCount += 1
					fmt.Println("不正期")
				}

			}
		}
	}
	fmt.Printf("回答数:%d\n",answerCount)
	fmt.Printf("正解数:%d\n",correctAnswerCount)
	fmt.Printf("正解率:%d\n%",(correctAnswerCount/answerCount)*100)

	return 1
}

func start_game(r io.Reader) chan struct{} {
	fmt.Println("-----------Start Game-----------")
	fmt.Println("--------------------------------")
	fmt.Println("------Please type anything------")
	fmt.Println(">")
	sch := make(chan struct{})
	go func() {
		scanner := bufio.NewScanner(r)
		for scanner.Scan() {
			sch <- struct{}{}
			break
		}
		close(sch)
	}()
	return sch

}

func input(r io.Reader) chan string {
	ch := make(chan string)
	go func() {
		scanner := bufio.NewScanner(r)
		for scanner.Scan() {
			ch <- scanner.Text()
		}
		close(ch)
	}()
	return ch

}
