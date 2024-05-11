-- name: CreateProduct :exec


INSERT INTO product (name, price, image_url, description)
VALUES ($1, $2, $3, $4);

-- name: GetProductByID :one

SELECT * FROM product WHERE id = $1;


-- name: GetAllProducts :many

SELECT * FROM product;
