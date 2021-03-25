package security

import "golang.org/x/crypto/bcrypt"

//Hash recebe uma string e insere um hash
func Hash(secret string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(secret), bcrypt.DefaultCost)
}

//VerifyPassword compara uma senha e um hash
func VerifyPassword(secretHash, secretString string) error {
	return bcrypt.CompareHashAndPassword([]byte(secretHash), []byte(secretString))
}
