package igboapi

import (
	"context"
	"fmt"
)

type TermService interface {
	Search(ctx context.Context, params *TermParams) (*[]SearchTermResponse, error)
	Find(ctx context.Context, termID string) (*Term, error)
}

type termService struct {
	client *Client
}

// Find implements TermService.
func (t *termService) Find(ctx context.Context, termID string) (*Term, error) {
	var term Term
	err := t.client.get(ctx, "words/"+termID, nil, &term)
	if err != nil {
		return nil, fmt.Errorf("get id %s failed: %w", termID, err)
	}
	return &term, nil
}

// Search implements TermService.
func (t *termService) Search(ctx context.Context, params *TermParams) (*[]SearchTermResponse, error) {
	var sr []SearchTermResponse
	err := t.client.get(ctx, "words", StructToMap(params), &sr)
	if err != nil {
		return nil, fmt.Errorf("get word %s failed: %w", params.Keyword, err)
	}
	return &sr, nil
}
// https://igboapi.com/api/v1/words?keyword=bia