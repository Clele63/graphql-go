package module

import (
	"strconv"

	"github.com/dgraph-io/badger/v4"
)

func (m *UsersModule) NextID(txn *badger.Txn) (int, error) {
	item, err := txn.Get([]byte(m.NamespaceSeq()))
	if err == badger.ErrKeyNotFound {
		if err := txn.Set([]byte(m.NamespaceSeq()), []byte("1")); err != nil {
			return 1, err
		}
		return 1, nil
	} else if err != nil {
		return 0, err
	}

	var current int
	err = item.Value(func(v []byte) error {
		current, _ = strconv.Atoi(string(v))
		return nil
	})
	if err != nil {
		return 0, err
	}

	newID := current + 1
	if err := txn.Set([]byte(m.NamespaceSeq()), []byte(strconv.Itoa(newID))); err != nil {
		return 0, err
	}
	return newID, nil
}
