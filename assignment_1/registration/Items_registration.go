package registration

import (
	"example.com/packages/Golang_KBTU/assignment_1/packages"
)

type dbItems struct {
	*packages.DatabaseItems
}

func (d *dbItems) RegisterItem(item packages.Item) {
	d.Items = append(d.Items, item)
}