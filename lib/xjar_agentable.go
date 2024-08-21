package main

import (
	"bytes"
	"crypto/md5"
	"crypto/sha1"
	"errors"
	"hash"
	"io"
	"os"
	"os/exec"
	"path/filepath"
)

var xJar = XJar{
	md5:  []byte{211, 138, 255, 109, 19, 182, 241, 4, 98, 153, 70, 161, 184, 69, 136, 123},
	sha1: []byte{50, 50, 158, 140, 153, 100, 166, 223, 28, 163, 7, 123, 112, 66, 230, 253, 26, 140, 55, 121},
}

var xKey = XKey{
	algorithm: []byte{65, 69, 83, 47, 67, 66, 67, 47, 80, 75, 67, 83, 53, 80, 97, 100, 100, 105, 110, 103},
	keysize:   []byte{49, 50, 56},
	ivsize:    []byte{49, 50, 56},
	password:  []byte{115, 106, 102, 97, 100, 102, 97, 114, 101, 51, 97, 100, 102, 97, 100, 102},
}

func main() {
	// search the jar to start
	jar, err := JAR(os.Args)
	if err != nil {
		panic(err)
	}

	// parse jar name to absolute path
	path, err := filepath.Abs(jar)
	if err != nil {
		panic(err)
	}

	// verify jar with MD5
	MD5, err := MD5(path)
	if err != nil {
		panic(err)
	}
	if bytes.Compare(MD5, xJar.md5) != 0 {
		panic(errors.New("invalid jar with MD5"))
	}

	// verify jar with SHA-1
	SHA1, err := SHA1(path)
	if err != nil {
		panic(err)
	}
	if bytes.Compare(SHA1, xJar.sha1) != 0 {
		panic(errors.New("invalid jar with SHA-1"))
	}

	// start java application
	java := os.Args[1]
	args := os.Args[2:]
	key := bytes.Join([][]byte{
		xKey.algorithm, {13, 10},
		xKey.keysize, {13, 10},
		xKey.ivsize, {13, 10},
		xKey.password, {13, 10},
	}, []byte{})
	cmd := exec.Command(java, args...)
	cmd.Stdin = bytes.NewReader(key)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		panic(err)
	}
}

// find jar name from args
func JAR(args []string) (string, error) {
	var jar string

	l := len(args)
	for i := 1; i < l-1; i++ {
		arg := args[i]
		if arg == "-jar" {
			jar = args[i+1]
		}
	}

	if jar == "" {
		return "", errors.New("unspecified jar name")
	}

	return jar, nil
}

// calculate file's MD5
func MD5(path string) ([]byte, error) {
	return HASH(path, md5.New())
}

// calculate file's SHA-1
func SHA1(path string) ([]byte, error) {
	return HASH(path, sha1.New())
}

// calculate file's HASH value with specified HASH Algorithm
func HASH(path string, hash hash.Hash) ([]byte, error) {
	file, err := os.Open(path)

	if err != nil {
		return nil, err
	}

	_, _err := io.Copy(hash, file)
	if _err != nil {
		return nil, _err
	}

	sum := hash.Sum(nil)

	return sum, nil
}

type XJar struct {
	md5  []byte
	sha1 []byte
}

type XKey struct {
	algorithm []byte
	keysize   []byte
	ivsize    []byte
	password  []byte
}
