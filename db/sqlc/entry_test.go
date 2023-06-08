package db

import (
	"context"
	"github.com/stretchr/testify/require"
	"simpleBank/db/util"
	"testing"
)

func createRandomEntry(t *testing.T, account1 Account) Entry {
	randomAmount := util.RandomMoney()
	arg := CreateEntryParams{
		AccountID: account1.ID,
		Amount:    randomAmount,
	}

	entry, err := testQueries.CreateEntry(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, entry)

	require.Equal(t, account1.ID, entry.AccountID)
	require.Equal(t, randomAmount, entry.Amount)

	return entry
}

func TestCreateEntry(t *testing.T) {
	account1 := createRandomAccount(t)
	createRandomEntry(t, account1)
}

func TestGetEntry(t *testing.T) {
	account1 := createRandomAccount(t)
	entry1 := createRandomEntry(t, account1)

	entry2, err := testQueries.GetEntry(context.Background(), entry1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, entry2)

	require.Equal(t, entry1.ID, entry2.ID)
	require.Equal(t, entry1.Amount, entry2.Amount)
}

func TestListEntries(t *testing.T) {
	account1 := createRandomAccount(t)
	for i := 0; i < 10; i++ {
		createRandomEntry(t, account1)
	}

	arg := ListEntriesParams{
		AccountID: account1.ID,
		Limit:     5,
		Offset:    5,
	}

	entries, err := testQueries.ListEntries(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, entries, 5)

	for _, entry := range entries {
		require.NotEmpty(t, entry)
	}

}
