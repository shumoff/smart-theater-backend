package recommender_grpc

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"time"

	pb "github.com/shumoff/smart-theater-backend/protobuffs"
)

type RecommenderClient struct {
	Address string
}

func (rc *RecommenderClient) GetRelevantMovies(userId int32, offset int32, limit int32) ([]int32, error) {
	var movieIds []int32

	conn, err := grpc.Dial(rc.Address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return nil, fmt.Errorf("could not connect to recommender: %w", err)
	}
	defer conn.Close()
	c := pb.NewRecommenderClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.RecommendMovie(ctx, &pb.RelevantMovieRequest{UserId: userId, Offset: offset, Limit: limit})
	if err != nil {
		return nil, fmt.Errorf("could not get recommendation: %w", err)
	}

	relevantMovies := r.GetRecommendations()
	for _, relevantMovie := range relevantMovies {
		movieIds = append(movieIds, relevantMovie.Id)
	}

	return movieIds, nil
}

func (rc *RecommenderClient) GetSimilarMovies(movieId int32, offset int32, limit int32) ([]int32, error) {
	var movieIds []int32

	conn, err := grpc.Dial(rc.Address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return nil, fmt.Errorf("could not connect to recommender: %w", err)
	}
	defer conn.Close()
	c := pb.NewRecommenderClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.RecommendSimilarMovie(ctx, &pb.SimilarMovieRequest{MovieId: movieId, Offset: offset, Limit: limit})
	if err != nil {
		return nil, fmt.Errorf("could not get similar movies: %w", err)
	}

	similarMovies := r.GetRecommendations()
	for _, similarMovie := range similarMovies {
		movieIds = append(movieIds, similarMovie.Id)
	}

	return movieIds, nil
}

func (rc *RecommenderClient) GetRelevantSimilarMovies(userId int32, movieId int32, offset int32, limit int32) ([]int32, error) {
	var movieIds []int32

	conn, err := grpc.Dial(rc.Address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return nil, fmt.Errorf("could not connect to recommender: %w", err)
	}
	defer conn.Close()
	c := pb.NewRecommenderClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.RecommendRelevantSimilarMovie(
		ctx,
		&pb.RelevantSimilarMovieRequest{UserId: userId, MovieId: movieId, Offset: offset, Limit: limit},
	)
	if err != nil {
		return nil, fmt.Errorf("could not get relevant similar movies: %w", err)
	}

	relevantSimilarMovies := r.GetRecommendations()
	for _, relevantSimilarMovie := range relevantSimilarMovies {
		movieIds = append(movieIds, relevantSimilarMovie.Id)
	}

	return movieIds, nil
}

func (rc *RecommenderClient) SaveNewRating(userId int32, movieId int32, rating int32) error {
	conn, err := grpc.Dial(rc.Address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return fmt.Errorf("could not connect to recommender: %w", err)
	}
	defer conn.Close()
	c := pb.NewRecommenderClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	_, err = c.SaveNewRating(ctx, &pb.NewRatingRequest{UserId: userId, MovieId: movieId, Rating: rating})
	if err != nil {
		return fmt.Errorf("could not save rating: %w", err)
	}

	return nil
}
