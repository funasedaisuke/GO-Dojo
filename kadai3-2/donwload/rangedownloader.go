package download

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"strconv"
)

const tempDir ="dlTmp" 
type Downloader struct{
	Args []string
	procs int
    url string
	name string

}

func New() *Downloader {
	return &Downloader{Args: os.Args[1:]}
}



func (d *Downloader) run()int{
	//os.ModePerm permissionのこと、defaultは077 
if err :=os.MkdirAll(tempDir,os.ModePerm); err != nil{
	fmt.Println(err)
	return 1
}

if err := d.parseCommandLine(); err != nil {
		fmt.Println(err)
		return 1
}
len,err := d.getContentLength()
if err != nil{
		fmt.Println(err)
		return 1
}



}

func (d *Downloader) getContentLength()(int,error){
	res,err := http.Head(d.url)
	if err!=nil{
		return 0,errors.Wrapf(err, "failed to access the site you provided: %s", d.url)
	}
	if res.Header.Get("Accept-Ranges") != "bytes"{
		return 0, errors.New("this site doesn't support a range request")
	}

	len,err := strconv.Atoi(res.Header.Get("Content-Length"))
	if err != nil{
		return 0, errors.Wrap(err, "failed to get Content-Length")
	}
	fmt.Printf("total length: %d bytes\n", len)
	return len,nil


}