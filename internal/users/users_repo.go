package users

import (
	"database/sql"
	"fmt"
)

type UserRepository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func scanRowIntoUser(row *sql.Rows) (*User, error) {
	user := new(User)

	err := row.Scan(
		&user.ID,
		&user.Fullname,
		&user.Email,
		&user.PhoneNumber,
		&user.Password,
		&user.ImageURL,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserRepository) GetAllUser() (users []*User, err error) {
	rows, err := s.db.Query("SELECT * FROM users")
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		u, err := scanRowIntoUser(rows)
		if err != nil {
			return nil, err
		}
		users = append(users, u)
	}

	return users, nil
}

func (s *UserRepository) GetUserByEmail(email string) (*User, error) {
	rows, err := s.db.Query("SELECT * FROM users WHERE email = $1", email)
	if err != nil {
		return nil, err
	}

	u := new(User)
	for rows.Next() {
		u, err = scanRowIntoUser(rows)
		if err != nil {
			return nil, err
		}
	}

	if u.Email == "" {
		return nil, fmt.Errorf("User not found")
	}

	return u, nil
}

func (s *UserRepository) GetUserByID(id int) (*User, error) {
	rows, err := s.db.Query("SELECT * FROM users WHERE id = $1", id)
	if err != nil {
		return nil, err
	}

	u := new(User)
	for rows.Next() {
		u, err = scanRowIntoUser(rows)
		if err != nil {
			return nil, err
		}
	}

	if u.ID == 0 {
		return nil, fmt.Errorf("User not found")
	}

	return u, nil
}

func (s *UserRepository) CreateUser(user *CreateUserPayload) error {
	_, err := s.db.Exec("INSERT INTO users (fullname, email, phone_number, password, image_url) VALUES ($1, $2, $3, $4, $5)", user.Fullname, user.Email, user.PhoneNumber, user.Password, user.ImageURL)
	if err != nil {
		return err
	}

	return nil
}

func (s *UserRepository) UpdateUser(user *UpdateUserPayload) error {
	fmt.Println(user)
	_, err := s.db.Exec("UPDATE users SET fullname = $1, email = $2, phone_number = $3,  password = $4, image_url = $5 WHERE id = $6", user.Fullname, user.Email, user.PhoneNumber, user.Password, user.ImageURL, user.ID)
	if err != nil {
		return err
	}

	return nil
}

func (s *UserRepository) DeleteUser(id int) error {
	_, err := s.db.Exec("DELETE FROM users WHERE id = $1", id)
	if err != nil {
		return err
	}

	return nil
}
