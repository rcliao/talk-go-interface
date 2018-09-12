package main

import "testing"

func TestGet(t *testing.T) {
	dao := dao.NewMemoryImpl()
	if dao.Get(1337).ID != 0 {
		t.Fatal("Should return empty DTO when getting id that does not exist")
	}
}
