package entity

type Flight [2]string

func (f *Flight) Valid() bool {
	if len(f[0]) != 3 {
		return false
	}
	if len(f[1]) != 3 {
		return false
	}
	return true
}

func (f *Flight) Depart() string {
	return f[0]
}

func (f *Flight) Arrive() string {
	return f[1]
}
