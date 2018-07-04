package summary

const StyleDefault = "default"
const StyleAws = "aws"

const OutputCsv = "csv"

var seperator = " | "

type Summary interface {
	AddEntry(entry string)
	Print()
}

func New(pattern string, options ...interface{}) Summary {
	var s Summary
	style := StyleDefault

	if len(options) == 1 {
		o, _ := options[0].(Options)
		if o.Style != "" {
			style = o.Style
		}

		if o.Output == OutputCsv {
			seperator = ";"
		}
	}

	switch style {
	case StyleAws:
		s = newAws()
		break
	default:
		s = newDefault(pattern)
		break
	}

	return s
}
