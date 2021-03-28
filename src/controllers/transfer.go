package controllers

import (
	"api/src/auth"
	"api/src/database"
	"api/src/model"
	"api/src/repositories"
	"api/src/responses"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

//CreateTransaction realiza uma transação entre a conta conectada e a conta de destino vinda do request
func CreateTransaction(w http.ResponseWriter, r *http.Request) {
	bodyR, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		responses.Error(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var transfer model.Transfers
	if erro = json.Unmarshal(bodyR, &transfer); erro != nil {
		responses.Error(w, http.StatusBadRequest, erro)
		return
	}

	tokenID, erro := auth.ExtractAccountId(r)
	if erro != nil {
		responses.Error(w, http.StatusUnauthorized, erro)
		return
	}

	transfer.FromAccountID = tokenID
	db, erro := database.Connect()
	if erro != nil {
		responses.Error(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	acctRepository := repositories.NewAccountRepository(db)
	fromAcct, erro := acctRepository.FindByID(tokenID)
	if erro != nil {
		responses.Error(w, http.StatusBadRequest, erro)
	}

	toAcct, erro := acctRepository.FindByID(transfer.ToAccountID)
	if erro != nil {
		responses.Error(w, http.StatusBadRequest, erro)
	}

	if erro := fromAcct.Withdraw(transfer.Amount); erro != nil {
		responses.Error(w, http.StatusBadRequest, erro)
		return
	}
	toAcct.Deposit(transfer.Amount)

	erro = acctRepository.UpdateAccountBalance(tokenID, fromAcct.Balance)
	if erro != nil {
		responses.Error(w, http.StatusInternalServerError, erro)
		return
	}
	acctRepository.UpdateAccountBalance(toAcct.ID, toAcct.Balance)
	if erro != nil {
		responses.Error(w, http.StatusInternalServerError, erro)
		return
	}

	transferRepository := repositories.NewTransferRepository(db)
	transfer.ID, erro = transferRepository.Save(transfer)
	if erro != nil {
		responses.Error(w, http.StatusInternalServerError, erro)
		return
	}

	responses.JSON(w, http.StatusCreated, transfer)

}

func GetAllCurrentAcctTransactions(w http.ResponseWriter, r *http.Request) {
	tokenID, erro := auth.ExtractAccountId(r)
	if erro != nil {
		responses.Error(w, http.StatusUnauthorized, erro)
		return
	}

	db, erro := database.Connect()
	if erro != nil {
		responses.Error(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	transferRepository := repositories.NewTransferRepository(db)
	transfers, erro := transferRepository.FindAll(tokenID)
	if erro != nil {
		responses.Error(w, http.StatusInternalServerError, erro)
		return
	}
	if transfers == nil {
		responses.JSON(w, http.StatusNotFound, "Não há transferencias para esta conta.")
		return
	}

	responses.JSON(w, http.StatusOK, transfers)

}
