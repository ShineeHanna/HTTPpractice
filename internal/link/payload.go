package link

type LinkCreateRequest struct { //то, что приходит и как его интерпретировать
	Url string `json:"url" validate:"required,url"`
}
