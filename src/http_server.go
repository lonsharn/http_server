package main

import "net/http"
import "html/template"
//import "io/ioutil"
import "fmt"
import "path/filepath"
import "os"


func ServeMyServer(w http.ResponseWriter, r *http.Request){
    return
}

func main(){
    filepath.Walk("./", func(path string, f os.FileInfo, err error) error{
        if !f.IsDir() && len(path) > 0{
            httpPath:="/"+path
            fmt.Println("http path=", httpPath, "disk path=", path)
            //ext:=filepath.Ext(path)
            http.HandleFunc(httpPath, func(w http.ResponseWriter, r *http.Request){
                /*bytes, err := ioutil.ReadFile(path)
                if err != nil{
                    fmt.Println("ReadFile Error:", err)
                    return
                }
                */
                if r.Method == "GET" {
                    t, err := template.ParseFiles(path)
                    if err != nil{
                        fmt.Println("template.Parsefiles", err)
                        return
                    }
                    t.Execute(w, nil)
                    /*
                    if httpPath == "/" {
                        bytes, err = ioutil.ReadFile("/");
                    }

                    if ext == ".html" {
                        written, err:=w.Write(bytes)
                        if err!= nil{
                            fmt.Println("Write Error:", err)
                            return
                        }
                        fmt.Println("Write", written, "Bytes")
                    }else if ext == ".png" || ext == ".jpg" {
                        http.ServeFile(w, r, path)
                        fmt.Println("Serve png file", path)
                    }else if ext == ".js" {
                        http.ServeFile(w, r, path)
                        fmt.Println("Serve js file", path)
                    }else if ext == ".css" {
                        http.ServeFile(w, r, path)
                        fmt.Println("Serve js file", path)
                    }else{
                        http.ServeFile(w, r, path)
                        fmt.Println("Serve file:", path)
                    }
                    */
                }else{
                    r.ParseForm()
                    fmt.Println(r.Form["first_name"])
                    fmt.Println(r.Form["last_name"])
                }

            })
        }
        return nil
    })

    fmt.Println(http.ListenAndServe(":8080", nil))
}

