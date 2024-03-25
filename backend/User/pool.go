package User

type pool struct {
	cache []Info
	index map[string]Info
}

func (p *pool) add(user Info) {
	p.cache = append(p.cache, user)
}

func (p *pool) remove(user Info) {
	for i, v := range p.cache {
		if v.ID == user.ID {
			p.cache = append(p.cache[:i], p.cache[i+1:]...)
			break
		}
	}
}

func (p *pool) getByName(name string) Info {
	if v, ok := p.index[name]; ok {
		return v
	}
	return Info{}
}
func (p *pool) get(id int) *Info {
	for _, v := range p.cache {
		if v.ID == id {
			return &v
		}
	}
	return nil
}
