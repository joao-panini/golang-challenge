package controllers

import (
	"api/src/database"
	"api/src/model"
	"api/src/repositories"
	"api/src/responses"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

//CreateAccount creates an account on the database
func CreateAccount(w http.ResponseWriter, r *http.Request) {

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
	fmt.Println(account)
	if erro = account.Prepare("cadastro"); erro != nil {
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
	_, erro = repository.FindByCPF(account.Cpf)
	if erro != nil {
		responses.Error(w, http.StatusConflict, erro)
	}

	account.ID, erro = repository.Save(account)
	if erro != nil {
		responses.Error(w, http.StatusInternalServerError, erro)
		return
	}

	responses.JSON(w, http.StatusCreated, account)
}

//GetAccounts returns all accounts saved on the database to the response
func GetAccounts(w http.ResponseWriter, r *http.Request) {
	db, erro := database.Connect()
	if erro != nil {
		responses.Error(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()
	repository := repositories.NewAccountRepository(db)
	accounts, erro := repository.FindAll()
	if erro != nil {
		responses.Error(w, http.StatusInternalServerError, erro)
		return
	}
	responses.JSON(w, http.StatusOK, accounts)
}

//GetAccountBalanceById retorna o saldo da conta passada no parametro
func GetAccountBalanceById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	accountID, erro := strconv.ParseUint(params["accountID"], 10, 64)
	if erro != nil {
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
	account, erro := repository.FindBalanceById(accountID)
	if erro != nil {
		responses.Error(w, http.StatusInternalServerError, erro)
		return
	}

	responses.JSON(w, http.StatusOK, account)
}

func GetAccountByID(w http.ResponseWriter, r *http.Request) {
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
	account, erro = repository.FindByID(account.ID)
	if erro != nil {
		responses.Error(w, http.StatusInternalServerError, erro)
		return
	}

	responses.JSON(w, http.StatusCreated, account)
}
