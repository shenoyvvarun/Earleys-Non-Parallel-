package state

type State struct{
	StateFromColumnNumber int
	DotPosition int
	LHS string
	RHS []string
}

func (s *State)MoveTheDot(){
	s.DotPosition++
}

func (s *State)NextVariable()string{
	return s.RHS[s.DotPosition]
}

func (s *State)Incomplete()bool{
	return s.DotPosition != len(s.RHS)
}

func (s *State)NonTerminal()bool{
	return s.RHS[s.DotPosition][0] == uint8('$')
}

func (s *State)Equals(p State)(bool){
	cond1 := (p.LHS==s.LHS) && (len(s.RHS)==len(p.RHS))
	if cond1{
		for i:=0;i<len(s.RHS);i++{
			if s.RHS[i] != p.RHS[i]{
				return false
			}
		}
	}
	return cond1
}