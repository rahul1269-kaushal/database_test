package main
import(
       "fmt"
       "net/http"
       "github.com/gorilla/mux"
       )
func main(){
       http.HandleFunc("/", handler)
       http.ListenAndServe("0.0.0.0:8080", nil)
       
}

func handler(w http.ResponseWriter, r *http.Request) {
 var name string
 var title string      
 fmt.println(w, "Enter title")
 fmt.scanf("%s", &title)
 fmt.println(w, "Enter name")
 fmt.scanf("%s", &name)
 fmt.println(w, "Hello", title, name)     
}
