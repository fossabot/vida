package meta

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMeta_SearchIMDB(t *testing.T) {
	meta := Meta{http.DefaultClient}

	suggestion, err := meta.SearchIMDB("avengers-endgame")
	assert.Nil(t, err)
	movie, err := suggestion.ToMovie("/test/path")
	assert.Nil(t, err)

	assert.Equal(t, "Avengers: Endgame", movie.Title)
	assert.Equal(t, int64(2019), movie.Year)
	assert.Equal(t, "tt4154796", movie.IMDBID)
}
