package http_helper

type HttpHelper interface {
	Fetch(url string) ([]byte, error)
	PutJsonData(url string, data []byte) ([]byte, error)
}
