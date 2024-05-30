
package main

import (
    "context"
    "log"
    "net"

    "github.com/xhimanshuz/Totality-Corp-assignment/user"
    "google.golang.org/grpc"
    "google.golang.org/grpc/codes"
)

type server struct {
    user.UnimplementedUserServiceServer
    users []user.User
}

func (s *server) GetUserDetails(ctx context.Context, req *user.UserIDRequest) (*user.UserDetailsResponse, error) {
    for _, u := range s.users {
        if u.Id == req.Id {
            return &user.UserDetailsResponse{User: &u}, nil
        }
    }
    return nil, grpc.Errorf(codes.NotFound, "user not found")
}

func (s *server) GetUsersDetails(ctx context.Context, req *user.UserIDsRequest) (*user.UsersDetailsResponse, error) {
    var users []*user.User
    for _, id := range req.Ids {
        for _, u := range s.users {
            if u.Id == id {
                users = append(users, &u)
                break
            }
        }
    }
    return &user.UsersDetailsResponse{Users: users}, nil
}

func (s *server) SearchUsers(ctx context.Context, req *user.SearchRequest) (*user.UsersDetailsResponse, error) {
    var users []*user.User
    for _, u := range s.users {
        if (req.City == "" || u.City == req.City) && (req.Phone == 0 || u.Phone == req.Phone) && (req.Married == false || u.Married == req.Married) {
            users = append(users, &u)
        }
    }
    return &user.UsersDetailsResponse{Users: users}, nil
}

func main() {
    lis, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }
    s := grpc.NewServer()
    user.RegisterUserServiceServer(s, &server{
        users: []user.User{
            {Id: 1, Fname: "Steve", City: "LA", Phone: 1234567890, Height: 5.8, Married: true},
        },
    })
    log.Printf("server listening at %v", lis.Addr())
    if err := s.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}
