package model

type (
	Item struct {
		ID          uint64 `json:"id"`
		Name        string `json:"name"`
		Description string `json:"description"`
		Price       uint    `json:"price"`
		Picture     string `json:"picture"`
	}

	ItemFilter struct {
		Name        string `query:"name" validate:"omitempty,max=64"`
		Description string `query:"description" validate:"omitempty,max=128"`
		ItemPaginate
	}

	ItemPaginate struct{
		Page int64 `query:"page" validate:"required,min=1"`
		Size int64 `query:"size" validate:"required,min=1,max=20"`
	}

	ItemResult struct{
		Items []*Item `json:"items"`
		Paginate PaginateResult `json:"paginate"`
	}

	PaginateResult struct{
		Page int64 `json:"page"`
		TotalPage int64 `json:"total_page"`
	}
)