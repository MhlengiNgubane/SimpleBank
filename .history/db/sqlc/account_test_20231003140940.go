package db

import "testing"

func TestCreateAccount(t *testing.T) {
	arg := CreateAccountParams{
		Owner: "Asanda",
		Balance: 100,
		Currency: "ZAR",
	}

	Account, err :=
}