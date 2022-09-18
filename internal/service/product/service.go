package product

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	dataprod "github.com/HermanSetiawan77777/JakMall/datasource/product"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func getData() map[string]*dataprod.Product {
	file, err := ioutil.ReadFile("../datasource/product/products.json")
	if err != nil {
		fmt.Println("Error when opening file: ", err)
	}
	var payload []*dataprod.Product
	err = json.Unmarshal(file, &payload)
	if err != nil {
		fmt.Println("Error during Unmarshal file: ", err)
	}
	result := make(map[string]*dataprod.Product)
	for _, v := range payload {
		result[v.Name] = v
	}
	return result
}

func GetDetail(name string) *dataprod.Product {
	data := getData()
	payload := data[cases.Title(language.Und).String(name)]
	return payload
}
