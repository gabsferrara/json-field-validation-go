package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type User struct {
	Id   int64     `json:"id"`
	Name string    `json:"name"`
	Age  Field[uint8] `json:"age"`
}

type Field[T any] struct {
	Valid bool
	Val   T
}

// Custom unmarshal method to handle Field type
func (f *Field[T]) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		f.Valid = false
		return nil
	}
	f.Valid = true
	return json.Unmarshal(data, &f.Val)
}

func main() {
	// Carregar e parsear o arquivo JSON
	fileName := "user.json"
	print(fileName)

	fileNameWithout := "user_without_age.json"
	print(fileNameWithout)
	
}

//https://pkg.go.dev/encoding/json#Unmarshaler.UnmarshalJSON

func print(fileName string) {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Erro ao abrir o arquivo:", err)
		return
	}
	defer file.Close()

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println("Erro ao ler o arquivo:", err)
		return
	}

	var user User
	err = json.Unmarshal(bytes, &user)
	if err != nil {
		fmt.Println("Erro ao parsear o JSON:", err)
		return
	}

	// Imprimir todos os campos
	fmt.Printf("Id: %d\n", user.Id)
	fmt.Printf("Name: %s\n", user.Name)
	if user.Age.Valid {
		fmt.Printf("Age: %d\n", user.Age.Val)
	} else {
		fmt.Println("Age: Not Provided")
	}
}