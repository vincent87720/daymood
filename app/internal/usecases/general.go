package usecases

import "fmt"

type UsecaseError struct {
	Usecase string
	Code    int
	Message string
}

func (usecaseError *UsecaseError) Error() string {
	return fmt.Sprintf("[Usecase] %v | Code %v | %v", usecaseError.Usecase, usecaseError.Code, usecaseError.Message)
}
