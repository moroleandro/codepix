package usecase

import (
    "github.com/moroleandro/poc-codepix/codepix/domain/model"
)

type PixUseCase struct {
    PixKeyRepository model.PixKeyRepositoryInterface
}

func (p *PixUseCase) RegisterKey(key string, kind string, accountId string) (*model.PixKey, error) {
    account, err := p.PixKeyRepository.FindAccount(accountId)
    if err != nil {
        return nil, err
    }

    pixKey, err := mode.NewPixKey(kind, account, key)
    if err != nil {
        return nil, err
    }

    p.PixKeyRepository.RegisterKey(pixKey)
    if pixKey.ID == "" {
        return nil, err
    }

    return pixKey, nil
}

func (p *PixUseCase) FindKey(key string, kind string) (*model.PixKey, error) {
    pixKey, err := p.PixKeyRepository.FindKeyByKind(key, kind)
    if err != null : nil, err
    return pixKey, nil
}
