 **client** package provides an abstraction layer over native golang features used to send HTTP Rest API requests.
 
# Example
```go
req, err := client.NewRequestBuilder().
		SetMethod(client.GET{}).
		SetURLTemplate("https://example.com/api/resource/{id}/subresource/{sub_id}").
		AddPathParam("id", "123").
		AddPathParam("sub_id", "456").
		AddHeader("Content-Type", "application/json").
		AddHeader("Authorization", "Bearer example_token").
		AddParam("param1", "value1").
		AddParam("param2", "value2").
		SetBody(`{
            "name": "John Doe",
            "age": 30
        }`).
		Build()

	if err != nil {
		log.Fatalf("Error building request: %v", err)
	}

	// Send the request
	resp := client.SendRequest(req)
	fmt.Println(resp.StatusCode)