package controllers

import (
	"api/src/auth"
	"api/src/database"
	"api/src/model"
	"api/src/repositories"
	"api/src/responses"
	"api/src/security"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

//Login autentica o usuario na aplicação
func Login(w http.ResponseWriter, r *http.Request) {
	bodyR, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		responses.Error(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var account model.Account
	if erro = json.Unmarshal(bodyR, &account); erro != nil {
		responses.Error(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := database.Connect()
	if erro != nil {
		responses.Error(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repository := repositories.NewAccountRepository(db)
	dbAccount, erro := repository.GetAccountByCPF(account.Cpf)
	if erro != nil {
		responses.Error(w, http.StatusInternalServerError, erro)
		return
	}

	if erro = security.VerifyPassword(dbAccount.Secret, account.Secret); erro != nil {
		responses.Error(w, http.StatusUnauthorized, erro)
		return
	}

	token, erro := auth.CreateToken(dbAccount.ID)
	if erro != nil {
		responses.Error(w, http.StatusInternalServerError, erro)
		return
	}
	w.Write([]byte(token))
}
