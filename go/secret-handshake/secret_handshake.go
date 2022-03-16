package secret

var events = map[uint]string{
	0x1: "wink",
	0x2: "double blink",
	0x4: "close your eyes",
	0x8: "jump",
}

func Handshake(code uint) []string {
	result := make([]string, 0)
	var eventIds []uint
	if code&uint(16) == 0x10 {
		// reverse
		eventIds = []uint{8, 4, 2, 1}
	} else {
		eventIds = []uint{1, 2, 4, 8}
	}

	for i := range eventIds {
		if event := code & eventIds[i]; event > 0x0 {
			result = append(result, events[event])
		}
	}
	return result
}
