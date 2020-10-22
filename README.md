# Create a user

`http POST http://localhost:8080/api/users/auth email="john@gmail.com" password="secret"`

Example of response :

```
HTTP/1.1 200 OK
Content-Length: 195
Content-Type: application/json; charset=utf-8
Date: Sun, 28 Apr 2019 18:41:43 GMT

{
    "code": 200,
    "expire": "2019-04-28T21:41:43+02:00",
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NTY0ODA1MDMsIm9yaWdfaWF0IjoxNTU2NDc2OTAzfQ.lzFH1CJen8lzHQApu3wwqX-CqpfRSRClt8rqHPhe5AU"
}
```