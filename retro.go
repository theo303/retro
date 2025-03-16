package retro

func getSticky(ss []*Sticky, id string) *Sticky {
	for _, s := range ss {
		if s.Id == id {
			return s
		}
	}
	return nil
}
