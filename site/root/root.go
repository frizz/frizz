package root

// ke: {"package": {"notest":true}}

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
	<meta name="go-import" content="kego.io git https://github.com/kego/ke">
	<meta name="go-source" content="kego.io https://github.com/kego/ke/tree/master https://github.com/kego/ke/tree/master{/dir} https://github.com/kego/ke/blob/master{/dir}/{file}#L{line}">
	<meta http-equiv="refresh" content="0; url=https://github.com/kego/ke">
</head>`)
}
