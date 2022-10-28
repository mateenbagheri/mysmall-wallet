package service

import "github.com/mateenbagheri/mysmall-wallet/models"

func SelectUserBalanceByID(user_id int) (models.UserBalance, error) {
	var balance models.UserBalance

	result, err := Mysql.Query(
		`
		SELECT sum(T.amount) AS balance
		FROM users AS U
			INNER JOIN transactions AS T
		    ON T.user_id = U.user_id
		WHERE U.user_id = ?
		`, user_id,
	)

	if err != nil {
		return balance, err
	}

	if result.Next() {
		err := result.Scan(&balance.Balance)
		if err != nil {
			return balance, err
		}
	}

	return balance, nil
}
