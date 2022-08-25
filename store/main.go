package main

import (
	"encoding/json"
	"net/http"
	"store/inventory"
	"store/store"
)

func main() {
	i, err := inventory.NewFileInventory("data/inventory.txt")
	if err != nil {
		panic(err)
	}
	defer i.Close()
	
	s := store.New(i)
	http.HandleFunc("/",greet)
	http.HandleFunc("/product",func(w http.ResponseWriter, r *http.Request) {
		str:= r.URL.Query().Get("name")
		value,b:=s.FindProduct(str)
		mmm:=map[string]interface{}{
			"product":value,
			"exists":b,
		}
		json.NewEncoder(w).Encode(mmm)
	})

	http.ListenAndServe(":8080",nil)
}
