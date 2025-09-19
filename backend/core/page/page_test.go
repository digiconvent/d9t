package pagination_test

import (
	"testing"

	pagination "github.com/digiconvent/d9t/core/page"
)

func TestPageStruct(t *testing.T) {
	type X struct {
		Prop int
	}

	items := []*X{
		{Prop: 1},
		{Prop: 2},
		{Prop: 4},
		{Prop: 8},
	}

	page := pagination.New(1, len(items), 3, items)

	if len(page.Items) != 3 {
		t.Fatal("Expected 3 items")
	}
	page = pagination.New(1, len(items), 2, items)
	if len(page.Items) != 2 {
		t.Fatal("Expected 2 items")
	}

	page = pagination.New(1, 2, 4, items)
	if page.ItemsCount != 4 {
		t.Fatal("Expected count of 4 items")
	}
}
