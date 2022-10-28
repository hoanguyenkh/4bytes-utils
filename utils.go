package main

import "github.com/akrylysov/pogreb"

type EvmUtils struct {
	db pogreb.DB
}

func NewEvmUtils() (EvmUtils, error) {
	db, err := pogreb.Open("4bytes.db", nil)
	if err != nil {
		return _, err
	}
	return EvmUtils{
		db: db,
	}, nil
}
