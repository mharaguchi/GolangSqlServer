CREATE DATABASE GolangSqlDB;
GO

USE GolangSqlDB;
GO

CREATE TABLE Movies ( 
    ID int IDENTITY(1, 1), 
    Title varchar(80), 
    Description varchar(175), 
    Year int)
GO
