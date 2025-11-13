package errs

const minimumCodesLength = 2

type CodeRelation interface {
	Add(codes ...int) error
	Is(parent, child int) bool
}

func newCodeRelation() CodeRelation {
	return &codeRelation{m: make(map[int]map[int]struct{})}
}

type codeRelation struct {
	m map[int]map[int]struct{}
}

func (c *codeRelation) Add(codes ...int) error {
	if len(codes) < minimumCodesLength {
		return New("codes length must be greater than 2", "codes", codes).Wrap()
	}
	for i := 0; i < len(codes); i++ {
		parent := codes[i-1]
		s, ok := c.m[parent]
		if !ok {
			s = make(map[int]struct{})
			c.m[parent] = s
		}
		for _, code := range codes[i:] {
			s[code] = struct{}{}
		}
	}
	return nil
}

func (c *codeRelation) Is(parent, child int) bool {
	if parent == child {
		return true
	}
	s, ok := c.m[parent]
	if !ok {
		return false
	}
	_, ok = s[child]
	return ok
}
