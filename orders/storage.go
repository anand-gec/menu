package main

import "context"

type store struct {
	//add hare our mongoDB
}

func NewStore() *store {
	return &store{}
}

func (s *store) Create(context.Context)error{
	return nil
}