use bank_db;

CREATE TABLE `accounts` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `cpf` varchar(255) NOT NULL UNIQUE,
  `secret` varchar(255) NOT NULL,
  `balance` float NOT NULL DEFAULT 0,
  `created_at` timestamp NOT NULL DEFAULT now()
);

CREATE TABLE `transfers` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `from_account_id` int,
  `to_account_id` int,
  `amount` float NOT NULL,
  `created_at` timestamp DEFAULT now()
);

ALTER TABLE `transfers` ADD FOREIGN KEY (`from_account_id`) REFERENCES `accounts` (`id`);

ALTER TABLE `transfers` ADD FOREIGN KEY (`to_account_id`) REFERENCES `accounts` (`id`);

CREATE INDEX `accounts_index_0` ON `accounts` (`cpf`);

CREATE INDEX `transfers_index_1` ON `transfers` (`from_account_id`);

CREATE INDEX `transfers_index_2` ON `transfers` (`to_account_id`);

CREATE INDEX `transfers_index_3` ON `transfers` (`from_account_id`, `to_account_id`);

insert into accounts (name,cpf,secret,balance) values ('joao','10573104921','$2a$10$5VevS1xpxCyqtNVMRL5Q2uGYvCJrOuLskA68SFb9B1cP/1.5m3Atq',1000.0);
insert into accounts (name,cpf,secret,balance) values ('paulo','09347584924','$2a$10$5VevS1xpxCyqtNVMRL5Q2uGYvCJrOuLskA68SFb9B1cP/1.5m3Atq',1000.0);


