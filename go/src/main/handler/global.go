package handler

import (
	"errors"
)


func Login(password string)error{
	if (password == "zboub"){
		return nil
	}else {
		return errors.New("wrong !!")
	}
		
	 

}
