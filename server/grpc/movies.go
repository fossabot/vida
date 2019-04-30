package grpc

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"

	"github.com/gangachris/vida/db"
	"github.com/gangachris/vida/meta"
	"github.com/gangachris/vida/models"
	"github.com/gangachris/vida/pb"
)

type moviesRequestServer struct {
	store *db.Storage
}

func (m *moviesRequestServer) SearchMovies(req *pb.SearchMovieRequest, server pb.MoviesRequests_SearchMoviesServer) error {
	movieCh := make(chan models.Movie)
	errCh := make(chan error)
	doneCh := make(chan struct{})

	go meta.SearchMoviesFromDir(req.GetPath(), movieCh, doneCh, errCh)

	for {
		select {
		case movie := <-movieCh:
			if err := server.Send(MovieToPbMovie(movie)); err != nil {
				return err
			}
		case err := <-errCh:
			if err != nil {
				return err
			}
		case <-doneCh:
			return nil
		}
	}
}

func (m *moviesRequestServer) ListMovies(ctx context.Context, e *empty.Empty) (*pb.SearchMovieResponse, error) {
	// select all movies from the database and list them
	movies, err := models.AllMovies(ctx, m.store)
	if err != nil {
		return nil, err
	}

	moviesResponse := make([]*pb.Movie, len(movies))
	for idx, movie := range movies {
		moviesResponse[idx] = MovieToPbMovie(movie)
	}

	response := &pb.SearchMovieResponse{
		Movies: moviesResponse,
	}

	return response, nil
}

// MovieToPbMovie converts a movie model to a protocol buffer movie
// conversion 🙄
func MovieToPbMovie(m models.Movie) *pb.Movie {
	return &pb.Movie{
		Id:          m.ID,
		ImdbId:      m.IMDBID,
		Title:       m.Title,
		Synopsis:    m.Synopsis,
		ImageUrl:    m.ImageURL,
		TrailerUrl:  m.TrailerURL,
		Starring:    m.Starring,
		Duration:    m.Duration,
		Year:        m.Year,
		ReleaseDate: m.ReleaseDate,
		Search:      m.Search,
		PlaybackUri: m.PlaybackURI,
		CreatedAt:   m.CreatedAt,
		UpdatedAt:   m.UpdatedAt,
		ImdbJson:    m.IMDBJSON, // not sure whether we need this
	}
}
