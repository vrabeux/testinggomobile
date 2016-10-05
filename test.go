package test

import (
	"github.com/vrabeux/testinggomobile/testgomobcpp"
)

func Hello() error {
	testinggomobcpp.Hello("Hello")
	return nil
}
