package repository

const (
	getUserByID = `
		SELECT *
		FROM users
		WHERE user_id = $1;
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

	deleteUser = `
		DELETE FROM users
		WHERE user_id = $1;
	`
)
