package comment

import (
    "context"
    "errors"
    "fmt"
    "go/types"
)

var (
    ErrFetchingComment = errors.New("Failed to fetch Comment by id")
    ErrNotImplemented = errors.New("Error not implemente d")
)

type Comment struct {
    ID string
    Slug string
    Body string
    Author string
}

type Store interface{
    GetComment(context.Context, string) (Comment, error)
}

type Service struct {
    Store Store
}

func NewService(store Store) *Service {
    return &Service{
        Store: store,
    }
}

func (s *Service) GetComment(ctx context.Context, ID string) (Comment, error) {
    fmt.Println("retrieving a comment")

    cmt, err := s.Store.GetComment(ctx, ID)
    if err!= nil {
        fmt.Println(err)
        return Comment{}, ErrFetchingComment
    }
    return cmt , nil
}

func (s *Service) UpdateComment(ctx context.Context, cmt Comment ) error {
    return ErrNotImplemented
}


func (s *Service) DeleteComment(ctx context.Context, id string ) error {
    return ErrNotImplemented
}

func (s *Service) CreateComment(ctx context.Context, cmt Comment )  (Comment, error) {
    return Comment{}, ErrNotImplemented
}