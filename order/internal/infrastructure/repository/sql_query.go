package repository

const (
	getFullOrderByID = `
		SELECT
			orders.order_id,
			orders.user_id,
			orders.status,
			orders.total_price,
			orders.shipping_cost,
			orders.delivery_address,
			orders.delivery_type,
			orders.created_at,
			orders.updated_at,
			cartlines.cartline_id,
			cartlines.quantity,
			products.product_id,
			products.name,
			products.description,
			products.price
		FROM
			orders
		JOIN
			cartlines ON $1 = cartlines.order_id
		JOIN
			products ON cartlines.cartline_id = products.cartline_id;
	`

	getAllUserOrders = `
		SELECT
			orders.order_id,
			orders.user_id,
			orders.status,
			orders.total_price,
			orders.shipping_cost,
			orders.delivery_address,
			orders.delivery_type,
			orders.created_at,
			orders.updated_at,
			cartlines.cartline_id,
			cartlines.quantity,
			products.product_id,
			products.name,
			products.description,
			products.price
		FROM
			orders
		JOIN
			cartlines ON orders.order_id = cartlines.order_id
		JOIN
			products ON cartlines.cartline_id = products.cartline_id
		WHERE orders.user_id = $1;
	`

	createOrder = `
		INSERT INTO orders (
			order_id,
			user_id,
			status,
			total_price,
			shipping_cost,
			delivery_address,
			delivery_type,
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
			$9
		);
	`

	createCartline = `
		INSERT INTO cartlines (
			cartline_id,
			order_id,
			quantity
		) VALUES (
			$1,
			$2,
			$3
		);
	`

	createProduct = `
		INSERT INTO cartlines (
			product_id,
			cartline_id,
			name,
			description,
			price
		) VALUES (
			$1,
			$2,
			$3,
			$4,
			$5
		);
	`

	cancelOrder = `
		DELETE FROM orders
		WHERE order_id = $1;
	`
)
