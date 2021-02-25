package main

import "net/http"

func addCookie(response http.ResponseWriter, name, value string) {
	// expire := time.Now().Add(ttl)
	cookie := http.Cookie{
		Name:  name,
		Value: value,
		// Expires: expire,
	}
	http.SetCookie(response, &cookie)
}
