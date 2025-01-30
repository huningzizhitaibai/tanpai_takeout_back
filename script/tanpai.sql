create database if not exists `tanpai`;
use `tanpai`;

drop table if exists `user_basic`;
create table `user_basic` (
    `id` bigint not null auto_increment comment '主键',
    `username` varchar(30) not null comment '用户名',
    `password` varchar(100) not null comment '密码',
    `type` int not null comment '用户类型，0-管理员，1-用户，2-商户，3-骑手',
    `real_name` varchar(50) comment '用户真实姓名，非必须',
    primary key (`id`)
);