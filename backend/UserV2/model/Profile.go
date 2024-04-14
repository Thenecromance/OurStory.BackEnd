package model

type Profile struct {
}

func (p *Profile) GetProfile(username string) *Profile {
	return p
}

func NewProfile() *Profile {
	return &Profile{}
}
