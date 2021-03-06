package bboltdbcache

import (
	"bytes"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/sniperkit/xcache/backend/local/boltdb/bbolt"
)

func TestBoltDBCache(t *testing.T) {
	tempDir, err := ioutil.TempDir("", "httpcache")
	if err != nil {
		t.Fatalf("TempDir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	cache, err := bbolt.New(filepath.Join(tempDir, "db"))
	if err != nil {
		t.Fatalf("bbolt.New(): %v", err)
	}

	key := "testKey"
	_, ok := cache.Get(key)
	if ok {
		t.Fatal("retrieved key before adding it")
	}

	val := []byte("some bytes")
	cache.Set(key, val)

	retVal, ok := cache.Get(key)
	if !ok {
		t.Fatal("could not retrieve an element we just added")
	}
	if !bytes.Equal(retVal, val) {
		t.Fatal("retrieved a different value than what we put in")
	}

	cache.Delete(key)

	_, ok = cache.Get(key)
	if ok {
		t.Fatal("deleted key still present")
	}
}
