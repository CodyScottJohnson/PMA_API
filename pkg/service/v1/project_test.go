package v1

import (
	"context"
	"fmt"

	"Product-Management-API/pkg/api/v1"
	"Product-Management-API/pkg/database"
	"testing"
)

func gsTest(t *testing.T) {

}
func Test_Create(t *testing.T) {
	ctx := context.Background()
	database, err := db.ConnectDB("")
	if err != nil {
		fmt.Println(err)
	}
	ProjectService := NewProjectServiceServer(database)
	pr := v1.Project{
		Company:     "Microsoft",
		Background:  "Blue",
		Image:       "test.png",
		Title:       "What We Did",
		Description: "And More",
	}

	req := &v1.CreateRequest{
		Api:     apiVersion,
		Project: &pr,
	}
	ProjectService.Create(ctx, req)
}

//go test Product-Management-API/pkg/service/v1 -run Test_Read -v
func Test_Read(t *testing.T) {
	ctx := context.Background()
	database, err := db.ConnectDB("")
	if err != nil {
		fmt.Println(err)
	}
	ProjectService := NewProjectServiceServer(database)
	req := &v1.ReadRequest{
		Api: apiVersion,
		Id:  2,
	}
	res, _ := ProjectService.Read(ctx, req)
	fmt.Println(res)
	t.Log("Done")

}

//go test Product-Management-API/pkg/service/v1 -run Test_ReadAll -v
func Test_ReadAll(t *testing.T) {
	ctx := context.Background()
	database, err := db.ConnectDB("")
	if err != nil {
		fmt.Println(err)
	}
	ProjectService := NewProjectServiceServer(database)
	req := &v1.ReadAllRequest{
		Api: apiVersion,
	}
	res, _ := ProjectService.ReadAll(ctx, req)
	fmt.Println(res)
	t.Log("Done")

}
