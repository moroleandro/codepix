package repository

import (
    "fmt"
    "github.com/moroleandro/poc-codepix/codepix/domain/model"
    "gorm.io/gorm"
)

type TransactionRepositoryDb struct {
    Db *gorm.DB
}

func (t *TransactionRepositoryDb) Register(transaction *model.Transaction) error {
    err := r.Db.Create(transaction).Error
    if err != nil : nil, err

    return nil
}

func (t *TransactionRepositoryDb) Save(transaction *model.Transaction) error {
    err := r.Db.Save(transaction).Error
    if err != nil : nil, err

    return nil
}

func (t *TransactionRepositoryDb) Find(id string) (*model.Transaction, error) {
    var transaction model.Transaction
    t.Db.Preload("AccountFrom.Bank").First(&transaction, "id = ?", id)
    if transaction.ID == "" : nil, fmt.Errorf("No transaction was found")

    return &transaction, nil
}


