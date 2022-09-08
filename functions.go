package function

import (
	"context"
	"net/http"

	"cloud.google.com/go/bigquery"
	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
	"google.golang.org/api/option"
)

func init() {
	functions.HTTP("HelloWorld", helloWorld)
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	c, err := bigquery.NewClient(ctx, "", option.WithoutAuthentication(), option.WithEndpoint("http://localhost:9000"))
	if err != nil {
		panic(err)
	}

	// panic
	err = c.Dataset("my_dataset").Create(ctx, nil)
	if err != nil {
		panic(err)
	}

	// err = c.Dataset("my_dataset").Table("my_table").Create(ctx, &bigquery.TableMetadata{})
	// if err != nil {
	// 	panic(err)
	// }

	// ins := c.Dataset("my_dataset").Table("my_table").Inserter()
	// err = ins.Put(ctx, []struct{ Name string }{{Name: "Joselito"}, {Name: "Viveiros"}})
	// if err != nil {
	// 	panic(err)
	// }

	// query := c.Query("SELECT * FROM my_dataset.my_table")
	// job, err := query.Run(ctx)
	// if err != nil {
	// 	panic(err)
	// }

	// it, err := job.Read(ctx)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(it)
}
