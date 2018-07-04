package summary

const StyleDefault = "default"
const StyleAws = "aws"

type Summary interface {
	AddEntry(entry string)
	Print()
}

func New(pattern string, style ...interface{}) Summary {
	var s Summary
	t := StyleDefault

	if len(style) == 1 {
		t, _ = style[0].(string)
	}

	switch t {
	case StyleAws:
		s = newAws()
		break
	default:
		s = newDefault(pattern)
		break
	}

	return s
}
