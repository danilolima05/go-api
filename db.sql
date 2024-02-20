CREATE DATABASE catalog;

USE `catalog`;

CREATE TABLE categories (
    id VARCHAR(36), 
    `name` VARCHAR(255), 
    PRIMARY KEY (`id`)
);

INSERT INTO `catalog`.`categories` (`id`,`name`) VALUES ("6b4c28f4-6831-495a-9444-19c93452faa3","Relogios");
INSERT INTO `catalog`.`categories` (`id`,`name`) VALUES ("7c0ca0d4-ff23-4bd7-b131-c563067c4b43","Eletronicos");

CREATE TABLE products (
    id VARCHAR(36), 
    `description` TEXT,
    `name` VARCHAR(255), 
    price DECIMAL(10,2), 
    category_id VARCHAR(36), 
    image_url VARCHAR(255), 
    PRIMARY KEY (`id`),
    FOREIGN KEY `fk_products_category_idx` (`category_id`) REFERENCES categories (`id`)
);

INSERT INTO `catalog`.`products` (`id`,`name`,`description`,`price`,`category_id`,`image_url`)
VALUES ("7f8c9d8e-9f0a-1b2c-3d4e-5f6g7h8i9j0k","Product 1","Description 1", 100, "6b4c28f4-6831-495a-9444-19c93452faa3", "http://google.com/products/1.png"),
("66805003-f9a2-4772-b577-d5ff66207707","Product 2","Description 2", 200, "6b4c28f4-6831-495a-9444-19c93452faa3", "http://google.com/products/2.png");

