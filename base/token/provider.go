package token

type Provider interface {
	Parse(str string) (map[string]interface{}, error)
	New(data map[string]interface{}, minutes uint) (string, error)
}
