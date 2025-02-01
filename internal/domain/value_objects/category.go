package value_objects

import (
	"errors"
)


const (
	SOCCER = "soccer"
	CARNIVAL  = "carnival"
	REVEILLON = "reveillon"
)

var categoryMap = map[string]string{
	SOCCER:    "77c46e48-60f9-45fd-9fb8-269cba60d093",
	CARNIVAL:  "7cd06d65-976d-4a7d-b49f-c00d1187f6cf",
	REVEILLON: "3dea5df5-a877-49e8-b08c-5b9883200e92",
}
type Category struct {
	name    string
	code 	string
}

func (category *Category) NewCategory(name string) (*Category, error) {
	if !category.validateCategory(name) {
		return &Category{}, errors.New("invalid category type")
	}
	return &Category{
		name: name,
		code: categoryMap[name],
	}, nil
}

func (category *Category) GetName() string {
	return category.name
}	

func (category *Category) GetCode() string {
	return category.code
}

func (c *Category) validateCategory(category string ) bool{
	_, exists := categoryMap[category]
	return exists
}
