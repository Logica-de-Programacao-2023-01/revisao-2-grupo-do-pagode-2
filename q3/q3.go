package q3

import (
	"errors"
	"fmt"
)

//Você está desenvolvendo um sistema de gerenciamento de estoque para uma loja.
//Cada produto possui um código único, nome, preço unitário e quantidade em estoque.
//Você decidiu usar uma struct para representar as informações de cada produto.
//Agora, você precisa implementar uma função chamada "updateStock"
//que recebe como parâmetro um ponteiro para um objeto do
//tipo Product e um mapa que representa as vendas realizadas, onde as chaves são os códigos dos produtos vendidos e os
//valores são as quantidades vendidas. A função deve atualizar a quantidade em estoque do produto com base nas vendas
//realizadas.
//Sua tarefa é escrever a função "updateStock" e usá-la para atualizar a quantidade em estoque dos produtos vendidos.

type Product struct {
	Code     string
	Name     string
	Price    float64
	Quantity int
}

func UpdateStock(product *Product, sales map[string]int) error {
	if product == nil {
		return errors.New("Product pointer is nil")
	}

	for code, quantitySold := range sales {
		if quantitySold < 0 {
			return errors.New("Invalid negative quantity sold")
		}
		if code != product.Code {
			return errors.New("Produto errado.")
		}
		if quantitySold <= product.Quantity {
			product.Quantity -= quantitySold
		} else {
			return fmt.Errorf("Not enough stock for product with code %s", code)
		}
	}
	return nil
}
