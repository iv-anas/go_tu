package hw

import (
	"fmt"
	"errors"
	// "randomFormat"
	// "bytes"

)

func Hello(msg string) (string , error){
	if msg == "" {
        return msg, errors.New("empty name")
    }
	
	// var buffer bytes.Buffer
	message :=fmt.Sprintf("hello world")
	
	
	if msg!=message{
		
		return msg, errors.New("empty name")
	}
    return message, nil
}

