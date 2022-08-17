package word

import 
(
	"fmt"
	"os"
)

func GetWords()([]string,error){
		
	
	//ファイルを読み込む
	file,err :=os.Open("./words.txt")
	//一行ずつ読み込む
	if err != nil{
		fmt.Print(err)
		return nil,err
		
	}
	defer file.Close()
	fmt.Print(file)
	//mapを使って、重複をなくす
	//stringに入れる
	return nil,err
}