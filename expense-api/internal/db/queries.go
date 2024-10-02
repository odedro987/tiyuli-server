package db

const NewExpenseQuery = `
	INSERT INTO expenses 
		(user_id, name, note, types, payment_date, currency_code, amount) 
	VALUES (?, ?, ?, ?, ?, ?, ?);
`
