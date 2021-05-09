package Service

import (
	"context"
	"errors"
	"fmt"
	"moonservice/Interface"
	"moonservice/Model"
	"time"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/db"

	"google.golang.org/api/option"
)

type FirebaseService struct {
	Client *db.Client
}

func SetupFirebaseService(credentialFile string, projectID string, dbUrl string) (Interface.IRTDB, error) {
	opt := option.WithCredentialsFile(credentialFile)
	config := &firebase.Config{
		ProjectID:   projectID,
		DatabaseURL: dbUrl,
	}
	app, err := firebase.NewApp(context.Background(), config, opt)
	if err != nil {
		errText := fmt.Sprintf("error initializing app: %v", err)
		return nil, errors.New(errText)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	if cancel != nil {
		defer cancel()
	}

	cli, err := app.Database(ctx)
	if err != nil {
		errText := fmt.Sprintf("Error initializing database client: %v", err)
		return nil, errors.New(errText)
	}

	return &FirebaseService{
		Client: cli,
	}, nil
}

func (fs *FirebaseService) GetMoonCoinFromRTDB() (mcModel Model.RTDBMoonCoinModel, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	if cancel != nil {
		defer cancel()
	}
	ref := fs.Client.NewRef("/MoonCoin")

	err = ref.Get(ctx, &mcModel)
	if err != nil {
		return mcModel, err
	}

	return mcModel, nil
}

func (fs *FirebaseService) DecreaseMoonCoinToRTDB(moonCoin float64, exchangeRate float64) (mcModel Model.RTDBMoonCoinModel, err error) {

	mooncoinFromDB, err := fs.GetMoonCoinFromRTDB()
	if err != nil {
		return mcModel, err
	}
	if mooncoinFromDB.Remaining == 0 {
		err = errors.New("MOON is not enough to buy, 0 MOON remaining")
		return mcModel, err
	}
	if mooncoinFromDB.Remaining < moonCoin {
		errMsg := fmt.Sprintf("MOON is not enough to buy, %.15f MOON remaining", mooncoinFromDB.Remaining)
		err = errors.New(errMsg)
		return mcModel, err
	}
	mooncoinFromDB.Remaining = mooncoinFromDB.Remaining - moonCoin
	mooncoinFromDB.ExchangeRate = (1 / exchangeRate)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	if cancel != nil {
		defer cancel()
	}
	ref := fs.Client.NewRef("/MoonCoin")
	err = ref.Set(ctx, &mooncoinFromDB)
	if err != nil {
		return mcModel, err
	}

	return mcModel, nil
}
