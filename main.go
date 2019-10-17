package main

import (
	"fmt"
	"errors"
	"os"
)

func main() {
	err := errors.New("errore 1")

	// wrap con %w
	err2 := fmt.Errorf("errore 2 %w", err)
	err3 := fmt.Errorf("errore 3 %w", err2)

	fmt.Println("err")
	fmt.Printf("%v\n\n", err)

	fmt.Println("err2")
	fmt.Printf("%v\n\n", err2)

	fmt.Println("err3")
	fmt.Printf("%v\n\n", err3)

	// Unwrap

	fmt.Println("errors.Unwrap(err3)")
	fmt.Printf("%v\n\n", errors.Unwrap(err3))

	// Is

	fmt.Println("errors.Is(err3, err)")
	bo := errors.Is(err3, err)
	fmt.Printf("%t\nerr3: %v\nerr: %v\n\n", bo, err3, err)

	fmt.Println("errors.Is(err3, err2)")
	bo = errors.Is(err3, err2)
	fmt.Printf("%t\nerr3: %v\nerr2: %v\n\n", bo, err3, err2)

	fmt.Println("errors.Is(err, err3)")
	bo = errors.Is(err, err3)
	fmt.Printf("%t\nerr: %v\nerr3: %v\n\n", bo, err, err3)
	
	// As

	fmt.Println("errors.As(err2, &err3)")
	bo = errors.As(err2, &err3)
	fmt.Printf("%t\nerr3: %v\nerr2: %v\n\n", bo, err3, err2)

	fmt.Println("errors.As(err3, &err2)")
	bo = errors.As(err3, &err2)
	fmt.Printf("%t\nerr3: %v\nerr2: %v\n\n", bo, err3, err2)

	file, myerr1 := os.Open("i-do-not-exist")
	if myerr1 != nil {
		myerr2 := fmt.Errorf("sono myerr2\n%w", myerr1)
		var patherr *os.PathError
		if errors.As(myerr2, &patherr) {
			fmt.Println(patherr)
		} else {
			fmt.Println(":(")
		}
	}
	defer file.Close()


}
