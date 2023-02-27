package categoriesModel

/* Table Definition */
type Categories struct {
	IdCategory    int64  `json:"id_category"`
	CategoryName  string `json:"category_name"`
	CategoryDesc  string `json:"category_desc"`
	CategoryScope string `json:"category_scope"`
	DateCreated   string `json:"date_created"`
	DateUpdated   string `json:"date_updated"`
	Status        string `json:"status"`
}

func (Categories) TableName() string {
	return "kbr_categories"
}
