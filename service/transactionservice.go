package service

import (
	"database/sql"

	"github.com/mateenbagheri/mysmall-wallet/models"
)

func InsertTransaction(t models.Transaction) (models.Reference, error) {
	var result models.Reference

	stmt, err := Mysql.Prepare(
		`
		INSERT INTO transactions
		SET
			transaction_id=?,
			user_id=?,
			amount=?,
			transaction_date=?;
		`,
	)

	if err != nil {
		return result, err
	}

	res, err := stmt.Exec(
		t.TransactionID,
		t.UserID,
		t.Amount,
		t.TransactionDate,
	)

	if err != nil {
		return result, err
	}

	// retrieving the Last Inserted ID
	lid, err := res.LastInsertId()

	if err != nil {
		return result, err
	}

	result.TransactionID = lid

	return result, nil
}

func SelectDailyAmount(now string) (sql.NullFloat64, error) {
	var dailyAmount sql.NullFloat64
	result, err := Mysql.Query(
		`
		SELECT sum(T.amount) AS amount
		FROM transactions AS T
		WHERE T.transaction_date = ?
		`, now,
	)

	if err != nil {
		return dailyAmount, err
	}

	if result.Next() {
		err := result.Scan(&dailyAmount)
		if err != nil {
			return dailyAmount, err
		}
	}

	return dailyAmount, nil
}
