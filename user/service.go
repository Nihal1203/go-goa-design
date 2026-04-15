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

func (u *Service) AddPerson(ctx context.Context, p *userpackage.Person) ([]byte, error) {

	// Step 1: Check if person already exists (by email)
	checkQuery := `SELECT id FROM persons WHERE email = $1`

	var existingID int64
	err := u.db.QueryRow(ctx, checkQuery, p.Email).Scan(&existingID)

	if err != nil && err != pgx.ErrNoRows {
		// Some real DB error occurred
		return nil, fmt.Errorf("failed to check existing person: %w", err)
	}

	if err == nil {
		// Person already exists
		return nil, fmt.Errorf("person with email %s already exists with id %d", *p.Email, existingID)
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
		return nil, fmt.Errorf("failed to insert person: %w", err)
	}

	// Step 3: Return response as bytes
	response := fmt.Sprintf(`{"message": "person added successfully", "id": %d}`, newID)
	return []byte(response), nil
}
