package grammar

/* RuleCollection
	S -> A B C
	S -> G H
	gets Converted to {s} => [ [A, B, C] , [G, H] ]
*/

import(
	"strings"
	"regexp"
	"errors"
	
)
type RuleCollection struct{
	RuleCollectionVal [][]string
}

func (r *RuleCollection)AppendRule(rule[]string){
	r.RuleCollectionVal = append(r.RuleCollectionVal,rule)
}

type Grammar struct{
	rules map[string](*RuleCollection) // Each grammar rule will be in the form rules[$S] => [$S SUM $M, $S MUL $M]
}

func (g *Grammar)GetGrammarFromString(grammarString string)(error){ 
	//Make a new Map for the grammar
	g.rules =  make(map[string](*RuleCollection))
	
	//Clear unecessary whitespaces
	re := regexp.MustCompile("(\\s)+|(\\n|\\t)")
	grammarString = re.ReplaceAllString(strings.Trim(grammarString," "), " ")
	
	//Fill the map with grammar rules
	rules := strings.Split(grammarString,";")
	for _,rule := range rules{
		LHSRHSPair := strings.Split(rule,"->")
		LHSRHSPair[1] =strings.Trim(LHSRHSPair[1]," ")
		LHSRHSPair[0] =strings.Trim(LHSRHSPair[0]," ")
		if LHSRHSPair[0][0] != uint8('$'){
			return errors.New("Grammar incorrect near"+ LHSRHSPair[0])
		}
		if g.rules[LHSRHSPair[0]] != nil{
			g.rules[LHSRHSPair[0]].AppendRule(strings.Split(LHSRHSPair[1]," "))
		}else{
			g.rules[LHSRHSPair[0]] = &RuleCollection{[][]string{strings.Split(LHSRHSPair[1]," ")}}
		}
	}
	return nil
}

func (g Grammar)GetRulesForNonTerminal(NonTerminal string)(*RuleCollection){
	return g.rules[NonTerminal]
}