package etcdv3

import (
	"testing"
	"time"

	"github.com/microsvs/libkv"
	"github.com/microsvs/libkv/store"
	"github.com/microsvs/libkv/testutils"
	"github.com/stretchr/testify/assert"
)

var (
	client = "localhost:4001"
)

func makeEtcdClient(t *testing.T) store.Store {
	kv, err := New(
		[]string{client},
		&store.Config{
			ConnectionTimeout: 3 * time.Second,
			//Username:          "test",
			//Password:          "very-secure",
		},
	)

	if err != nil {
		t.Fatalf("cannot create store: %v", err)
	}

	return kv
}

func TestRegister(t *testing.T) {
	Register()

	kv, err := libkv.NewStore(store.ETCDV3, []string{client}, nil)
	assert.NoError(t, err)
	assert.NotNil(t, kv)

	if _, ok := kv.(*Etcd); !ok {
		t.Fatal("Error registering and initializing etcd")
	}
}

func TestEtcdStore(t *testing.T) {
	cli := makeEtcdClient(t)

	testutils.RunTestCommon(t, cli)
	testutils.RunTestAtomic(t, cli)
	testutils.RunTestWatch(t, cli)
	//testutils.RunTestLock(t, kv)
	//testutils.RunTestLockTTL(t, kv, lockKV)
	//testutils.RunTestTTL(t, cli, cli)
	testutils.RunCleanup(t, cli)
}
