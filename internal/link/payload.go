package link

type LinkCreateRequest struct { //то, что приходит и как его интерпретировать
	Url string `json:"url" validate:"required,url"`
}

type LinkUpdateRequest struct { //то, что приходит и как его интерпретировать
	Url  string `json:"url" validate:"required,url"`
	Hash string `json:"hash"`
}
