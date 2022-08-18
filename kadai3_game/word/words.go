package word

import 
(
	"fmt"
	"os"
	"bufio"
)

const questionCount int = 30
func GetWords()([]string,error){
		
	//ファイルを読み込む
	file,err :=os.Open("./words.txt")
	//一行ずつ読み込む
	if err != nil{
		fmt.Print(err)
		return nil,err
		
	}
	defer file.Close()
	//mapを使って、重複をなくす

	lines:= make(map[string]struct{})
	scanner  := bufio.NewScanner(file)
	for scanner.Scan(){
		wd := scanner.Text()
		lines[wd]=struct{}{}
	}
	//stringに入れる
	words:=make([]string,0)
	count :=0
	for key,_ := range lines{
		words =append(words,key)
		count++
		if count == questionCount{
			break
		}

	}


	return words,nil
}