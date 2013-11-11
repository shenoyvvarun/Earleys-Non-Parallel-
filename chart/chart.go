package chart

import(
	"./state"
)

type StateCollection struct{
	cols []state.State
}

func (sc StateCollection)At(pos int)(state.State){
	return sc.cols[pos]
}

func (sc *StateCollection)AddStateToColumn(s state.State){
	for _,val := range sc.cols{
		if s.Equals(val){
			return
		}
	}
	sc.cols = append(sc.cols,s)
}

func (c *StateCollection)Len()(int){
	return len(c.cols)
}

type Chart struct{
	rows []StateCollection
}

func (c *Chart)AddStateToColumn(s state.State, column int){
	if column >= len(c.rows){
		//New Column
		c.rows = append(c.rows,StateCollection{[]state.State{s}})
	}else{
		c.rows[column].AddStateToColumn(s)
	}
}

func (c Chart)At(column int)(StateCollection){
	return c.rows[column]
}

func (c Chart)LenAt(column int)(int){
	return c.rows[column].Len()
}

func (c Chart)Len()int{
	return len(c.rows)
}