package repository

const (
	getUserCart = `
		SELECT
			carts.user_id,
			carts.created_at,
			carts.updated_at
		FROM carts
		WHERE carts.user_id = $1;
	`

	getFullUserCart = `
		SELECT
			carts.user_id,
			carts.created_at,
			carts.updated_at,
			cartlines.user_id,
			cartlines.product_id,
			cartlines.name,
			cartlines.quantity,
			cartlines.created_at,
			cartlines.updated_at
		FROM carts
		JOIN cartlines ON $1 = cartlines.user_id;
	`

	createCart = `
		INSERT INTO carts (
			user_id,
			created_at,
			updated_at
		) VALUES (
			$1,
			$2,
			$3
		);
	`

	updateCart = `
		UPDATE carts
		SET updated_at = $1
		WHERE user_id = $2;
	`

	createCartline = `
		INSERT INTO cartlines (
			user_id,
			product_id,
			name,
			quantity,
			created_at,
			updated_at
		) VALUES (
			$1,
			$2,
			$3,
			$4,
			$5,
			$6
		);
	`

	updateCartline = `
		UPDATE cartlines
		SET
			name = $1,
			quantity = $2,
			updated_at = $3
		WHERE user_id = $4 AND product_id = $5;
	`

	deleteCart = `
		DELETE FROM carts
		WHERE user_id = $1;
	`

	deleteCartline = `
		DELETE FROM cartlines
		WHERE user_id = $1 AND product_id = $2;
	`

	deleteCartCartlines = `
		DELETE FROM cartlines
		WHERE user_id = $1;
	`
)
