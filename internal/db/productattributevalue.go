package db

import (
	"context"

	"github.com/go-pg/pg/v9"
	"github.com/go-pg/pg/v9/orm"
)

type ProductAttributeValueRepo struct {
	db      orm.DB
	filters map[string][]Filter
	sort    map[string][]SortField
	join    map[string][]string
}

// NewProductAttributeValueRepo returns new repository
func NewProductAttributeValueRepo(db orm.DB) ProductAttributeValueRepo {
	return ProductAttributeValueRepo{
		db:      db,
		filters: map[string][]Filter{},
		sort: map[string][]SortField{
			Tables.ProductAttributeValue.Name: {{Column: Columns.ProductAttributeValue.ID, Direction: SortDesc}},
		},
		join: map[string][]string{
			Tables.ProductAttributeValue.Name: {TableColumns, Columns.ProductAttributeValue.Product, Columns.ProductAttributeValue.Attribute},
		},
	}
}

// WithTransaction is a function that wraps ProductAttributeValueRepo with pg.Tx transaction.
func (pavr ProductAttributeValueRepo) WithTransaction(tx *pg.Tx) ProductAttributeValueRepo {
	pavr.db = tx
	return pavr
}

// WithEnabledOnly is a function that adds "statusId"=1 as base filter.
func (pavr ProductAttributeValueRepo) WithEnabledOnly() ProductAttributeValueRepo {
	f := make(map[string][]Filter, len(pavr.filters))
	for i := range pavr.filters {
		f[i] = make([]Filter, len(pavr.filters[i]))
		copy(f[i], pavr.filters[i])
		f[i] = append(f[i], StatusEnabledFilter)
	}
	pavr.filters = f

	return pavr
}

/*** ProductAttributeValue ***/

// FullProductAttributeValue returns full joins with all columns
func (pavr ProductAttributeValueRepo) FullProductAttributeValue() OpFunc {
	return WithColumns(pavr.join[Tables.ProductAttributeValue.Name]...)
}

// DefaultProductAttributeValueSort returns default sort.
func (pavr ProductAttributeValueRepo) DefaultProductAttributeValueSort() OpFunc {
	return WithSort(pavr.sort[Tables.ProductAttributeValue.Name]...)
}

// ProductAttributeValueByID is a function that returns ProductAttributeValue by ID(s) or nil.
func (pavr ProductAttributeValueRepo) ProductAttributeValueByID(ctx context.Context, id string, ops ...OpFunc) (*ProductAttributeValue, error) {
	return pavr.OneProductAttributeValue(ctx, &ProductAttributeValueSearch{ID: &id}, ops...)
}

// OneProductAttributeValue is a function that returns one ProductAttributeValue by filters. It could return pg.ErrMultiRows.
func (pavr ProductAttributeValueRepo) OneProductAttributeValue(ctx context.Context, search *ProductAttributeValueSearch, ops ...OpFunc) (*ProductAttributeValue, error) {
	obj := &ProductAttributeValue{}
	err := buildQuery(ctx, pavr.db, obj, search, pavr.filters[Tables.ProductAttributeValue.Name], PagerTwo, ops...).Select()

	switch err {
	case pg.ErrMultiRows:
		return nil, err
	case pg.ErrNoRows:
		return nil, nil
	}

	return obj, err
}

// ProductAttributeValuesByFilters returns ProductAttributeValue list.
func (pavr ProductAttributeValueRepo) ProductAttributeValuesByFilters(ctx context.Context, search *ProductAttributeValueSearch, pager Pager, ops ...OpFunc) (productAttributeValues []ProductAttributeValue, err error) {
	err = buildQuery(ctx, pavr.db, &productAttributeValues, search, pavr.filters[Tables.ProductAttributeValue.Name], pager, ops...).Select()
	return
}

// CountProductAttributeValues returns count
func (pavr ProductAttributeValueRepo) CountProductAttributeValues(ctx context.Context, search *ProductAttributeValueSearch, ops ...OpFunc) (int, error) {
	return buildQuery(ctx, pavr.db, &ProductAttributeValue{}, search, pavr.filters[Tables.ProductAttributeValue.Name], PagerOne, ops...).Count()
}

// AddProductAttributeValue adds ProductAttributeValue to DB.
func (pavr ProductAttributeValueRepo) AddProductAttributeValue(ctx context.Context, productAttributeValue *ProductAttributeValue, ops ...OpFunc) (*ProductAttributeValue, error) {
	q := pavr.db.ModelContext(ctx, productAttributeValue)
	applyOps(q, ops...)
	_, err := q.Insert()

	return productAttributeValue, err
}

// UpdateProductAttributeValue updates ProductAttributeValue in DB.
func (pavr ProductAttributeValueRepo) UpdateProductAttributeValue(ctx context.Context, productAttributeValue *ProductAttributeValue, ops ...OpFunc) (bool, error) {
	q := pavr.db.ModelContext(ctx, productAttributeValue).WherePK()
	applyOps(q, ops...)
	res, err := q.Update()
	if err != nil {
		return false, err
	}

	return res.RowsAffected() > 0, err
}

// DeleteProductAttributeValue deletes ProductAttributeValue from DB.
func (pavr ProductAttributeValueRepo) DeleteProductAttributeValue(ctx context.Context, id string) (deleted bool, err error) {
	productAttributeValue := &ProductAttributeValue{ID: id}

	res, err := pavr.db.ModelContext(ctx, productAttributeValue).WherePK().Delete()
	if err != nil {
		return false, err
	}

	return res.RowsAffected() > 0, err
}
