package entities

// Movie represents a movie entity
type Movie struct {
	ID          int64  `db:"id"` // struct tags
	IMDBID      string `db:"imdb_id"`
	Title       string `db:"title"`
	Synopsis    string `db:"synopsis"`
	ImageURL    string `db:"image_url"`
	TrailerURL  string `db:"trailer_url"`
	PlaybackURI string `db:"playback_uri"`
	Starring    string `db:"starring"`
	Duration    string `db:"duration"`
	Year        int64  `db:"year"`
	ReleaseDate int64  `db:"release_date"`
	CreatedAt   int64  `db:"created_at"`
	UpdatedAt   int64  `db:"update_at"`
	IMDBJSON    string `db:"imdb_json"`
	Search      string `db:"search"`
}
