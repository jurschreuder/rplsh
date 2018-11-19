Random Projection Location Sensitive Hash
=========================================

Random Projection LSH is a way to create a binary hash from a vector, with closer vectors having a higher chance of having the same hash.

### steps

 - create a random vector with the length of the hash
 - calculate the dot product of the to-be-encoded hash with the random vector
 - with output: All positive values become 1, all negative values become 0

# initialize a new random hash table
To hash vectors of []float64 with length 128.
Create a hashes of 16 bits
```
ht := NewHashTable(128, 16)
```

# create a []byte hash
```
vecs := make([]float64, 128)

bs := ht.Hash(vecs)
```

# create a uint hash
supports: uint8, uint16, uint32, uint64

```
vecs := make([]float64, 128)

bs := ht.HashUint16(vecs)
```

# save hash table for consistent results
```
err := ht.Save("save_file.js")
```

# load hash table from save file
```
hs, err := ht.Load("save_file.js")
```
