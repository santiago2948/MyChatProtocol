package client

import ("net")

type Client struct {
	Conexion net.Conn
	Nickname string
}