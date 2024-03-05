package rogo

import (
	"github.com/carlmjohnson/requests"
)

type Data[T any] struct {
	Data []T `json:"data"`
}

type Error struct {
	Errors []struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	}
}

type HasClient interface {
	getClient() *Client
}

type Pagination[T any] struct {
	Data           []T `json:"data"`
	client         *Client
	NextCursor     string `json:"nextPageCursor"`
	PreviousCursor string `json:"previousPageCursor"`
	URL            string
}

func (p *Pagination[T]) paginate(cursor string) (*Pagination[T], error) {
	client := p.client

	var pagination Pagination[T]
	err := requests.
		New(client.config).
		BaseURL(p.URL).
		Param("cursor", cursor).
		ToJSON(&pagination).
		Fetch(ctx)

	return &pagination, err
}

func (p *Pagination[T]) Next() (*Pagination[T], error) {
	return p.paginate(p.NextCursor)
}

func (p *Pagination[T]) Previous() (*Pagination[T], error) {
	return p.paginate(p.PreviousCursor)

}
