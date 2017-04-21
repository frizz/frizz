package root

// notest

import (
	"fmt"
	"net/http"
)

func init() {
	http.HandleFunc("/", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, `<!DOCTYPE html>
<html>
<head>
	<meta http-equiv="Content-Type" content="text/html; charset=utf-8"/>
	<meta name="go-import" content="frizz.io git https://github.com/frizz/frizz">
	<meta name="go-source" content="frizz.io https://github.com/frizz/frizz/tree/master https://github.com/frizz/frizz/tree/master{/dir} https://github.com/frizz/frizz/blob/master{/dir}/{file}#L{line}">
	<meta http-equiv="refresh" content="0; url=https://github.com/frizz/frizz">
</head>`)
}
