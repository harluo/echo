package core

type Router interface {
	Routes(*Server) []Route
}
