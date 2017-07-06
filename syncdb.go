package main

func sync_redis_to_mysql(config accumulator_config, sqlconn *SQLConn, redisconn RedisConn) {
	_dbg("sync_redis_to_mysql")

	oldhset := config.RedisHashSetName
	// TODO:newhset name should append procno+timestamp for concurrency
	newhset := config.RedisHashSetName + config.RedisHashShuffleSuffix

	if !redisconn.Exists(oldhset) {
		_log("old,new hset name:", oldhset, newhset, "-oldhset not exist,continue")
		return
	}
	redisconn.Rename(oldhset, newhset)

	toupdate := map[string]string{}
	hkeys := redisconn.Hkeys(newhset)

	for _, key := range hkeys {
		toupdate[key], _ = redisconn.Hget(newhset, key)
	}

	// n
	// TODO:need more error handle

	// update table set field=field+value where id=key
	table := config.MysqlTable
	field := config.MysqlField
	id := config.MysqlKey
	sqlstr := "update " + table + " set " + field + "=" + field + "+? where " + id + "=?"
	_dbg(sqlstr, sqlconn)

	//TODO: batch n to reduce sql operation timeout
	tx, e := sqlconn.db.Begin()
	if e != nil {
		_err(e)
	}
	for k, v := range toupdate {
		_, e := tx.Exec(sqlstr, v, k)
		if e != nil {
			_err(e)
		}
	}

	e = tx.Commit()
	if e != nil {
		_err(e)
	}

	//last step
	redisconn.Del(newhset)
}
