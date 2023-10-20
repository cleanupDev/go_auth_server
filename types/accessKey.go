package responseTypes

type AccessKey string

func NewAccessKey(key *string) *AccessKey {
	// TODO: implement access key logic
	AccessKey := AccessKey(*key)
	return &AccessKey
}
