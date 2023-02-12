package item_filter

import "example.com/packages/Golang_KBTU/assignment_1/packages"

type items struct{
	*packages.DatabaseItems
}

func (s *items) FilterByPrice(price float64) []packages.Item {   
	result := make([]packages.Item, 0)
	for _, item := range s.Items {      
		if item.Price <= price {result = append(result, item)}
	}   
	return result
 }