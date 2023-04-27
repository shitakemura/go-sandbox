// package main

// import "fmt"

// type Status int

// const (
// 	InvalidLogin Status = iota + 1
// 	NotFound
// )

// type StatusError struct {
// 	Status  Status
// 	Message string
// }

// func (se StatusError) Error() string {
// 	return se.Message
// }

// func login(uid, pwd string) error {
// 	return nil
// }

// func getData(file string) ([]byte, error) {
// 	return []byte{}, nil
// }

// func LoginAndGetData(uid, pwd, file string) ([]byte, error) {
// 	err := login(uid, pwd)
// 	if err != nil {
// 		return nil, StatusError{
// 			Status:  InvalidLogin,
// 			Message: fmt.Sprintf("invalid credentials for user %s", uid),
// 		}
// 	}
// 	data, err := getData(file)
// 	if err != nil {
// 		return nil, StatusError{
// 			Status:  NotFound,
// 			Message: fmt.Sprintf("fils %s not found", file),
// 		}
// 	}
// 	return data, nil
// }
