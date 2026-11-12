package req

import (
	"net/http"
	"pack/pkg"
)

func HandleBody[T any](w *http.ResponseWriter, r *http.Request) (*T, error) { //обработать body
	body, err := Decode[T](r.Body)
	if err != nil {
		pkg.Json(*w, err.Error(), 402)
		return nil, err
	}
	err = IsValidte(body)
	if err != nil {
		pkg.Json(*w, err.Error(), 402)
		return nil, err
	}
	return &body, nil
}
