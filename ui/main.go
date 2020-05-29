package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
)

type Video struct {
	Condition string `json:"condition"`
	FilePath  string `json:"file"`
}

type Videos struct {
	Count  int     `json:"count"`
	Videos []Video `json:"videos`
	Error  string  `json:"error"`
}

func main() {

	m := martini.Classic()
	m.Use(render.Renderer(
		render.Options{
			Directory: "tmpl",
			Charset:   "UTF-8",
		}))

	m.Use(martini.Static("css"))
	m.Use(martini.Static("js"))

	m.Get("/", func(r render.Render) {

		svc := s3.New(session.New(&aws.Config{
			Region: aws.String("us-east-1"),
		}))

		input := &s3.ListObjectsInput{
			// Bucket:  aws.String("kvs-img"),
			Bucket:  aws.String("kvs-img"),
			MaxKeys: aws.Int64(1000),
		}

		result, err := svc.ListObjects(input)

		if err != nil {
			if aerr, ok := err.(awserr.Error); ok {
				switch aerr.Code() {
				case s3.ErrCodeNoSuchBucket:
					fmt.Println(s3.ErrCodeNoSuchBucket, aerr.Error())
				default:
					fmt.Println(aerr.Error())
				}
			} else {
				// Print the error, cast err to awserr.Error to get the Code and
				// Message from an error.
				fmt.Println(err.Error())
			}
			return
		}

		type URL struct {
			Timestamp string
			Link      string
		}

		list := []URL{}

		for _, item := range result.Contents {

			// rtsp_20200521-070250.txt
			tsRaw := []rune(*item.Key)

			if strings.HasSuffix(*item.Key, ".png") {

				ts := fmt.Sprintf("%s-%s-%s %s:%s:%s %s",
					string(tsRaw[5:9]),
					string(tsRaw[9:11]),
					string(tsRaw[11:13]),
					string(tsRaw[14:16]),
					string(tsRaw[16:18]),
					string(tsRaw[18:20]),
					string(tsRaw[21:24]))

				str := *item.Key
				list = append(list, URL{
					Timestamp: ts,
					Link:      fmt.Sprintf("https://kvs-img.s3.amazonaws.com/%s", str),
				})
			}
		}

		r.HTML(http.StatusOK, "home", list)
	})

	m.Get("/health_check", func(r render.Render) {

		type HC struct {
			Result string `json: "result"`
		}
		r.JSON(200, &HC{Result: "ok"})
	})

	m.Get("/samplejson", func(r render.Render) {
		json := map[string]interface{}{"hoge": 100, "key": [3]int{1, 2, 3}}
		r.JSON(200, json)
	})

	m.Get("/videos", func(params martini.Params, res http.ResponseWriter, req *http.Request, r render.Render) {

		// keys, ok := req.URL.Query()["duration"]

		// if !ok || len(keys[0]) < 1 {
		// 	json := &Videos{
		// 		Count:  0,
		// 		Videos: nil,
		// 		Error:  "Duration Value is missing. Please specify duration in hour unit.",
		// 	}
		// 	r.JSON(406, json)

		// 	return
		// }

		svc := s3.New(session.New(&aws.Config{
			Region: aws.String("us-east-1"),
		}))

		input := &s3.ListObjectsInput{
			Bucket:  aws.String("kvs-img"),
			MaxKeys: aws.Int64(1000),
		}

		result, err := svc.ListObjects(input)
		if err != nil {
			if aerr, ok := err.(awserr.Error); ok {
				switch aerr.Code() {
				case s3.ErrCodeNoSuchBucket:
					fmt.Println(s3.ErrCodeNoSuchBucket, aerr.Error())
				default:
					fmt.Println(aerr.Error())
				}
			} else {
				// Print the error, cast err to awserr.Error to get the Code and
				// Message from an error.
				fmt.Println(err.Error())
			}
			return
		}

		// t1 := time.Now().UTC().Add(time.Duration(-1*12) * time.Hour)
		// f1 := fmt.Sprintf("rtsp_%d%02d%02d-%02d%02d%02d", t1.Year(), t1.Month(), t1.Day(), t1.Hour(), t1.Minute(), t1.Second())

		// f1 = "20200521-071054"
		vs := &Videos{}

		for _, item := range result.Contents {

			// if *item.Key > f1 {

			// if strings.HasSuffix(*item.Key, ".webm") {
			vs.Videos = append(vs.Videos, *&Video{
				Condition: "Without Helmet",
				FilePath:  fmt.Sprintf("https://kvs-img.s3.amazonaws.com/%s", *item.Key),
			})
			// }
			log.Println(*item.Key)
			// }
		}

		r.JSON(200, vs)
	})

	m.RunOnAddr(":3000")
}
