package utils

import (
	"encoding/json"
	"io"
)

type JSONResponse struct{
	Data any
}

func WriteJSON(w io.Writer, data any){
	res := JSONResponse{Data: data}
	json.NewEncoder(w).Encode(res)
}