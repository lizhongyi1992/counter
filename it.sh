#!/bin/bash

set -e

MYSQL="mysql -uroot -proot -e "

$MYSQL"drop database if exists test;"
$MYSQL"create database if not exists test default charset utf8;"
$MYSQL"create table if not exists test.tshare(no varchar(64) not null primary key, playtimes int) default charset=utf8;"

$MYSQL"insert into test.tshare (no,playtimes) values('testid1',0);"
$MYSQL"insert into test.tshare (no,playtimes) values('testid2',0);"
$MYSQL"insert into test.tshare (no,playtimes) values('testid3',0);"

for i in {1..10};do
    curl -XPOST 'localhost:8888/incr_video_views?video_id=testid1'
done

for i in {1..20};do
    curl -XPOST 'localhost:8888/incr_video_views?video_id=testid2'
done

for i in {1..30};do
    curl -i -XPOST 'localhost:8888/incr_video_views?video_id=testid3'
done

sleep 2
$MYSQL"select * from test.tshare;"

for i in {1..10};do
    curl -XPOST 'localhost:8888/incr_video_views?video_id=testid1'
done

for i in {1..20};do
    curl -XPOST 'localhost:8888/incr_video_views?video_id=testid2'
done

for i in {1..30};do
    curl -i -XPOST 'localhost:8888/incr_video_views?video_id=testid3'
done
