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
	(full_name, phone, password_hash, location)
	VALUES (?, ?, ?, ?)
	`

	_, err := r.db.Exec(
		query,
		farmer.FullName,
		farmer.Phone,
		farmer.PasswordHash,
		farmer.Location,
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

//func (r *FarmerRepository) GetByID(id int) (*models.Farmer, error)
