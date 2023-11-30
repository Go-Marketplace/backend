package repository

const (
	getAllProducts = `
		SELECT
			products.product_id,
			products.user_id,
			products.category_id,
			products.name,
			products.description,
			products.price,
			products.quantity,
			products.moderated,
			products.created_at,
			products.updated_at
		FROM products
		WHERE products.moderated = false;
	`

	getProductByID = `
		SELECT
			products.product_id,
			products.user_id,
			products.category_id,
			products.name,
			products.description,
			products.price,
			products.quantity,
			products.moderated,
			products.created_at,
			products.updated_at
		FROM products
		WHERE product_id = $1;
	`

	getAllCategoryProducts = `
		SELECT
			products.product_id,
			products.user_id,
			products.category_id,
			products.name,
			products.description,
			products.price,
			products.quantity,
			products.moderated,
			products.created_at,
			products.updated_at
		FROM products
		WHERE category_id = $1 AND products.moderated = false;
	`

	getAllUserProducts = `
		SELECT
			products.product_id,
			products.user_id,
			products.category_id,
			products.name,
			products.description,
			products.price,
			products.quantity,
			products.moderated,
			products.created_at,
			products.updated_at
		FROM products
		WHERE user_id = $1;
	`

	createProduct = `
		INSERT INTO products (
			product_id,
			user_id,
			category_id,
			name,
			description,
			price,
			quantity,
			moderated,
			created_at,
			updated_at
		) VALUES (
			$1,
			$2,
			$3,
			$4,
			$5,
			$6,
			$7,
			$8,
			$9,
			$10
		);
	`

	updateProduct = `
		UPDATE products
		SET
			category_id = $1,
			name = $2,
			description = $3,
			price = $4,
			quantity = $5,
			updated_at = $6
		WHERE product_id = $7;
	`

	moderateProduct = `
		UPDATE products
		SET
			moderated = $1,
			updated_at = $2
		WHERE product_id = $3;
	`

	deleteProduct = `
		DELETE FROM products
		WHERE product_id = $1;
	`

	getCategoryByID = `
		SELECT
			categories.category_id,
			categories.name,
			categories.description
		FROM categories
		WHERE categories.category_id = $1;
	`

	getAllCategories = `
		SELECT
			categories.category_id,
			categories.name,
			categories.description
		FROM categories;
	`
)
