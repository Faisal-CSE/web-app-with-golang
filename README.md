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

We need to add a function on the server side to handle this form.

```shell
http.HandleFunc("/upload", upload)

// upload logic
func upload(w http.ResponseWriter, r *http.Request) {
   	fmt.Println("method:", r.Method)
   	if r.Method == "GET" {
       	crutime := time.Now().Unix()
       	h := md5.New()
       	io.WriteString(h, strconv.FormatInt(crutime, 10))
       	token := fmt.Sprintf("%x", h.Sum(nil))

       	t, _ := template.ParseFiles("upload.gtpl")
       	t.Execute(w, token)
   	} else {
       	r.ParseMultipartForm(32 << 20)
       	file, handler, err := r.FormFile("uploadfile")
       	if err != nil {
           	fmt.Println(err)
           	return
       	}
       	defer file.Close()
       	fmt.Fprintf(w, "%v", handler.Header)
       	f, err := os.OpenFile("./test/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
       	if err != nil {
           	fmt.Println(err)
           	return
       	}
       	defer f.Close()
       	io.Copy(f, file)
   	}
}
```

---
#### periodic task management using golang 
##### First install the gocorn package to use periodic task management 

```shell
go get github.com/jasonlvhit/gocron
```
then execute those line of codes
```go
package main

import (
	"fmt"
	"github.com/jasonlvhit/gocron"
)

func myTask() {
	fmt.Println("This task will run periodically")
}
func executeCronJob() {
    gocron.Every(1).Second().Do(myTask)
    <- gocron.Start()
}
func main() {
    go executeCronJob()
}
```














