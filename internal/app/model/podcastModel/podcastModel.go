package podcastModel

/* Table Definition */
type Podcast struct {
	IdPodcast       string `json:"id_podcast"`
	IdShow          string `json:"id_show"`
	PodcastTitle    string `json:"podcast_title"`
	PodcastDesc     string `json:"podcast_desc"`
	PodcastCategory string `json:"podcast_category"`
	PublishDate     string `json:"publish_date"`
	UpdatedDate     string `json:"updated_date"`
	PublishTime     string `json:"publish_time"`
	EpisodeNumber   string `json:"episode_number"`
	EpisodeType     string `json:"episode_type"`
	EpisodeCover    string `json:"episode_cover"`
	Tags            string `json:"tags"`
	Author          string `json:"author"`
	Featured        string `json:"featured"`
	Status          string `json:"status"`
	Audio           string `json:"audio"`
	Oss             string `json:"oss"`
	Script          string `json:"script"`
}

func (Podcast) TableName() string {
	return "kbr_podcasts"
}
