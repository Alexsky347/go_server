package repository

import (
	"context"
	"go-server/internal/models"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository struct {
	DB *pgxpool.Pool
}

func (r *UserRepository) FindAll() ([]interface{}, error) {
	rows, err := r.DB.Query(context.Background(), "SELECT * FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []interface{}
	for rows.Next() {
		var user models.User
		err = rows.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.IsActive, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func (r *UserRepository) FindByID(id int) (models.User, error) {
	row := r.DB.QueryRow(context.Background(), "SELECT * FROM users WHERE id = $1", id)

	var user models.User
	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.IsActive, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return models.User{}, nil
		}
		return models.User{}, err
	}

	return user, nil
}

func (r *UserRepository) FindByEmail(email string) (interface{}, error) {
	row := r.DB.QueryRow(context.Background(), "SELECT * FROM users WHERE email = $1", email)

	var user models.User
	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.IsActive, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return user, nil
}

func (r *UserRepository) Update(id int, user models.User) error {
	_, err := r.DB.Exec(context.Background(), "UPDATE users SET name = $1, email = $2, password = $3, is_active = $4, updated_at = $5 WHERE id = $6", user.Name, user.Email, user.Password, user.IsActive, user.UpdatedAt, id)
	if err != nil {
		return err
	}

	return nil
}

func (r *UserRepository) Delete(id int) error {
    _, err := r.DB.Exec(context.Background(), "DELETE FROM users WHERE id = $1", id)
    if err != nil {
        return err
    }

    return nil
}

func (r *UserRepository) Create(data interface{}) (interface{}, error) {
    user, ok := data.(models.User)
    if !ok {
        return nil, errors.New("invalid data")
    }

    err := r.DB.QueryRow(context.Background(), "INSERT INTO users (name, email, password, is_active, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id", user.Name, user.Email, user.Password, user.IsActive, user.CreatedAt, user.UpdatedAt).Scan(&user.ID)
    if err != nil {
        return nil, err
    }

    return user, nil
}
