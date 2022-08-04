package main

import (
	"fmt"
	"game"
	"os"
)



func main(){
defer func(){
	if err:= recover(); err != nil{
		fmt.Fprintf(os.Stderr,"Error:¥n%s¥n",err)
		os.Exit(1)
	}
}()
cli := game.Game{OutStream:os.Stdout , ErrStreamos:os.Stderr}
os.Exit(cli.Run())
}
 