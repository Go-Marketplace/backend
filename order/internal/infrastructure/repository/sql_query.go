package repository

const (
	getFullOrderByID = `
		SELECT
			orders.order_id,
			orders.user_id,
			orders.total_price,
			orders.created_at,
			orders.updated_at,
			orderlines.orderline_id,
			orderlines.order_id,
			orderlines.product_id,
			orderlines.name,
			orderlines.price,
			orderlines.quantity,
			orderlines.status,
			orderlines.created_at,
			orderlines.updated_at
		FROM
			orders
		JOIN
			orderlines ON $1 = orderlines.order_id;
	`

	getAllUserOrders = `
		SELECT
			orders.order_id,
			orders.user_id,
			orders.total_price,
			orders.created_at,
			orders.updated_at,
			orderlines.orderline_id,
			orderlines.order_id,
			orderlines.product_id,
			orderlines.name,
			orderlines.price,
			orderlines.quantity,
			orderlines.status,
			orderlines.created_at,
			orderlines.updated_at
		FROM
			orders
		JOIN
			orderlines ON orders.order_id = orderlines.order_id
		WHERE orders.user_id = $1;
	`

	createOrder = `
		INSERT INTO orders (
			order_id,
			user_id,
			total_price,
			created_at,
			updated_at
		) VALUES (
			$1,
			$2,
			$3,
			$4,
			$5
		);
	`

	createOrderline = `
		INSERT INTO orderlines (
			orderline_id,
			order_id,
			product_id,
			name,
			price,
			quantity,
			status,
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

	deleteOrder = `
		DELETE FROM orders
		WHERE order_id = $1;
	`
)
