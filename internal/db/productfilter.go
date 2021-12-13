package db

import (
	"context"

	"github.com/go-pg/pg/v9"
	"github.com/go-pg/pg/v9/orm"
)

type ProductFilterRepo struct {
	db      orm.DB
	filters map[string][]Filter
	sort    map[string][]SortField
	join    map[string][]string
}

// NewProductFilterRepo returns new repository
func NewProductFilterRepo(db orm.DB) ProductFilterRepo {
	return ProductFilterRepo{
		db:      db,
		filters: map[string][]Filter{},
		sort: map[string][]SortField{
			Tables.ProductFilter.Name: {{Column: Columns.ProductFilter.ID, Direction: SortDesc}},
		},
		join: map[string][]string{
			Tables.ProductFilter.Name: {TableColumns},
		},
	}
}

// WithTransaction is a function that wraps ProductFilterRepo with pg.Tx transaction.
func (pfr ProductFilterRepo) WithTransaction(tx *pg.Tx) ProductFilterRepo {
	pfr.db = tx
	return pfr
}

// WithEnabledOnly is a function that adds "statusId"=1 as base filter.
func (pfr ProductFilterRepo) WithEnabledOnly() ProductFilterRepo {
	f := make(map[string][]Filter, len(pfr.filters))
	for i := range pfr.filters {
		f[i] = make([]Filter, len(pfr.filters[i]))
		copy(f[i], pfr.filters[i])
		f[i] = append(f[i], StatusEnabledFilter)
	}
	pfr.filters = f

	return pfr
}

/*** ProductFilter ***/

// FullProductFilter returns full joins with all columns
func (pfr ProductFilterRepo) FullProductFilter() OpFunc {
	return WithColumns(pfr.join[Tables.ProductFilter.Name]...)
}

// DefaultProductFilterSort returns default sort.
func (pfr ProductFilterRepo) DefaultProductFilterSort() OpFunc {
	return WithSort(pfr.sort[Tables.ProductFilter.Name]...)
}

// ProductFilterByID is a function that returns ProductFilter by ID(s) or nil.
func (pfr ProductFilterRepo) ProductFilterByID(ctx context.Context, id int, ops ...OpFunc) (*ProductFilter, error) {
	return pfr.OneProductFilter(ctx, &ProductFilterSearch{ID: &id}, ops...)
}

// OneProductFilter is a function that returns one ProductFilter by filters. It could return pg.ErrMultiRows.
func (pfr ProductFilterRepo) OneProductFilter(ctx context.Context, search *ProductFilterSearch, ops ...OpFunc) (*ProductFilter, error) {
	obj := &ProductFilter{}
	err := buildQuery(ctx, pfr.db, obj, search, pfr.filters[Tables.ProductFilter.Name], PagerTwo, ops...).Select()

	switch err {
	case pg.ErrMultiRows:
		return nil, err
	case pg.ErrNoRows:
		return nil, nil
	}

	return obj, err
}

// ProductFiltersByFilters returns ProductFilter list.
func (pfr ProductFilterRepo) ProductFiltersByFilters(ctx context.Context, search *ProductFilterSearch, pager Pager, ops ...OpFunc) (productFilters []ProductFilter, err error) {
	err = buildQuery(ctx, pfr.db, &productFilters, search, pfr.filters[Tables.ProductFilter.Name], pager, ops...).Select()
	return
}

// CountProductFilters returns count
func (pfr ProductFilterRepo) CountProductFilters(ctx context.Context, search *ProductFilterSearch, ops ...OpFunc) (int, error) {
	return buildQuery(ctx, pfr.db, &ProductFilter{}, search, pfr.filters[Tables.ProductFilter.Name], PagerOne, ops...).Count()
}

// AddProductFilter adds ProductFilter to DB.
func (pfr ProductFilterRepo) AddProductFilter(ctx context.Context, productFilter *ProductFilter, ops ...OpFunc) (*ProductFilter, error) {
	q := pfr.db.ModelContext(ctx, productFilter)
	applyOps(q, ops...)
	_, err := q.Insert()

	return productFilter, err
}

// UpdateProductFilter updates ProductFilter in DB.
func (pfr ProductFilterRepo) UpdateProductFilter(ctx context.Context, productFilter *ProductFilter, ops ...OpFunc) (bool, error) {
	q := pfr.db.ModelContext(ctx, productFilter).WherePK()
	applyOps(q, ops...)
	res, err := q.Update()
	if err != nil {
		return false, err
	}

	return res.RowsAffected() > 0, err
}

// DeleteProductFilter deletes ProductFilter from DB.
func (pfr ProductFilterRepo) DeleteProductFilter(ctx context.Context, id int) (deleted bool, err error) {
	productFilter := &ProductFilter{ID: id}

	res, err := pfr.db.ModelContext(ctx, productFilter).WherePK().Delete()
	if err != nil {
		return false, err
	}

	return res.RowsAffected() > 0, err
}
