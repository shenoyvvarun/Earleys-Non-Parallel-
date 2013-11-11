package main

import(
	"fmt"
	"./grammar"
	"./parser"
	"strings"
)

func main(){
		g:= `$ROOT->$S;$VP->$V $NP;$NP->papa;$V->ate;
		$S->$NP $VP;$VP->$VP $PP;$N->caviar;$P->with;
		$NP->$Det $N;$PP->$P $NP;$N->spoon;$Det->the;
		$NP->$NP $PP;$Det->a`
		var gram grammar.Grammar
		gram.GetGrammarFromString(g)
		chart,_ := parser.Parse(&gram,strings.Split("with papa ate the caviar spoon with a spoon"," "))
		fmt.Println(chart)
}