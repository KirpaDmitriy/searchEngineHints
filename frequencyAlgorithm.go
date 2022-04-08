package main
import (
    "fmt"
    "os"
    "unicode"
    "strings"
    "bufio"
    "sort"
)

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func countCombinations(n_combinations int, sentences [][]string) map[string]map[string]int{
    frequencyCounter := make(map[string]map[string]int)

    for _, sentence := range sentences[:]{
        for pos, _ := range sentence[:]{
            if pos < len(sentence) - n_combinations{
                wordsCombo := ""
                for keyWordOffset := 0; keyWordOffset < n_combinations; keyWordOffset += 1{
                    wordsCombo += sentence[pos + keyWordOffset]
                }
                if frequencyCounter[wordsCombo] == nil{
                    frequencyCounter[wordsCombo] = map[string]int{sentence[pos + n_combinations]: 1}
                } else{
                    frequencyCounter[wordsCombo][sentence[pos + n_combinations]] += 1
                }
            }
        }
    }
    return frequencyCounter
}

func extractWordsFromQuery(n_combinations int, query string) string{
    queryWords := []string{""}
    wordsCounter := 0

    for _, c := range query[:]{
        if unicode.IsLetter(c){
            queryWords[wordsCounter] += strings.ToLower(string(c))
        } else {
            queryWords = append(queryWords, "")
            wordsCounter += 1
        }
    }

    var notEmptyQueryWords []string
    var lastQueryWords string
    for i := 0; i < len(queryWords); i += 1{
        if len(queryWords[i]) > 0{
            notEmptyQueryWords = append(notEmptyQueryWords, queryWords[i])
        }
    }

    for i := max(0, len(notEmptyQueryWords) - n_combinations); i < len(notEmptyQueryWords); i += 1{
        lastQueryWords += notEmptyQueryWords[i]
    }
    return lastQueryWords
}

func main() {
    file, err := os.ReadFile("Пугачев.txt")
    if err != nil{
        fmt.Println(err)
        os.Exit(1)
    }

    data := string(file[:])

    sentences := [][]string{{""}}
    sentencesCounter := 0

    for _, c := range data[:]{
        if unicode.IsLetter(c){
            currentWord := len(sentences[sentencesCounter]) - 1
            sentences[sentencesCounter][currentWord] += strings.ToLower(string(c))
        } else{
            if (c == '.' || c == '?' || c == '!' || c == ';'){
                if len(sentences[sentencesCounter]) > 0{
                    tmp := []string{""}
                    sentences = append(sentences, tmp)
                    sentencesCounter += 1
                }
            } else {
                currentSentenceLength := len(sentences[sentencesCounter])
                if currentSentenceLength > 0{
                    if len(sentences[sentencesCounter][currentSentenceLength - 1]) != 0{
                        sentences[sentencesCounter] = append(sentences[sentencesCounter], "")
                    }
                }
            }
        }
    }

    n_combinations := 1
    frequencyCounter := countCombinations(n_combinations, sentences)

    in := bufio.NewReader(os.Stdin)
    query, err := in.ReadString('\n')

    lastQueryWords := extractWordsFromQuery(n_combinations, query)

    preAnswer := frequencyCounter[lastQueryWords]

    var keys []string
    for key := range preAnswer{
        keys = append(keys, key)
    }

    sort.SliceStable(keys, func(l, r int) bool{
        return preAnswer[keys[l]] > preAnswer[keys[r]]
    })

    var answer []string

    for n, key := range keys{
        if n > 4{
            break
        }
        answer = append(answer, key)
    }

    fmt.Printf("%v\n", answer)
}