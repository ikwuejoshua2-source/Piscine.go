package piscine

func ForEach(callback func (int)), list [] int) {
	for_, v:= range list {
		callback (v)
	}
}