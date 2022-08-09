package game

import (
	"bufio"
	"fmt"
	"io"
	"main/word"
	"math"
	"os"
	"time"
)

type Game struct {
OutStream,ErrStream io.Writer
}

func (g *Game)Run()int {
	sch := start(os.Stdin)
	<-sch
	ch :=  input(os.Stdin)

	words,err := word.GetWords()
	if err != nil{
		fmt.Print(err)
		return 1
	}
	timerChan := time.NewTimer(30 * time.Second).C
	correctAnswerCount := 0
	answerCount := 0


Outer:

	for _,word := range words{
		fmt.Printf("type:%s\n",word)
        Inner:
		for {
			fmt.Print(">")
		select{
			
		case <-timerChan:
			break Outer

		case s := <-ch:
			if s == word{
				fmt.Println("正解")
				correctAnswerCount++
				answerCount++
				break Inner

			} else{
             fmt.Println("不正解")
			 answerCount++
			 fmt.Printf("type: %s\n", word)
			}
		}
         }
    }  
	fmt.Println("\n終了")
	fmt.Printf("正解数: %d words\n", correctAnswerCount)
	correctness := float64(0)
	if answerCount != 0 {
		correctness = roundPlus(float64(correctAnswerCount) / float64(answerCount) * 100, 2)
	}
	fmt.Printf("正確さ: %v %%\n", correctness)
	return 0

	} 

func round(f float64)float64{
	return math.Floor(f + .5)
}

func roundPlus(f float64, places int) (float64) {
	shift := math.Pow(10, float64(places))
	return round(f * shift) / shift;
}



func start(r io.Reader)(chan struct{}){
	fmt.Println("--------- Typing Game Start---------")
	fmt.Println("--------- PLease Type any key---------")
	sch := make(chan struct{})
	go func() {
		scanner:= bufio.NewScanner(r)
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
	go func() {
		scanner := bufio.NewScanner(r)
	    for scanner.Scan(){
		   ch <- scanner.Text()
     	}
		 close(ch)

	}()
	return ch

}