// Random Projection Local Sensitive Hashing
// By Jurriaan Schreuder

package rplsh

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"math"
	"math/rand"
)

type HashTable struct {
	vecsLen     int
	hashLen     int
	projections [][]float64
}

func NewHashTable(vecsLen, hashLen int) *HashTable {
	ht := HashTable{
		vecsLen:     vecsLen,
		hashLen:     hashLen,
		projections: make([][]float64, hashLen),
	}
	for i := range ht.projections {
		ps := make([]float64, vecsLen)
		for j := range ps {
			ps[j] = rand.NormFloat64()
		}
		ht.projections[i] = ps
	}
	return &ht
}

func (ht *HashTable) Save(path string) error {
	js, _ := json.Marshal(&ht.projections)
	return ioutil.WriteFile(path, js, 0644)
}

func Load(path string) (*HashTable, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return &HashTable{}, err
	}
	ps := [][]float64{}
	err = json.Unmarshal(data, &ps)
	if err != nil {
		return &HashTable{}, err
	}
	if len(ps) < 1 {
		return &HashTable{}, errors.New("invalid save file")
	}
	if len(ps[0]) < 1 {
		return &HashTable{}, errors.New("invalid save file")
	}

	hs := &HashTable{
		vecsLen:     len(ps[0]),
		hashLen:     len(ps),
		projections: ps,
	}
	return hs, nil
}

func (ht *HashTable) Hash(vecs []float64) []byte {
	out := make([]byte, int(math.Ceil(float64(ht.hashLen)/8.)))
	bN := 0 // byte no
	iN := 0 // bit no
	l := len(ht.projections)
	for i, pr := range ht.projections {
		if dot(pr, vecs) > 0 {
			out[bN] += uint8(1)
		}
		if i < l-1 {
			out[bN] = out[bN] << 1
		}
		// determine next byte index to write
		iN++
		if iN == 8 {
			bN++
			iN = 0
		}
	}
	return out
}

func (ht *HashTable) HashUint8(vecs []float64) uint8 {
	bs := ht.Hash(vecs)
	return uint8(bs[0])
}

func (ht *HashTable) HashUint16(vecs []float64) uint16 {
	bs := ht.Hash(vecs)
	tt := uint16(bs[0])
	for i := 1; i < len(bs) && i < 2; i++ {
		tt |= uint16(bs[i]) << uint(i*8)
	}
	return tt
}

func (ht *HashTable) HashUint32(vecs []float64) uint32 {
	bs := ht.Hash(vecs)
	tt := uint32(bs[0])
	for i := 1; i < len(bs) && i < 4; i++ {
		tt |= uint32(bs[i]) << uint(i*8)
	}
	return tt
}

func (ht *HashTable) HashUint64(vecs []float64) uint64 {
	bs := ht.Hash(vecs)
	tt := uint64(bs[0])
	for i := 1; i < len(bs) && i < 8; i++ {
		tt |= uint64(bs[i]) << uint(i*8)
	}
	return tt
}

// dot calculates dot product, doesnt check if slices are the same length
func dot(vals1, vals2 []float64) float64 {
	out := 0.
	for i := range vals1 {
		out += vals1[i] * vals2[i]
	}
	return out
}
