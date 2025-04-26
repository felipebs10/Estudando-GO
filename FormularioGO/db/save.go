package db

import "context"

func Insert(Collection string, data any) error {
	client, ctx := getConnection()
	defer client.Disconnect(ctx)

	c := client.Database("webform").Collection(Collection)

	_, err := c.InsertOne(context.Background(), data)

	return err
}
