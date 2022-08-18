package game

import (
	"bufio"
	"fmt"
	"io"
	"main/word"
	"os"
	"time"
	"math"
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
	fmt.Print("終了")
	fmt.Printf("回答数:%d\n",answerCount)
	fmt.Printf("正解数:%d\n",correctAnswerCount)
	correctness := float64(answerCount)
	if correctness != 0{
		correctness := roundPlus((float64(correctAnswerCount)/float64(answerCount))*100,2)
	}

	fmt.Printf("正解率:%d\n",correctness)
	return 1
}

func round(f float64)float64{
	return math.floor(f +0.5)
}

func roundPlus(f float64 ,digit int)float64{
	shift := math.Pow(10,float64(digit))
	return round(f * shift) / shift
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
