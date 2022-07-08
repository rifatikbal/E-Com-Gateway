***RegisterUserAPI***
```
web: POST api/user/register
```

Request Body

```
{
   "name":  "string", // ex: mahfuz
   "email": "string" // ex: "mahfuz@gmail.com",
   "address": "string" //  ex: "House no 332, shahjadpur,dhaka",
   "password": "string" // ex: 1234
}

```

***Response***
```
status code : 201
```

```
{
   "success": "true"
}
```

***Error***
{
"error": string
}



***LoginUserAPI***

```
web: POST api/user/register
```

Request Body

```
{
   "email": "string" // ex: "mahfuz@gmail.com",
   "password": "string" // ex: 1234
}
```


