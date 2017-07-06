#!/bin/bash

set -e

MYSQL="mysql -uroot -proot -e "

$MYSQL"drop database if exists test;"
$MYSQL"create database if not exists test default charset utf8;"
$MYSQL"create table if not exists test.tshare(no int not null primary key, playtimes int) default charset=utf8;"

$MYSQL"insert into test.tshare (no,playtimes) values(1,0);"
$MYSQL"insert into test.tshare (no,playtimes) values(2,0);"
$MYSQL"insert into test.tshare (no,playtimes) values(3,0);"

for i in {1..10};do
    curl -XPOST 'localhost:8888/incr_video_views?video_id=1'
done

for i in {1..20};do
    curl -XPOST 'localhost:8888/incr_video_views?video_id=2'
done

for i in {1..30};do
    curl -XPOST 'localhost:8888/incr_video_views?video_id=3'
done

sleep 2
$MYSQL"select * from test.tshare;"

for i in {1..10};do
    curl -XPOST 'localhost:8888/incr_video_views?video_id=1'
done

for i in {1..20};do
    curl -XPOST 'localhost:8888/incr_video_views?video_id=2'
done

for i in {1..30};do
    curl -XPOST 'localhost:8888/incr_video_views?video_id=3'
done

sleep 2
$MYSQL"select * from test.tshare;"
