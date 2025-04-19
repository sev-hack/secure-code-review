# SSRF (Server-Side Request Forgery)

Although the code tries to prevent access to internal resources via `checkBlacklistedURL` function, checks can be bypassed:
- Using DNS rebinding (where a domain returns different IPs depending on the source of the request).
- Using of IPv6 addresses or alternate forms of IP records.
- Using domains that resolve to resolved IPs but redirect (HTTP 3xx) to internal addresses.

Also the code doesn't check the URL scheme, which allows the use of `file://`, `ftp://` and other schemes
