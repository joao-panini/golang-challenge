package main

import (
	"api/src/config"

	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

//CreateSecretKey funcao que foi usada para gerar o secret code do token
// func CreateSecretKey() {
// 	key := make([]byte, 64)

// 	if _, erro := rand.Read(key); erro != nil {
// 		log.Fatal(erro)
// 	}
// 	fmt.Println(key)
// 	keyBase64 := base64.StdEncoding.EncodeToString(key)
// 	fmt.Println(keyBase64)

// }

func main() {
	config.Carregar()
	r := router.CreateRouter()
	fmt.Printf("Escutando na porta %d\n", config.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), r))
}
