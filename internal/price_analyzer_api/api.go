package price_analyzer_api

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	api "price_analyzer_prototype/api"
	"price_analyzer_prototype/internal/db"
	"strings"
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
	q := `select id, name, code, updated_at, category_id, is_active from product where is_active=true %s`
	if req.CategoryId != nil {
		whereClause := fmt.Sprintf("category_id = %s", req.CategoryId)
		q = fmt.Sprintf(q, "and %s", whereClause)
	}
	if strings.Contains(q, "%s") {
		q = strings.ReplaceAll(q, "%s", "")
	}

	var productList []*db.Product
	if err := a.db.SelectContext(ctx, &productList, q); err != nil {
		return nil, err
	}

	var productListAPI []*api.Product
	for i, _ := range productList {
		productListAPI = append(productListAPI, productList[i].ToProductAPI())
	}

	return &api.ProductListResponse{
		Items: productListAPI,
	}, nil
}

func (a *ProductAnalyzerAPI) GetCategoryList(ctx context.Context, req *api.CategoryListRequest) (*api.CategoryListResponse, error) {
	q := `select id, name, code, is_active from category where is_active=true`
	var categoryList []*db.Category
	if err := a.db.SelectContext(ctx, &categoryList, q); err != nil {
		return nil, err
	}

	var categoryListAPI []*api.Category
	for i, _ := range categoryList {
		categoryListAPI = append(categoryListAPI, categoryList[i].ToAPI())
	}

	return &api.CategoryListResponse{
		Categories: categoryListAPI,
	}, nil
}
func (a *ProductAnalyzerAPI) GetProductsFilter(ctx context.Context, req *api.GetProductsFilterRequest) (*api.GetProductsFilterResponse, error) {
	return nil, nil
}

func (a *ProductAnalyzerAPI) AddProductsFilter(ctx context.Context, req *api.AddProductsFilterRequest) (*api.AddProductsFilterResponse, error) {
	return nil, nil
}
