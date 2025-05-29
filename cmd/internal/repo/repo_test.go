package repo

import (
	"github.com/stretchr/testify/require"
	"muzz_challenge/cmd/internal/types"
	"muzz_challenge/pkg/uid"
	"os"
	"testing"
)

var (
	repo *Repo
	//sqlDb *sqlx.DB
)

func TestMain(m *testing.M) {
	// Setup code if needed

	////////////////////////////////
	/// Start MySQL container //////
	////////////////////////////////

	//_db, cont, err := testcontainerx.StartMySqlContainer()
	//if err != nil {
	//	fmt.Printf("start mysql container: %s\n", err)
	//	os.Exit(1)
	//}
	//sqlDb = _db

	////////////////////////////////
	// Start DynamoDB container ////
	////////////////////////////////

	//dynamoCli := testcontainerx.StartDynamoDBContainer()
	//err := testcontainerx.CreateTable(dynamoCli)
	//if err != nil {
	//	fmt.Printf("create table: %s\n", err)
	//	os.Exit(1)
	//}

	repo = New()

	code := m.Run()
	os.Exit(code)

	// Teardown code if needed
	//testcontainerx.StopMySqlContainer(context.Background(), sqlDb, cont)
}

var user1 = &types.User{
	Id: uid.MakeAccount(),
}

var users = map[int]*types.User{
	1: {
		Id: uid.MakeAccount(),
	},
	2: {
		Id: uid.MakeAccount(),
	},
	3: {
		Id: uid.MakeAccount(),
	},
	4: {
		Id: uid.MakeAccount(),
	},
	5: {
		Id: uid.MakeAccount(),
	},
}

func TestRepo(t *testing.T) {
	t.Run("PutDecision", func(t *testing.T) {
		for _, user := range users {
			err := repo.PutDecision(user.Id.String(), &types.Decision{
				// TODO - Implement a sensible decision structure here
				// Each user should like user1
			})
			require.NoError(t, err)

			// TODO - Add any other assertions you feel are required to prove this works as expected
		}
	})

	t.Run("CountLikedYou", func(t *testing.T) {
		count, err := repo.CountLikedYou(user1.Id.String())
		require.NoError(t, err)
		require.Equal(t, len(users), count)
	})

	t.Run("ListLikedYou", func(t *testing.T) {
		likes, token, err := repo.ListLikedYou(user1.Id.String(), 6, nil)
		require.NoError(t, err)
		require.Nil(t, token)
		require.Equal(t, len(users), len(likes))

		// TODO - Add any further sensible assertions here

		likes, token, err = repo.ListLikedYou(user1.Id.String(), 3, nil)
		require.NoError(t, err)
		// Next page token
		require.NotNil(t, token)
		require.Equal(t, 3, len(likes))

		likes, token, err = repo.ListLikedYou(user1.Id.String(), 3, token)
		require.NoError(t, err)
		// No next page
		require.Nil(t, token)
		// Just final like returned
		require.Len(t, likes, 2)
	})

	t.Run("PutDecisionFromUser1", func(t *testing.T) {
		err := repo.PutDecision(user1.Id.String(), &types.Decision{
			// TODO - Like User 1
		})

		require.NoError(t, err)

		err = repo.PutDecision(user1.Id.String(), &types.Decision{
			// TODO - Like User 2
		})
		require.NoError(t, err)
	})

	t.Run("ListNewLikedYou", func(t *testing.T) {
		likes, token, err := repo.ListLikedYou(user1.Id.String(), 5, nil)
		require.NoError(t, err)
		require.Nil(t, token)
		require.Equal(t, 3, len(likes))
		for _, like := range likes {
			require.NotNil(t, like)

			// TODO - Assert Decisions from user 1 and user 2 have not been returned
		}

		// TODO - Add any further sensible assertions here

		likes, token, err = repo.ListLikedYou(user1.Id.String(), 2, nil)
		require.NoError(t, err)
		// Next page token
		require.NotNil(t, token)
		require.Equal(t, 2, len(likes))

		likes, token, err = repo.ListLikedYou(user1.Id.String(), 2, token)
		require.NoError(t, err)
		// No next page
		require.Nil(t, token)
		// Just final like returned
		require.Len(t, likes, 1)
	})
}
