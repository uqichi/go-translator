package translator

import (
	"bytes"
	"encoding/binary"
	"encoding/gob"
)

type ProductEntity struct {
	ID         int
	Name       string
	Price      int
	Cost       int
	CategoryID int
	Deleted    bool
}

type CategoryEntity struct {
	ID      int
	Name    string
	Deleted bool
}

func (p *ProductEntity) ToModelUsingGob() (*ProductModel, error) {
	b := bytes.Buffer{}
	err := gob.NewEncoder(&b).Encode(p)
	if err != nil {
		return nil, err
	}
	ret := new(ProductModel)
	err = gob.NewDecoder(&b).Decode(ret)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

func (p *ProductEntity) ToModelUsingBinary() (*ProductModel, error) {
	b := bytes.Buffer{}
	err := binary.Write(&b, binary.LittleEndian, p)
	if err != nil {
		return nil, err
	}
	ret := new(ProductModel)
	err = binary.Read(&b, binary.LittleEndian, ret)
	if err != nil {
		return nil, err
	}
	return ret, nil
}
