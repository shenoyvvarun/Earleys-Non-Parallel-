package scanner

import(
	"../../chart/state"
	"../../chart"
)

func Scan(parseChart *chart.Chart,s state.State,word string, pos int){
	if word == s.NextVariable(){
		parseChart.AddStateToColumn(state.State{pos,s.DotPosition+1,s.LHS,s.RHS[:]},pos+1)
	}
}