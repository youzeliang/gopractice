package main

import (
	"encoding/json"
	"fmt"
)

var a = `[{
		"publisher": "人民教育出版社",
		"term": "上学期",
		"version": "202408301131",
		"full_name": "一年级上语文（人教统编版）",
		"book_id": "tape1a_002001",
		"grade": "一年级",
		"book_version": "人教统编版",
		"subject": "语文",
		"cover_url": "https://d.namibox.com/static/third_sdk/book_cover1/tape1a_002001.jpg",
		"book_edition": ""
	}, {
		"publisher": "人民教育出版社",
		"term": "上学期",
		"version": "202406180300",
		"full_name": "一年级上语文（人教五四版）",
		"book_id": "tape1a_002010",
		"grade": "一年级",
		"book_version": "人教五四版",
		"subject": "语文",
		"cover_url": "https://d.namibox.com/static/third_sdk/book_cover1/tape1a_002010.jpg",
		"book_edition": ""
	}, {
		"publisher": "人民教育出版社",
		"term": "上学期",
		"version": "202409031226",
		"full_name": "一年级上语文（人教统编版）-24版",
		"book_id": "tape1a_002011",
		"grade": "一年级",
		"book_version": "人教统编版",
		"subject": "语文",
		"cover_url": "https://d.namibox.com/static/third_sdk/book_cover1/tape1a_002011.jpg",
		"book_edition": "2024新课标"
	}, {
		"publisher": "人民教育出版社",
		"term": "上学期",
		"version": "202409031250",
		"full_name": "一年级上语文（人教五四版）-24版",
		"book_id": "tape1a_002013",
		"grade": "一年级",
		"book_version": "人教五四版",
		"subject": "语文",
		"cover_url": "https://d.namibox.com/static/third_sdk/book_cover1/tape1a_002013.jpg",
		"book_edition": "2024新课标"
	}, {
		"publisher": "人民教育出版社",
		"term": "下学期",
		"version": "202407120936",
		"full_name": "一年级下语文（人教统编版）",
		"book_id": "tape1b_002001",
		"grade": "一年级",
		"book_version": "人教统编版",
		"subject": "语文",
		"cover_url": "https://d.namibox.com/static/third_sdk/book_cover1/tape1b_002001.jpg",
		"book_edition": ""
	}, {
		"publisher": "人民教育出版社",
		"term": "下学期",
		"version": "202407120914",
		"full_name": "一年级下语文（人教五四版）",
		"book_id": "tape1b_002010",
		"grade": "一年级",
		"book_version": "人教五四版",
		"subject": "语文",
		"cover_url": "https://d.namibox.com/static/third_sdk/book_cover1/tape1b_002010.jpg",
		"book_edition": ""
	}, {
		"publisher": "人民教育出版社",
		"term": "上学期",
		"version": "202408231002",
		"full_name": "二年级上语文（人教统编版）",
		"book_id": "tape2a_002001",
		"grade": "二年级",
		"book_version": "人教统编版",
		"subject": "语文",
		"cover_url": "https://d.namibox.com/static/third_sdk/book_cover1/tape2a_002001.jpg",
		"book_edition": ""
	}, {
		"publisher": "人民教育出版社",
		"term": "上学期",
		"version": "202408300716",
		"full_name": "二年级上语文（人教五四版）",
		"book_id": "tape2a_002010",
		"grade": "二年级",
		"book_version": "人教五四版",
		"subject": "语文",
		"cover_url": "https://d.namibox.com/static/third_sdk/book_cover1/tape2a_002010.jpg",
		"book_edition": ""
	}, {
		"publisher": "人民教育出版社",
		"term": "下学期",
		"version": "202407121031",
		"full_name": "二年级下语文（人教统编版）",
		"book_id": "tape2b_002001",
		"grade": "二年级",
		"book_version": "人教统编版",
		"subject": "语文",
		"cover_url": "https://d.namibox.com/static/third_sdk/book_cover1/tape2b_002001.jpg",
		"book_edition": ""
	}, {
		"publisher": "人民教育出版社",
		"term": "下学期",
		"version": "202407160628",
		"full_name": "二年级下语文（人教五四版）",
		"book_id": "tape2b_002010",
		"grade": "二年级",
		"book_version": "人教五四版",
		"subject": "语文",
		"cover_url": "https://d.namibox.com/static/third_sdk/book_cover1/tape2b_002010.jpg",
		"book_edition": ""
	}, {
		"publisher": "人民教育出版社",
		"term": "上学期",
		"version": "202409030005",
		"full_name": "三年级上语文（人教统编版）",
		"book_id": "tape3a_002001",
		"grade": "三年级",
		"book_version": "人教统编版",
		"subject": "语文",
		"cover_url": "https://d.namibox.com/static/third_sdk/book_cover1/tape3a_002001.jpg",
		"book_edition": ""
	}, {
		"publisher": "人民教育出版社",
		"term": "上学期",
		"version": "202408300747",
		"full_name": "三年级上语文（人教五四版）",
		"book_id": "tape3a_002010",
		"grade": "三年级",
		"book_version": "人教五四版",
		"subject": "语文",
		"cover_url": "https://d.namibox.com/static/third_sdk/book_cover1/tape3a_002010.jpg",
		"book_edition": ""
	}, {
		"publisher": "人民教育出版社",
		"term": "下学期",
		"version": "202407150812",
		"full_name": "三年级下语文（人教统编版）",
		"book_id": "tape3b_002001",
		"grade": "三年级",
		"book_version": "人教统编版",
		"subject": "语文",
		"cover_url": "https://d.namibox.com/static/third_sdk/book_cover1/tape3b_002001.jpg",
		"book_edition": ""
	}, {
		"publisher": "人民教育出版社",
		"term": "下学期",
		"version": "202407150738",
		"full_name": "三年级下语文（人教五四版）",
		"book_id": "tape3b_002010",
		"grade": "三年级",
		"book_version": "人教五四版",
		"subject": "语文",
		"cover_url": "https://d.namibox.com/static/third_sdk/book_cover1/tape3b_002010.jpg",
		"book_edition": ""
	}, {
		"publisher": "人民教育出版社",
		"term": "上学期",
		"version": "202408231325",
		"full_name": "四年级上语文（人教统编版）",
		"book_id": "tape4a_002001",
		"grade": "四年级",
		"book_version": "人教统编版",
		"subject": "语文",
		"cover_url": "https://d.namibox.com/static/third_sdk/book_cover1/tape4a_002001.jpg",
		"book_edition": ""
	}, {
		"publisher": "人民教育出版社",
		"term": "上学期",
		"version": "202408300821",
		"full_name": "四年级上语文（人教五四版）",
		"book_id": "tape4a_002010",
		"grade": "四年级",
		"book_version": "人教五四版",
		"subject": "语文",
		"cover_url": "https://d.namibox.com/static/third_sdk/book_cover1/tape4a_002010.jpg",
		"book_edition": ""
	}, {
		"publisher": "人民教育出版社",
		"term": "下学期",
		"version": "202407150926",
		"full_name": "四年级下语文（人教统编版）",
		"book_id": "tape4b_002001",
		"grade": "四年级",
		"book_version": "人教统编版",
		"subject": "语文",
		"cover_url": "https://d.namibox.com/static/third_sdk/book_cover1/tape4b_002001.jpg",
		"book_edition": ""
	}, {
		"publisher": "人民教育出版社",
		"term": "下学期",
		"version": "202407150842",
		"full_name": "四年级下语文（人教五四版）",
		"book_id": "tape4b_002010",
		"grade": "四年级",
		"book_version": "人教五四版",
		"subject": "语文",
		"cover_url": "https://d.namibox.com/static/third_sdk/book_cover1/tape4b_002010.jpg",
		"book_edition": ""
	}, {
		"publisher": "人民教育出版社",
		"term": "上学期",
		"version": "202409030037",
		"full_name": "五年级上语文（人教统编版）",
		"book_id": "tape5a_002001",
		"grade": "五年级",
		"book_version": "人教统编版",
		"subject": "语文",
		"cover_url": "https://d.namibox.com/static/third_sdk/book_cover1/tape5a_002001.jpg",
		"book_edition": ""
	}, {
		"publisher": "人民教育出版社",
		"term": "上学期",
		"version": "202408300906",
		"full_name": "五年级上语文（人教五四版）",
		"book_id": "tape5a_002010",
		"grade": "五年级",
		"book_version": "人教五四版",
		"subject": "语文",
		"cover_url": "https://d.namibox.com/static/third_sdk/book_cover1/tape5a_002010.jpg",
		"book_edition": ""
	}, {
		"publisher": "人民教育出版社",
		"term": "下学期",
		"version": "202407151047",
		"full_name": "五年级下语文（人教统编版）",
		"book_id": "tape5b_002001",
		"grade": "五年级",
		"book_version": "人教统编版",
		"subject": "语文",
		"cover_url": "https://d.namibox.com/static/third_sdk/book_cover1/tape5b_002001.jpg",
		"book_edition": ""
	}, {
		"publisher": "人民教育出版社",
		"term": "下学期",
		"version": "202407160701",
		"full_name": "五年级下语文（人教五四版）",
		"book_id": "tape5b_002010",
		"grade": "五年级",
		"book_version": "人教五四版",
		"subject": "语文",
		"cover_url": "https://d.namibox.com/static/third_sdk/book_cover1/tape5b_002010.jpg",
		"book_edition": ""
	}, {
		"publisher": "人民教育出版社",
		"term": "上学期",
		"version": "202408290335",
		"full_name": "六年级上语文（人教统编版）",
		"book_id": "tape6a_002001",
		"grade": "六年级",
		"book_version": "人教统编版",
		"subject": "语文",
		"cover_url": "https://d.namibox.com/static/third_sdk/book_cover1/tape6a_002001.jpg",
		"book_edition": ""
	}, {
		"publisher": "人民教育出版社",
		"term": "上学期",
		"version": "202406180845",
		"full_name": "六年级上语文（人教五四版）",
		"book_id": "tape6a_002010",
		"grade": "六年级",
		"book_version": "人教五四版",
		"subject": "语文",
		"cover_url": "https://d.namibox.com/static/third_sdk/book_cover1/tape6a_002010.jpg",
		"book_edition": ""
	}, {
		"publisher": "人民教育出版社",
		"term": "上学期",
		"version": "202409031339",
		"full_name": "六年级上语文（人教五四版）-24版",
		"book_id": "tape6a_002013",
		"grade": "六年级",
		"book_version": "人教五四版",
		"subject": "语文",
		"cover_url": "https://d.namibox.com/static/third_sdk/book_cover1/tape6a_002013.jpg",
		"book_edition": "2024新课标"
	}, {
		"publisher": "人民教育出版社",
		"term": "下学期",
		"version": "202407160742",
		"full_name": "六年级下语文（人教统编版）",
		"book_id": "tape6b_002001",
		"grade": "六年级",
		"book_version": "人教统编版",
		"subject": "语文",
		"cover_url": "https://d.namibox.com/static/third_sdk/book_cover1/tape6b_002001.jpg",
		"book_edition": ""
	}, {
		"publisher": "人民教育出版社",
		"term": "下学期",
		"version": "202407121332",
		"full_name": "六年级下语文（人教五四版）",
		"book_id": "tape6b_002010",
		"grade": "六年级",
		"book_version": "人教五四版",
		"subject": "语文",
		"cover_url": "https://d.namibox.com/static/third_sdk/book_cover1/tape6b_002010.jpg",
		"book_edition": ""
	}, {
		"publisher": "人民教育出版社",
		"term": "上学期",
		"version": "202408300116",
		"full_name": "七年级上语文（人教统编版）",
		"book_id": "tape7a_002001",
		"grade": "初一",
		"book_version": "人教统编版",
		"subject": "语文",
		"cover_url": "https://d.namibox.com/static/third_sdk/book_cover1/tape7a_002001.jpg",
		"book_edition": ""
	}, {
		"publisher": "人民教育出版社",
		"term": "上学期",
		"version": "202408300947",
		"full_name": "七年级上语文（人教五四版）",
		"book_id": "tape7a_002010",
		"grade": "初一",
		"book_version": "人教五四版",
		"subject": "语文",
		"cover_url": "https://d.namibox.com/static/third_sdk/book_cover1/tape7a_002010.jpg",
		"book_edition": ""
	}, {
		"publisher": "人民教育出版社",
		"term": "上学期",
		"version": "202409031158",
		"full_name": "七年级上语文（人教统编版）-24版",
		"book_id": "tape7a_002011",
		"grade": "初一",
		"book_version": "人教统编版",
		"subject": "语文",
		"cover_url": "https://d.namibox.com/static/third_sdk/book_cover1/tape7a_002011.jpg",
		"book_edition": "2024新课标"
	}, {
		"publisher": "人民教育出版社",
		"term": "下学期",
		"version": "202407011706",
		"full_name": "七年级下语文（人教统编版）",
		"book_id": "tape7b_002001",
		"grade": "初一",
		"book_version": "人教统编版",
		"subject": "语文",
		"cover_url": "https://d.namibox.com/static/third_sdk/book_cover1/tape7b_002001.jpg",
		"book_edition": ""
	}, {
		"publisher": "人民教育出版社",
		"term": "下学期",
		"version": "202407011616",
		"full_name": "七年级下语文（人教五四版）",
		"book_id": "tape7b_002010",
		"grade": "初一",
		"book_version": "人教五四版",
		"subject": "语文",
		"cover_url": "https://d.namibox.com/static/third_sdk/book_cover1/tape7b_002010.jpg",
		"book_edition": ""
	}, {
		"publisher": "人民教育出版社",
		"term": "上学期",
		"version": "202408231215",
		"full_name": "八年级上语文（人教统编版）",
		"book_id": "tape8a_002001",
		"grade": "初二",
		"book_version": "人教统编版",
		"subject": "语文",
		"cover_url": "https://d.namibox.com/static/third_sdk/book_cover1/tape8a_002001.jpg",
		"book_edition": ""
	}, {
		"publisher": "人民教育出版社",
		"term": "上学期",
		"version": "202408301026",
		"full_name": "八年级上语文（人教五四版）",
		"book_id": "tape8a_002010",
		"grade": "初二",
		"book_version": "人教五四版",
		"subject": "语文",
		"cover_url": "https://d.namibox.com/static/third_sdk/book_cover1/tape8a_002010.jpg",
		"book_edition": ""
	}, {
		"publisher": "人民教育出版社",
		"term": "下学期",
		"version": "202407011808",
		"full_name": "八年级下语文（人教统编版）",
		"book_id": "tape8b_002001",
		"grade": "初二",
		"book_version": "人教统编版",
		"subject": "语文",
		"cover_url": "https://d.namibox.com/static/third_sdk/book_cover1/tape8b_002001.jpg",
		"book_edition": ""
	}, {
		"publisher": "人民教育出版社",
		"term": "下学期",
		"version": "202407012017",
		"full_name": "八年级下语文（人教五四版）",
		"book_id": "tape8b_002010",
		"grade": "初二",
		"book_version": "人教五四版",
		"subject": "语文",
		"cover_url": "https://d.namibox.com/static/third_sdk/book_cover1/tape8b_002010.jpg",
		"book_edition": ""
	}, {
		"publisher": "人民教育出版社",
		"term": "上学期",
		"version": "202408231255",
		"full_name": "九年级上语文（人教统编版）",
		"book_id": "tape9a_002001",
		"grade": "初三",
		"book_version": "人教统编版",
		"subject": "语文",
		"cover_url": "https://d.namibox.com/static/third_sdk/book_cover1/tape9a_002001.jpg",
		"book_edition": ""
	}, {
		"publisher": "人民教育出版社",
		"term": "上学期",
		"version": "202408301100",
		"full_name": "九年级上语文（人教五四版）",
		"book_id": "tape9a_002010",
		"grade": "初三",
		"book_version": "人教五四版",
		"subject": "语文",
		"cover_url": "https://d.namibox.com/static/third_sdk/book_cover1/tape9a_002010.jpg",
		"book_edition": ""
	}, {
		"publisher": "人民教育出版社",
		"term": "下学期",
		"version": "202407011909",
		"full_name": "九年级下语文（人教统编版）",
		"book_id": "tape9b_002001",
		"grade": "初三",
		"book_version": "人教统编版",
		"subject": "语文",
		"cover_url": "https://d.namibox.com/static/third_sdk/book_cover1/tape9b_002001.jpg",
		"book_edition": ""
	}, {
		"publisher": "人民教育出版社",
		"term": "下学期",
		"version": "202407011839",
		"full_name": "九年级下语文（人教五四版）",
		"book_id": "tape9b_002010",
		"grade": "初三",
		"book_version": "人教五四版",
		"subject": "语文",
		"cover_url": "https://d.namibox.com/static/third_sdk/book_cover1/tape9b_002010.jpg",
		"book_edition": ""
	}]`

type Book struct {
	Publisher   string `json:"publisher"`
	Term        string `json:"term"`
	Version     string `json:"version"`
	FullName    string `json:"full_name"`
	BookId      string `json:"book_id"`
	Grade       string `json:"grade"`
	BookVersion string `json:"book_version"`
	CoverUrl    string `json:"cover_url"`
	Subject     string `json:"subject"`
}

type ReverseGrade struct {
	Grade string        `json:"grade"`
	Terms []ReverseTerm `json:"terms"`
}

type ReverseTerm struct {
	Term         string   `json:"term"`
	BookVersions []string `json:"book_versions"`
}

func main() {

	var books []Book
	err := json.Unmarshal([]byte(a), &books)
	if err != nil {
		fmt.Println(err)
		return
	}
	tags := map[string]ReverseGrade{}
	for _, book := range books {
		if _, ok := tags[book.Grade]; !ok {
			tags[book.Grade] = ReverseGrade{
				Grade: book.Grade,
				Terms: []ReverseTerm{
					{
						Term: book.Term,
						BookVersions: []string{
							book.BookVersion,
						},
					},
				},
			}
		} else {
			var hasTermIndex = -1
			for i, tagTerm := range tags[book.Grade].Terms {
				if tagTerm.Term == book.Term {
					hasTermIndex = i
					break
				}
			}
			if hasTermIndex >= 0 {
				var hasBookVersion bool
				for _, existBookVersion := range tags[book.Grade].Terms[hasTermIndex].BookVersions {
					if existBookVersion == book.BookVersion {
						hasBookVersion = true
						break
					}
				}
				if !hasBookVersion {
					tags[book.Grade].Terms[hasTermIndex].BookVersions = append(tags[book.Grade].Terms[hasTermIndex].BookVersions, book.BookVersion)
				}
			} else {
				gradeTag := tags[book.Grade]
				gradeTag.Terms = append(gradeTag.Terms, ReverseTerm{
					Term: book.Term,
					BookVersions: []string{
						book.BookVersion,
					},
				})
				tags[book.Grade] = gradeTag
			}
		}
	}
	var realReverseTag []ReverseGrade
	orderGradeSlice := [...]string{"一年级", "二年级", "三年级", "四年级", "五年级", "六年级", "初一", "初二", "初三", "高一", "高二", "高三"}
	for _, orderGrade := range orderGradeSlice {
		if reverseGrade, ok := tags[orderGrade]; ok {
			realReverseTag = append(realReverseTag, reverseGrade)
		}
	}

	fmt.Println(realReverseTag)

}
