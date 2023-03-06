package main

import "net/http"

type Server interface {
	Route(pattern string, handleFunc http.HandlerFunc)
	Start(address string) error
}

// sdkHttpServer 基于http库实现
type sdkHttpServer struct {
	Name string
}

// Route 注册路由
func (s *sdkHttpServer) Route(pattern string, handleFunc http.HandlerFunc) {
	//TODO implement me
	panic("implement me")
}

func (s *sdkHttpServer) Start(address string) error {
	//TODO implement me
	panic("implement me")
}
