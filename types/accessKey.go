package responseTypes

type AccessKey string

func NewAccessKey() *AccessKey {
	// TODO: implement access key logic
	key := "Not implemented yet"
	AccessKey := AccessKey(key)
	return &AccessKey
}
