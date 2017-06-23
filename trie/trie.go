// Copyright (c) 2017, The zkawa Authors.
// All rights reserved.
//
// Author: Zheng Gonglin <scaugrated@gmail.com>
// Created: 2017/06/23

package trie

import (
    // "log"
    "errors"
)

var (
    errorKeyExisted = errors.New("key was existed")
    errorKeyNotExisted = errors.New("key was not existed")
    errorTrieSearchFailed = errors.New("search trie exception")
)

type node struct {
    index       byte
    children    map[byte]*node
    isEnd       bool
    value       interface{}
}

type Trie struct {
    root        *node
}

func newNode(index byte, value interface{}) (trieNode *node) {
    return &node{
        index:      index,
        children:   make(map[byte]*node),
        isEnd:      false,
        value:      value,
    }
}

func NewTrie() (*Trie, error) {
    return &Trie{
        root:   newNode(0, nil),
    }, nil
}

func (t *Trie) search(key []byte, isNew bool) (fatherNode *node, lastNode *node, err error) {
    size := len(key)

    fatherNode = nil
    lastNode = t.root
    err = nil

    for i := 0; i < size; i++ {
        b := key[i]
        cnode, ok := lastNode.children[b]
        if !ok {
            if isNew {
                cnode = newNode(b, nil)
                lastNode.children[b] = cnode
            } else {
                return nil, nil, errorKeyNotExisted
            }
        }
        fatherNode = lastNode
        lastNode = cnode
    }
    return fatherNode, lastNode, err
}

func (t *Trie) Insert(key []byte, value interface{}) (err error) {
    var lastNode *node = nil

    _, lastNode, err = t.search(key, true)
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

    _, lastNode, err = t.search(key, false)
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

    _, lastNode, err = t.search(key, false)
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
    var fatherNode *node = nil

    fatherNode, lastNode, err = t.search(key, false)
    if err != nil {
        return err
    }

    if lastNode == nil || !lastNode.isEnd {
        return errorKeyNotExisted
    }

    if fatherNode == nil && len(key) > 0{
        return errorTrieSearchFailed
    }

    lastNode.isEnd = false
    lastNode.value = nil

    return nil
}
