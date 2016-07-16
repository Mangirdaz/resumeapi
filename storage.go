package main

import (
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/docker/libkv"
	"github.com/docker/libkv/store"
	consul "github.com/docker/libkv/store/consul"
	"time"
)

type Note struct {
	Path string
	Key  string
	Note string
}

type Notes struct {
	Note []Note
}

// LibKVBackend - libkv container
type LibKVBackend struct {
	Store store.Store
}

// Put - puts object into kv store
func (l *LibKVBackend) Put(namespace string, key string, value []byte) error {
	return l.Store.Put(namespace+"/"+key, value, nil)
}

func (l *LibKVBackend) GetAll(namespace string) (notes Notes, err error) {
	pair, err := l.Store.List(namespace + "/")
	if err != nil {
		log.Error("Directory not found: ", err)
	} else {
		items := []Note{}
		box := Notes{items}
		for _, v := range pair {
			item := Note{
				Path: "notes",
				Key:  v.Key,
				Note: string(v.Value[:]),
			}
			box.AddItem(item)
		}
		log.Info(box)

		return box, nil
	}
	return Notes{}, err
}

func (l *LibKVBackend) Get(namespace string, key string) (notes Notes, err error) {
	pair, err := l.Store.Get(namespace + "/" + key)
	if err != nil {
		log.Warning("Key not found")
	} else {
		items := []Note{}
		box := Notes{items}

		item := Note{
			Path: "notes",
			Key:  pair.Key,
			Note: string(pair.Value[:]),
		}
		box.AddItem(item)
		log.Info(box)
		return box, nil
	}
	return Notes{}, fmt.Errorf("error: key not found ")
}

func (box *Notes) AddItem(item Note) []Note {
	box.Note = append(box.Note, item)
	return box.Note
}

func NewLibKVBackend() (storage LibKVBackend) {
	log.Info("Create New LibKV backend")
	config := InitKeyValueStorageConfig()
	consul.Register()

	client := config.Ip + ":" + config.Port

	log.Info("Init New store")
	kv, err := libkv.NewStore(
		store.CONSUL, // or "consul"
		[]string{client},
		&store.Config{
			ConnectionTimeout: 10 * time.Second,
		},
	)
	if err != nil {
		log.Fatal("Cannot create store consul")
	}

	var backend LibKVBackend
	backend.Store = kv
	log.Info("Store init")
	return backend

}
