--- book

create database bms;
use bms;
create  table book(
id bigint(20) auto_increment Primary key,
title varchar(20) not null,
price double(10,2) not null
) engine=Innodb default charset=utf8mb4;