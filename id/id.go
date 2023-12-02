package id

type Generate func() string

func dummyGenerator() string {
	return "dummy"
}

func Generator() Generate {
	return dummyGenerator
}
