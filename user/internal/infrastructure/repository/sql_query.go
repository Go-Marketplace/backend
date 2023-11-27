package repository

const (
	getUserByID = `
		SELECT *
		FROM users
		WHERE user_id = $1;
	`

	getUserByEmail = `
		SELECT *
		FROM users
		WHERE email = $1;
	`

	getAllUsers = `
		SELECT *
		FROM users;
	`

	createUser = `
		INSERT INTO users (
			user_id,
			first_name,
			last_name,
			password,
			email,
			address,
			phone,
			role,
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

	updateUser = `
		UPDATE users
		SET
			first_name = $1,
			last_name = $2,
			address = $3,
			phone = $4,
			updated_at = $5
		WHERE user_id = $6;
	`

	changeUserRole = `
		UPDATE users
		SET
			role = $1
		WHERE user_id = $2;
	`

	deleteUser = `
		DELETE FROM users
		WHERE user_id = $1;
	`
)
