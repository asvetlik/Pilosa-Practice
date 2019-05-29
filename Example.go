package main

import (
	"fmt"
	"time"

	"github.com/pilosa/go-pilosa"
)

func main() {
	var err error

	client := pilosa.DefaultClient()
	schema, err := client.Schema()
	theindex := schema.Index("theindex")
	field1 := theindex.Field("field1")
	field2 := theindex.Field("field2")
	field3 := theindex.Field("field3")
	field4 := theindex.Field("field4")
	field5 := theindex.Field("field5")

	err = client.SyncSchema(schema)

	response, err := client.Query(field1.Row(0))

	// Field1: simply adds numbers to the field then returns them
	for r := 1; r <= 10; r++ {
		response, err = client.Query(field1.Set(2, r))
		if err != nil {
			fmt.Println(err)
		}
	}
	response, err = client.Query(field1.Row(0))
	fmt.Println("Field1: ", response.Result().Row().Columns)
	if err != nil {
		fmt.Println(err)
	}

	// Field2: adds characters to the field and returns
	for i := 'A'; i < 'z'; i++ {
		response, err = client.Query(field2.Set(0, i))
		if err != nil {
			fmt.Println(err)
		}
	}
	response, err = client.Query(field2.Row(0))
	fmt.Println("Field2: ", response.Result().Row().Columns)
	if err != nil {
		fmt.Println(err)
	}

	// Field3: displays how to load and clear a field
	fmt.Println("Field3: ")
	response, err = client.Query(theindex.BatchQuery(field3.Set(0, 22), field3.Set(1, 24), field3.Set(2, 37)))
	if err != nil {
		fmt.Println(err)
	}
	for q := 0; q < 3; q++ {
		response, err = client.Query(field3.Row(q))
		if err != nil {
			fmt.Println(err)
		}
		fmt.Print(response.Result().Row().Columns)
		response, err = client.Query(field3.ClearRow(q))
		if err != nil {
			fmt.Println(err)
		}
	}
	fmt.Println()
	for g := 0; g < 3; g++ {
		response, err = client.Query(field3.Row(g))
		if err != nil {
			fmt.Println(err)
		}
		fmt.Print(" ", response.Result().Row().Columns, " ")
	}
	fmt.Println()

	// Field4: displays how to set a timestamp
	//		Note: I don't think there is a way to call just the time stamp
	fmt.Println("Field4: ")
	for h := 0; h < 3; h++ {
		for k := 0; k < 13; k++ {
			response, err = client.Query(field4.SetTimestamp(h, k, time.Date(2019, time.May, h, k, 0, 0, 0, time.UTC)))
			if err != nil {
				fmt.Println(err)
			}
		}
	}
	for u := 0; u < 3; u++ {
		response, err = client.Query(field4.Row(u))
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(response.Result().Row().Columns)
	}

	// Field5: displays the application of the Min/Max functions
	//		Note: There is no difference between the Min/Max
	fmt.Println("Field5: ")
	for z := 0; z < 13; z++ {
		response, err = client.Query(field5.Set(0, z))
		if err != nil {
			fmt.Println(err)
		}
	}
	response, err = client.Query(field5.Row(0))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(response.Result().Row().Columns)
	response, err = client.Query(field5.Min(field5.Row(0)))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Min: ", response.Result().Value())
	response, err = client.Query(field5.Max(field5.Row(0)))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Max: ", response.Result().Value())
}
