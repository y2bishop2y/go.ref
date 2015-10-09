// Copyright 2015 The Vanadium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package bcrypter implements mechanisms for encrypting
// messages under blessing-pattern policies.
//
// A message encrypted under a specific blessing pattern can
// only be decrypted by principals possessing a blessing
// matching the pattern.
package bcrypter

import (
	"v.io/v23/context"
	"v.io/v23/security"

	"v.io/x/lib/ibe"
)

// The cryptographic scheme used for encrypting messages for blessing
// patterns.
type Scheme int32

// Plaintext represents the plaintext used by the encryption operation.
type Plaintext ibe.Plaintext

// Ciphertext represents the ciphertext generated by the encryption
// operation.
type Ciphertext struct {
	// Scheme is the cryptographic scheme used.
	Scheme Scheme
	// Ciphertexts is a map from a hash (the scheme would fix the hash
	// function) of a blessing pattern to an encryption of a plaintext
	// message for that pattern.
	Ciphertexts map[string]*ibe.Ciphertext
}

// Encrypter is the interface for encrypting a plaintext message under a
// set of blessing-patterns.
//
// Typically, this interface would be used for encrypting a symmetric key
// which in turn would be used to encrypt an arbitrary length data message.
type Encrypter interface {
	// Encrypt encrypts the provided 'plaintext' for the provided
	// blessing patterns, and returns the ciphertext (or nil if an
	// error is returned).
	Encrypt(ctx *context.T, forPatterns []security.BlessingPattern, plaintext *Plaintext) (*Ciphertext, error)
}

// Decrypter is the interface for decrypting ciphertexts produced by an
// encrypter.
type Decrypter interface {
	// Decrypt decrypts the provided 'ciphertext' and returns the corresponding
	// plaintext (or nil if an error is returned).
	Decrypt(ctx *context.T, ciphertext *Ciphertext) (*Plaintext, error)
}
