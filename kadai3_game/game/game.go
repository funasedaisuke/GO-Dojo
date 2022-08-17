package game

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type Game struct{
	OutStream ,ErrStream io.Writer
}


func(g *Game)Run()int{
	//GAMEを始める
	st_ch :=start_game(os.Stdout)
	<-st_ch
	//問題を入力する
	//ch :=inputWords()
	
	return 1
} 

func start_game(r io.Reader)(chan struct{}){
	fmt.Println("-----------Start Game-----------")
	fmt.Println("----------------------------")
	fmt.Println("------Please type anything------")
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
