package test

import (
	"log"
	"net/http"
	"encoding/json"
	"bitbucket.org/avcl/testinggomobile/testgomobcpp"
)

func Hello() error {

	testinggomobcpp.Hello("Hello")

	log.Println("WTF m")

	url := "http://163.172.155.202:1234/smartshows?userId=3682cf35-2f9f-408e-ba1f-1fdd8bed1fc4-1464814334308"
	r, err := http.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	obj := struct {
		Smartshows []struct{
			ID string `json:"id"`
		} `json:"smartshows"`
	}{}

	err = json.NewDecoder(r.Body).Decode(&obj)
	for _, o := range  obj.Smartshows {
		log.Println(o)
	}
	return err
}