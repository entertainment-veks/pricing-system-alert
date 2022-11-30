package repository

import (
	"errors"
	"pricing-system-alert-service/domain"

	"github.com/golang-collections/collections/stack"
)

type DataBase struct {
	BTC stack.Stack
	ETH stack.Stack
	BNB stack.Stack
}

func NewDataBase() *DataBase {
	return &DataBase{
		BTC: *stack.New(),
		ETH: *stack.New(),
		BNB: *stack.New(),
	}
}

func (r *DataBase) AddNoteBTC(note *domain.PriceNote) {
	r.BTC.Push(note)
}

func (r *DataBase) GetLastNoteBTC() (*domain.PriceNote, error) {
	if v, ok := r.BTC.Peek().(*domain.PriceNote); ok {
		return v, nil
	}
	return nil, errors.New("casting error, there is non *domain.PriceNote value at in-memory-db.go DataBase.BTC stack")
}

func (r *DataBase) AddNoteETH(note *domain.PriceNote) {
	r.ETH.Push(note)
}

func (r *DataBase) GetLastNoteETH() (*domain.PriceNote, error) {
	if v, ok := r.ETH.Peek().(*domain.PriceNote); ok {
		return v, nil
	}
	return nil, errors.New("casting error, there is non *domain.PriceNote value at in-memory-db.go DataBase.ETH stack")
}

func (r *DataBase) AddNoteBNB(note *domain.PriceNote) {
	r.BNB.Push(note)
}

func (r *DataBase) GetLastNoteBNB() (*domain.PriceNote, error) {
	if v, ok := r.BNB.Peek().(*domain.PriceNote); ok {
		return v, nil
	}
	return nil, errors.New("casting error, there is non *domain.PriceNote value at in-memory-db.go DataBase.BNB stack")
}
