package game

import (
	"bufio"
	"fmt"
	"io"
	"os"
)


type Game struct{
	utStream , ErrStreamos io.Writer
}

func main(){

cli := game.Game(OutStream:os.Stdout , ErrStreamos:os.Stderr)

 //配列に問題を入れる
 //配列の数内のランダムな数を返す
 //時間内
 var questions =[]string{
"test12",
"test22",
"test32",
"test42",
"test13",
"test23",
"test33",
"test43",
"test14",
"test24",
"test34",
"test44",
 }


 var score int = 0
 for _,q := range questions{
 
 fmt.Printf("入力してください:%s",q)
 fmt.Println("")
 
 scanner := bufio.NewScanner(os.Stdin)
 scanner.Scan()
 in := scanner.Text()

 if in == q {
	 score ++
	 fmt.Printf("現在の得点:%d",score)
	 fmt.Println("")
 }

}
fmt.Printf("最終スコアの得点:%d",score)


}