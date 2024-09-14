package main

import (
	"encoding/json"
	"fmt"
)

var a = `[ {
		"publisher": "江苏译林出版社",
		"term": "上学期",
		"version": "202408140608",
		"full_name": "三年级上英语（译林版）",
		"book_id": "tape3a_001002",
		"grade": "三年级",
		"book_version": "译林版",
		"subject": "英语",
		"cover_url": "https://d.namibox.com/digitbook/tape3a_001002/title.jpg?ts=20210617190402",
		"book_edition": ""
	}, {
		"publisher": "江苏译林出版社",
		"term": "上学期",
		"version": "202408130744",
		"full_name": "三年级上英语（译林版）-24版",
		"book_id": "tape3a_001012",
		"grade": "三年级",
		"book_version": "译林版",
		"subject": "英语",
		"cover_url": "https://d.namibox.com/digitbook/tape3a_001012/title.jpg?ts=20240806154019",
		"book_edition": "2024新课标"
	}, {
		"publisher": "人民教育出版社",
		"term": "上学期",
		"version": "202407030615",
		"full_name": "三年级上英语（人教新起点）",
		"book_id": "tape3a_002002",
		"grade": "三年级",
		"book_version": "人教新起点",
		"subject": "英语",
		"cover_url": "https://d.namibox.com/static/third_sdk/book_cover1/tape3a_002002.jpg",
		"book_edition": ""
	}, {
		"publisher": "人民教育出版社",
		"term": "上学期",
		"version": "202406190626",
		"full_name": "三年级上英语（人教PEP版）",
		"book_id": "tape3a_002003",
		"grade": "三年级",
		"book_version": "人教PEP版",
		"subject": "英语",
		"cover_url": "https://d.namibox.com/static/third_sdk/book_cover1/tape3a_002003.jpg",
		"book_edition": ""
	}, {
		"publisher": "人民教育出版社",
		"term": "上学期",
		"version": "202310220022",
		"full_name": "三年级上英语（人教精通版）",
		"book_id": "tape3a_002004",
		"grade": "三年级",
		"book_version": "人教精通版",
		"subject": "英语",
		"cover_url": "https://d.namibox.com/static/third_sdk/book_cover1/tape3a_002004.jpg",
		"book_edition": ""
	}, {
		"publisher": "人民教育出版社",
		"term": "上学期",
		"version": "202409031138",
		"full_name": "三年级上英语（人教PEP版）-24版",
		"book_id": "tape3a_002013",
		"grade": "三年级",
		"book_version": "人教PEP版",
		"subject": "英语",
		"cover_url": "https://d.namibox.com/static/third_sdk/book_cover1/tape3a_002013.jpg",
		"book_edition": "2024新课标"
	}, {
		"publisher": "人民教育出版社",
		"term": "上学期",
		"version": "202409030926",
		"full_name": "三年级上英语（人教精通版）-24版",
		"book_id": "tape3a_002014",
		"grade": "三年级",
		"book_version": "人教精通版",
		"subject": "英语",
		"cover_url": "https://d.namibox.com/static/third_sdk/book_cover1/tape3a_002014.jpg",
		"book_edition": "2024新课标"
	}, {
		"publisher": "北京师范大学出版社",
		"term": "上学期",
		"version": "202310220057",
		"full_name": "三年级上英语（北师大一年级起）",
		"book_id": "tape3a_003002",
		"grade": "三年级",
		"book_version": "北师大一年级起",
		"subject": "英语",
		"cover_url": "https://d.namibox.com/digitbook/tape3a_003002/title.jpg?ts=20210727104235",
		"book_edition": ""
	}, {
		"publisher": "北京师范大学出版社",
		"term": "上学期",
		"version": "202408200401",
		"full_name": "三年级上英语（北师大三年级起）",
		"book_id": "tape3a_003003",
		"grade": "三年级",
		"book_version": "北师大三年级起",
		"subject": "英语",
		"cover_url": "https://d.namibox.com/digitbook/tape3a_003003/title.jpg?ts=20210727091947",
		"book_edition": ""
	}, {
		"publisher": "北京师范大学出版社",
		"term": "上学期",
		"version": "202408241034",
		"full_name": "三年级上英语（北师大三年级起）-24版",
		"book_id": "tape3a_003013",
		"grade": "三年级",
		"book_version": "北师大三年级起",
		"subject": "英语",
		"cover_url": "https://d.namibox.com/digitbook/tape3a_003013/title.jpg?ts=20240801135944",
		"book_edition": "2024新课标"
	}, {
		"publisher": "外语教学与研究出版社",
		"term": "上学期",
		"version": "202310231357",
		"full_name": "三年级上英语（外研版一年级起点）",
		"book_id": "tape3a_004001",
		"grade": "三年级",
		"book_version": "外研版一年级起点",
		"subject": "英语",
		"cover_url": "https://d.namibox.com/digitbook/tape3a_004001/title.jpg?ts=20220901161437",
		"book_edition": ""
	}, {
		"publisher": "外语教学与研究出版社",
		"term": "上学期",
		"version": "202310240857",
		"full_name": "三年级上英语（外研版三年级起点）",
		"book_id": "tape3a_004002",
		"grade": "三年级",
		"book_version": "外研版三年级起点",
		"subject": "英语",
		"cover_url": "https://d.namibox.com/digitbook/tape3a_004002/title.jpg?ts=20220901161713",
		"book_edition": ""
	}, {
		"publisher": "外语教学与研究出版社",
		"term": "上学期",
		"version": "202408140906",
		"full_name": "三年级上英语（外研版Join In）",
		"book_id": "tape3a_004008",
		"grade": "三年级",
		"book_version": "外研社Join in",
		"subject": "英语",
		"cover_url": "https://d.namibox.com/digitbook/tape3a_004008/title.jpg?ts=20230821175312",
		"book_edition": ""
	}, {
		"publisher": "外语教学与研究出版社",
		"term": "上学期",
		"version": "202408300404",
		"full_name": "三年级上英语（外研版三年级起点）-24版",
		"book_id": "tape3a_004012",
		"grade": "三年级",
		"book_version": "外研版三年级起点",
		"subject": "英语",
		"cover_url": "https://d.namibox.com/digitbook/tape3a_004012/title.jpg?ts=20240726175434",
		"book_edition": "2024新课标"
	}, {
		"publisher": "外语教学与研究出版社",
		"term": "上学期",
		"version": "202408140833",
		"full_name": "三年级上英语（外研版Join In）-24版",
		"book_id": "tape3a_004018",
		"grade": "三年级",
		"book_version": "外研社Join in",
		"subject": "英语",
		"cover_url": "https://d.namibox.com/digitbook/tape3a_004018/title.jpg?ts=20240813151841",
		"book_edition": "2024新课标"
	}, {
		"publisher": "陕西旅游出版社",
		"term": "上学期",
		"version": "202310301035",
		"full_name": "三年级上英语（陕旅版）",
		"book_id": "tape3a_005002",
		"grade": "三年级",
		"book_version": "陕旅版",
		"subject": "英语",
		"cover_url": "https://d.namibox.com/digitbook/tape3a_005002/title.jpg?ts=20220826112034",
		"book_edition": ""
	}, {
		"publisher": "陕西旅游出版社",
		"term": "上学期",
		"version": "202408301659",
		"full_name": "三年级上英语（陕旅版）-24版",
		"book_id": "tape3a_005012",
		"grade": "三年级",
		"book_version": "陕旅版",
		"subject": "英语",
		"cover_url": "https://d.namibox.com/digitbook/tape3a_005012/title.jpg?ts=20240805003505",
		"book_edition": "2024新课标"
	}, {
		"publisher": "湖南少年儿童出版社",
		"term": "上学期",
		"version": "202405201025",
		"full_name": "三年级上英语（湘少版）",
		"book_id": "tape3a_006002",
		"grade": "三年级",
		"book_version": "湘少版",
		"subject": "英语",
		"cover_url": "https://f.namibox.com/oms/book_cover/2087325/ss_1721820236113.png?ts=20240724192418",
		"book_edition": ""
	}, {
		"publisher": "湖南教育出版社",
		"term": "上学期",
		"version": "202312261157",
		"full_name": "三年级上英语（湘鲁版）",
		"book_id": "tape3a_007002",
		"grade": "三年级",
		"book_version": "湘鲁版",
		"subject": "英语",
		"cover_url": "https://d.namibox.com/digitbook/tape3a_007002/title.jpg?ts=20220825173754",
		"book_edition": ""
	}, {
		"publisher": "湖南教育出版社",
		"term": "上学期",
		"version": "202408301408",
		"full_name": "三年级上英语（湘鲁版）-24版",
		"book_id": "tape3a_007012",
		"grade": "三年级",
		"book_version": "湘鲁版",
		"subject": "英语",
		"cover_url": "https://d.namibox.com/digitbook/tape3a_007012/title.jpg?ts=20240827173949",
		"book_edition": "2024新课标"
	}, {
		"publisher": "上海教育出版社",
		"term": "上学期",
		"version": "202310231011",
		"full_name": "三年级上英语（上教社牛津全国版）",
		"book_id": "tape3a_008002",
		"grade": "三年级",
		"book_version": "牛津英语全国版",
		"subject": "英语",
		"cover_url": "https://d.namibox.com/digitbook/tape3a_008002/title.jpg?ts=20210803153228",
		"book_edition": ""
	}, {
		"publisher": "上海教育出版社",
		"term": "上学期",
		"version": "202409011542",
		"full_name": "三年级上英语（上教社牛津全国版）-24版",
		"book_id": "tape3a_008012",
		"grade": "三年级",
		"book_version": "牛津英语全国版",
		"subject": "英语",
		"cover_url": "https://d.namibox.com/digitbook/tape3a_008012/title.jpg?ts=20240808145005",
		"book_edition": "2024新课标"
	}, {
		"publisher": "北京出版社",
		"term": "上学期",
		"version": "202310231252",
		"full_name": "三年级上英语（北京出版社）",
		"book_id": "tape3a_009002",
		"grade": "三年级",
		"book_version": "北京版",
		"subject": "英语",
		"cover_url": "https://d.namibox.com/digitbook/tape3a_009002/title.jpg?ts=20220909145517",
		"book_edition": ""
	}, {
		"publisher": "北京出版社",
		"term": "上学期",
		"version": "202408241050",
		"full_name": "三年级上英语（北京出版社）-24版",
		"book_id": "tape3a_009012",
		"grade": "三年级",
		"book_version": "北京版",
		"subject": "英语",
		"cover_url": "https://d.namibox.com/digitbook/tape3a_009012/title.jpg?ts=20240822151340",
		"book_edition": "2024新课标"
	}, {
		"publisher": "教育科学出版社",
		"term": "上学期",
		"version": "202310220137",
		"full_name": "三年级上英语（教科版）",
		"book_id": "tape3a_010002",
		"grade": "三年级",
		"book_version": "教科版",
		"subject": "英语",
		"cover_url": "https://d.namibox.com/digitbook/tape3a_010002/title.jpg?ts=20220816172950",
		"book_edition": ""
	}, {
		"publisher": "教育科学出版社",
		"term": "上学期",
		"version": "202310220147",
		"full_name": "三年级上英语（教科EEC）",
		"book_id": "tape3a_010003",
		"grade": "三年级",
		"book_version": "教科EEC",
		"subject": "英语",
		"cover_url": "https://d.namibox.com/digitbook/tape3a_010003/title.jpg?ts=20220816172609",
		"book_edition": ""
	}, {
		"publisher": "教育科学出版社",
		"term": "上学期",
		"version": "202409031008",
		"full_name": "三年级上英语（教科版）-24版",
		"book_id": "tape3a_010012",
		"grade": "三年级",
		"book_version": "教科版",
		"subject": "英语",
		"cover_url": "https://d.namibox.com/digitbook/tape3a_010012/title.jpg?ts=20240812124115",
		"book_edition": "2024新课标"
	}, {
		"publisher": "科学普及出版社",
		"term": "上学期",
		"version": "202310220157",
		"full_name": "三年级上英语（科普版）",
		"book_id": "tape3a_011002",
		"grade": "三年级",
		"book_version": "科普版",
		"subject": "英语",
		"cover_url": "https://d.namibox.com/digitbook/tape3a_011002/title.jpg?ts=20220825172055",
		"book_edition": ""
	}, {
		"publisher": "科学普及出版社",
		"term": "上学期",
		"version": "202408300320",
		"full_name": "三年级上英语（科普版）-24版",
		"book_id": "tape3a_011012",
		"grade": "三年级",
		"book_version": "科普版",
		"subject": "英语",
		"cover_url": "https://d.namibox.com/digitbook/tape3a_011012/title.jpg?ts=20240820091306",
		"book_edition": "2024新课标"
	}, {
		"publisher": "河北教育出版社",
		"term": "上学期",
		"version": "202310220205",
		"full_name": "三年级上英语（冀教版一年级起）",
		"book_id": "tape3a_012002",
		"grade": "三年级",
		"book_version": "冀教版一年级起",
		"subject": "英语",
		"cover_url": "https://d.namibox.com/digitbook/tape3a_012002/title.jpg?ts=20220817165215",
		"book_edition": ""
	}, {
		"publisher": "河北教育出版社",
		"term": "上学期",
		"version": "202310220212",
		"full_name": "三年级上英语（冀教版三年级起）",
		"book_id": "tape3a_012003",
		"grade": "三年级",
		"book_version": "冀教版三年级起",
		"subject": "英语",
		"cover_url": "https://d.namibox.com/digitbook/tape3a_012003/title.jpg?ts=20220817165652",
		"book_edition": ""
	}, {
		"publisher": "河北教育出版社",
		"term": "上学期",
		"version": "202408260734",
		"full_name": "三年级上英语（冀教版三年级起）-24版",
		"book_id": "tape3a_012013",
		"grade": "三年级",
		"book_version": "冀教版三年级起",
		"subject": "英语",
		"cover_url": "https://d.namibox.com/digitbook/tape3a_012013/title.jpg?ts=20240822123934",
		"book_edition": "2024新课标"
	}, {
		"publisher": "福建教育出版社",
		"term": "上学期",
		"version": "202408060824",
		"full_name": "三年级上英语（闽教版）-24版",
		"book_id": "tape3a_013005",
		"grade": "三年级",
		"book_version": "闽教版英语",
		"subject": "英语",
		"cover_url": "https://d.namibox.com/digitbook/tape3a_013005/title.jpg?ts=20240726172546",
		"book_edition": "2024新课标"
	}, {
		"publisher": "广东人民出版社",
		"term": "上学期",
		"version": "202310220220",
		"full_name": "三年级上英语（粤人版）",
		"book_id": "tape3a_014002",
		"grade": "三年级",
		"book_version": "粤人版",
		"subject": "英语",
		"cover_url": "https://d.namibox.com/digitbook/tape3a_014002/title.jpg?ts=20230508093203",
		"book_edition": ""
	}, {
		"publisher": "辽宁师范大学出版社",
		"term": "上学期",
		"version": "202310220231",
		"full_name": "三年级上英语（辽宁师大版）",
		"book_id": "tape3a_017002",
		"grade": "三年级",
		"book_version": "辽宁师大版",
		"subject": "英语",
		"cover_url": "https://d.namibox.com/digitbook/tape3a_017002/title.jpg?ts=20221108191633",
		"book_edition": ""
	}, {
		"publisher": "辽宁师范大学出版社",
		"term": "上学期",
		"version": "202310220240",
		"full_name": "三年级上快乐英语（辽宁师大版）",
		"book_id": "tape3a_017003",
		"grade": "三年级",
		"book_version": "辽宁师大快乐英语",
		"subject": "英语",
		"cover_url": "https://d.namibox.com/digitbook/tape3a_017003/title.jpg?ts=20211123143625",
		"book_edition": ""
	}, {
		"publisher": "辽宁师范大学出版社",
		"term": "上学期",
		"version": "202310220245",
		"full_name": "三年级上快乐英语（辽宁师大市场版）",
		"book_id": "tape3a_017004",
		"grade": "三年级",
		"book_version": "辽宁师大快乐英语市场版",
		"subject": "英语",
		"cover_url": "https://d.namibox.com/digitbook/tape3a_017004/title.jpg?ts=20211123161200",
		"book_edition": ""
	}, {
		"publisher": "辽宁师范大学出版社",
		"term": "上学期",
		"version": "202408310542",
		"full_name": "三年级上英语（辽宁师大版）-24版",
		"book_id": "tape3a_017012",
		"grade": "三年级",
		"book_version": "辽宁师大版",
		"subject": "英语",
		"cover_url": "https://d.namibox.com/digitbook/tape3a_017012/title.jpg?ts=20240827173857",
		"book_edition": "2024新课标"
	}, {
		"publisher": "清华大学出版社",
		"term": "上学期",
		"version": "202310220251",
		"full_name": "三年级上英语（清华大学版）",
		"book_id": "tape3a_018002",
		"grade": "三年级",
		"book_version": "清华大学版",
		"subject": "英语",
		"cover_url": "https://d.namibox.com/digitbook/tape3a_018002/title.jpg?ts=20210819211135",
		"book_edition": ""
	}, {
		"publisher": "四川教育出版社",
		"term": "上学期",
		"version": "202310220303",
		"full_name": "三年级上英语（川教版三年级起）",
		"book_id": "tape3a_019003",
		"grade": "三年级",
		"book_version": "川教三年级起",
		"subject": "英语",
		"cover_url": "https://d.namibox.com/digitbook/tape3a_019003/title.jpg?ts=20220818222503",
		"book_edition": ""
	}, {
		"publisher": "四川教育出版社",
		"term": "上学期",
		"version": "202408300550",
		"full_name": "三年级上英语（川教版三年级起）-24版",
		"book_id": "tape3a_019013",
		"grade": "三年级",
		"book_version": "川教三年级起",
		"subject": "英语",
		"cover_url": "https://d.namibox.com/digitbook/tape3a_019013/title.jpg?ts=20240806153538",
		"book_edition": "2024新课标"
	}, {
		"publisher": "重庆大学出版社",
		"term": "上学期",
		"version": "202310220312",
		"full_name": "三年级上英语（重大版）",
		"book_id": "tape3a_020002",
		"grade": "三年级",
		"book_version": "重大版",
		"subject": "英语",
		"cover_url": "https://d.namibox.com/digitbook/tape3a_020002/title.jpg?ts=20210818104100",
		"book_edition": ""
	}, {
		"publisher": "重庆大学出版社",
		"term": "上学期",
		"version": "202408300341",
		"full_name": "三年级上英语（重大版）-24版",
		"book_id": "tape3a_020012",
		"grade": "三年级",
		"book_version": "重大版",
		"subject": "英语",
		"cover_url": "https://d.namibox.com/digitbook/tape3a_020012/title.jpg?ts=20240827173321",
		"book_edition": "2024新课标"
	}, {
		"publisher": "培生朗文",
		"term": "上学期",
		"version": "202310220321",
		"full_name": "三年级上新朗文小学英语",
		"book_id": "tape3a_023002",
		"grade": "三年级",
		"book_version": "新朗文小学英语LWTE",
		"subject": "英语",
		"cover_url": "https://d.namibox.com/digitbook/tape3a_023002/title.jpg?ts=20210819155236",
		"book_edition": ""
	}, {
		"publisher": "培生朗文",
		"term": "上学期",
		"version": "202310220330",
		"full_name": "三年级上新朗文小学英语（Gold升级版）",
		"book_id": "tape3a_023003",
		"grade": "三年级",
		"book_version": "新朗文小学英语LWTE Gold",
		"subject": "英语",
		"cover_url": "https://d.namibox.com/digitbook/tape3a_023003/title.jpg?ts=20210617194907",
		"book_edition": ""
	}, {
		"publisher": "培生朗文",
		"term": "上学期",
		"version": "202310220343",
		"full_name": "三年级上新思维小学英语",
		"book_id": "tape3a_023009",
		"grade": "三年级",
		"book_version": "新思维小学英语NWTE",
		"subject": "英语",
		"cover_url": "https://d.namibox.com/digitbook/tape3a_023009/title.jpg?ts=20210729151733",
		"book_edition": ""
	}, {
		"publisher": "山东科学技术出版社",
		"term": "上学期",
		"version": "202310220355",
		"full_name": "三年级上英语（鲁科版）",
		"book_id": "tape3a_025002",
		"grade": "三年级",
		"book_version": "鲁科版",
		"subject": "英语",
		"cover_url": "https://d.namibox.com/digitbook/tape3a_025002/title.jpg?ts=20220816181740",
		"book_edition": ""
	}, {
		"publisher": "山东科学技术出版社",
		"term": "上学期",
		"version": "202408300308",
		"full_name": "三年级上英语（鲁科版）-24版",
		"book_id": "tape3a_025012",
		"grade": "三年级",
		"book_version": "鲁科版",
		"subject": "英语",
		"cover_url": "https://d.namibox.com/digitbook/tape3a_025012/title.jpg?ts=20240815192026",
		"book_edition": "2024新课标"
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
	_ = json.Unmarshal([]byte(a), &books)

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
					gradeTag := tags[book.Grade] // 拷贝原来的值
					gradeTag.Terms[hasTermIndex].BookVersions = append(gradeTag.Terms[hasTermIndex].BookVersions, book.BookVersion)
					tags[book.Grade] = gradeTag // 更新 map 中的值
				}
			} else {
				gradeTag := tags[book.Grade] // 拷贝原来的值
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
