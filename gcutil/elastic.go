package gcutil

import (
	"context"
	"encoding/json"
	"github.com/olivere/elastic"
)

var elasticUrl = "http://35.204.247.102//elasticsearch/"

//var elasticUrl = "http://10.164.0.2:9200"

func GetAllCourses(from int) (allCourses []Course, totalPages int64) {
	client := GetElasticCon(elasticUrl)

	query := elastic.NewMatchAllQuery()

	searchResult, err := client.Search().
		Index("courses").
		Query(query).
		Sort("timestamp", false).
		From(from).
		Size(15).
		Pretty(true).
		Do(context.Background())

	if err != nil {
		panic(err)
	}
	//testing page count
	totalPages = searchResult.Hits.TotalHits / 15

	if searchResult.Hits.TotalHits > 0 {

		for _, hit := range searchResult.Hits.Hits {
			var course Courses
			json.Unmarshal(*hit.Source, &course.Course)
			allCourses = append(allCourses, course.Course)
		}
	}
	return
}

func GetCatCourses(s string) []Course {
	client := GetElasticCon(elasticUrl)

	query := elastic.NewMatchQuery("cat_slug", s)

	//query := elastic.NewMatchQuery("cat-slug", s)

	searchResult, err := client.Search().
		Index("courses").
		Query(query).
		Sort("timestamp", false).
		From(0).
		Size(1000).
		Pretty(true).
		Do(context.Background())

	if err != nil {
		panic(err)
	}

	var allCourses []Course
	if searchResult.Hits.TotalHits > 0 {

		for _, hit := range searchResult.Hits.Hits {
			var course Courses
			json.Unmarshal(*hit.Source, &course.Course)
			allCourses = append(allCourses, course.Course)
		}
	}
	return allCourses
}

func SearchCourses(s string) []Course {

	client := GetElasticCon(elasticUrl)

	//query := elastic.NewTermQuery("title", s)
	query := elastic.NewMatchQuery("title", s)

	searchResult, err := client.Search().
		Index("courses").
		Query(query).
		Sort("timestamp", false).
		From(0).
		Size(1000).
		Pretty(true).
		Do(context.Background())

	if err != nil {
		panic(err)
	}

	var allCourses []Course
	if searchResult.Hits.TotalHits > 0 {

		for _, hit := range searchResult.Hits.Hits {
			var course Courses
			json.Unmarshal(*hit.Source, &course.Course)
			allCourses = append(allCourses, course.Course)
		}
	}

	return allCourses
}

func GetRelatedCourses(s string) []Course {
	client := GetElasticCon(elasticUrl)

	//query := elastic.NewTermQuery("title", s)
	query := elastic.NewMatchQuery("category", s)

	searchResult, err := client.Search().
		Index("courses").
		Query(query).
		Sort("timestamp", false).
		From(0).
		Size(3).
		Pretty(true).
		Do(context.Background())

	if err != nil {
		panic(err)
	}

	var allCourses []Course
	if searchResult.Hits.TotalHits > 0 {

		for _, hit := range searchResult.Hits.Hits {
			var course Courses
			json.Unmarshal(*hit.Source, &course.Course)
			allCourses = append(allCourses, course.Course)
		}
	}

	return allCourses

}

func GetCourse(id string) (course Course) {
	client := GetElasticCon(elasticUrl)

	getCourse, _ := client.Get().Index("courses").Type("course").Id(id).Do(context.Background())
	json.Unmarshal(*getCourse.Source, &course)
	return
}

func GetElasticCon(url string) *elastic.Client {
	client, err := elastic.NewSimpleClient(elastic.SetSniff(false), elastic.SetURL(url), elastic.SetBasicAuth("user", ""))
	//client, err := elastic.NewSimpleClient(elastic.SetSniff(false), elastic.SetURL(url))
	if err != nil {
		panic(err)
	}
	return client
}
