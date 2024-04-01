package leveldb

import (
	"testing"

	go_fuzz_utils "github.com/trailofbits/go-fuzz-utils"
)

func GetTypeProvider(data []byte) (*go_fuzz_utils.TypeProvider, error) {
	tp, err := go_fuzz_utils.NewTypeProvider(data)
	if err != nil {
		return nil, err
	}
	err = tp.SetParamsStringBounds(0, 1024)
	if err != nil {
		return nil, err
	}
	err = tp.SetParamsSliceBounds(0, 4096)
	if err != nil {
		return nil, err
	}
	err = tp.SetParamsBiases(0, 0, 0, 0)
	if err != nil {
		return nil, err
	}
	return tp, nil
}

func Fuzz_Nosy_LevelDBStore_Delete__(f *testing.F) {
	f.Fuzz(func(t *testing.T, path string, key []byte) {
		d, err := NewLevelDBStore(path)
		if err != nil {
			return
		}
		d.Delete(key)
	})
}

func Fuzz_Nosy_LevelDBStore_DeleteBatch__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var path string
		fill_err = tp.Fill(&path)
		if fill_err != nil {
			return
		}
		var keys [][]byte
		fill_err = tp.Fill(&keys)
		if fill_err != nil {
			return
		}

		d, err := NewLevelDBStore(path)
		if err != nil {
			return
		}
		d.DeleteBatch(keys)
	})
}

func Fuzz_Nosy_LevelDBStore_Get__(f *testing.F) {
	f.Fuzz(func(t *testing.T, path string, key []byte) {
		d, err := NewLevelDBStore(path)
		if err != nil {
			return
		}
		d.Get(key)
	})
}

func Fuzz_Nosy_LevelDBStore_NewIterator__(f *testing.F) {
	f.Fuzz(func(t *testing.T, path string, prefix []byte) {
		d, err := NewLevelDBStore(path)
		if err != nil {
			return
		}
		d.NewIterator(prefix)
	})
}

func Fuzz_Nosy_LevelDBStore_Put__(f *testing.F) {
	f.Fuzz(func(t *testing.T, path string, key []byte, value []byte) {
		d, err := NewLevelDBStore(path)
		if err != nil {
			return
		}
		d.Put(key, value)
	})
}

func Fuzz_Nosy_LevelDBStore_WriteBatch__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var path string
		fill_err = tp.Fill(&path)
		if fill_err != nil {
			return
		}
		var keys [][]byte
		fill_err = tp.Fill(&keys)
		if fill_err != nil {
			return
		}
		var values [][]byte
		fill_err = tp.Fill(&values)
		if fill_err != nil {
			return
		}

		d, err := NewLevelDBStore(path)
		if err != nil {
			return
		}
		d.WriteBatch(keys, values)
	})
}
