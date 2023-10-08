package db

import (
	"context"
	"testing"
	
	"github.com/stretchr/testify"
)

func TestCreateAccount(t *testing.T) {
	arg := CreateAccountParams{
		Owner: "Asanda",
		Balance: 100,
		Currency: "ZAR",
	}

	Account, err := testQueries.CreateAccount(context.Background(), arg)

}