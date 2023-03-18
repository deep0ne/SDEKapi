package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"

	"github.com/deep0ne/SDEKapi/data"
	"github.com/deep0ne/SDEKapi/utils"
)

func main() {
	// парсим из командной строки логин и пароль для доступа, а так же флаг тестового режима
	username := flag.String("username", "", "setting username")
	password := flag.String("password", "", "setting password")
	utils.TestMode = flag.Bool("testmode", true, "setting mode to test")

	flag.Parse()

	if *utils.TestMode {
		utils.CalcURL = "https://api.edu.cdek.ru/v2/calculator/tarifflist"
	} else {
		utils.CalcURL = "https://api.cdek.ru/v2/calculator/tarifflist"
	}

	if *username == "" || *password == "" {
		log.Fatal("Обязательно нужен логин и пароль для доступа.")
	}

	utils.SetUsername(*username)
	utils.SetPassword(*password)

	sizeOfPackage := data.NewPackage(50000, 100, 50, 30)
	b, err := utils.Calculate(data.AddressFrom, data.AddressTo, sizeOfPackage)
	if err != nil {
		log.Fatal(err)
	}

	// демонстрация результата
	res, err := json.MarshalIndent(b, "", "\t")
	if err != nil {
		log.Fatal("Не удалось провести сериализацию")
	}
	fmt.Println(string(res))
}
