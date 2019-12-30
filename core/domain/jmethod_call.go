package domain

type JMethodCall struct {
	Package           string
	Type              string
	Class             string
	MethodName        string
	Parameters        []string
	StartLine         int
	StartLinePosition int
	StopLine          int
	StopLinePosition  int
}

func NewJMethodCall() JMethodCall {
	return *&JMethodCall{
		Package:           "",
		Type:              "",
		Class:             "",
		MethodName:        "",
		Parameters:        nil,
		StartLine:         0,
		StartLinePosition: 0,
		StopLine:          0,
		StopLinePosition:  0,
	}
}

func (c *JMethodCall) BuilFullMethodName() string {
	return c.Package + "." + c.Class + "." + c.MethodName
}