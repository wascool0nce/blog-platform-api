package service

import "github.com/wascool0nce/blog-platform-api/intenal/model"

type PostRepository interface {
	CreatePost(title, content, category string) (*model.Post, error)
}

type PostService struct {
	postRepo PostRepository
}

func NewPostService(postRepo PostRepository) *PostService {
	return &PostService{
		postRepo: postRepo,
	}
}
func (s *PostService) CreatePost(title, content, category string) (*model.Post, error) {
	post, err := s.postRepo.CreatePost(title, content, category)
	if err != nil {
		return nil, err
	}

	return post, nil
}
