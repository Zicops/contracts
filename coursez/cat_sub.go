package coursez

import "github.com/scylladb/gocqlx/v2/table"

//create table coursez.cat_sub_mapping(
//    category varchar
//    sub_category varchar,
//    PRIMARY KEY (category)
//);

var CatSubMapMeta = table.Metadata{
	Name: "cat_sub_mapping",
	Columns: []string{
		"category",
		"sub_category",
	},
	PartKey: []string{
		"category",
	},
	SortKey: []string{},
}

var CatSubMapTable = table.New(CatSubMapMeta)

type CatSubMap struct {
	Category    string `db:"category"`
	SubCategory string `db:"sub_category"`
}
