package filemanager

import (
	"bufio"
	"errors"
	"os"
)
func ReadLines(path string)([]string,error){
	file,err:=os.Open(path)
	if err!=nil{
		return nil,errors.New("failed to read file")
	}
	sc:=bufio.NewScanner(file)
	var lines []string
	for sc.Scan(){
     lines=append(lines, sc.Text())
	}
    err=sc.Err()
	if err!=nil{
		file.Close()
		return nil,errors.New("read line failed")
	}
	file.Close()
	return lines,nil

}