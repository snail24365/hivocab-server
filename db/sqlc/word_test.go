package db

import (
	"context"
	"testing"

	"github.com/snail24365/hivocab-server/util"
	"github.com/stretchr/testify/require"
)
func TestCreateWord(t *testing.T) {
 	testWord := util.RandomString(3)
	resultWord, err := testQueries.CreateWord(context.Background(), testWord)
	
	require.NoError(t, err)
 	require.Equal(t, testWord, resultWord.Spelling)
	
}

func TestGetWordBySpelling(t *testing.T) {
	testWord := util.RandomString(3)
	targetWord, err := testQueries.CreateWord(context.Background(), testWord)
	require.NoError(t, err)

	finalResult, err := testQueries.GetWordBySpelling(context.Background(), targetWord.Spelling)
	require.NoError(t, err)
	require.Equal(t, targetWord.Spelling, finalResult.Spelling)
}

func TestListWordByPage(t *testing.T) {
	/* doesn't work because of test dependency.
	var testWords []string
	for i := 0; i < 4; i++ {
		testWord := util.RandomString(3)
		testWords = append(testWords, testWord)
		testQueries.CreateWord(context.Background(), testWord)
	}
	limit := 2
	offset := 1
	result, err := testQueries.ListWordByPage(context.Background(), ListWordByPageParams{int32(limit),int32(offset)})
	require.Equal(t, testWords[1], result[0].Spelling)
	require.Equal(t, testWords[2], result[1].Spelling)
	require.Equal(t, len(testWords), limit)
	require.NoError(t, err)
	*/
}