-- Active: 1692670144809@@127.0.0.1@3306@golang_restful_api

CREATE DATABASE golang_restful_api ;

USE golang_restful_api;

CREATE TABLE
    category(
        id INT PRIMARY KEY AUTO_INCREMENT,
        name VARCHAR(200) NOT NULL
    ) ENGINE = InnoDB;

DESC category ;

SELECT * FROM category;

CREATE DATABASE golang_restful_api_test ;

USE golang_restful_api_test;

CREATE TABLE
    category(
        id INT PRIMARY KEY AUTO_INCREMENT,
        name VARCHAR(200) NOT NULL
    ) engine = InnoDB;

DESC category;

SELECT * FROM category;