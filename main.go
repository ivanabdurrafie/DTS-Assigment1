package main

import (
	"assignment-1/Person"
	"fmt"
	"os"
	"reflect"
	"strconv"
)

func main() {
	absen := os.Args[1]

	intAbsen, _ := strconv.Atoi(absen)

	if intAbsen >= len(Person.Students) {
		fmt.Println("data tidak ada")
	}
	for i, v := range Person.Students {
		if intAbsen == i {
			fmt.Printf("%+v \n \n", v)
			fmt.Println(reflect.ValueOf(v).Type().Field(0).Name, v.Nama)
			fmt.Println(reflect.ValueOf(v).Type().Field(1).Name, v.Alamat)
			fmt.Println(reflect.ValueOf(v).Type().Field(2).Name, v.Pekerjaan)
			fmt.Println(reflect.ValueOf(v).Type().Field(3).Name, v.Alasan)
		}
	}

}
