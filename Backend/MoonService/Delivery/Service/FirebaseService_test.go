package Service

import (
	"fmt"
	"testing"

	"github.com/gofiber/fiber/v2/utils"
)

func Test_SetupFirebaseService_Success(t *testing.T) {

	creFile := "./mooncoinrtdb-firebase-adminsdk-yioeo-36db39fffa.json"
	projectID := "mooncoinrtdb"
	dbUrl := "https://mooncoinrtdb-default-rtdb.firebaseio.com"

	rtdbClient, err := SetupFirebaseService(creFile, projectID, dbUrl)

	utils.AssertEqual(t, nil, err)
	fmt.Println(rtdbClient)
}

func Test_SetupFirebaseService_InitAppErrorWithCreFile(t *testing.T) {

	creFile := "./mooncoinrtawdawdawffawdb-firebase-adminsdk-yioeo-36db39fffa.json"
	projectID := "mooncoinrtdb"
	dbUrl := "https://mooncoinrtdb-default-rtdb.firebaseio.com"

	rtdbClient, _ := SetupFirebaseService(creFile, projectID, dbUrl)

	utils.AssertEqual(t, nil, rtdbClient)

}

func Test_SetupFirebaseService_InitAppErrorWithDBURL(t *testing.T) {

	creFile := "./mooncoinrtawdawdawffawdb-firebase-adminsdk-yioeo-36db39fffa.json"
	projectID := "mooncoinrtdb"
	dbUrl := "https://mooncoinrtdb-default-rtdb.fireawdawdawdawfafwbaseio.com"

	rtdbClient, _ := SetupFirebaseService(creFile, projectID, dbUrl)

	utils.AssertEqual(t, nil, rtdbClient)

}

func Test_GetMoonCoinFromRTDB(t *testing.T) {
	creFile := "./mooncoinrtdb-firebase-adminsdk-yioeo-36db39fffa.json"
	projectID := "mooncoinrtdb"
	dbUrl := "https://mooncoinrtdb-default-rtdb.firebaseio.com"

	rtdbClient, err := SetupFirebaseService(creFile, projectID, dbUrl)

	utils.AssertEqual(t, nil, err)

	mcModel, err := rtdbClient.GetMoonCoinFromRTDB()

	utils.AssertEqual(t, nil, err)
	fmt.Println(mcModel)
}

func Test_DecreaseMoonCoinToRTDB(t *testing.T) {
	creFile := "./mooncoinrtdb-firebase-adminsdk-yioeo-36db39fffa.json"
	projectID := "mooncoinrtdb"
	dbUrl := "https://mooncoinrtdb-default-rtdb.firebaseio.com"

	rtdbClient, err := SetupFirebaseService(creFile, projectID, dbUrl)

	utils.AssertEqual(t, nil, err)

	moonCoin := 1.0
	exchangeRate := 50.0

	mcModel, err := rtdbClient.DecreaseMoonCoinToRTDB(moonCoin, exchangeRate)

	utils.AssertEqual(t, nil, err)
	fmt.Println(mcModel)
}
