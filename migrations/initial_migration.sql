USE dokichallenge

CREATE TABLE `user`(
	`user_id` INT NOT NULL AUTO_INCREMENT,
	`user_name` VARCHAR(30) NOT NULL,
	PRIMARY KEY (`user_id`),
);

CREATE TABLE `transactions`(
	`transaction_id` BIGINT NOT NULL AUTO_INCREMENT,
	`user_id` INT NOT NULL,
	`amount` DECIMAl(18, 2) NOT NULL,
    `transaction_date` DATE NOT NULL, 
	PRIMARY KEY(`transaction_id`),
	FOREIGN KEY(`user_id`) REFERENCES `users`(`user_id`)
);
