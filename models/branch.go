package models

import (
	"github.com/shouryagautam/bankdeploy/database"
	"context"
	"errors"

	"github.com/go-pg/pg/v10/orm"
	"github.com/google/uuid"
)

type Branch struct {
	ID uint
	Address string 
	BankID uint `pg:"on_delete:CASCADE"`
	Bank *Bank `pg:"rel:has-one"`
	IFSC_CODE uuid.UUID `pg:"type:uuid"`
	Account []*Account `pg:"rel:has-many"`
	Customer []*Customer `pg:"rel:has-many"`
}

func (branch *Branch) Save() (*Branch, error) {
	_, insertErr := database.Db.Model(branch).Returning("*").Insert()

	if insertErr != nil {
		return nil,insertErr
	}

	return branch, nil
}

func (branch *Branch) BeforeInsert (context context.Context) (context.Context,error) {

	branch.IFSC_CODE = uuid.New()
	return context,nil

}



func FindAllBranches() ([]Branch,error) {
	var branches []Branch
	getErr := database.Db.Model(&branches).
		Select()

	if getErr != nil {
		return nil,getErr
	}

	return branches,nil
}


func FindBranchByID(id uint) (*Branch, error){
	var output Branch
	getErr := database.Db.Model(&output).
		Where("id = ?",id).
		Select()

	if getErr != nil {
		return &Branch{},getErr
	}

	return &output,nil
}

func FindAllBranchesByBankID(id uint) ([]Branch,error) {
	var branches []Branch
	getErr := database.Db.Model(&branches).
		Where("bank_id =?",id).
		Select()

	if getErr != nil {
		return nil,getErr
	}

	return branches,nil
}

func DeleteAllBranches()  error {
	var branch Branch

	opts := &orm.DropTableOptions{
		IfExists: true,
		Cascade: true,
	}

	deleteErr := database.Db.Model(&branch).DropTable(opts)
	if deleteErr != nil {
		return deleteErr
	}

	return nil
}

func DeleteBranchByID(id uint) (*Branch, error) {
	var branch Branch
	_, deleteErr := database.Db.Model(&branch).Where("id=?",id).Returning("*").Delete(&branch)
	if deleteErr != nil {
		return nil,deleteErr
	}

	return &branch,nil
}

func (branch *Branch) Update() (*Branch, error)  {
	tx, txErr := database.Db.Begin()
	if txErr != nil {
		return nil,txErr
	}

	updateResult, updateErr := tx.Model(branch).WherePK().Returning("*").UpdateNotZero(branch)

	
	if updateErr != nil {
		tx.Rollback()
		return nil,updateErr
	}
	
	if updateResult.RowsAffected() == 0 {
		tx.Rollback()
		return nil, errors.New("no record updated")
	}
	
	tx.Commit()
	
	return branch,nil
}