package ratings

import "example.com/packages/Golang_KBTU/assignment_1/packages"

type items struct{
	*packages.DatabaseItems
}

func (s *items) GiveRating(item packages.Item, rating float64) {   
	for i, it := range s.Items {
		if it.Name == item.Name { 
			s.Items[i].Rating = rating
			break     
		}
    }
}