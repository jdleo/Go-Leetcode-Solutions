package main

import "strconv"

type Codec struct {
	urls []string // array to hold url mappings
	ptr  int      // pointer for current url's mapping
}

func Constructor() Codec {
	return Codec{[]string{}, 0}
}

// Encodes a URL to a shortened URL.
func (this *Codec) encode(longUrl string) string {
	// encode current url, store url, increase pointer
	this.urls = append(this.urls, longUrl)
	this.ptr++
	return strconv.Itoa(this.ptr - 1)
}

// Decodes a shortened URL to its original URL.
func (this *Codec) decode(shortUrl string) string {
	// get url at this ptr
	ptr, _ := strconv.Atoi(shortUrl)
	return this.urls[ptr]
}
