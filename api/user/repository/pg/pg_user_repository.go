package pg

import (
	"context"
	"database/sql"
	"errors"

	"github.com/Lorengamboa/whitelabel-travel/user/domain"
	"github.com/google/uuid"
)

type pgUserRepository struct {
	Conn *sql.DB
}

func NewPgUserRepository(conn *sql.DB) domain.UserRepository {
	return &pgUserRepository{conn}
}

func (pg *pgUserRepository) GetById(ctx context.Context, id uuid.UUID) (*domain.User, error) {
	query := `
	SELECT
		u.*, p.*
	FROM
		users u
		LEFT JOIN user_profile p ON p.user_id = u.id
	WHERE
		u.is_active = true AND u.id = $1
	`

	user := domain.User{
		Profile: &domain.Profile{},
	}

	err := pg.Conn.QueryRowContext(ctx, query, id).Scan(&user.ID,
		&user.Email, &user.Password.Hash, &user.FirstName, &user.LastName, &user.IsActive, &user.Role, &user.Thumbnail, &user.DateJoined, &user.Profile.ID, &user.Profile.UserID, &user.Profile.PhoneNumber, &user.Profile.BirthDate,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("record not found")
		} else {
			return nil, err
		}
	}

	return &user, nil
}

func (pg *pgUserRepository) GetAll(ctx context.Context) ([]domain.User, error) {
	query := `
		SELECT 
			u.*, p.* 
		FROM 
			users u 
			LEFT JOIN user_profile p ON p.user_id = u.id 
		WHERE 
			u.is_active = true
	`
	var users []domain.User

	rows, err := pg.Conn.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		user := domain.User{
			Profile: &domain.Profile{},
		}

		err := rows.Scan(&user.ID, &user.Email, &user.Password.Hash, &user.FirstName, &user.LastName, &user.IsActive, &user.Role, &user.Thumbnail, &user.DateJoined, &user.Profile.ID, &user.Profile.UserID, &user.Profile.PhoneNumber, &user.Profile.BirthDate)

		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}
