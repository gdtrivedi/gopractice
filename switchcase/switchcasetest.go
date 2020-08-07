package switchcase

type SwitchCaseIn struct {
	Hostnames []string
	Install   *string
}

func SwitchCaseTest(in SwitchCaseIn) {
	switch len(in.Hostnames) {
	case 0:

	}
}
