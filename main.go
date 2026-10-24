package main

import (
	"fmt"
	"net/http"

	"github.com/corazawaf/coraza/v3"
)

func main() {
	waf, _ := coraza.NewWAF(coraza.NewWAFConfig().
		WithDirectives(`SecRuleEngine On`))

	http.HandleFunc("/upload", func(w http.ResponseWriter, r *http.Request) {
		// VULNERABLE: Coraza before 3.8.1 bypasses FILES rules via filename* charset
		// An attacker can send filename*=iso-8859-1''shell.php to bypass WAF
		tx := waf.NewTransaction()
		defer tx.ProcessLogging()
		tx.ProcessRequestHeaders()
		// ... handle upload
		fmt.Fprintf(w, "Upload accepted")
	})

	http.ListenAndServe(":8080", nil)
}
