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

func AddProductToCart(w http.ResponseWriter, r *http.Request) {
	request, _ := ioutil.ReadAll(r.Body)

	var product model.Cart

	json.Unmarshal(request, &product)
	db, _ := database.Connection()
	defer db.Close()

	for _, products := range product.Products {
		statement, err := db.Prepare("insert into tb_cart (idtb_product, quantity) values (?,?)")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		defer statement.Close()
		fmt.Println(products.ID_product)
		fmt.Println(products.Quantity)
		statement.Exec(products.ID_product, products.Quantity)
	}
	//fmt.Println(products.ID_product, product.Quantity)
	//statement.Exec(product.ID, products.Quantity)

	w.WriteHeader(http.StatusCreated)
	return

}
func GetCartById(w http.ResponseWriter, r *http.Request) {
	paramters := mux.Vars(r)
	ID, _ := strconv.ParseUint(paramters["id"], 10, 32)

	db, _ := database.Connection()
	row, err := db.Query("select idtb_cart, quantity from tb_cart where idtb_cart = ?", ID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var product model.Detail
	if row.Next() {
		row.Scan(&product.ID_product, &product.Quantity)
	}

	json.NewEncoder(w).Encode(product)
	w.WriteHeader(http.StatusOK)
	return
}
