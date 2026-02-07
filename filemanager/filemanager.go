package filemanager

import (
	"bufio"
	"errors"
	"os"
	"encoding/json"
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

func WriteJSON(path string,data interface{})error{
	file ,err:=os.Create(path)
	if err!=nil{
		return errors.New("failed create file")
	}
	//scanner:=bufio.NewScanner(file)

	encoder:=json.NewEncoder(file)
	err=encoder.Encode(data)

	if err!=nil{
		return errors.New("failed to convert data to json")
	}

	file.Close()
	return nil


	
}