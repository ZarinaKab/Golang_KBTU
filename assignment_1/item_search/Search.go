package itemsearch

import (
	"strings"

	"example.com/packages/Golang_KBTU/assignment_1/packages"
)

type Items struct{
	*packages.DatabaseItems
}

func (items *Items)ItemSearch(it string) *packages.Item{
	for _, i := range items.Items{
		if strings.Contains(i.Name, it) {
			return &i
		}
	}
	return nil
}