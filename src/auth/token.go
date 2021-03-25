package auth

import (
	"api/src/config"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

//CreateToken cria o token da sessao do usuario
func CreateToken(accountID uint64) (string, error) {
	permissions := jwt.MapClaims{}
	permissions["authorized"] = true
	permissions["exp"] = time.Now().Add(time.Hour * 6).Unix()
	permissions["accountID"] = accountID

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permissions)
	return token.SignedString([]byte(config.SecretKey))

}

//ValidateToken verifica se o token passado no request é valido
func ValidateToken(r *http.Request) error {
	tokenString := extractToken(r)
	token, erro := jwt.Parse(tokenString, getVerifyKey)

	if erro != nil {
		return erro
	}

	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return nil
	}
	return errors.New("invalid token")

}

func extractToken(r *http.Request) string {
	token := r.Header.Get("Authorization")

	if len(strings.Split(token, " ")) == 2 {
		return strings.Split(token, " ")[1]
	}
	return ""
}

func getVerifyKey(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("metodo de assinatura inesperado! %v", token.Header["alg"])
	}
	return config.SecretKey, nil
}

//ExtractAccountId retorna accountID do token
func ExtractAccountId(r *http.Request) (uint64, error) {
	tokenString := extractToken(r)
	token, erro := jwt.Parse(tokenString, getVerifyKey)
	if erro != nil {
		return 0, erro
	}
	if permissions, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		accountID, erro := strconv.ParseUint(fmt.Sprintf("%.0f", permissions["accountID"]), 10, 64)
		if erro != nil {
			return 0, erro
		}
		return accountID, nil
	}
	return 0, errors.New("token invalido")
}
