package controller

import (
	"encoding/json"
	"fmt"
	"go-api-meli/database"
	"go-api-meli/model"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	request, _ := ioutil.ReadAll(r.Body)

	var product model.Product

	json.Unmarshal(request, &product)
	db, _ := database.Connection()

	statement, err := db.Prepare("insert into tb_product (title, price, quantity) values (?,?,?)")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	defer statement.Close()

	statement.Exec(product.Title, product.Price, product.Quantity)

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("Product successfully registered.")))

}

func GetProducts(w http.ResponseWriter, r *http.Request) {
	db, _ := database.Connection()
	defer db.Close()

	rows, err := db.Query("select * from tb_product")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	defer rows.Close()

	var products []model.Product
	for rows.Next() {
		var product model.Product
		rows.Scan(&product.ID, &product.Title, &product.Price, &product.Quantity)
		products = append(products, product)
	}

	json.NewEncoder(w).Encode(products)
	w.WriteHeader(http.StatusOK)

}

func GetProductById(w http.ResponseWriter, r *http.Request) {
	paramters := mux.Vars(r)
	ID, _ := strconv.ParseUint(paramters["productID"], 10, 32)

	db, _ := database.Connection()
	row, err := db.Query("select idtb_product, title, price, quantity from tb_product where idtb_product = ?", ID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var product model.Product
	if row.Next() {
		row.Scan(&product.ID, &product.Title, &product.Price, &product.Quantity)
	}

	json.NewEncoder(w).Encode(product)
	w.WriteHeader(http.StatusOK)
	return
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	paramters := mux.Vars(r)
	ID, _ := strconv.ParseInt(paramters["productID"], 10, 32)
	request, _ := ioutil.ReadAll(r.Body)

	var product model.Product

	json.Unmarshal(request, &product)

	db, _ := database.Connection()

	defer db.Close()

	statement, err := db.Prepare("update tb_product set title = ?, price = ?, quantity = ? where idtb_product = ? ")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	defer statement.Close()
	statement.Exec(&product.Title, &product.Price, &product.Quantity, ID)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("Product updated successfully.")))
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	paramters := mux.Vars(r)
	ID, _ := strconv.ParseInt(paramters["productID"], 10, 32)

	db, _ := database.Connection()

	defer db.Close()
	statement, err := db.Prepare("delete from tb_product where idtb_product = ? ")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	defer statement.Close()
	statement.Exec(ID)

	w.WriteHeader(http.StatusNoContent)
	w.Write([]byte(fmt.Sprintf("Product deleted successfully.")))

}
