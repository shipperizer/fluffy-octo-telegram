package jwks

import (
	"fmt"
	"io/ioutil"
	"sync"

	"github.com/lestrrat-go/jwx/jwk"
	log "github.com/sirupsen/logrus"
)

func LoadKeyFromFile(filename string) ([]byte, error) {
	key, err := ioutil.ReadFile(filename)
	return key, err
}

func LoadJWKs(pubPath, privPath string) (jwk.Set, jwk.Key, error) {
	var set jwk.Set
	var key jwk.Key

	var wg sync.WaitGroup
	errs := make([]error, 0)

	wg.Add(2)

	go func() {
		var err error

		set, err = jwk.ReadFile(pubPath)

		if err != nil {
			log.Errorf("cannot find the file %s: %v", pubPath, err)
			errs = append(errs, err)
		}

		wg.Done()
	}()

	go func() {
		var err error

		keyRaw, err := LoadKeyFromFile(privPath)

		if err != nil {
			log.Errorf("cannot find the file %s: %v...moving on", privPath, err)
			// errs = append(errs, err)
			return
		}

		key, err = jwk.ParseKey(keyRaw)

		if err != nil {
			log.Errorf("error parsing key %s: %v...moving on", keyRaw, err)
			// errs = append(errs, err)
		}

		wg.Done()
	}()

	wg.Wait()

	for _, e := range errs {
		if e != nil {
			return set, key, fmt.Errorf("error fetching keys: %v", errs)
		}
	}

	return set, key, nil
}
