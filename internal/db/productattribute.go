package db

import (
	"context"

	"github.com/go-pg/pg/v9"
	"github.com/go-pg/pg/v9/orm"
)

type ProductAttributeRepo struct {
	db      orm.DB
	filters map[string][]Filter
	sort    map[string][]SortField
	join    map[string][]string
}

// NewProductAttributeRepo returns new repository
func NewProductAttributeRepo(db orm.DB) ProductAttributeRepo {
	return ProductAttributeRepo{
		db:      db,
		filters: map[string][]Filter{},
		sort: map[string][]SortField{
			Tables.ProductAttribute.Name: {{Column: Columns.ProductAttribute.ID, Direction: SortDesc}},
		},
		join: map[string][]string{
			Tables.ProductAttribute.Name: {TableColumns},
		},
	}
}

// WithTransaction is a function that wraps ProductAttributeRepo with pg.Tx transaction.
func (par ProductAttributeRepo) WithTransaction(tx *pg.Tx) ProductAttributeRepo {
	par.db = tx
	return par
}

// WithEnabledOnly is a function that adds "statusId"=1 as base filter.
func (par ProductAttributeRepo) WithEnabledOnly() ProductAttributeRepo {
	f := make(map[string][]Filter, len(par.filters))
	for i := range par.filters {
		f[i] = make([]Filter, len(par.filters[i]))
		copy(f[i], par.filters[i])
		f[i] = append(f[i], StatusEnabledFilter)
	}
	par.filters = f

	return par
}

/*** ProductAttribute ***/

// FullProductAttribute returns full joins with all columns
func (par ProductAttributeRepo) FullProductAttribute() OpFunc {
	return WithColumns(par.join[Tables.ProductAttribute.Name]...)
}

// DefaultProductAttributeSort returns default sort.
func (par ProductAttributeRepo) DefaultProductAttributeSort() OpFunc {
	return WithSort(par.sort[Tables.ProductAttribute.Name]...)
}

// ProductAttributeByID is a function that returns ProductAttribute by ID(s) or nil.
func (par ProductAttributeRepo) ProductAttributeByID(ctx context.Context, id int, ops ...OpFunc) (*ProductAttribute, error) {
	return par.OneProductAttribute(ctx, &ProductAttributeSearch{ID: &id}, ops...)
}

// OneProductAttribute is a function that returns one ProductAttribute by filters. It could return pg.ErrMultiRows.
func (par ProductAttributeRepo) OneProductAttribute(ctx context.Context, search *ProductAttributeSearch, ops ...OpFunc) (*ProductAttribute, error) {
	obj := &ProductAttribute{}
	err := buildQuery(ctx, par.db, obj, search, par.filters[Tables.ProductAttribute.Name], PagerTwo, ops...).Select()

	switch err {
	case pg.ErrMultiRows:
		return nil, err
	case pg.ErrNoRows:
		return nil, nil
	}

	return obj, err
}

// ProductAttributesByFilters returns ProductAttribute list.
func (par ProductAttributeRepo) ProductAttributesByFilters(ctx context.Context, search *ProductAttributeSearch, pager Pager, ops ...OpFunc) (productAttributes []ProductAttribute, err error) {
	err = buildQuery(ctx, par.db, &productAttributes, search, par.filters[Tables.ProductAttribute.Name], pager, ops...).Select()
	return
}

// CountProductAttributes returns count
func (par ProductAttributeRepo) CountProductAttributes(ctx context.Context, search *ProductAttributeSearch, ops ...OpFunc) (int, error) {
	return buildQuery(ctx, par.db, &ProductAttribute{}, search, par.filters[Tables.ProductAttribute.Name], PagerOne, ops...).Count()
}

// AddProductAttribute adds ProductAttribute to DB.
func (par ProductAttributeRepo) AddProductAttribute(ctx context.Context, productAttribute *ProductAttribute, ops ...OpFunc) (*ProductAttribute, error) {
	q := par.db.ModelContext(ctx, productAttribute)
	applyOps(q, ops...)
	_, err := q.Insert()

	return productAttribute, err
}

// UpdateProductAttribute updates ProductAttribute in DB.
func (par ProductAttributeRepo) UpdateProductAttribute(ctx context.Context, productAttribute *ProductAttribute, ops ...OpFunc) (bool, error) {
	q := par.db.ModelContext(ctx, productAttribute).WherePK()
	applyOps(q, ops...)
	res, err := q.Update()
	if err != nil {
		return false, err
	}

	return res.RowsAffected() > 0, err
}

// DeleteProductAttribute deletes ProductAttribute from DB.
func (par ProductAttributeRepo) DeleteProductAttribute(ctx context.Context, id int) (deleted bool, err error) {
	productAttribute := &ProductAttribute{ID: id}

	res, err := par.db.ModelContext(ctx, productAttribute).WherePK().Delete()
	if err != nil {
		return false, err
	}

	return res.RowsAffected() > 0, err
}
