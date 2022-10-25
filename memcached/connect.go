package memcached

import (
	"os"

	"github.com/bradfitz/gomemcache/memcache"
	"github.com/ecsaas/memcache-define-config/DEFINE_VARIABLES/MEMCACHED_ENV"
)

var MC *memcache.Client

func Memcached() {
	if MC == nil {
		MC = &memcache.Client{}
	}
	*MC = *memcache.New(os.Getenv(MEMCACHED_ENV.MCD_SERVER))
}

func Set(mc *memcache.Client, Key string, Value []byte) error {
	return mc.Set(&memcache.Item{Key: Key, Value: Value})
}

func Delete(mc *memcache.Client, Key string) error {
	return mc.Delete(Key)
}

func Touch(mc *memcache.Client, Key string, Seconds int) error {
	return mc.Touch(Key, int32(Seconds))
}

func Get(mc *memcache.Client, Key string) (checkCache bool, Value []byte) {
	it, err := mc.Get(Key)
	if err == nil && it != nil {
		checkCache = it.Key == Key
		Value = it.Value
	}
	return
}

//checkCache, value := memcached.Get(e.MC, "key")
//fmt.Println(checkCache, string(value))
//
//fmt.Println(memcached.Set(e.MC, "key", []byte(fmt.Sprintf("%d", time.Now().UnixNano()))))
//fmt.Println(memcached.Touch(e.MC, "key", 5))
//
//checkCache, value = memcached.Get(e.MC, "key")
//fmt.Println(checkCache, string(value))
