
package main

import (
    "context"
    "testing"

    "github.com/xhimanshuz/Totality-Corp-assignment/user"
    "google.golang.org/grpc"
    "google.golang.org/grpc/credentials/insecure"
)

func TestGetUserDetails(t *testing.T) {
    conn, err := grpc.Dial(":50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
    if err != nil {
        t.Fatalf("did not connect: %v", err)
    }
    defer conn.Close()
    c := user.NewUserServiceClient(conn)

    res, err := c.GetUserDetails(context.Background(), &user.UserIDRequest{Id: 1})
    if err != nil {
        t.Errorf("could not get user details: %v", err)
    }
    if res.User.Fname != "Steve" {
        t.Errorf("expected Steve, got %s", res.User.Fname)
    }
}

func TestSearchUsers(t *testing.T) {
    conn, err := grpc.Dial(":50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
    if err != nil {
        t.Fatalf("did not connect: %v", err)
    }
    defer conn.Close()
    c := user.NewUserServiceClient(conn)

    res, err := c.SearchUsers(context.Background(), &user.SearchRequest{City: "LA"})
    if err != nil {
        t.Errorf("could not search users: %v", err)
    }
    if len(res.Users) == 0 {
        t.Errorf("expected users, got none")
    }
}
