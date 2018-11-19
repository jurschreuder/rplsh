// Random Projection Local Sensitive Hashing
// By Jurriaan Schreuder

package rplsh

import (
	"testing"
)

func TestRplsh(t *testing.T) {
	ht := NewHashTable(5, 32)
	t.Logf("%08b\n", ht.Hash([]float64{0, 1, 0, 1, 0}))
	t.Logf("%08b\n", ht.Hash([]float64{1, 1, 0, 1, 0}))
	t.Logf("%08b\n", ht.Hash([]float64{1, 0, 0, 1, 0}))
	t.Logf("%08b\n", ht.Hash([]float64{1, 0, 1, 1, 0}))
	t.Logf("%08b\n", ht.Hash([]float64{1, 0, 1, 0, 0}))
	t.Logf("%08b\n", ht.Hash([]float64{1, 0, 1, 0, 1}))

	t.Log(ht.HashUint8([]float64{1, 0, 1, 0, 1}))
	t.Log(ht.HashUint16([]float64{1, 0, 1, 0, 1}))
	t.Log(ht.HashUint32([]float64{1, 0, 1, 0, 1}))
	t.Log(ht.HashUint64([]float64{1, 0, 1, 0, 1}))
}

func TestSave(t *testing.T) {
	ht := NewHashTable(5, 64)
	t.Log("before save", ht.HashUint64([]float64{1, 0, 1, 0, 1}))
	err := ht.Save("test_save.js")
	if err != nil {
		t.Fatal(err)
	}
	ht2, err := Load("test_save.js")
	if err != nil {
		t.Fatal(err)
	}
	t.Log("after save", ht2.HashUint64([]float64{1, 0, 1, 0, 1}))
}
