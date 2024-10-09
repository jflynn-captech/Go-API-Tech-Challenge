package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"time"

	"google.golang.org/grpc"
	"jf.go.techchallenge.data/internal/applog"
	"jf.go.techchallenge.data/internal/config"
	"jf.go.techchallenge.data/internal/database"
	"jf.go.techchallenge.data/internal/repository"
	pb "jf.go.techchallenge.data/protodata"
)

// func (s *server) ListTasks(ctx context.Context, req *pb.ListTasksRequest) (*pb.ListTasksResponse, error) {
// 	s.mu.Lock()
// 	defer s.mu.Unlock()

// 	var tasks []*pb.Task
// 	for id, title := range s.tasks {
// 		tasks = append(tasks, &pb.Task{Id: id, Title: title})
// 	}

// 	return &pb.ListTasksResponse{Tasks: tasks}, nil
// }

func generateID() string {
	// Simple ID generation logic (use a more robust solution in production)
	return fmt.Sprintf("%d", time.Now().UnixNano())
}

func main() {

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	logger := applog.New(log.New(os.Stdout, "\r\n", log.LstdFlags))

	config, err := config.New(logger)

	if err != nil {
		logger.Fatal("Failed to load config", err)
	}

	db, err := database.New(config, logger)

	if err != nil {
		logger.Fatal("Failed to load config", err)
	}

	personRepository := repository.NewPerson(db, logger)

	s := grpc.NewServer()
	pb.RegisterPersonRepositoryServer(s, personRepository)

	log.Println("Server is running on port 50051...")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
