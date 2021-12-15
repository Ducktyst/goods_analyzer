package price_analyzer_api

import (
	"context"
	"github.com/jmoiron/sqlx"
	"price_analyzer_prototype/api"
	"price_analyzer_prototype/internal/db"
)

// реализация PriceAnalyzeServer
type ProductAnalyzerAPI struct {
	api.UnimplementedPriceAnalyzeServer
	db *sqlx.DB
}

func NewProductAnalyzerAPI(db *sqlx.DB) *ProductAnalyzerAPI {
	return &ProductAnalyzerAPI{
		db: db,
	}
}

func (a *ProductAnalyzerAPI) GetProductList(ctx context.Context, req *api.ProductListRequest) (*api.ProductListResponse, error) {
	q := `select id, name, code, updated_at, category_id from product where is_active=true`
	var productList []*db.Product
	if err := a.db.SelectContext(ctx, productList, q); err != nil {
		return nil, err
	}

	var productListAPI []*api.Product
	for i, _ := range productList {
		productListAPI = append(productListAPI, productList[i].ToProductAPI())
	}

	return &api.ProductListResponse{
		Items:                productListAPI,
		XXX_NoUnkeyedLiteral: struct{}{},
		XXX_unrecognized:     nil,
		XXX_sizecache:        0,
	}, nil
}
func (a *ProductAnalyzerAPI) GetCategoryList(ctx context.Context, req *api.CategoryListRequest) (*api.CategoryListResponse, error) {
	return nil, nil
}
func (a *ProductAnalyzerAPI) GetProductsFilter(ctx context.Context, req *api.GetProductsFilterRequest) (*api.GetProductsFilterResponse, error) {
	return nil, nil
}

func (a *ProductAnalyzerAPI) AddProductsFilter(ctx context.Context, req *api.AddProductsFilterRequest) (*api.AddProductsFilterResponse, error) {
	return nil, nil
}
