package nihalpkg

import (
	"context"
	"errors"

	"github.com/Nihal1203/go-goa-design/gen/nihal"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

type Service struct {
	db *pgx.Conn
}

func NewService(db *pgx.Conn) *Service {
	return &Service{
		db,
	}
}

func (s *Service) AddBorrower(c context.Context, n *nihal.Borrower) (*nihal.AddBorrowerResult, error) {
	_, err := s.db.Exec(c,
		`INSERT INTO borrower (first_name, last_name, age, address, email)
		 VALUES ($1, $2, $3, $4, $5)`,
		n.FirstName, n.LastName, n.Age, n.Address, n.Email,
	)
	if err != nil {
		var pgErr *pgconn.PgError
		// unique_violation = 23505
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			statusCode := int32(409)
			return nil, &nihal.BorrowerExists{Message: "borrower with this email already exists", StatusCode: &statusCode}
		}
		statusCode := int32(500)
		return nil, &nihal.InternalServerError{Message: err.Error(), StatusCode: &statusCode}
	}

	statusCode := int32(201)
	message := "borrower added successfully"
	return &nihal.AddBorrowerResult{StatusCode: &statusCode, Message: &message}, nil
}
