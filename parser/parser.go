package parser

import(
	"../chart"
	"../grammar"
	"../chart/state"
	"fmt"
	"./predictor"
	"./scanner"
	"./reducer"
	"strings"
)

func Parse(g *grammar.Grammar,words []string)(parseChart chart.Chart,err error){
	//Add initial state
	parseChart.AddStateToColumn(state.State{0,0,"$ROOT",[]string{"$S"}},0)
	var i,word = 0,""

	//handle parseError
	defer func(){
		if r:= recover();r!=nil{
			fmt.Println("ParseError At : "+strings.Join(words[:i-1]," ")+" ^^"+words[i-1]+"^^ ")		
			}
	}();
	for i,word = range words{
		for k:=0;k<parseChart.LenAt(i);k++{
				state := parseChart.At(i).At(k)
				if state.Incomplete(){
					if state.NonTerminal(){
						predictor.Predict(&parseChart,state,i,g)
					}else{
						scanner.Scan(&parseChart,state,word,i)
					}
				}else{
					reducer.Reduce(&parseChart,state,i)
				}
		}
	}
	return
}