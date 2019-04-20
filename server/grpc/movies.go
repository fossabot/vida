package grpc

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"

	"github.com/gangachris/vida/db"
	"github.com/gangachris/vida/models"
	"github.com/gangachris/vida/pb"
)

type moviesRequestServer struct {
	store *db.Storage
}

func (m *moviesRequestServer) SearchMovies(req *pb.SearchMovieRequest, server pb.MoviesRequests_SearchMoviesServer) error {
	// do the search
	// send to the database
	// stream results
	panic("implement me")
}

func (m *moviesRequestServer) ListMovies(ctx context.Context, e *empty.Empty) (*pb.SearchMovieResponse, error) {
	// select all movies from the database and list them
	movies, err := models.AllMovies(ctx, m.store)
	if err != nil {
		return nil, err
	}

	// conversion 🙄
	moviesResponse := make([]*pb.Movie, len(movies))
	for idx, movie := range movies {
		m := &pb.Movie{
			Id:          movie.ID,
			ImdbId:      movie.IMDBID,
			Title:       movie.Title,
			Synopsis:    movie.Synopsis,
			ImageUrl:    movie.ImageURL,
			TrailerUrl:  movie.TrailerURL,
			Starring:    movie.Starring,
			Duration:    movie.Duration,
			Year:        movie.Year,
			ReleaseDate: movie.ReleaseDate,
			Search:      movie.Search,
			PlaybackUri: movie.PlaybackURI,
			CreatedAt:   movie.CreatedAt,
			UpdatedAt:   movie.UpdatedAt,
			ImdbJson:    movie.IMDBJSON, // not sure whether we need this
		}
		moviesResponse[idx] = m
	}

	response := &pb.SearchMovieResponse{
		Movies: moviesResponse,
	}

	return response, nil
}
