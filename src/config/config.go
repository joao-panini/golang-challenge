package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	ConnectionString = ""
	Port             = 0
	SecretKey        []byte
)

//Carregar le o arquivo .env e inicializa as variaveis
func Carregar() {
	var erro error

	if erro = godotenv.Load(); erro != nil {
		log.Fatal(erro)
	}

	Port, erro = strconv.Atoi(os.Getenv("API_PORT"))
	if erro != nil {
		fmt.Println(erro)
	}
	ConnectionString = fmt.Sprintf("%s:%s@(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("MYSQL_HOST"),
		os.Getenv("DB_NAME"),
	)
	fmt.Println(ConnectionString)

	SecretKey = []byte(os.Getenv(string("SECRET_KEY")))

}
