package showModel

/* Table Definition */
type Show struct {
	IdShow      int    `json:"id_show"`
	ShowName    string `json:"show_name"`
	ShowDesc    string `json:"show_desc"`
	ShowArtwork string `json:"show_artwork"`
	ShowType    string `json:"show_type"`
	Explicit    string `json:"explicit"`
	Category    string `json:"category"`
	Tags        string `json:"tags"`
	DateCreated string `json:"date_created"`
	DateUpdated string `json:"date_updated"`
	Status      string `json:"status"`
}

func (Show) TableName() string {
	return "kbr_shows"
}
