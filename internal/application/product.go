package application

import (
	"errors"
	"github.com/asaskevich/govalidator"
)

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

type ProductInterface interface {
	IsValid() (bool, error)
	Enable() error
	Disable() error
	GetId() string
	GetStatus() string
	GetPrice() float64
}

const (
	DISABLE = "disable"
	ENABLE  = "enable"
)

type Product struct {
	ID     string  `valid:"uuidv4"`
	Name   string  `valid:"required"`
	Price  float64 `valid:"float,optional"`
	Status string  `valid:"required"`
}

func (p *Product) IsValid() (bool, error) {
	if p.Name == "" {
		p.Status = DISABLE
	}

	if p.Price < 0 {
		return false, errors.New("preco precisa ser maior que 0")
	}

	if p.Status != DISABLE && p.Status != ENABLE {
		return false, errors.New("Status precisa ser enable ou disable")
	}
	_, err := govalidator.ValidateStruct(p)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (p *Product) Disable() error {
	if p.Price == 0 {
		p.Status = DISABLE
		return nil
	}
	return errors.New("Prico precisa ser 0")
}

func (p *Product) Enable() error {
	if p.Price > 0 {
		p.Status = ENABLE
		return nil
	}
	return errors.New("Price precisa ser maior que 0")
}

func (p *Product) GetId() string {
	return p.ID
}

func (p *Product) GetStatus() string {
	return p.Status
}

func (p *Product) GetPrice() float64 {
	return p.Price
}
