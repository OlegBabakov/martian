// Copyright 2015 Google Inc. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package martianhttp

import (
	"crypto/x509"
	"encoding/pem"
	"net/http"
)

type authorityHandler struct {
	cert []byte
}

// NewAuthorityHandler returns an http.Handler that will present the client
// with the CA certificate to use in browser.
func NewAuthorityHandler(ca *x509.Certificate) http.Handler {
	return &authorityHandler{
		cert: pem.EncodeToMemory(&pem.Block{
			Type:  "CERTIFICATE",
			Bytes: ca.Raw,
		}),
	}
}

// ServeHTTP writes the CA certificate in PEM format to the client.
func (h *authorityHandler) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	rw.Header().Set("Content-Type", "application/x-x509-ca-cert")
	if _, err := rw.Write(h.cert); err != nil {
		// gologger.Debug().Msgf("%s\n", err)
	}
}
