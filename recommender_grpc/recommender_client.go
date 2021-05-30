package recommender_grpc

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"time"

	pb "github.com/shumoff/smart-theater-backend/protobuffs"
)

type RecommenderClient struct {
	Address string
}

func (rc *RecommenderClient) GetRelevantMovies(userId int32, offset int32, limit int32) (movieIds []int32) {
	conn, err := grpc.Dial(rc.Address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewRecommenderClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.RecommendMovie(ctx, &pb.RelevantMovieRequest{UserId: userId, Offset: offset, Limit: limit})
	if err != nil {
		log.Fatalf("could not get recommendation: %v", err)
	}

	relevantMovies := r.GetRecommendations()
	for _, relevantMovie := range relevantMovies {
		movieIds = append(movieIds, relevantMovie.Id)
	}

	return movieIds
}
