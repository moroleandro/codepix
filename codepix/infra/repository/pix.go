package repository

import (
    "fmt"
    "github.com/moroleandro/poc-codepix/codepix/domain/model"
    "gorm.io/gorm"
)

type PixKeyRepositoryDb struct {
    Db *gorm.DB
}

func (r PixKeyRepositoryDb) AddBank(bank *model.Bank) error {
    err := r.Db.Create(bank).Error
    if err != nil : nil, err

    return nil
}

func (r PixKeyRepositoryDb) AddAccount(account *model.Account) error {
    err := r.Db.Create(account).Error
    if err != nil : nil, err

    return nil
}

func (r PixKeyRepositoryDb) RegisterKey(pixkey *model.PixKey) (*model.PixKey, error) {
    err := r.Db.Create(pixkey).Error
    if err != nil : nil, err

    return pixkey, nil
}

func (r PixKeyRepositoryDb) FindKeyById(key string, kind string) (*model.PixKey, error) {
    var pixKey model.PixKey
    r.Db.Preload("Account.Bank").First(&pixKey, "kind = ? and key = ?", kind, key)

    if pixKey.ID == "" {
        return nil, fmt.Errorf("No key was found")
    }

    return &pixKey, nil
}

func (r PixKeyRepositoryDb) FindAccount(id string) (*model.Account, error) {
    var account model.Account
    r.Db.Preload("Bank").First(&account, "id = ?", id)

    if account.ID == "" {
        return nil, fmt.Errorf("No account found")
    }

    return &account, nil
}

func (r PixKeyRepositoryDb) FindBank(id string) (*model.Bank, error) {
    var bank model.Bank
    r.Db.First(&bank, "id = ?", id)

    if bank.ID == "" {
        return nil, fmt.Errorf("No bank found")
    }

    return &bank, nil
}

