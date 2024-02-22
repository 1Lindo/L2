/*
Программа принимает массив слов на русском языке в кодировке UTF-8, находит все множества анаграмм и выводит их
в формате ключ-значение, где ключ - первое встретившееся слово из множества, а значение - массив слов из этого множества.
Слова приводятся к нижнему регистру, множества из одного элемента не попадают в результат,
и каждое слово встречается только один раз.

*/

package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	dictionary := []string{"пятак", "пятка", "листок", "тяпка", "слиток", "столик"}
	anagramSets := findAnagramSets(dictionary)

	for key, value := range anagramSets {
		fmt.Printf("%s: %v\n", key, value)
	}
}

func findAnagramSets(words []string) map[string][]string {

	/*Создаются два словаря:
	anagramMap: словарь, где ключ - отсортированное слово, значение - слайс слов с теми же буквами.
	anagramSets: словарь, где ключ - первое слово из группы анаграмм, значение - группа анаграмм.*/
	anagramSets := make(map[string][]string)
	anagramMap := make(map[string][]string)

	/*Итерируемся по слайсу слов и каждое преобразуем каждое в отсортированную строку, а затем добавляем в anagramMap.
	Если в anagramMap уже есть другие слова с такой же сортировкой, они добавляются в группу анаграмм.*/
	for _, word := range words {
		sortedWord := sortString(word)
		anagramMap[sortedWord] = append(anagramMap[sortedWord], word)
	}

	//Итерируется по anagramMap и формируется anagramSets, оставляя только группы анаграмм с более чем одним словом.
	for _, anagrams := range anagramMap {
		if len(anagrams) > 1 {
			firstWord := anagrams[0]
			sort.Strings(anagrams)
			anagramSets[firstWord] = anagrams
		}
	}

	return anagramSets
}

/*
Тут слово преобразуется в сортированную по алфавиту строку.
Это необходимо для того, чтобы сравнивать слова с одинаковыми буквами, но в разном порядке.
*/
func sortString(s string) string {
	sChars := strings.Split(s, "")
	sort.Strings(sChars)
	return strings.Join(sChars, "")
}
