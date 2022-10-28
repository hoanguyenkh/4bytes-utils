package evm

import (
	"github.com/akrylysov/pogreb"
	"github.com/akrylysov/pogreb/fs"
	"strings"
)

type EvmUtils struct {
	db *pogreb.DB
}

func NewEvmUtils() (EvmUtils, error) {
	db, err := pogreb.Open("4bytes.db", &pogreb.Options{FileSystem: fs.OSMMap})
	if err != nil {
		return EvmUtils{}, err
	}
	return EvmUtils{
		db: db,
	}, nil
}

func (utils *EvmUtils) GetSignature(signature string) (string, error) {
	key := signature
	if strings.HasPrefix(key, "0x") {
		key = key[2:]
	}
	value, err := utils.db.Get([]byte(key))
	if err != nil {
		return "", err
	}
	return string(value), nil
}
