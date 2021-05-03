package main

import (
	"archive/zip"
	"bytes"
	"errors"
	"fmt"
	"os"
)

func calcRemainderAndMod(numerator, denominator int) (int, int, error) {
	if denominator == 0 {
		return 0, 0, errors.New("denominator is 0")
	}
	return numerator / denominator, numerator % denominator, nil
}

type error interface {
	Error() string
}

func doubleEven(i int) (int, error) {
	if i%2 != 0 {
		return 0, fmt.Errorf("%d isn't an even number", i)
	}
	return i * 2, nil
}

type Status int

const (
	InvalidLogin Status = iota + 1
	NotFound
)

type StatusErr struct {
	Status  Status
	Message string
}

func (se StatusErr) Error() string {
	return se.Message
}

//func LoginAndGetData(uid, pwd, file string) ([]byte, error) {
//	login := func(uid, pwd string) {}
//	getData := func(file string) {}
//	err := login(uid, pwd)
//	if err != nil {
//		return nil, StatusErr{
//			Status:  InvalidLogin,
//			Message: fmt.Sprintf("invalid credentials for user %s", uid),
//		}
//	}
//	data, err := getData(file)
//	if err != nil {
//		return nil, StatusErr{
//			Status:  NotFound,
//			Message: fmt.Sprintf("file %s not found", file),
//		}
//	}
//}

func generateError(flag bool) error {
	var genErr StatusErr
	if flag {
		genErr = StatusErr{
			Status: NotFound,
		}
	}
	return genErr
}

func main() {

	err1 := generateError(true)
	fmt.Println(err1 != nil)
	err2 := generateError(false)
	fmt.Println(err2 != nil)

	data := []byte("This is not a zip file")
	notAZipFile := bytes.NewReader(data)
	_, err := zip.NewReader(notAZipFile, int64(len(data)))

	if err == zip.ErrFormat {
		fmt.Println("Told you so")
	}

	numerator, denominator := 20, 3
	remainder, mod, err := calcRemainderAndMod(numerator, denominator)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(remainder, mod)

}
