-- name: AddToCart :exec

INSERT INTO user_cart (user_id, product_id, quantity)
VALUES ($1, $2, $3);

-- name: GetUserCart :many
SELECT
    uc.user_id,
    uc.product_id,
    uc.quantity,
    p.name AS product_name,
    p.description AS product_description,
    p.price AS product_price,
    p.image_url as product_image
FROM
    user_cart uc
JOIN
    product p ON uc.product_id = p.id
WHERE
    uc.user_id = $1;

