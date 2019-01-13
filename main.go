package main

import (
	"gc/gcutil"
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
	"strconv"
	"strings"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", index)
	r.HandleFunc("/course/{id}/{prettyurl}/", single)
	r.HandleFunc("/page/{id}", page)
	r.HandleFunc("/category/{category}", category)
	r.HandleFunc("/goto/{id}", udemyUrl)
	r.HandleFunc("/result", result)
	r.HandleFunc("/privacy", privacy)
	http.Handle("/", r)
	//http.Handle("/css", http.FileServer(http.Dir("css/")))
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css"))))

	http.ListenAndServe(":8080", nil)
	//http.ListenAndServe(":9000", nil)
}

func privacy(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "index.html", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	var gcIndex gcutil.GCIndex
	gcIndex.Courses, gcIndex.Paging.TotalPages = gcutil.GetAllCourses(0)
	gcIndex.Paging.ActivePage = 1
	gcIndex.Paging.NextPage = 2
	gcIndex.Paging.PrevPage = 0
	gcIndex.Title = "GreatCourses.co - Get the latest discounted and free UDEMY courses"
	gcIndex.Description = "GreatCourses.co - all the latest and valid udemy coupon codes and udemy special offers in one place."
	gcIndex.Canonical = "/"

	tpl.ExecuteTemplate(w, "index.html", gcIndex)
}

func page(w http.ResponseWriter, r *http.Request) {
	pullID := strings.Split(r.URL.RequestURI(), "/")
	intPage, _ := strconv.Atoi(pullID[2])
	from := (intPage - 1) * 15
	var gcIndex gcutil.GCIndex
	gcIndex.Courses, gcIndex.Paging.TotalPages = gcutil.GetAllCourses(from)
	gcIndex.Paging.ActivePage = intPage
	gcIndex.Paging.NextPage = intPage + 1
	gcIndex.Paging.PrevPage = intPage - 1
	gcIndex.Paging.PpbtPage = gcIndex.Paging.TotalPages - 2
	gcIndex.Paging.PbtPage = gcIndex.Paging.TotalPages - 1
	gcIndex.Title = "GreatCourses.co - Get the latest discounted and free UDEMY courses - Page " + pullID[2]
	gcIndex.Description = "GreatCourses.co - all the latest and valid udemy coupon codes and udemy special offers in one place - Page " + pullID[2]
	gcIndex.Canonical = r.URL.Path

	tpl.ExecuteTemplate(w, "page.html", gcIndex)
}

func single(w http.ResponseWriter, r *http.Request) {
	pullID := strings.Split(r.URL.RequestURI(), "/")
	var gcSingle gcutil.GCSingle
	gcSingle.Course = gcutil.GetCourse(pullID[2])
	gcSingle.Title = "[" + gcSingle.Course.DiscountRate + "% OFF] " + gcSingle.Course.Title + " | Udemy Coupon"
	gcSingle.Description = gcSingle.Course.Description
	gcSingle.Canonical = r.URL.Path
	gcSingle.GoToUrl = "/goto/" + gcSingle.Course.ID

	gcSingle.Related = gcutil.GetRelatedCourses(gcSingle.Course.Category)

	tpl.ExecuteTemplate(w, "single.html", gcSingle)
}

func udemyUrl(w http.ResponseWriter, r *http.Request) {
	pullID := strings.Split(r.URL.RequestURI(), "/")
	var gcGoTo gcutil.GCGoTo
	gcGoTo.Course = gcutil.GetCourse(pullID[2])
	udemyUrl := "https://click.linksynergy.com/deeplink?id=HgwMTFqkx/0&mid=39197&murl=" + gcGoTo.Course.RemoteUrl + "?couponCode=" + gcGoTo.Course.CouponCode
	http.Redirect(w, r, udemyUrl, 301)
}

func result(w http.ResponseWriter, r *http.Request) {

	var gcIndex gcutil.GCIndex

	gcIndex.Courses = gcutil.SearchCourses(r.FormValue("search"))

	tpl.ExecuteTemplate(w, "search.html", gcIndex)
}

func category(w http.ResponseWriter, r *http.Request) {

	pullID := strings.Split(r.URL.RequestURI(), "/")

	var gcIndex gcutil.GCIndex
	gcIndex.Courses = gcutil.GetCatCourses(pullID[2])

	tpl.ExecuteTemplate(w, "category.html", gcIndex)
}
