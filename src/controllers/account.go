package controllers

import (
	"api/src/database"
	"api/src/model"
	"api/src/repositories"
	"api/src/responses"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

//CreateAccount cria uma conta no banco de dados
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
	account.ID, erro = repository.Create(account)
	if erro != nil {
		responses.Error(w, http.StatusInternalServerError, erro)
		return
	}

	responses.JSON(w, http.StatusCreated, account)
}

func GetAccounts(w http.ResponseWriter, r *http.Request) {
	db, erro := database.Connect()
	if erro != nil {
		responses.Error(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()
	repository := repositories.NewAccountRepository(db)
	accounts, erro := repository.GetAllAccounts()
	if erro != nil {
		responses.Error(w, http.StatusInternalServerError, erro)
		return
	}

	responses.JSON(w, http.StatusOK, accounts)

	w.Write([]byte("buscando accounts"))
}

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
	account, erro := repository.GetAccountBalanceById(accountID)
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
	account, erro = repository.GetAccountByID(account.ID)
	if erro != nil {
		responses.Error(w, http.StatusInternalServerError, erro)
		return
	}

	responses.JSON(w, http.StatusCreated, account)
}

func UpdateAccount(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	accountID, erro := strconv.ParseUint(params["accountID"], 10, 64)
	if erro != nil {
		responses.Error(w, http.StatusBadRequest, erro)
		return
	}

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

	if erro = account.Prepare("edit"); erro != nil {
		responses.Error(w, http.StatusInternalServerError, erro)
		return
	}

	db, erro := database.Connect()
	if erro != nil {
		responses.Error(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repository := repositories.NewAccountRepository(db)
	if erro = repository.UpdateAccount(accountID, account); erro != nil {
		responses.Error(w, http.StatusInternalServerError, erro)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)

}
