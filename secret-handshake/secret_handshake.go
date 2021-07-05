package secret

var events = []string{"wink", "double blink", "close your eyes", "jump"}

func Handshake(code uint) (hs []string) {
	for i := 0; i < 4; i++ {
		if code&(1<<i) > 0 {
			hs = append(hs, events[i])
		}
	}

	if code&16 > 0 {
		for i := 0; i < len(hs)/2; i++ {
			hs[i], hs[len(hs)-i-1] = hs[len(hs)-i-1], hs[i]
		}
	}

	return
}
