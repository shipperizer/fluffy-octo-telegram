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

func LoadJWKs(pubPath, privPath string) (jwk.Set, jwk.Set, error) {
	var pub jwk.Set
	var prv jwk.Set

	var wg sync.WaitGroup
	errs := make([]error, 0)

	wg.Add(2)

	go func() {
		var err error

		pub, err = jwk.ReadFile(pubPath)

		if err != nil {
			log.Errorf("cannot find the file %s: %v", pubPath, err)
			errs = append(errs, err)
		}

		wg.Done()
	}()

	go func() {
		var err error

		prv, err = jwk.ReadFile(privPath)

		if err != nil {
			log.Errorf("cannot find the file %s: %v", pubPath, err)
			errs = append(errs, err)
		}

		wg.Done()
	}()

	wg.Wait()

	for _, e := range errs {
		if e != nil {
			return pub, prv, fmt.Errorf("error fetching keys: %v", errs)
		}
	}

	return pub, prv, nil
}
