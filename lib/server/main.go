package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

// Product представляет товар
type Product struct {
	ID          int     `json:"id"`
	ImageURL    string  `json:"imageUrl"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}

// Пример списка товаров
var products = []Product{
	{ID: 1, ImageURL: "https://basket-15.wbbasket.ru/vol2322/part232297/232297447/images/c246x328/1.webp", Name: "Мяч баскетбольный Molten GF7X 7", Description: "Мяч баскетбольный Molten GF7X 7 размер профессиональный", Price: 1196},
	{ID: 2, ImageURL: "https://basket-04.wbbasket.ru/vol634/part63450/63450038/images/c246x328/2.webp", Name: "Баскетбольный мяч FAKE", Description: "Баскетбольный мяч FAKE для уличного баскетбола размер 7", Price: 925},
	{ID: 3, ImageURL: "https://basket-04.wbbasket.ru/vol496/part49601/49601561/images/c246x328/2.webp", Name: "Баскетбольный мяч", Description: "Баскетбольный мяч", Price: 2152},
	{ID: 4, ImageURL: "https://basket-12.wbbasket.ru/vol1704/part170446/170446253/images/c246x328/1.webp", Name: "Баскетбольный мяч светоотражающий", Description: "Баскетбольный мяч светоотражающий 7 размер для улицы и зала", Price: 1801},
	{ID: 5, ImageURL: "https://basket-13.wbbasket.ru/vol1986/part198660/198660598/images/c246x328/2.webp", Name: "Баскетбольный мяч Challenger размер 7", Description: "Баскетбольный мяч Challenger размер 7", Price: 1399},
	{ID: 6, ImageURL: "https://basket-12.wbbasket.ru/vol1786/part178682/178682377/images/c516x688/1.webp", Name: "Spalding мяч баскетбольный ", Description: "Баскетбольный мяч 7 размер для улицы и зала", Price: 1674},
	{ID: 7, ImageURL: "https://basket-14.wbbasket.ru/vol2154/part215413/215413727/images/c246x328/1.webp", Name: "ECOBALL Replica размер 7", Description: "Баскетбольный мяч профессиональный ECOBALL Replica размер 7", Price: 1999},
	{ID: 8, ImageURL: "https://basket-03.wbbasket.ru/vol414/part41461/41461806/images/c516x688/1.webp", Name: "Баскетбольный мяч SHOT для уличного баскетбола размер 7", Description: "Баскетбольный мяч SHOT для уличного баскетбола размер 7", Price: 842},
	{ID: 9, ImageURL: "https://basket-16.wbbasket.ru/vol2417/part241752/241752426/images/c516x688/2.webp", Name: "KGUMIHO баскетбольный мяч", Description: "Баскетбольный мяч 7 розовый", Price: 1568},
	{ID: 10, ImageURL: "https://basket-12.wbbasket.ru/vol1704/part170444/170444976/images/c246x328/1.webp", Name: "Баскетбольный мяч 7 размер для улицы и зала", Description: "Баскетбольный мяч 7 размер для улицы и зала", Price: 1653},
	{ID: 11, ImageURL: "https://basket-16.wbbasket.ru/vol2469/part246956/246956303/images/c516x688/1.webp", Name: "City-ride", Description: "Мяч баскетбольный, размер 7, диаметр 25см", Price: 775},
}

// Получить все продукты
func getProductsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}

// Создать продукт
func createProductHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var newProduct Product
	err := json.NewDecoder(r.Body).Decode(&newProduct)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	newProduct.ID = len(products) + 1
	products = append(products, newProduct)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newProduct)
}

// Получить продукт по ID
func getProductByIDHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Path[len("/products/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Product ID", http.StatusBadRequest)
		return
	}

	for _, product := range products {
		if product.ID == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(product)
			return
		}
	}

	http.Error(w, "Product not found", http.StatusNotFound)
}

// Обновить продукт
func updateProductHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	idStr := r.URL.Path[len("/products/update/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Product ID", http.StatusBadRequest)
		return
	}

	var updatedProduct Product
	err = json.NewDecoder(r.Body).Decode(&updatedProduct)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	for i, product := range products {
		if product.ID == id {
			products[i] = updatedProduct
			products[i].ID = id
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(products[i])
			return
		}
	}

	http.Error(w, "Product not found", http.StatusNotFound)
}

// Удалить продукт
func deleteProductHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	idStr := r.URL.Path[len("/products/delete/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Product ID", http.StatusBadRequest)
		return
	}

	for i, product := range products {
		if product.ID == id {
			products = append(products[:i], products[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}

	http.Error(w, "Product not found", http.StatusNotFound)
}

func main() {
	http.HandleFunc("/products", getProductsHandler)
	http.HandleFunc("/products/create", createProductHandler)
	http.HandleFunc("/products/", getProductByIDHandler)
	http.HandleFunc("/products/update/", updateProductHandler)
	http.HandleFunc("/products/delete/", deleteProductHandler)

	fmt.Println("Server is running on port 8080!")
	http.ListenAndServe(":8080", nil)
}
