package user

import (
	"context"
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"errors"
	"fmt"
	"strings"

	"golang.org/x/crypto/argon2"
)

// Argon2Params contain the required for running the Argon2 hashing algorithm. Make these numbers big but not too big
// See https://argon2-cffi.readthedocs.io/en/stable/parameters.html
type Argon2Params struct {
	Memory      uint32
	Iterations  uint32
	Parallelism uint8
	SaltLength  uint32
	KeyLength   uint32
}

var (
	// ErrInvalidHash is returned when the given encoded password is in the wrong format
	ErrInvalidHash = errors.New("The encoded hash is not in the correct format")
	// ErrIncompatibleVersion is returned when the given passwaord was hashed with a different version of argon2
	ErrIncompatibleVersion = errors.New("Incompatible version of argon2")
)

// Argon2ParamsCtx exists so we can change the hashing parameters during testing to speed up the process
type Argon2ParamsCtx struct{}

func getArgon2Params(ctx context.Context) *Argon2Params {
	// Establish the parameters to use for Argon2.
	p, ok := ctx.Value(Argon2ParamsCtx{}).(*Argon2Params)
	if !ok {
		// If the context does not have any parameters set, use these default production ready parameters
		return &Argon2Params{
			Memory:      64 * 1024,
			Iterations:  3,
			Parallelism: 2,
			SaltLength:  16,
			KeyLength:   32,
		}
	}
	return p
}

// HashPassword will generate a bcrypt hash using the default cost of 14
func HashPassword(ctx context.Context, password string) (encodedHash string, err error) {
	p := getArgon2Params(ctx)
	// Generate a cryptographically secure random salt.
	salt, err := generateRandomBytes(p.SaltLength)
	if err != nil {
		return "", err
	}

	// Pass the plaintext password, salt and parameters to the argon2.IDKey
	// function. This will generate a hash of the password using the Argon2id
	// variant.
	hash := argon2.IDKey([]byte(password), salt, p.Iterations, p.Memory, p.Parallelism, p.KeyLength)

	// Base64 encode the salt and hashed password.
	b64Salt := base64.RawStdEncoding.EncodeToString(salt)
	b64Hash := base64.RawStdEncoding.EncodeToString(hash)

	// Return a string using the standard encoded hash representation.
	encodedHash = fmt.Sprintf("$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s", argon2.Version, p.Memory, p.Iterations, p.Parallelism, b64Salt, b64Hash)

	return encodedHash, nil
}

// PasswordMatchesHash compares the plain text password with the given hash. Returns true if they match and false otherwise
func PasswordMatchesHash(password string, hash string) bool {
	matches, err := comparePasswordAndHash(password, hash)
	if err != nil {
		fmt.Printf("comparePasswordAndHash returned err %+v\n", err)
		return false
	}
	return matches
}

func generateRandomBytes(n uint32) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}

	return b, nil
}

func comparePasswordAndHash(password string, encodedHash string) (match bool, err error) {
	// Extract the parameters, salt and derived key from the encoded password
	// hash.
	p, salt, hash, err := decodeHash(encodedHash)
	if err != nil {
		return false, err
	}

	// Derive the key from the other password using the same parameters.
	otherHash := argon2.IDKey([]byte(password), salt, p.Iterations, p.Memory, p.Parallelism, p.KeyLength)

	// Check that the contents of the hashed passwords are identical. Note
	// that we are using the subtle.ConstantTimeCompare() function for this
	// to help prevent timing attacks.
	if subtle.ConstantTimeCompare(hash, otherHash) == 1 {
		return true, nil
	}
	return false, nil
}

func decodeHash(encodedHash string) (p *Argon2Params, salt, hash []byte, err error) {
	vals := strings.Split(encodedHash, "$")
	if len(vals) != 6 {
		return nil, nil, nil, ErrInvalidHash
	}

	var version int
	_, err = fmt.Sscanf(vals[2], "v=%d", &version)
	if err != nil {
		return nil, nil, nil, err
	}
	if version != argon2.Version {
		return nil, nil, nil, ErrIncompatibleVersion
	}

	p = &Argon2Params{}
	_, err = fmt.Sscanf(vals[3], "m=%d,t=%d,p=%d", &p.Memory, &p.Iterations, &p.Parallelism)
	if err != nil {
		return nil, nil, nil, err
	}

	salt, err = base64.RawStdEncoding.DecodeString(vals[4])
	if err != nil {
		return nil, nil, nil, err
	}
	p.SaltLength = uint32(len(salt))

	hash, err = base64.RawStdEncoding.DecodeString(vals[5])
	if err != nil {
		return nil, nil, nil, err
	}
	p.KeyLength = uint32(len(hash))

	return p, salt, hash, nil
}
