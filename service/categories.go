package service

import (
	"simple/user/auth/database"
	"simple/user/auth/model"
	"strconv"
	// "simple/user/auth/model"
)

type CategoryDTO struct {
	Name string `json:"name" binding:"required"`
}

func (category *CategoryDTO) InsertCategory() error {
	query := `INSERT INTO category (name) VALUES ($1)`
	_, err := database.DB.Exec(query, &category.Name)
	if err != nil {
		return err
	}

	return nil
}

func FindAll() ([]model.Category, error) {
	query := `SELECT id, name, created_at FROM category`
	rows, err := database.DB.Query(query)
	if err != nil {
		return []model.Category{}, err
	}

	var categories []model.Category

	for rows.Next() {
		var category model.Category
		err := rows.Scan(&category.ID, &category.Name, &category.Created_at)
		if err != nil {
			return []model.Category{}, err
		}
		categories = append(categories, category)
	}

	return categories, nil
}

func (category *CategoryDTO) UpdateOne(id string) (int64, error) {
	intID, _ := strconv.ParseInt(id, 10, 64)

	query := `UPDATE category 
	         SET name =$1, updated_at = CURRENT_TIMESTAMP WHERE id = $2
	`
	result, err := database.DB.Exec(query, &category.Name, intID)
	if err != nil {
		return 0, err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}

	return rows, nil
}
