package models

import "layeredsample/storage"

type CustomerModel struct {
	db storage.DB
}

func (c *CustomerModel) Save() {

}
