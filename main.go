package main

import "time"

var deployTime string = time.Now().UTC().Format("2006-01-02 15:04:05") + " UTC"

const version string = "1.0.0"
const dbName string = "Member"
const redisPool int = 0
