package service

import (
	"API/internal/models"
	"API/internal/repository"
	"context"
	"fmt"
	"math/rand"
)

type LinkService struct {
	repo repository.LinkRepository
}

func NewLinkService(repo repository.LinkRepository) *LinkService {
	return &LinkService{repo: repo}
}

func generateCode() string {
	const letters = "abcdefghijklmnopqrstuvwxyz0123456789"
	code := make([]byte, 6)

	for i := 0; i < 6; i++ {
		code[i] = letters[rand.Intn(len(letters))]
	}
	return string(code)
}

func (s *LinkService) CreateLink(ctx context.Context, original_url string, telegramID int64) (*models.Link, error) {
	if original_url == "" {
		return nil, fmt.Errorf("url is required")
	}

	link := &models.Link{
		OriginalURL: original_url,
		ShortCode:   generateCode(),
		TelegramID: telegramID,
	}

	if err := s.repo.Create(ctx, link); err != nil {
		return nil, err
	}
	return link, nil
}

func (s *LinkService) GetAllLinks(ctx context.Context, telegramID int64) ([]*models.Link, error) {
	links, err := s.repo.GetAll(ctx, telegramID)
	if err != nil {
		return nil, err
	}
	return links, nil
}

func (s *LinkService) GetByCode(ctx context.Context, code string) (*models.Link, error) {
	link, err := s.repo.GetByCode(ctx, code)
	if err != nil {
		return nil, err
	}
	return link, nil
}
func (s *LinkService) DeleteLink(ctx context.Context, code string, telegramID int64) error {
	err := s.repo.Delete(ctx, code, telegramID)
	if err != nil {
		return err
	}
	return nil
}
func (s *LinkService) IncrementClicks(ctx context.Context, code string) error {
	err := s.repo.IncrementClicks(ctx, code)
	if err != nil {
		return err
	}
	return nil
}
