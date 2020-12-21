package test

import (
	"testing"
	"go_study/c36_test/model"
)

func TestStore(t *testing.T){
	monster := model.Monster{"tom", 23, "swimming"}
	err := monster.Store("m1.txt")
	if err != nil{
		t.Fatalf("monster.Store err: %s\n",err)
	}
	t.Logf("monster.Store succeed")
}

func TestReadStore(t *testing.T){
	monster := model.Monster{"tom", 23, "swimming"}
	store, err := monster.ReadStore("m1.txt")
	if err != nil{
		t.Fatalf("monster.ReadStore err: %s\n",err)
	}
	t.Logf("monster.ReadStore succeed\n%+v\n",store)
}

func TestStore2(t *testing.T){
	monster := model.Monster{"tom", 23, "swimming"}
	err := monster.Store2("m2.txt")
	if err != nil{
		t.Fatalf("monster.Store err: %s\n",err)
	}
	t.Logf("monster.Store succeed")
}

func TestReadStore2(t *testing.T){
	monster := model.Monster{"tom", 23, "swimming"}
	store, err := monster.ReadStore2("m2.txt")
	if err != nil{
		t.Fatalf("monster.ReadStore err: %s\n",err)
	}
	t.Logf("monster.ReadStore succeed\n%+v\n",store)
}