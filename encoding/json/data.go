package json

func mapPayload() map[string]string {
	return map[string]string{
		"Bedfordshire":       "England",
		"Buckinghamshire":    "England",
		"Cambridgeshire":     "England",
		"Cheshire":           "England",
		"Cleveland":          "England",
		"Cornwall":           "England",
		"Cumbria":            "England",
		"Derbyshire":         "England",
		"Devon":              "England",
		"Dorset":             "England",
		"Durham":             "England",
		"East Sussex":        "England",
		"Essex":              "England",
		"Gloucestershire":    "England",
		"Greater London":     "England",
		"Greater Manchester": "England",
		"Hampshire":          "England",
		"Hertfordshire":      "England",
		"Kent":               "England",
		"Lancashire":         "England",
		"Leicestershire":     "England",
		"Lincolnshire":       "England",
		"Merseyside":         "England",
		"Norfolk":            "England",
		"North Yorkshire":    "England",
		"Northamptonshire":   "England",
		"Northumberland":     "England",
		"Nottinghamshire":    "England",
		"Oxfordshire":        "England",
		"Shropshire":         "England",
		"Somerset":           "England",
		"South Yorkshire":    "England",
		"Staffordshire":      "England",
		"Suffolk":            "England",
		"Surrey":             "England",
		"Tyne and Wear":      "England",
		"Warwickshire":       "England",
		"West Berkshire":     "England",
		"West Midlands":      "England",
		"West Sussex":        "England",
		"West Yorkshire":     "England",
		"Wiltshire":          "England",
		"Worcestershire":     "England",
	}
}

func mapTextMarshalPayload() map[*textStruct]string {
	// TODO: more
	return map[*textStruct]string{
		{[]byte("Bedfordshire")}:    "England",
		{[]byte("Buckinghamshire")}: "England",
		{[]byte("Cambridgeshire")}:  "England",
		{[]byte("Cheshire")}:        "England",
		{[]byte("Cleveland")}:       "England",
		{[]byte("Cornwall")}:        "England",
	}
}
