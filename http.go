package main
  
import (
        "net/http"
		"fmt"
       )

func main() {
    http.HandleFunc("/",myResponse)
   http.ListenAndServe("127.0.0.1:8888",nil)
}

func myResponse(w http.ResponseWriter,r* http.Request)  {
    w.Write([]byte("<html><center> <font size=\"40\">Hello！ I am go service started by 梁志成!</font></center></html>"))
	fmt.Println("A client has just visited!")
}