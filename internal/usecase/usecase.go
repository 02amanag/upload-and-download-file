package usecase

type UsecaseStruct struct {
}

func NewUsecaseStruct() *UsecaseStruct {
	return &UsecaseStruct{}
}

func (u *UsecaseStruct) Healthy(param string) string {
	param += "added from usecase"
	return param
}
