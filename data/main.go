package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	pb "jf.go.techchallenge.data/protodata"
	"jf.go.techchallenge.data/repository"

	"google.golang.org/grpc"
)

type PersonRepositoryServer struct {
	pb.UnimplementedPersonServiceServer
	repository repository.Person
}

func (s *PersonRepositoryServer) FindAllPeople(ctx context.Context, req *pb.Filters) (*pb.PersonList, error) {
	peoples := []*pb.Person{}

	people, err := s.repository.FindAll(req.Filters)

	if err != nil {
		return nil, err
	}

	for _, person := range people {
		peoples = append(peoples, &pb.Person{
			ID:        uint64(person.ID),
			Guid:      person.Guid,
			FirstName: person.FirstName,
			LastName:  person.LastName,
			Email:     person.Email,
			Age:       uint32(person.Age),
		})
	}
	return &pb.PersonList{
		People: peoples,
	}, nil
}

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

    

	// lis, err := net.Listen("tcp", ":50051")
	// if err != nil {
	// 	log.Fatalf("failed to listen: %v", err)
	// }

	// s := grpc.NewServer()
	// pb.RegisterTaskServiceServer(s, &server{tasks: make(map[string]string)})

	// log.Println("Server is running on port 50051...")
	// if err := s.Serve(lis); err != nil {
	// 	log.Fatalf("failed to serve: %v", err)
	// }
}
