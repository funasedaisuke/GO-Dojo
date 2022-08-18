package game

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"main/word"
)

type Game struct{
	OutStream ,ErrStream io.Writer
}


func(g *Game)Run()int{
	//GAMEを始める
	st_ch :=start_game(os.Stdout)
	<-st_ch
	//答えを入力する
	//ch :=inputWords()
	ch := input(os.Stdout)
	fmt.Print(ch)
	
	//問題を入力
	words,err := word.GetWords()
	if err !=nil{
		return 0
	}
	fmt.Print("--------")
	fmt.Print(words)


	
	return 1
} 

func start_game(r io.Reader)(chan struct{}){
	fmt.Println("-----------Start Game-----------")
	fmt.Println("--------------------------------")
	fmt.Println("------Please type anything------")
	fmt.Println(">")
	sch := make(chan struct{})
	go func(){
		scanner := bufio.NewScanner(r)
		for scanner.Scan(){
			sch <- struct{}{}
			break
		}
		close(sch)
	}()
	return sch

}

func input(r io.Reader)(chan string){
	ch := make(chan string)
	go func(){
		scanner := bufio.NewScanner(r)
		for scanner.Scan(){
			ch <- scanner.Text()
		}
		close(ch)
	}()
	return ch

}
