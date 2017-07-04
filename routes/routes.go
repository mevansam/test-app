package routes

import "github.com/tedsuo/rata"

const (
	Hello   = "HELLO"
	Version = "VERSION"
	Env     = "ENV"
	Exit    = "EXIT"
	Index   = "INDEX"
	Port    = "PORT"
)

var Routes = rata.Routes{
	{Path: "/", Method: "GET", Name: Hello},
	{Path: "/version", Method: "GET", Name: Version},
	{Path: "/env", Method: "GET", Name: Env},
	{Path: "/exit", Method: "GET", Name: Exit},
	{Path: "/index", Method: "GET", Name: Index},
	{Path: "/port", Method: "GET", Name: Port},
}
