package wiki

import (
	"errors"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
)

type Page struct {
	Title string
	Body  []byte
}

var templates *template.Template = template.Must(template.ParseFiles("templates/view.html", "templates/edit.html"))
var validPath = regexp.MustCompile("^/(view|edit|save)/([a-zA-Z0-9]+)$")

func generalHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi There, I lovs %s", r.URL.Path[1:])
}

func StartGeneralServer(port int) {
	http.HandleFunc("/", generalHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func StartFileViewServer() {
	http.HandleFunc("/view/", handlerFuncMaker(fileViewHandler))
	http.Handle("/edit/", editHandler{})
	http.HandleFunc("/save/", handlerFuncMaker(saveHandler))
	log.Fatal(http.ListenAndServe(":8080", nil))
}

//view handler func
func fileViewHandler(w http.ResponseWriter, r *http.Request, title string) {
	page, err := Load(title)
	if err != nil {
		fmt.Printf("Error loading file %s\n%s", title, err)
		//fmt.Fprintf(w, "Error loading file %s", title)
		//return
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}
	renderTemplate(w, r, "view", page)

	//fmt.Fprintf(w, "<h1>%s</h1><div>%s<div>", page.Title, page.Body)
}

// save handler
func saveHandler(w http.ResponseWriter, r *http.Request, title string) {
	body := r.FormValue("body")
	page := &Page{Title: title, Body: []byte(body)}
	err := page.Save()
	if err != nil {
		fmt.Println("Error saving file ", title)
		//fmt.Fprintf(w, "Error saving file %s", title)
		http.Error(w, "Error saving file "+title+err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

//edit handler
type editHandler struct{}

func (h editHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	title, err := extractTitleFromURLPath(w, r)
	if err != nil {
		return
	}
	page, err := Load(title)
	if err != nil {
		fmt.Println("Error loading file ", title, "\n", err)
		//fmt.Fprintf(w, "Error loading file %s\n", title)
		//return
		page = &Page{
			Title: title,
			Body:  []byte(""),
		}
	}
	renderTemplate(w, r, "edit", page)
	// fmt.Fprintf(w, "<h1>Editing %s</h1>"+
	// "<form action=\"/save/%s\" method=\"POST\">"+
	// "<textarea name=\"body\">%s</textarea><br>"+
	// "<input type=\"submit\" value=\"Save\">"+
	// "</form>",
	// page.Title, page.Title, page.Body)
}

func extractTitleFromURLPath(w http.ResponseWriter, r *http.Request) (string, error) {
	m := validPath.FindStringSubmatch(r.URL.Path)
	if m == nil {
		http.NotFound(w, r)
		return "", errors.New("Invalid page title")
	}
	return m[2], nil
}

func renderTemplate(w http.ResponseWriter, r *http.Request, templateName string, page *Page) {
	fileName := templateName + ".html"
	err := templates.ExecuteTemplate(w, fileName, page)
	if err != nil {
		fmt.Printf("Error executing template file %s. page: %v \n %s", fileName, page, err)
		//fmt.Fprintf(w, "Error executing template file %s", fileName)
		http.Error(w, "Error executing template file "+fileName+"\n"+err.Error(), http.StatusInternalServerError)
		return
	}
}

func handlerFuncMaker(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		title, err := extractTitleFromURLPath(w, r)
		if err != nil {
			http.NotFound(w, r)
			return
		}
		fn(w, r, title)
	}
}

// save save page
func (p *Page) Save() error {
	fileName := "files/" + p.Title + ".txt"
	if _, err := os.Stat("files"); os.IsNotExist(err) {
		os.MkdirAll("files", 0700)
	}

	return ioutil.WriteFile(fileName, p.Body, 0600)
}

// load load file
func Load(title string) (*Page, error) {
	fileName := "files/" + title + ".txt"
	body, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error loading file ", fileName)
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}
