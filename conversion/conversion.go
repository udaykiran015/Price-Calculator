package conversion

import (
	"errors"
	"strconv"
)
func StringToFloat(s []string)([]float64,error){
	var floats []float64
	for _,line:=range s{
		fprice,err:=strconv.ParseFloat(line, 64)
		if err!=nil{
			return nil,errors.New("Failed to convert string to number")
		}
		floats=append(floats,fprice)
	}
	return floats,nil
}