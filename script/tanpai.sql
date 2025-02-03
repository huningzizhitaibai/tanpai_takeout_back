create database if not exists `tanpai`;
use `tanpai`;

drop table if exists `user_basic`;
create table `user_basic` (
    `id` bigint not null auto_increment comment '主键',
    `username` varchar(30) not null comment '用户名',
    `password` varchar(100) not null comment '密码',
    `type` int not null comment '用户类型，0-管理员，1-用户，2-商户，3-骑手',
    primary key (`id`)
);

drop table if exists `user`;
create table `user`(
    `id` bigint not null auto_increment,
    `username` varchar(30) not null ,
    `password` varchar(100) not null ,
    primary key (`id`)
);

drop table if exists `shop`;
create table `shop`(
    `id` bigint not null auto_increment,
    `username` varchar(30) not null ,
    `password` varchar(100) not null ,
    `real_name` varchar(20) not null comment '法人真实姓名',
    `id_number` varchar(20) not null comment '身份证号',
    `certificate_for_food` varchar(100) not null comment '食品安全证的存储路径',
    `id_card1` varchar(100) not null comment '身份证正面存储路径',
    `id_card2` varchar(100) not null comment '身份证反面存储路径',
    `certificate_for_shop` varchar(100) not null comment '营业执照',
    `is_verify` boolean default false comment '是否完成认证',
    primary key (`id`)
);

drop table if exists `deliver`;
create table `deliver` (
    `id` bigint not null auto_increment,
    `username` varchar(30) not null ,
    `password` varchar(100) not null ,
    `real_name` varchar(20) not null ,
    `studer_card` varchar(100),
    `is_student` boolean default false,
    `is_verify` boolean default  false,
    primary key (`id`)
);

drop table if exists `controller`;
create table  `controller` (
    `id` bigint not null auto_increment,
    `username` varchar(30) not null ,
    `password` varchar(100) not null ,
    `real_name` varchar(20) not null ,
    `id_card1` varchar(100) not null ,
    `id_card2` varchar(100) not null ,
    `invite_code` varchar(100) not null ,
    primary key (`id`)
);
