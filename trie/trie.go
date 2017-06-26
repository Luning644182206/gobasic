// Copyright (c) 2017, The gobasic Authors.
// All rights reserved.
//
// Author: Zheng Gonglin <scaugrated@gmail.com>
// Created: 2017/06/23

package trie

import (
    "errors"
)

var (
    errorKeyExisted         = errors.New("key was existed")
    errorKeyNotExisted      = errors.New("key was not existed")
    // errorTrieSearchFailed   = errors.New("search trie exception")
)

type node struct {
    index       byte
    father      *node
    children    map[byte]*node
    isEnd       bool
    value       interface{}
}

type Trie struct {
    root        *node
}

func newNode(father *node, index byte, value interface{}) (trieNode *node) {
    return &node{
        index:      index,
        father:     father,
        children:   make(map[byte]*node),
        isEnd:      false,
        value:      value,
    }
}

func NewTrie() (t *Trie) {
    return &Trie{
        root:   newNode(nil, 0, nil),
    }
}

func (t *Trie) Insert(key []byte, value interface{}) (err error) {
    var lastNode *node = nil

    lastNode, err = t.search(key, true)
    if err != nil {
        return err
    }

    if lastNode.isEnd {
        return errorKeyExisted
    }

    lastNode.isEnd = true
    lastNode.value = value

    return nil
}

func (t *Trie) Update(key []byte, value interface{}) (err error) {
    var lastNode *node = nil

    lastNode, err = t.search(key, false)
    if err != nil {
        return err
    }

    if lastNode == nil || !lastNode.isEnd {
        return errorKeyNotExisted
    }

    lastNode.value = value
    return nil
}

func (t *Trie) Find(key []byte) (value interface{}, err error) {
    var lastNode *node = nil

    lastNode, err = t.search(key, false)
    if err != nil {
        return nil, err
    }

    if lastNode == nil || !lastNode.isEnd {
        return nil, errorKeyNotExisted
    }

    return lastNode.value, nil
}

func (t *Trie) Delete(key []byte) (err error) {
    var lastNode *node = nil

    lastNode, err = t.search(key, false)
    if err != nil {
        return err
    }

    if lastNode == nil || !lastNode.isEnd {
        return errorKeyNotExisted
    }

    lastNode.isEnd = false
    lastNode.value = nil

    if len(lastNode.children) == 0 && lastNode.father != nil {
        delete(lastNode.father.children, lastNode.index)
    }

    return nil
}

func (t *Trie) Keys() (keys [][]byte) {
    buffer := make([]byte, 0, 1024)
    pkeys := &keys
    pbuffer := &buffer

    if t.root.isEnd {
        *pkeys = append(*pkeys, []byte{})
    }
    for _, next := range(t.root.children) {
        t.dfs(next, 0, pbuffer, pkeys)
    }
    return keys
}

func (t *Trie) dfs(ptr *node, depth int, pbuffer *[]byte, pkeys *[][]byte) {
    if len(*pbuffer) > depth {
        (*pbuffer)[depth] = ptr.index
    } else {
        *pbuffer = append(*pbuffer, ptr.index)
    }

    if ptr.isEnd {
        key := make([]byte, depth+1)
        copy(key, (*pbuffer)[:depth+1])
        *pkeys = append(*pkeys, key)
    }

    for _, next := range ptr.children {
        t.dfs(next, depth+1, pbuffer, pkeys)
    }
}

func (t *Trie) search(key []byte, isNew bool) (lastNode *node, err error) {
    size := len(key)

    lastNode = t.root
    err = nil

    for i := 0; i < size; i++ {
        b := key[i]
        cnode, ok := lastNode.children[b]
        if !ok {
            if isNew {
                cnode = newNode(lastNode, b, nil)
                lastNode.children[b] = cnode
            } else {
                return nil, errorKeyNotExisted
            }
        }
        lastNode = cnode
    }
    return lastNode, err
}

