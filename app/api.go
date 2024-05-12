package app

import (
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/xnpltn/azshop/database"
	"github.com/xnpltn/azshop/utls"
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
	price := c.FormValue("price")
	tempfile, err := os.CreateTemp("assets/products/", "image-"+c.FormValue("name")+"-*.jpg")
	if err != nil{
		fmt.Println("error occured", err)
	}

	pp, err := strconv.ParseFloat(price, 64)
	if err!= nil{
		fmt.Println(err)
	}

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
		Name: c.FormValue("name"),
		Price: fmt.Sprintf("%f", pp),
		ImageUrl: fmt.Sprintf("/static%s", tempfile.Name()[6:]),
		Description: c.FormValue("description"),
	}

	if 	err := a.db.CreateProduct(c.Request().Context(), product) ;err!= nil{
		fmt.Println(err)

		return c.HTML(200, "<p>Something went wrong</p>")
	}
	return c.HTML(200, "Product Added succesfully")
}



func (a *Application)AddToCarT() echo.HandlerFunc{
	return func(c echo.Context) error {
		id, err := uuid.Parse(c.FormValue("product_id"))
		if err!= nil{
			return c.HTML(200, "<script> alert('Something went wrong')</script>")
		}
		user, _ := utls.CheckCredentials(c, a.db)
		cart := database.AddToCartParams{
			UserID: user.ID,
			ProductID: id,
			Quantity: 1,
		}
		err = a.db.AddToCart(c.Request().Context(), cart)
		if err!= nil{
			fmt.Println(err)
			return c.HTML(200, "<script> alert('Something went wrong')</script>")
		}
		return c.HTML(200, "<script> alert('Added successfully')</script>")
	}
}


func (a *Application) DeleteProduct(c echo.Context) error{
	c.FormParams()
	id, err := uuid.Parse(c.FormValue("product_id")) 
	if err!= nil{
		fmt.Println(err)
		return c.HTML(200, "something went wrong")
	}
	err = a.db.DeleteProductByID(c.Request().Context(), id)
	if err!= nil{
		return c.HTML(200, "something went wrong")
	}
	return c.HTML(202, "deleted")
}

