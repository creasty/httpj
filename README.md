httpj
=====

Simple http client wrapper for JSON API


Usage
-----

```go
type User struct {
	Id int `json:"id"`
	Name string `json:"name"`
}

resp, err := httpj.Default().NewRequest("https://www.example.com/users/123").
	SetHeader("Authorization", "Barer xxxxxxxxxxxxxxx").
	Patch(User{Name: "John"})
if err != nil {
	panic(err)
}
defer resp.Body.Close()

if resp.IsSuccess() {
	user := User{}
	if err := resp.Bind(&user); err == nil {
		fmt.Println(user)
	}
}
```


License
-------

This project is copyright by [Creasty](http://creasty.com), released under the MIT license.  
See `LICENSE.txt` file for details.
