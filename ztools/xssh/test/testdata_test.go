// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// IMPLEMENTATION NOTE: To avoid a package loop, this file is in three places:
// ssh/, ssh/agent, and ssh/test/. It should be kept in sync across all three
// instances.

package test

import (
	"crypto/rand"
	"fmt"

	"github.com/baykovr/zgrab/ztools/xssh"
	"github.com/baykovr/zgrab/ztools/xssh/testdata"
)

var (
	testPrivateKeys map[string]interface{}
	testSigners     map[string]xssh.Signer
	testPublicKeys  map[string]xssh.PublicKey
)

func init() {
	var err error

	n := len(testdata.PEMBytes)
	testPrivateKeys = make(map[string]interface{}, n)
	testSigners = make(map[string]xssh.Signer, n)
	testPublicKeys = make(map[string]xssh.PublicKey, n)
	for t, k := range testdata.PEMBytes {
		testPrivateKeys[t], err = xssh.ParseRawPrivateKey(k)
		if err != nil {
			panic(fmt.Sprintf("Unable to parse test key %s: %v", t, err))
		}
		testSigners[t], err = xssh.NewSignerFromKey(testPrivateKeys[t])
		if err != nil {
			panic(fmt.Sprintf("Unable to create signer for test key %s: %v", t, err))
		}
		testPublicKeys[t] = testSigners[t].PublicKey()
	}

	// Create a cert and sign it for use in tests.
	testCert := &xssh.Certificate{
		Nonce:           []byte{},                       // To pass reflect.DeepEqual after marshal & parse, this must be non-nil
		ValidPrincipals: []string{"gopher1", "gopher2"}, // increases test coverage
		ValidAfter:      0,                              // unix epoch
		ValidBefore:     xssh.CertTimeInfinity,          // The end of currently representable time.
		Reserved:        []byte{},                       // To pass reflect.DeepEqual after marshal & parse, this must be non-nil
		Key:             testPublicKeys["ecdsa"],
		SignatureKey:    testPublicKeys["rsa"],
		Permissions: xssh.Permissions{
			CriticalOptions: map[string]string{},
			Extensions:      map[string]string{},
		},
	}
	testCert.SignCert(rand.Reader, testSigners["rsa"])
	testPrivateKeys["cert"] = testPrivateKeys["ecdsa"]
	testSigners["cert"], err = xssh.NewCertSigner(testCert, testSigners["ecdsa"])
	if err != nil {
		panic(fmt.Sprintf("Unable to create certificate signer: %v", err))
	}
}
