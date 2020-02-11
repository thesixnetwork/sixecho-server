package main

import (
	"context"
	"fmt"
	"strings"

	"github.com/olivere/elastic"
)

const (
	accountAlias     string = "ssc_accounts"
	transactionAlias string = "ssc_transactions"
	imageAlias       string = "ssc_images"
	textAlias        string = "ssc_texts"
	hourlyUpload     string = "ssc_hourly_uploads"
	blockNumAlias    string = "ssc_blocknum"
	dialyUpload      string = "ssc_daily_uploads"
	errorAlias       string = "ssc_errors"
	oldVersion       string = "v1"
	nextVersion      string = "v2"
)

var (
	ctx        = context.Background()
	elasticURL = "https://search-es-six-zunsizmfamv7eawswgdvwmyd6u.ap-southeast-1.es.amazonaws.com"
	client, _  = elastic.NewClient(elastic.SetURL(elasticURL), elastic.SetSniff(false),
		elastic.SetHealthcheck(false),
		// elastic.SetHttpClient(signingClient),
		//elastic.SetErrorLog(log.New(os.Stderr, "", log.LstdFlags)),
		//elastic.SetInfoLog(log.New(os.Stdout, "", log.LstdFlags)))
	)
)

func removeAlias(client *elastic.Client, alias string) {
	if hourlyUpload == alias || dialyUpload == alias {
		_, err := client.Alias().Remove(fmt.Sprintf("%s_%s", strings.TrimSuffix(alias, "s"), oldVersion), alias).Do(ctx)
		if err != nil {
			fmt.Println(err.Error())
		}
	} else {
		_, err := client.Alias().Remove(fmt.Sprintf("%s_%s", alias, oldVersion), alias).Do(ctx)
		if err != nil {
			fmt.Println(err.Error())
		}
	}
}

func addAlias(client *elastic.Client, alias string) {
	if hourlyUpload == alias || dialyUpload == alias {
		_, err := client.Alias().Add(fmt.Sprintf("%s_%s", strings.TrimSuffix(alias, "s"), nextVersion), alias).Do(ctx)
		if err != nil {
			fmt.Println(err.Error())
		}

	} else {
		_, err := client.Alias().Add(fmt.Sprintf("%s_%s", alias, nextVersion), alias).Do(ctx)
		if err != nil {
			fmt.Println(err.Error())
		}
	}
}

func showAlias(client *elastic.Client, alias string) {
	result, err := client.CatAliases().Alias(alias).Do(ctx)
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("%#v\n", result)
}

func main() {
	allIndex := []string{
		accountAlias,
		transactionAlias,
		imageAlias,
		textAlias,
		hourlyUpload,
		blockNumAlias,
		dialyUpload,
		errorAlias,
	}
	for _, element := range allIndex {
		showAlias(client, element)
	}
	for _, element := range allIndex {
		removeAlias(client, element)
	}
	for _, element := range allIndex {
		addAlias(client, element)
	}
}
