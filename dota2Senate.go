package main

import "fmt"

func main() {
	fmt.Println(predictPartyVictory("RD"))
}

func predictPartyVictory(senate string) string {
	var seqCurrent byte
	seq := 0
	senateByte := make([]byte, len(senate))
	for i := 0; i < len(senate); i++ {
		senateByte[i] = senate[i]
	}
	return predictCalc(&senateByte, &seqCurrent, &seq)
}

func predictCalc(senate *[]byte, seqCurrent *byte, seq *int) string {
	countR := 0
	countD := 0
	for i := 0; i < len(*senate); i++ {
		switch (*senate)[i] {
		case 'R':
			countR++
		case 'D':
			countD++
		}
		if *seq == 0 {
			*seqCurrent = (*senate)[i]
			*seq++
		} else if (*senate)[i] == *seqCurrent {
			*seq++
		} else {
			*seq--
			*senate = append((*senate)[:i], (*senate)[i+1:]...)
			i--
		}
	}
	if countD == 0 {
		return "Radiant"
	}
	if countR == 0 {
		return "Dire"
	}
	return predictCalc(senate, seqCurrent, seq)
}
