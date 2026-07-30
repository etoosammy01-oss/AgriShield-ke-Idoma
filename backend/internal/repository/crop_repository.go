package repository

import (
	"backend/internal/models"
	"database/sql"
)

type CropRepository struct {
	db *sql.DB
}

func NewCropRepository(db *sql.DB) *CropRepository {
	return &CropRepository{db: db}
}

func (r *CropRepository) Create(crop *models.Crop) error {
	query := `
	INSERT INTO crops (farmer_id, name, quantity, unit, location, price_per_unit, listed_for_sale, image_url)
	VALUES (?, ?, ?, ?, ?, ?, ?, ?)
	`
	listed := 0
	if crop.ListedForSale {
		listed = 1
	}
	res, err := r.db.Exec(query, crop.FarmerID, crop.Name, crop.Quantity, crop.Unit, crop.Location, crop.PricePerUnit, listed, crop.ImageURL)
	if err != nil {
		return err
	}
	if id, err := res.LastInsertId(); err == nil {
		crop.ID = int(id)
	}
	return nil
}

// ListByFarmer returns everything a farmer has in storage (listed or not).
func (r *CropRepository) ListByFarmer(farmerID int) ([]models.Crop, error) {
	query := `
	SELECT id, farmer_id, name, quantity, unit, location, price_per_unit, listed_for_sale, image_url, created_at, updated_at
	FROM crops WHERE farmer_id = ? ORDER BY created_at DESC
	`
	rows, err := r.db.Query(query, farmerID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var crops []models.Crop
	for rows.Next() {
		var c models.Crop
		var listed int
		if err := rows.Scan(&c.ID, &c.FarmerID, &c.Name, &c.Quantity, &c.Unit, &c.Location, &c.PricePerUnit, &listed, &c.ImageURL, &c.CreatedAt, &c.UpdatedAt); err != nil {
			return nil, err
		}
		c.ListedForSale = listed == 1
		crops = append(crops, c)
	}
	return crops, rows.Err()
}

// ListAvailable returns everything currently listed for sale, across all farmers,
// with quantity remaining — this is what buyers see on the Marketplace page.
func (r *CropRepository) ListAvailable() ([]models.Crop, error) {
	query := `
	SELECT crops.id, crops.farmer_id, crops.name, crops.quantity, crops.unit, crops.location,
	       crops.price_per_unit, crops.listed_for_sale, crops.image_url, crops.created_at, crops.updated_at,
	       farmers.full_name
	FROM crops
	JOIN farmers ON farmers.id = crops.farmer_id
	WHERE crops.listed_for_sale = 1 AND crops.quantity > 0
	ORDER BY crops.created_at DESC
	`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var crops []models.Crop
	for rows.Next() {
		var c models.Crop
		var listed int
		if err := rows.Scan(&c.ID, &c.FarmerID, &c.Name, &c.Quantity, &c.Unit, &c.Location, &c.PricePerUnit, &listed, &c.ImageURL, &c.CreatedAt, &c.UpdatedAt, &c.SellerName); err != nil {
			return nil, err
		}
		c.ListedForSale = listed == 1
		crops = append(crops, c)
	}
	return crops, rows.Err()
}

func (r *CropRepository) GetByID(id int) (*models.Crop, error) {
	query := `
	SELECT id, farmer_id, name, quantity, unit, location, price_per_unit, listed_for_sale, image_url, created_at, updated_at
	FROM crops WHERE id = ?
	`
	var c models.Crop
	var listed int
	err := r.db.QueryRow(query, id).Scan(&c.ID, &c.FarmerID, &c.Name, &c.Quantity, &c.Unit, &c.Location, &c.PricePerUnit, &listed, &c.ImageURL, &c.CreatedAt, &c.UpdatedAt)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	c.ListedForSale = listed == 1
	return &c, nil
}

func (r *CropRepository) ReduceQuantity(cropID int, amount float64) error {
	_, err := r.db.Exec(`UPDATE crops SET quantity = quantity - ?, updated_at = CURRENT_TIMESTAMP WHERE id = ?`, amount, cropID)
	return err
}
