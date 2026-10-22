# nilepatch-coraza-bypass

Go app using Coraza WAF before 3.8.1 (CVE-2026-114101).

## Vulnerability

Coraza before 3.8.1 extracts multipart upload filenames with Go's `mime.ParseMediaType`,
which silently ignores RFC 5987 `filename*` parameters whose charset is not utf-8 or us-ascii.
Attackers can send a decoy filename with `filename*=iso-8859-1''shell.php` to bypass FILES-based
WAF rules, letting dangerous uploads reach backends.

## Expected fix

Upgrade Coraza to 3.8.1 or later in go.mod.

## Build

```sh
go mod tidy
go build
```
