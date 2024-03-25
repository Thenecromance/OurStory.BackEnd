package Wish

import "time"

type Data struct {
	Id     int    `db:"id"`
	detail string `db:"detail"`
}

type Info struct {
	Id    int       `db:"id"`
	Title string    `db:"title"`
	Date  time.Time `db:"time"`
}
