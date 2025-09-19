package pagination

import "slices"

type Page[T any] struct {
	Page         int  `json:"page"`
	ItemsCount   int  `json:"items_count"`
	ItemsPerPage int  `json:"items_per_page"`
	Items        []*T `json:"items"`
}

func New[T any](pageNumber, itemsCount, itemsPerPage int, items []*T) *Page[T] {
	if len(items) > itemsPerPage {
		var i int = 0
		items = slices.DeleteFunc(items, func(e *T) bool {
			i += 1
			return i > itemsPerPage
		})
	}
	if itemsCount < len(items) {
		itemsCount = len(items)
	}
	return &Page[T]{
		Items:        items,
		ItemsCount:   itemsCount,
		ItemsPerPage: itemsPerPage,
		Page:         pageNumber,
	}
}
