package repository

import (
	"database/sql"
	"simple-CRUD/pkg/entity"
)

type Repository struct {
	db *sql.DB
}

func (r *Repository) CreateUser(name, email string, age int, hashedPassword string) error {
	_, err := r.db.Exec("INSERT INTO users (name, email, age, hashedPassword) VALUES ($1, $2, $3, $4)", name, email, age)
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) GetLoginInfoByEmail(email string) (*entity.LoginInfo, error) {
	var pass string
	var IsAdmin bool
	var id int
	err := r.db.QueryRow("SELECT id, password, is_admin FROM users WHERE email = $1", email).Scan(&id, &pass, &IsAdmin)
	if err != nil {
		return nil, err
	}
	return &entity.LoginInfo{
		Password: pass,
		IsAdmin:  IsAdmin,
		Id:       id,
	}, nil
}

func (r *Repository) GetAllUsers() ([]entity.User, error) {
	rows, err := r.db.Query("SELECT id, name, email, age FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []entity.User
	for rows.Next() {
		var user entity.User
		if err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Age); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (r *Repository) GetUserByID(id int) (*entity.User, error) {
	row := r.db.QueryRow("SELECT name, age, email FROM users WHERE id = $1", id)

	var user entity.User
	if err := row.Scan(&user.Name, &user.Age, &user.Email); err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *Repository) UpdateUser(id int, name, email string, age int) error {
	_, err := r.db.Exec("UPDATE users SET name = $1, email = $2, age = $3 WHERE id = $4", name, email, age, id)
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) DeleteUser(id int) error {
	_, err := r.db.Exec("DELETE FROM users WHERE id = $1", id)
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) GetUserByEmail(email string) (*entity.User, error) {
	row := r.db.QueryRow("SELECT id, name, age FROM users WHERE email = $1", email)

	var user entity.User
	if err := row.Scan(&user.ID, &user.Name, &user.Age); err != nil {
		return nil, err
	}

	return &user, nil
}
