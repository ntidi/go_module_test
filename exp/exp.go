package exp

var jobs []map[string]string

func init() {
	jobs = []map[string]string{
		{
			"from":                "2010",
			"to":                  "2012",
			"company":             "ABC",
			"position":            "Developer",
			"programing language": "Ruby",
		},
		{
			"from":                "2012",
			"to":                  "now",
			"company":             "XYZ",
			"position":            "Architecture",
			"programing language": "Go, Java, Prolog, Smalltalk",
		},
	}
}

// Jobs return working experiences
func Jobs() []map[string]string {
	return jobs
}
