package service

import (
	"encoding/json"
	"fmt"
	"log"
	"simple/user/auth/database"
	"simple/user/auth/model"
	"strconv"
)

type ProductDTO struct {
	Name        string  `json:"name" validate:"required"`
	Description string  `json:"description"`
	Price       float64 `json:"price" validate:"required"`
	Currency    string  `json:"currency"`
	Quantity    int64   `json:"quantity"`
	Active      bool    `json:"active"`
	User_id     int64   `json:"user_id"`
	Category_id int64   `json:"category_id" validate:"required"`
	Updated_at  string  `json:"Updated_at"`
}

func (p *ProductDTO) SetDefaults() {
	if p.Currency == "" {
		p.Currency = "USD"
	}
	p.Active = true
}

func (product *ProductDTO) InsertProduct() error {

	product.SetDefaults()

	query := `INSERT INTO products 
	(name, description, price, currency, quantity, active, user_id, category_id)
	 VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`

	stmt, err := database.DB.Prepare(query)
	if err != nil {
		log.Printf("Error preparing: %v", err)
		return err
	}

	_, err = stmt.Exec(product.Name, product.Description, product.Price, product.Currency, product.Quantity, product.Active, product.User_id, product.Category_id)
	if err != nil {
		log.Printf("Error inserting product: %v", err)
		return err
	}

	return nil
}

func FindAllProducts() ([]model.Product, error) {
	query := `SELECT p.id, p.name, p.description, p.price, p.currency,
                p.quantity, p.active, p.user_id, p.created_at,

                (SELECT ROW_TO_JSON(category_obj) FROM(
                SELECT id, name, created_at FROM category WHERE id = p.category_id	   
                ) category_obj) AS category

                FROM products p;`

	rows, err := database.DB.Query(query)
	if err != nil {
		log.Printf("Error inserting product: %v", err)
		return []model.Product{}, err
	}

	defer rows.Close()

	var products []model.Product

	for rows.Next() {
		var product model.Product
		var categoryJSON []byte

		err := rows.Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.Currency, &product.Quantity, &product.Active, &product.User_id, &product.Created_at, &categoryJSON)
		if err != nil {
			log.Printf("Error selecting all product: %v", err)
			return []model.Product{}, err
		}

		err = json.Unmarshal(categoryJSON, &product.Category)
		if err != nil {
			log.Fatal(err)
		}

		products = append(products, product)
	}

	return products, nil
}

func FindOne(id string) (model.Product, error) {
	intID, _ := strconv.ParseInt(id, 10, 64)
	query := `SELECT p.id, p.name, p.description, p.price, p.currency,
	p.quantity, p.active, p.user_id, p.created_at,

	(SELECT ROW_TO_JSON(category_obj) FROM(
	SELECT id, name, created_at FROM category WHERE id = p.category_id	   
	) category_obj) AS category

   FROM products p
   WHERE id = $1;`

	var product model.Product

	//Declaring a byte slice to hold the raw JSON data.
	//Scanning the query results into the appropriate variables.
	//Parsing the JSON data into a structured Go type.
	//Handling any errors that occur during these processes.
	var categoryJson []byte
	err := database.DB.QueryRow(query, intID).Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.Currency, &product.Quantity, &product.Active, &product.User_id, &product.Created_at, &categoryJson)
	if err != nil {
		log.Printf("Error finding one product: %v", err)
		return model.Product{}, err
	}
	err = json.Unmarshal(categoryJson, &product.Category)
	if err != nil {
		log.Fatal(err)
	}

	return product, nil
}

// check if product with id exist
func ProductExist(id string) (bool, error) {
	intID, _ := strconv.ParseInt(id, 10, 64)

	var exist bool
	query := `SELECT EXISTS (SELECT FROM products WHERE id =$1)`
	err := database.DB.QueryRow(query, intID).Scan(&exist)
	if err != nil {
		log.Printf("Error product not exist: %v", err)
		return false, err
	}

	return exist, nil
}


func GetProductUserId(id string) (int64, error) {
	intID, _ := strconv.ParseInt(id, 10, 64)

	query := `SELECT user_id FROM products WHERE id = $1`

	var userId int64
	err := database.DB.QueryRow(query, intID).Scan(&userId)
	if err != nil {
		log.Printf("Error: %v", err)
		return 0, err
	}
	fmt.Println(userId)

	return userId, nil

}

func (product *ProductDTO) Update(id string) (int64, error) {
	intID, _ := strconv.ParseInt(id, 10, 64)

	product.SetDefaults()

	query := `UPDATE products 
	            SET name=$1, description=$2, price=$3, currency=$4, quantity=$5, active=$6, updated_at= CURRENT_TIMESTAMP
				WHERE id=$7
	  `
	stmt, err := database.DB.Exec(query, product.Name, product.Description, product.Price, product.Currency, product.Quantity, product.Active, intID)
	if err != nil {
		log.Printf("Error updating products: %v", err)
		return 0, err
	}

	result, err := stmt.RowsAffected()
	if err != nil {
		log.Printf("Error: %v", err)
		return 0, err
	}
	return result, nil
}

func Delete(id string) (int64, error) {
	intID, _ := strconv.ParseInt(id, 10, 64)

	query := `DELETE FROM products WHERE id=$1`
	stmt, err := database.DB.Exec(query, intID)
	if err != nil {
		log.Printf("Error updating products: %v", err)
		return 0, err
	}

	result, err := stmt.RowsAffected()
	if err != nil {
		log.Printf("Error: %v", err)
		return 0, err
	}
	return result, nil
}

func UserProducts(id string) ([]model.Product, error) {
	intID, _ := strconv.ParseInt(id, 10, 64)
	query := `SELECT p.id, p.name, p.description, p.price, p.currency,
	p.quantity, p.active, p.created_at,

	(SELECT ROW_TO_JSON(category_obj) FROM(
	SELECT id, name, created_at FROM category WHERE id = p.category_id	   
	) category_obj) AS category

   FROM products p
   WHERE user_id = $1;`

	row, err := database.DB.Query(query, intID)
	if err != nil {
		return nil, err
	}

	var products []model.Product

	for row.Next() {
		var product model.Product
		var categoryJSON []byte

		err := row.Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.Currency, &product.Quantity, &product.Active, &product.Created_at, &categoryJSON)
		if err != nil {
			return nil, err
		}
		_ = json.Unmarshal(categoryJSON, &product.Category)

		products = append(products, product)
	}

	return products, nil

}
