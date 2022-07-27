# web-application-with-golang
#### Simple web application using golang and postgres database 

---
#### Copy environment variable from .env.example to .env
```bash
cp .env.example .env
```

#### Project run command 
```bash
go run main.go
```

---

### CURL COMMANDS

---
#### Check apis with CURL (`POST API`)
```bash
curl -X POST [options] [URL]
```

#### Check apis with CURL (`GET API`)
```bash
curl -X GET [options] [URL]
```


#### way to make a POST request is to use the -d option. This causes curl to send the data using the application/x-www-form-urlencoded Content-Type.
```bash
curl -X POST -d 'name=XXXXXXX' -d 'email=XXXXXXX@example.com' https://example.com/example.go
```

or

```bash
curl -X POST -d 'name=XXXXXXX&email=XXXXXXX@example.com' https://example.com/example.go
```


#### To set a specific header or Content-Type use the -H option. The following command sets the POST request type to application/json and sends a JSON object:
```bash
curl -X POST -H "Content-Type: application/json" \
    -d '{"name": "XXXXXXX", "email": "XXXXXXX@example.com"}' \
    https://example.com/example.go
```

#### To POST a file with curl, simply add the @ symbol before the file location. The file can be an archive, image, document, etc.
```bash
curl -X POST -F 'image=@/home/user/Pictures/wallpaper.jpg' https://example.com/upload.go
```
---

#### FIle upload process
`Simple HTML Example Code`
```html
<html>
<head>
   	<title>Upload file</title>
</head>
<body>
<form enctype="multipart/form-data" action="http://127.0.0.1:9090/upload" method="post">
	<input type="file" name="uploadfile" />
	<input type="hidden" name="token" value="{{.}}"/>
	<input type="submit" value="upload" />
</form>
</body>
</html>
```









