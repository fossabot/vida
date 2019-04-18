package models

// Movie represents a movie
type Movie struct {
	ID          int64  `json:"id,omitempty"` // struct tags
	IMDBID      string `json:"imdb_id,omitempty"`
	Title       string `json:"title,omitempty"`
	Synopsis    string `json:"synopsis,omitempty"`
	ImageURL    string `json:"image_url,omitempty"`
	TrailerURL  string `json:"trailer_url,omitempty"`
	PlaybackURI string `json:"playback_uri,omitempty"`
	Duration    string `json:"duration,omitempty"`
	Year        int64  `json:"year,omitempty"`
	ReleaseDate int64  `json:"release_date,omitempty"`
	CreatedAt   int64  `json:"created_at,omitempty"`
	UpdatedAt   int64  `json:"updated_at,omitempty"`
	IMDBJSON    string `json:"imdbjson,omitempty"`
}
