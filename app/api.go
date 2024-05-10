package app

import (
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/xnpltn/azshop/database"
)





func(a *Application) CreateProduct(c echo.Context) error{
	f, err := c.FormFile("image")
	if err!= nil{
		fmt.Println(err)
	}
	if err := os.MkdirAll("assets/products", os.FileMode(0777)); err!= nil{
		if !errors.Is(err, os.ErrExist){
			fmt.Println(err)
		}
	}
	name := c.FormValue("name")
	price := c.FormValue("price")
	tempfile, err := os.CreateTemp("assets/products/", "image-"+name+"-*.jpg")
	if err != nil{
		fmt.Println("error occured", err)
	}

	pp, err := strconv.ParseFloat(price, 64)
	if err!= nil{
		fmt.Println(err)
	}

	fmt.Println("Parsed Float: ", pp)

	file, err := f.Open()
	if err!= nil{
		fmt.Println(err)
	}

	b, err := io.ReadAll(file)
	if err!= nil{
		fmt.Println(err)
	}
	
	if _, err = tempfile.Write(b); err!= nil{
		fmt.Println(err)
	}
	product := database.CreateProductParams{
		Name: name,
		Price: price,
		ImageUrl: fmt.Sprintf("/static%s", tempfile.Name()[6:]),
	}
	fmt.Println(product.ImageUrl)

	err = a.db.CreateProduct(c.Request().Context(), product)
	if err!= nil{
		fmt.Println(err)

		return c.HTML(200, "<p>Something went wrong</p>")
	}
	return c.HTML(200, "Product Added succesfully")
}


