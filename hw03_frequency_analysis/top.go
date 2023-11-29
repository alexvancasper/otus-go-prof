package hw03frequencyanalysis

import (
	"sort"
	"strings"
)

type freq struct {
	word     string
	frequent int
}

type FreqList []freq

func (f FreqList) Len() int { return len(f) }
func (f FreqList) Less(i, j int) bool {
	return f[i].frequent > f[j].frequent ||
		f[i].frequent == f[j].frequent && f[i].word < f[j].word
}
func (f FreqList) Swap(i, j int) { f[i], f[j] = f[j], f[i] }

func Top10(text string) []string {
	arrayWords := textSplitter(text)
	freqWord := freqCounter(arrayWords)
	fl := make(FreqList, len(freqWord))
	i := 0
	for key, value := range freqWord {
		fl[i] = freq{key, value}
		i++
	}
	sort.Sort(fl)
	result := make([]string, 0, 10)
	for idx, word := range fl {
		if idx == 10 {
			break
		}
		result = append(result, word.word)
	}

	return result
}

func textSplitter(text string) []string {
	textSplit := strings.Split(text, " ")
	result := make([]string, 0, len(textSplit))
	for _, word := range textSplit {
		wordSplit := strings.Split(word, "\n")
		for _, w := range wordSplit {
			ws := strings.Split(w, "\t")
			for _, v := range ws {
				if v == "" {
					continue
				}
				result = append(result, v)
			}
		}
	}
	return result
}

func freqCounter(wordArray []string) map[string]int {
	freqWord := make(map[string]int)
	for _, word := range wordArray {
		if val, ok := freqWord[word]; !ok {
			freqWord[word] = 1
		} else {
			val++
			freqWord[word] = val
		}
	}
	return freqWord
}
