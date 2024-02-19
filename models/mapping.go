package models

import (
	"github.com/shouryagautam/bankdeploy/database"
	"errors"

	"github.com/google/uuid"
)

type CustomerToAccount struct{
	ID uint
	AccountID uint `pg:"on_delete:CASCADE"`
	CustomerID uint `pg:"on_delete:CASCADE"`

	Customer *Customer `pg:"rel:has-one"`
    Account  *Account  `pg:"rel:has-one"`
}

func (mapping *CustomerToAccount) Save() error {
	_, insertErr := database.Db.Model(mapping).Returning("*").Insert()

	if insertErr != nil {
		return insertErr
	}

	return nil
}


func DeleteNomineeFromAccountByID(accNumber uuid.UUID,id uint) error {

	account,err := FindAccountByAccountNumber(accNumber)

	if err != nil {
		return err
	}

    count, countErr := database.Db.Model((*CustomerToAccount)(nil)).
        Where("account_id = ?", account.ID).
        Count()
    if countErr != nil {
        return countErr
    }

    // Check if there is only one customer mapped to the account
    if count <= 1 {
        return errors.New("the account must have atleast one customer")
    }

	var mapping CustomerToAccount
	_, deleteErr := database.Db.Model(&mapping).Where("customer_id=?",id).Where("account_id=?",account.ID).Delete()
	if deleteErr != nil {
		return deleteErr
	}

	return nil
}
