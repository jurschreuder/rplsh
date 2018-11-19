![gopher.png](gopher.png)

Random Projection Location Sensitive Hash in Go
=========================================

Random Projection LSH is a way to create a binary hash from a vector, with closer vectors having a higher chance of having the same hash.

### Steps

 - create a random vector with the length of the hash
 - calculate the dot product of the to-be-encoded hash with the random vector
 - with output: All positive values become 1, all negative values become 0

## New random hash table
To hash vectors of []float64 with length 128.\
Create a hashes of 16 bits.
```
ht := NewHashTable(128, 16)
```

## Create a []byte hash
```
vecs := make([]float64, 128)

bs := ht.Hash(vecs)
```

## Create a uint hash
Supports: uint8, uint16, uint32, uint64.

```
vecs := make([]float64, 128)

bs := ht.HashUint16(vecs)
```

## Save hash table for consistent results
```
err := ht.Save("save_file.js")
```

## Load hash table from save file
```
hs, err := ht.Load("save_file.js")
```
