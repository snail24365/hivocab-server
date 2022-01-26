package db

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
)

type DecodeWord struct {
	WordId int `json:"wordId"`
	Word string `json:"word"`
}

type DecodeUsecase struct {
	UsecaseId int `json:"useCaseId"`
	WordId int `json:"wordId"` 
	Description string `json:"usageDescription"`
}

type DecodeExample struct {
	ExampleId int `json:"exampleId"`
	UsecaseId int `json:"useCaseId"`
	Sentence string `json:"sentence"`
}

func InitializeDatabase(store Store) {
	if !isNeedInit(store) {
		return
	}
	
	store.DeleteAllExample(context.Background())
	store.DeleteAllUsecase(context.Background())
	store.DeleteAllWord(context.Background())
	insertVocab(store)
}

func isNeedInit(store Store) bool {
	const threshhold = 1000
	numWord , _ := store.CountAllWord(context.Background())
	numUsecase , _ := store.CountAllUsecase(context.Background())
	numExample , _ := store.CountAllExample(context.Background())
	return numWord < threshhold || numUsecase < threshhold || numExample < threshhold	
}

func insertVocab(store Store) {
	words, usecases, examples := parseVocab()

	for i := 0; i < len(words); i++ {
		word := Word {
			ID: int64(words[i].WordId),
			Spelling: words[i].Word,
		}
		_, err := store.InsertWord(context.Background(), InsertWordParams{
			ID: word.ID, Spelling: word.Spelling,
		})
		check(err)
	}
	
	// Usecase id를 랜덤하게 만들기 위하여 random lookUp 사용
	lookUp := make([]int, len(usecases))
  for i := 0; i < len(lookUp); i++ {
    lookUp[i] = i
  }
	rand.Shuffle(len(lookUp), func(i, j int) {
		lookUp[i], lookUp[j] = lookUp[j], lookUp[i]
	})

	for i := 0; i < len(usecases); i++ {
		usecase := Usecase{
			ID: int64(lookUp[usecases[i].UsecaseId]),
			WordID: int64(usecases[i].WordId),
			DescriptionSentence: usecases[i].Description,
		}
		_, err := store.InsertUsecase(context.Background(), InsertUsecaseParams{
			ID: usecase.ID,
			WordID: usecase.WordID,
			DescriptionSentence: usecase.DescriptionSentence,
		})
		check(err)
	}

	for i := 0; i < len(examples); i++ {
		example := Example{
			ID: int64(examples[i].ExampleId),
			UsecaseID: int64(lookUp[examples[i].UsecaseId]),
			Sentence: examples[i].Sentence,
		}
		_, err := store.InsertExample(context.Background(), InsertExampleParams{
			ID: example.ID,
			UsecaseID: example.UsecaseID,
			Sentence: example.Sentence,
		})
		check(err)
	}
}


func parseVocab() ([]DecodeWord, []DecodeUsecase, []DecodeExample){
	base := "resource/"
	wordUri := base + "word.json"
	usecaseUri := base + "useCase.json"
	exampleUri := base + "example.json"

	wordData, e1 := ioutil.ReadFile(wordUri)
	usecaseData, e2 := ioutil.ReadFile(usecaseUri)
	exampleData, e3 := ioutil.ReadFile(exampleUri)

	fmt.Println(e1,e2,e3)

	// enough size for all of the above files
	maxSize := 80000
	words := make([]DecodeWord, maxSize)
	usecases := make([]DecodeUsecase, maxSize)
	examples := make([]DecodeExample, maxSize)
	
	var err error
	err = json.Unmarshal([]byte(wordData), &words)
	check(err)

	err = json.Unmarshal([]byte(usecaseData), &usecases)
	check(err)

	err = json.Unmarshal([]byte(exampleData), &examples)
	check(err)
	
	return words, usecases, examples
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}