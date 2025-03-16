package retro

import (
	"slices"
)

func getSticky(ss []*Sticky, id string) *Sticky {
	for _, s := range ss {
		if s.Id == id {
			return s
		}
	}
	return nil
}

func riseSticky(ss []*Sticky, id string) *Sticky {
	for i, s := range ss {
		if s.Id == id {
			ss = slices.Delete(ss, i, i+1)
			ss = append(ss, s)
			return s
		}
	}
	return nil
}

// func formatStickies(ss []*Sticky) string {
// 	var sb strings.Builder
// 	for i, s := range ss {
// 		if i != 0 {
// 			sb.WriteString("; ")
// 		}
// 		sb.WriteString(fmt.Sprintf("%d: %s (%s)", i, s.Id, s.Content))
// 	}
// 	return sb.String()
// }
