package md5hasher

import (
	"crypto/md5"
	"fmt"
	"io"
	"net/http"
)

type client interface {
	Get(string) (*http.Response, error)
}

type hasher struct {
	client
}

// New creates hasher with specified client
func New(c client) hasher {
	return hasher{c}
}

// Hash gets md5 from url response body
func (h *hasher) Hash(url string) (string, error) {
	resp, err := h.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	hasher := md5.New()

	_, err = io.Copy(hasher, resp.Body)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", hasher.Sum(nil)), nil
}
