package repository

import "simple-CRUD/pkg/entity"

func (r *Repository) CreateUser(name, email string, age int) error {
	_, err := r.db.Exec("INSERT INTO users (name, email, age) VALUES ($1, $2, $3)", name, email, age)
	if err != nil {
		return err
	}

	return nil
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
