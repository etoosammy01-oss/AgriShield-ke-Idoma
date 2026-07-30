package repository

import (
	"backend/internal/models"
	"database/sql"
)

type FarmerRepository struct {
	db *sql.DB
}

func NewFarmerRepository(db *sql.DB) *FarmerRepository {
	return &FarmerRepository{
		db: db,
	}
}

func (r *FarmerRepository) Create(farmer *models.Farmer) error {
	query := `
	INSERT INTO farmers
	(full_name, phone, password_hash, location, role)
	VALUES (?, ?, ?, ?, ?)
	`

	_, err := r.db.Exec(
		query,
		farmer.FullName,
		farmer.Phone,
		farmer.PasswordHash,
		farmer.Location,
		farmer.Role,
	)

	return err
}

func (r *FarmerRepository) GetByPhone(phone string) (*models.Farmer, error) {

	query := `
	SELECT
		id,
		full_name,
		phone,
		password_hash,
		location,
		role,
		created_at,
		updated_at
	FROM farmers
	WHERE phone = ?
	`

	var farmer models.Farmer

	err := r.db.QueryRow(query, phone).Scan(
		&farmer.ID,
		&farmer.FullName,
		&farmer.Phone,
		&farmer.PasswordHash,
		&farmer.Location,
		&farmer.Role,
		&farmer.CreatedAt,
		&farmer.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return &farmer, nil
}

// GetByID fetches a farmer/buyer by their primary key. Used to load the
// logged-in user's details for the dashboard and profile pages.
func (r *FarmerRepository) GetByID(id int) (*models.Farmer, error) {

	query := `
	SELECT
		id,
		full_name,
		phone,
		password_hash,
		location,
		role,
		created_at,
		updated_at
	FROM farmers
	WHERE id = ?
	`

	var farmer models.Farmer

	err := r.db.QueryRow(query, id).Scan(
		&farmer.ID,
		&farmer.FullName,
		&farmer.Phone,
		&farmer.PasswordHash,
		&farmer.Location,
		&farmer.Role,
		&farmer.CreatedAt,
		&farmer.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return &farmer, nil
}
