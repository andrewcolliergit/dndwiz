package class

import "strings"

type Equipment []string

func (e Equipment) String() string {
	letter := []string{"A", "B", "C", "D"}
	var sb strings.Builder

	for i, item := range e {
		if i > 0 {
			sb.WriteString(", ")
		}
		if i < len(letter) {
			sb.WriteString("(")
			sb.WriteString(letter[i])
			sb.WriteString(") ")
		}
		sb.WriteString(item)
	}

	return sb.String()
}
