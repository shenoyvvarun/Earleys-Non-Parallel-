package reducer

import(
	"../../chart/state"
	"../../chart"
)

func Reduce(parseChart *chart.Chart,s state.State,i int){
	collection :=  parseChart.At(s.StateFromColumnNumber)
	for k:=0;k<collection.Len();k++{
		val := collection.At(k)
		if val.Incomplete() && val.NextVariable() == s.LHS{
			val.MoveTheDot()
			parseChart.AddStateToColumn(val,i)
		}
	}
}