package item_filter

import "example.com/packages/Golang_KBTU/assignment_1/packages"


func (s *items) FilterByRatings(rating float64) []packages.Item {   
	result := make([]packages.Item, 0)
	for _, item := range s.Items {      
		if item.Rating <= rating {result = append(result, item)}
	}   
	return result
 }