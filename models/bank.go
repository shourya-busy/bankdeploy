package models

import (
	"github.com/shouryagautam/bankdeploy/database"
	"errors"
	"fmt"

	"github.com/go-pg/pg/v10/orm"
)

type Bank struct {
	ID uint
	Name string 
	Branch []*Branch `pg:"rel:has-many"`
}

func (bank *Bank) Save() (*Bank, error) {
	_, insertErr := database.Db.Model(bank).Returning("*").Insert()

	if insertErr != nil {
		return nil,insertErr
	}

	return bank, nil
}

func FindAllBanks() ([]Bank,error) {
	var banks []Bank
	getErr := database.Db.Model(&banks).
		Select()


	if getErr != nil {
		return nil,getErr
	}

	
	return banks,nil
}

func FindBankByID(id uint) (*Bank, error){
	var bank Bank
	print(id)
	getErr := database.Db.Model(&bank).
		Where("id = ?",id).
		Select()
	
	fmt.Println(bank)
	if getErr != nil {
		return nil,getErr
	}

	return &bank,nil

}

func DeleteAllBanks()  error {
	var bank Bank

	opts := &orm.DropTableOptions{
		IfExists: true,
		Cascade: true,
	}

	deleteErr := database.Db.Model(&bank).DropTable(opts)
	if deleteErr != nil {
		return deleteErr
	}

	return nil
}

func DeleteBankByID(id uint) (*Bank, error) {
	var bank Bank
	_, deleteErr := database.Db.Model(&bank).Where("id=?",id).Returning("*").Delete(&bank)
	if deleteErr != nil {
		return nil,deleteErr
	}

	return &bank,nil
}

func (bank *Bank) Update() (*Bank, error)  {
	tx, txErr := database.Db.Begin()
	if txErr != nil {
		return nil,txErr
	}

	updateResult, updateErr := tx.Model(bank).WherePK().Returning("*").UpdateNotZero(bank)
	if updateErr != nil {
		tx.Rollback()
		return nil,updateErr
	}

	if updateResult.RowsAffected() == 0 {
		tx.Rollback()
		return nil, errors.New("no record updated")
	}

	tx.Commit()
	
	return bank,nil
}