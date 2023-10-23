package utils

func PrefixHTTP(url *string) {
	if (*url)[:4] != "http" {
		(*url) = "http://" + *url
	}
}
