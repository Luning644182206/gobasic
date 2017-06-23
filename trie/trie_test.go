// Copyright (c) 2017, The zkawa Authors.
// All rights reserved.
//
// Author: Zheng Gonglin <scaugrated@gmail.com>
// Created: 2017/06/23

package trie

import (
    "math/rand"
    "testing"
)

const (
    randomCharSet = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
    lengthOfRandomCharSet = len(randomCharSet)
)

func assertEqualByteSlice(xbs, ybs []byte, t *testing.T) {
    if len(xbs) != len(ybs) {
        t.Fatalf("byte slice length different: %d != %d\n", len(xbs), len(ybs))
    }
    for idx, b := range xbs {
        if b != ybs[idx] {
            t.Fatalf("index of [%d], '%s' != '%s'\n", idx, b, ybs[idx])
        }
    }
}

func randomBytes(length int) []byte {
    ret := make([]byte, 0, length)
    for i := 0; i < length; i++ {
        idx := rand.Intn(lengthOfRandomCharSet)
        ret = append(ret, randomCharSet[idx])
    }
    return ret
}

func TestComperhensive(t *testing.T) {
    optNums := 100000
    keys := make([][]byte, 0)
    values := make([][]byte, 0)
    dict := make(map[string][]byte)
    mytrie, err := NewTrie()
    for i := 0; i < optNums; i++ {
        opt := rand.Intn(4)
        if opt == 0 {
        // Trie.Insert
            keyLen := rand.Intn(6);
            key := randomBytes(keyLen)
            value := randomBytes(keyLen)
            keys = append(keys, key)
            values = append(values, value)

            _, ok := dict[string(key)]
            err = mytrie.Insert(key, value)
            if ok && err == nil {
                t.Fatalf("insert key failed, details: %v", err)
            }
            if !ok && err != nil {
                t.Fatalf("insert key failed, details: %v", err)
            }

            if !ok {
                dict[string(key)] = value
                // t.Logf("insert key '%s'\n", string(key))
            }
        } else if opt == 1 {
        // Trie.Update
            var key []byte
            var value []byte
            k := rand.Intn(2)
            if k == 1 && len(keys) > 0 {
                idx := rand.Intn(len(keys))
                key = keys[idx]
                value = values[idx]
            } else {
                keyLen := rand.Intn(6);
                key = randomBytes(keyLen)
                value = randomBytes(keyLen)
            }
            _, ok := dict[string(key)]
            err = mytrie.Update(key, value)
            if !ok && err == nil {
                t.Fatalf("update key failed, details: %v", err)
            }
            if ok && err != nil {
                t.Fatalf("update key failed, details: %v", err)
            }
            if ok {
                dict[string(key)] = value
            }
        } else if opt == 2 {
        // Trie.Find
            var key []byte
            k := rand.Intn(2)
            if k == 1 && len(keys) > 0 {
                idx := rand.Intn(len(keys))
                key = keys[idx]
            } else {
                keyLen := rand.Intn(6);
                key = randomBytes(keyLen)
            }
            v1, ok := dict[string(key)]
            v2, err := mytrie.Find(key)
            if !ok && err == nil {
                t.Fatalf("find key failed, details: %v", err)
            }
            if ok && err != nil {
                t.Fatalf("find key '%s' failed, details: %v", string(key), err)
            }
            if ok {
                v2_, ok := v2.([]byte)
                if !ok {
                    t.Fatal("value was not byte slice.")
                }
                assertEqualByteSlice(v1, v2_, t)
            }
        } else {
        // Trie.Delete
            var key []byte
            k := rand.Intn(2)
            if k == 1 && len(keys) > 0 {
                idx := rand.Intn(len(keys))
                key = keys[idx]
            } else {
                keyLen := rand.Intn(6);
                key = randomBytes(keyLen)
            }
            _, ok := dict[string(key)]
            err = mytrie.Delete(key)
            if !ok && err == nil {
                t.Fatalf("delete key failed, details: %v", err)
            }
            if ok && err != nil {
                t.Fatalf("delete key failed, details: %v", err)
            }
            if ok {
                delete(dict, string(key))
                // t.Logf("delete key '%s'\n", string(key))
            }
        }
    }
}
