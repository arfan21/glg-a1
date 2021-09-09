package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/arfan21/hacktiv8-assignment-1/data"
)

func addDummyData(db *data.ListBiodata) {
	data1 := data.NewBiodata("Arfan", "Karanganyar", "Mahasiswa")
	db.Add(data1)

	data2 := data.NewBiodata("Karomi chan", "Semarang", "Mahasiswa")
	db.Add(data2)

	data3 := data.NewBiodata("Loki", "Agard", "Dewa Penipu")
	db.Add(data3)

	data4 := data.NewBiodata("Thor", "Asgard", "Dewa Petir")
	db.Add(data4)
}

func main() {
	db := data.NewListBiodata()
	addDummyData(db)

	args := os.Args

	if len(args) == 1 {
		fmt.Println("Harap masukkan argumen angka. Misal -> go run main.go 2")
		return
	}

	arg := args[1]

	argInt, err := strconv.Atoi(arg)

	if err != nil {
		fmt.Println("Argumen yang anda masukkan bukan angka")
		return
	}

	data := db.GetDataByNo(argInt)

	if data == nil {
		fmt.Println("Data tidak ditemukan")
		return
	}

	data.PrintData()
}
