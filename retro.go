package retro

import (
	"slices"
)

func getSticky(ss []*Sticky, id string) *Sticky {
	for _, s := range ss {
		if s == nil {
			continue
		}
		if s.Id == id {
			return s
		}
	}
	return nil
}

func riseSticky(ss []*Sticky, id string) *Sticky {
	for i, s := range ss {
		if s == nil {
			continue
		}
		if s.Id == id {
			ss = append(slices.Delete(ss, i, i+1), s)
			return s
		}
	}
	return nil
}

func deleteSticky(ss []*Sticky, id string) {
	for i, s := range ss {
		if s == nil {
			continue
		}
		if s.Id == id {
			ss = slices.Delete(ss, i, i+1)
			ss = ss[:len(ss)-1]
			return
		}
	}
}

// func formatStickies(ss []*Sticky) string {
// 	var sb strings.Builder
// 	for i, s := range ss {
// 		if i != 0 {
// 			sb.WriteString("; ")
// 		}
// 		if s == nil {
// 			sb.WriteString(fmt.Sprintf("%d: nil", i))
// 			continue
// 		}
// 		sb.WriteString(fmt.Sprintf("%d: %s (%s)", i, s.Id, s.Content))
// 	}
// 	return sb.String()
// }
