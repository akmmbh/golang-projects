package utils

import (
	"encoding/json"
	"io"
	"net/http"
)
func ParseBody(r *http.Request, x interface{}){
	//x interface {} means any type we can put on it

	if body, err := io.ReadAll(r.Body); err == nil{//read all the body and slice into a array

		if err:= json.Unmarshal([]byte(body),x); err!=nil{// body ko unmarshel mtlb mai jisme chahu usme convert kar diya
			
			return 
		}
	}
}