package main

import (
	"fmt"
	"sort"
	"strings"
)

type Pair struct {
	Key   string
	Value int
}

type PairList []Pair

func (p PairList) Len() int {
	return len(p)
}

func (p PairList) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p PairList) Less(i, j int) bool {
	return p[i].Value < p[j].Value
}

func main() {
	input := "Mother,' he insisted. 'What's going on? Are we moving?' 'Come downstairs with me,' said Mother, leading the way towards the large dining room where the Fury had been to dinner the week before. 'We'll talk down there.' (page 3) Bruno ran downstairs and even passed her out on the staircase so that he was waiting in the dining room when she arrived. He looked at her without saying anything for a moment and thought to himself that she couldn't have applied her make-up correctly that morning because the rims of her eyes were more red than usual, like his own after he'd been causing chaos and got into trouble and ended up crying. 'Now, you don't have to worry, Bruno,' said Mother, sitting down in the chair where the beautiful blonde woman who had come to dinner with the Fury had sat and waved at him when Father closed the doors. 'In fact if anything it's going to be a great adventure.' 'What is?' he asked. 'Am I being sent away?' 'No, not just you,' she said, looking as if she might smile for a moment but thinking better of it. 'We all are. Your father and I, Gretel and you. All four of us.' Bruno thought about this and frowned. He wasn't particularly bothered if Gretel was being sent away because she was a Hopeless Case and caused nothing but trouble for him. But it seemed a little unfair that they all had to go with her. 'But where?' he asked. 'Where are we going exactly? Why can't we stay here?' 'Your father's job,' explained Mother. 'You know how important it is, don't you?' (page 4) 'Yes, of course,' said Bruno, nodding his head, because there were always so many visitors to the house - men in fantastic uniforms, women with typewriters that he had to keep his mucky hands off - and they were always very polite to Father and told each other that he was a man to watch and that the Fury had big things in mind for him. 'Well, sometimes when someone is very important,' continued Mother, 'The man who employs him asks him to go somewhere else because there's a very special job that needs doing there.' 'What kind of job?' asked Bruno, because if he was honest with"
	CheckWordUsage(input)
}

func CheckWordUsage(input string) {

	//splitting the input and
	formattedInput := strings.Split(input, " ")
	result := make(map[string]int)

	for _, v := range formattedInput {
		count := 1
		if result[string(v)] > 0 {
			count = result[string(v)] + 1
		}
		result[string(v)] = count
	}

	p := make(PairList, len(result))
	i := 0

	/*populating data in the PairList type variable
	and incrementing the value of array index */
	for k, v := range result {
		p[i] = Pair{k, v}
		i++
	}

	//Sorting data in ascending order
	sort.Sort(p)

	fmt.Printf("Top 10 most used words : \n")

	l := len(p)
	var limit int

	//limiting loop constraint to run max 10 times
	if l > 10 {
		limit = len(p) - 10
	}

	//printing the required result. Looping in reverse order to print words with max no. of usage
	for i := l - 1; i >= limit; i-- {
		fmt.Println(p[i].Key, p[i].Value)
	}
}
