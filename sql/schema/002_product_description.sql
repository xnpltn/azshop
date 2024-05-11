-- +goose Up

ALTER TABLE Product
ADD COLUMN description TEXT NOT NULL DEFAULT 'A great Product';


-- +goose Down
ALTER TABLE Product
DROP COLUMN description;
