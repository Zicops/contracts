package coursez

import "github.com/scylladb/gocqlx/v2/table"

var CatMeta = table.Metadata{
	Name: "category",
	Columns: []string{
		"name",
	},
	PartKey: []string{
		"name",
	},
	SortKey: []string{
		"name",
	},
}

var SubCatMeta = table.Metadata{
	Name: "sub_category",
	Columns: []string{
		"name",
	},
	PartKey: []string{
		"name",
	},
	SortKey: []string{
		"name",
	},
}

// CatTable is the table for the category table above.
var CatTable = table.New(CatMeta)
var SubCatTable = table.New(SubCatMeta)

// define struct Cat for table coursez.category
type Cat struct {
	Name string `db:"name"`
}
type SubCategory struct {
	Name string `db:"name"`
}
