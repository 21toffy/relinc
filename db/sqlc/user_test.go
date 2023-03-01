package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/21toffy/relinc/util"

	"github.com/stretchr/testify/require"
)

func CreateRandomUser(t *testing.T) User {
	password := util.SecondRandomString(10)
	hashedPassword, err := util.HashPassword(password)
	require.NoError(t, err)

	arg := CreateUserParams{
		FirstName:      util.RandomOwner(),
		LastName:       util.RandomOwner(),
		EmailAddress:   util.RandomOwner() + "@gmail.com",
		PhoneNumber:    util.RandomPhone(),
		Username:       util.RandomOwner(),
		Dob:            GetTime(),
		Password:       hashedPassword,
		Address:        "1 bulabai way",
		ProfilePicture: "https:www.tofunmiprofile.com.png",
		Gender:         "male",
	}
	user, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)
	require.Equal(t, arg.FirstName, user.FirstName)
	require.Equal(t, arg.LastName, user.LastName)
	require.Equal(t, arg.EmailAddress, user.EmailAddress)
	require.Equal(t, arg.PhoneNumber, user.PhoneNumber)
	require.Equal(t, arg.Username, user.Username)
	// require.Equal(t, arg.Dob, user.Dob)
	require.Equal(t, arg.Address, user.Address)
	require.Equal(t, arg.ProfilePicture, user.ProfilePicture)
	require.Equal(t, arg.Gender, user.Gender)
	require.NotZero(t, user.Dob)
	require.NotZero(t, user.ID)
	require.NotZero(t, user.CreatedAt)
	return user
}

func TestCreateUser(t *testing.T) {
	CreateRandomUser(t)
}

func TestGetAccount(t *testing.T) {
	user1 := CreateRandomUser(t)
	user2, err := testQueries.GetUser(context.Background(), user1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, user2)
	require.Equal(t, user1.ID, user2.ID)
	require.Equal(t, user1.FirstName, user2.FirstName)
	require.Equal(t, user1.LastName, user2.LastName)
	require.Equal(t, user1.EmailAddress, user2.EmailAddress)
	require.Equal(t, user1.PhoneNumber, user2.PhoneNumber)
	require.Equal(t, user1.Username, user2.Username)
	require.Equal(t, user1.Address, user2.Address)
	require.Equal(t, user1.ProfilePicture, user2.ProfilePicture)
	require.Equal(t, user1.Gender, user2.Gender)
	require.Equal(t, user1.Dob, user2.Dob)
	require.Equal(t, user1.CreatedAt, user2.CreatedAt)
	require.WithinDuration(t, user1.CreatedAt, user2.CreatedAt, time.Second)
}

func TestUserUpdate(t *testing.T) {
	user := CreateRandomUser(t)

	arg := UpdateUserParams{
		ID:           user.ID,
		FirstName:    util.RandomOwner(),
		LastName:     util.RandomOwner(),
		EmailAddress: util.RandomOwner() + "@relinc.com",
		PhoneNumber:  util.RandomPhone(),
		Username:     util.RandomOwner(),
		Dob:          GetTime(),
		Address:      "1 bulabai way",
	}
	user2, err := testQueries.UpdateUser(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, user2)
	require.Equal(t, arg.ID, user.ID)
	require.Equal(t, arg.FirstName, user2.FirstName)
	require.Equal(t, arg.LastName, user2.LastName)
	require.Equal(t, arg.EmailAddress, user2.EmailAddress)
	require.Equal(t, arg.PhoneNumber, user2.PhoneNumber)
	require.Equal(t, arg.Username, user2.Username)
	require.Equal(t, arg.Address, user2.Address)
	require.Equal(t, user.Dob, user2.Dob)
}

func TestUserDelete(t *testing.T) {
	user := CreateRandomUser(t)
	err := testQueries.DeleteUser(context.Background(), user.ID)
	require.NoError(t, err)
	user2, err := testQueries.GetUser(context.Background(), user.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, user2)
}

func TestUserList(t *testing.T) {
	for i := 0; i < 10; i++ {
		CreateRandomUser(t)
	}
	arg := ListUsersParams{
		Offset: 5,
		Limit:  5,
	}

	userc, err := testQueries.ListUsers(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, userc, 5)
	for _, user := range userc {
		require.NotEmpty(t, user)
	}
}
