package main

import (
	"testing"

	"github.com/vmihailenco/msgpack/v5"
)

//go:generate msgp
type Person struct {
	Name    string `msg:"name"`
	Address string `msg:"address"`
	Age     int    `msg:"age"`
}

var input = Person{
	Name:    "Harry Potter",
	Address: "4 Privet Drive, Surrey",
	Age:     11,
}

func BenchmarkMsgpack(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bytes, err := msgpack.Marshal(&input)
		checkErr(err)

		var output Person
		err = msgpack.Unmarshal(bytes, &output)
		checkErr(err)
	}
}

func BenchmarkMsgp(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bytes, err := input.MarshalMsg(nil)
		checkErr(err)

		var output Person

		_, err = output.UnmarshalMsg(bytes)
		checkErr(err)
	}
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
