package evm

import "github.com/akrylysov/pogreb"

type EvmUtils struct {
	db *pogreb.DB
}

func NewEvmUtils() (EvmUtils, error) {
	db, err := pogreb.Open("4bytes.db", nil)
	if err != nil {
		return EvmUtils{}, err
	}
	return EvmUtils{
		db: db,
	}, nil
}

func (utils *EvmUtils) GetSignature(signature string) (string, error) {
	value, err := utils.db.Get([]byte(signature))
	if err != nil {
		return "", err
	}
	return string(value), nil
}
