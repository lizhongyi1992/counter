#!/bin/bash

set -e

MYSQL="mysql -uroot -proot -e "

$MYSQL"drop database if exists test;"
$MYSQL"create database if not exists test default charset utf8;"
$MYSQL"create table if not exists test.tshare(no varchar(64) not null primary key, playtimes int) default charset=utf8;"

$MYSQL"insert into test.tshare (no,playtimes) values('testid1',0);"
$MYSQL"insert into test.tshare (no,playtimes) values('testid2',0);"
$MYSQL"insert into test.tshare (no,playtimes) values('testid3',0);"

curl -i -XPOST 'localhost:8888/incr_video_views?video_id=testid1'
curl -i -XPOST 'localhost:8888/incr_video_views?video_id=testid1'
curl -i -XPOST 'localhost:8888/incr_video_views?video_id=testid1'
curl -i -XPOST 'localhost:8888/incr_video_views?video_id=testid1'
curl -i -XPOST 'localhost:8888/incr_video_views?video_id=testid1'

$MYSQL"select playtimes from test.tshare where no='testid1';"

