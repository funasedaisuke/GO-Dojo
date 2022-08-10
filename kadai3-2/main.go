package main

import (
	"fmt"
	"os"
)

func main(){
	defer func(){
		if err := recover();ere != nil{
			fmt.Fprintf(os.Stderr,"Error\n%s\n",err)
			os.Exit(1)
		}

	}()
	cli:=rangedownloader.new()
	os.Exit(cli.Run())


}