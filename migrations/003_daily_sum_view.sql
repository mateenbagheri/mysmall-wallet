CREATE VIEW daily_transactions_sum_view AS
SELECT sum(T.amount) AS amount
FROM transactions AS T
WHERE T.transaction_date = CURDATE()
