package main

/*
	#include "sum.c"
*/
import "C"
import (
	"errors"
	"fmt"

	log "github.com/sirupsen/logrus"
)

//Myprinter is my Printer
func Myprinter(a ...interface{}) (n int, err error) {
	return fmt.Println(a)
}

func fib(n int) (int, error) {
	if n <= 0 {
		return n, errors.New("What the CAT")
	} else if n == 1 {
		return n, nil
	}
	res, err := fib(n - 1)
	return res * n, err
}

func main() {
	for i := 1; i <= 15; i++ {
		log.Info("Start Loop")
		n, X := fib(i)
		fmt.Println(n, X)
		Myprinter(n)
		println(C.sum(10, 50))
		log.Info("Finish Loop")
	}
	log.Warn("Bye Bye!")
	log.Fatal("Bye Bye!")
	//
	log.Info("Unreachable")
}
