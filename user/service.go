package user

import (
	"context"
	"fmt"

	userpackage "github.com/Nihal1203/go-goa-design/gen/user"
	"github.com/jackc/pgx/v5"
)

type Service struct {
	db *pgx.Conn
}

func NewService(db *pgx.Conn) *Service {
	return &Service{
		db,
	}
}

func (u *Service) GetUser(ctx context.Context, id string) (string, error) {
	return id + "1", nil
}

func (u *Service) PrintPerson(ctx context.Context, p *userpackage.Person) (res map[int32]*userpackage.Person, err error) {
	fmt.Println("Inside PrintPerson method")
	m := map[int32]*userpackage.Person{
		1: p,
	}

	return m, nil
}

func (u *Service) AddPerson(ctx context.Context, p *userpackage.Person) (*userpackage.AddPersonResponse, error) {

	// Step 1: Check if person already exists
	checkQuery := `SELECT id FROM persons WHERE email = $1`

	var existingID int64
	err := u.db.QueryRow(ctx, checkQuery, p.Email).Scan(&existingID)

	if err != nil && err != pgx.ErrNoRows {
		// Internal DB error
		return nil, &userpackage.InternalError{
			Message: fmt.Sprintf("failed to check existing person: %v", err),
		}
	}

	if err == nil {
		// Person already exists
		return nil, &userpackage.PersonAlreadyExists{
			Message: fmt.Sprintf("person with email %s already exists with id %d", *p.Email, existingID),
		}
	}

	// Step 2: Insert new person
	insertQuery := `
		INSERT INTO persons (name, age, mobile_no, email)
		VALUES ($1, $2, $3, $4)
		RETURNING id
	`

	var newID int64
	err = u.db.QueryRow(ctx, insertQuery,
		p.Name,
		p.Age,
		p.MobileNo,
		p.Email,
	).Scan(&newID)

	if err != nil {
		return nil, &userpackage.InternalError{
			Message: fmt.Sprintf("failed to insert person: %v", err),
		}
	}

	// Step 3: Success
	return &userpackage.AddPersonResponse{
		Success: true,
	}, nil
}

func (u *Service) GetPerson(ctx context.Context, payload *userpackage.GetPersonPayload) (*userpackage.Person, error) {

	query := `SELECT id, name, age, mobile_no, email FROM persons WHERE id=$1`

	var (
		id       int64
		name     string
		age      int64
		mobile   string
		email    string
	)

	err := u.db.QueryRow(ctx, query, payload.ID).Scan(
		&id,
		&name,
		&age,
		&mobile,
		&email,
	)

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, &userpackage.InternalError{
				Message: "person not found",
			}
		}

		return nil, &userpackage.InternalError{
			Message: fmt.Sprintf("db error: %v", err),
		}
	}

	// convert to Goa type (pointer fields)
	return &userpackage.Person{
		ID:       &id,
		Name:     &name,
		Age:      &age,
		MobileNo: &mobile,
		Email:    &email,
	}, nil
}
