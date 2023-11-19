package repository

const (
	getCartByID = `
		SELECT
			carts.cart_id,
			carts.user_id,
			carts.created_at,
			carts.updated_at
		FROM carts
		WHERE carts.cart_id = $1;
	`

	getUserCart = `
		SELECT
			carts.cart_id,
			carts.user_id,
			carts.created_at,
			carts.updated_at
		FROM carts
		WHERE carts.user_id = $1;
	`

	getFullCartByID = `
		SELECT
			carts.cart_id,
			carts.user_id,
			carts.created_at,
			carts.updated_at,
			cartlines.cartline_id,
			cartlines.cart_id,
			cartlines.product_id,
			cartlines.name,
			cartlines.quantity,
			cartlines.created_at,
			cartlines.updated_at
		FROM carts
		JOIN cartlines ON $1 = cartlines.cart_id;
	`

	getFullUserCart = `
		SELECT
			carts.cart_id,
			carts.user_id,
			carts.created_at,
			carts.updated_at,
			cartlines.cartline_id,
			cartlines.cart_id,
			cartlines.product_id,
			cartlines.name,
			cartlines.quantity,
			cartlines.created_at,
			cartlines.updated_at
		FROM carts
		JOIN cartlines ON carts.cart_id = cartlines.cart_id
		WHERE carts.user_id = $1;
	`

	createCart = `
		INSERT INTO carts (
			cart_id,
			user_id,
			created_at,
			updated_at
		) VALUES (
			$1,
			$2,
			$3,
			$4
		);
	`

	updateCart = `
		UPDATE carts
		SET updated_at = $1
		WHERE cart_id = $2;
	`

	createCartline = `
		INSERT INTO cartlines (
			cartline_id,
			cart_id,
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
			$6,
			$7
		);
	`

	updateCartline = `
		UPDATE cartlines
		SET
			name = $1,
			quantity = $2,
			updated_at = $3
		WHERE cartline_id = $4;
	`

	deleteCart = `
		DELETE FROM carts
		WHERE cart_id = $1;
	`

	deleteCartline = `
		DELETE FROM cartlines
		WHERE cartline_id = $1;
	`

	deleteCartCartlines = `
		DELETE FROM cartlines
		WHERE cart_id = $1;
	`
)
