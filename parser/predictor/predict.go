package predictor

import(
	"../../chart/state"
	"../../chart"
	"../../grammar"
)

func Predict(c *chart.Chart,s state.State,pos int,g *grammar.Grammar){
	collection :=  g.GetRulesForNonTerminal(s.NextVariable()).RuleCollectionVal
	for _,val:= range collection{
		c.AddStateToColumn(state.State{pos,0,s.NextVariable(),val},pos)
	}
}