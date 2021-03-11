package main

import "net/http"

//Route struct to define route def
type Route struct {
	Name string
	Method string
	Pattern string
	HandlerFunc http.HandlerFunc
}

//Routes slice/array
type Routes []Route


var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
	Route{
		"ProductIndex",
		"GET",
		"/products",
		ProductIndex,
	},
	Route{
		"ProductShow",
		"GET",
		"/products/{productId}",
		ProductShow,
	},
	Route{
		"ProductCreate",
		"POST",
		"/products",
		ProductCreate,
	},
	
}
