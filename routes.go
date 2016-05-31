package main

import "net/http"

type Routes []Route
type Route struct {
	HandlerFunc http.HandlerFunc
	Method      string
	Name        string
	Pattern     string
}

var routes = Routes{
	Route{
		Name:        "Index",
		Method:      "GET",
		Pattern:     "/",
		HandlerFunc: GetIndex},
	Route{
		Name:        "Lists",
		Method:      "GET",
		Pattern:     "/lists",
		HandlerFunc: GetLists},
	Route{
		Name:        "List",
		Method:      "GET",
		Pattern:     "/lists/{ListID}",
		HandlerFunc: GetList},
	Route{
		Name:        "ListItems",
		Method:      "GET",
		Pattern:     "/lists/{ListID}/items",
		HandlerFunc: GetListItems},
	Route{
		Name:        "ListItem",
		Method:      "GET",
		Pattern:     "/lists/{ListID}/items/{ItemID}",
		HandlerFunc: GetListItem}}
