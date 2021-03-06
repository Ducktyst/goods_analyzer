package db

import price_analyzer_api "price_analyzer_prototype/api"

type Category struct {
	tableName struct{} `pg:"category_old,alias:t,discard_unknown_columns"`

	ID     int    `pg:"id,pk" db:"id"`
	Name   string `pg:"name" db:"name"`
	Code   string `pg:"code,use_zero" db:"code"`
	Active bool   `pg:"active,use_zero" db:"is_active"`
	//Url       string     `pg:"url,use_zero"`
	//Path      string     `pg:"path,use_zero"`
	//CreatedAt time.Time  `pg:"created_at,use_zero"`
	//UpdatedAt *time.Time `pg:"updated_at"`
	//ParentID  *int       `pg:"parent_id"`
	//Priority  int        `pg:"priority,use_zero"`
}

type ProductFilter struct {
	tableName struct{} `pg:"product_filter,alias:t,discard_unknown_columns"`

	ID   int     `pg:"id,pk"`
	Name *string `pg:"name"`
	Code *string `pg:"code"`
	Type *string `pg:"type"`
}

type ProductAttribute struct {
	tableName struct{} `pg:"product_attribute,alias:t,discard_unknown_columns"`

	ID       int     `pg:"id,pk"`
	Name     *string `pg:"name"`
	Code     *string `pg:"code"`
	Multiple *bool   `pg:"multiple"`
	Type     *string `pg:"type"`
}

type ProductAttributeValue struct {
	tableName struct{} `pg:"product_attribute_value,alias:t,discard_unknown_columns"`

	ID          string  `pg:"id,pk,type:uuid"`
	ProductID   string  `pg:"product_id,type:uuid,use_zero"`
	AttributeID int     `pg:"attribute_id,use_zero"`
	Value       *string `pg:"value"`

	Product   *Product          `pg:"fk:product_id"`
	Attribute *ProductAttribute `pg:"fk:attribute_id"`
}

// TODO: toAPI in other package. This folder for autogenerated db-structs
func (c *Category) ToAPI() *price_analyzer_api.Category {
	return &price_analyzer_api.Category{
		Id:   int32(c.ID),
		Name: c.Name,
		Code: c.Code,
	}
}
