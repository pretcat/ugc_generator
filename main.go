package main

import (
	"fmt"

	"github.com/pretcat/ugc_generator/generator"
)

func main() {
	conf := generator.Config{
		Host:     "localhost",
		Port:     "5432",
		DbName:   "usg_test_task",
		User:     "postgres",
		Password: "7823",
	}
	gen, err := generator.New(conf)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	//if err := gen.GenerateBuildings(6000); err != nil {
	//	fmt.Println(err.Error())
	//	return
	//}
	if err := gen.GenerateCategories(4000); err != nil {
		fmt.Println(err.Error())
		return
	}
}
