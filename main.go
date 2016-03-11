package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {
	//Open DB
	db, err := sql.Open("sqlite3", "./person.db")
	if err != nil {
		fmt.Println("打开数据库出错")
		panic(err)
	}

	//插入数据
	stmt, err := db.Prepare("INSERT INTO person(firstname,lastname,birthday) values(?,?,?)")
	if err != nil {
		panic(err)
	}

	for i := 0; i < 20; i++ {
		firstname, lastname := "li"+strconv.Itoa(i), "huisong"+strconv.Itoa(i)
		birthday := randomBirthday()
		stmt.Exec(firstname, lastname, birthday)
	}

	//导出数据
	rows, err := db.Query("SELECT * FROM person")
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		var idcard int
		var firstname string
		var lastname string
		var birthday time.Time

		rows.Scan(&idcard, &firstname, &lastname, &birthday)
		fmt.Printf("身份证: %d\n", idcard)
		fmt.Printf("姓: %s\n", firstname)
		fmt.Printf("名字: %s\n", lastname)
		fmt.Printf("生日: %s\n", birthday.Format("2006-01-02"))
		//todo
		save(idcard, firstname, lastname)
	}

}
func randomBirthday() time.Time {
	year := rand.Intn(1015) + 1000
	month := rand.Intn(1) + 10
	day := "12"
	birthdaystr := strconv.Itoa(year) + "-" + strconv.Itoa(month) + "-" + day
	t, err := time.Parse("2016-01-02 15:04:02", birthdaystr)
	//日期处理有问题
	if err != nil {
		fmt.Println(err)
	}
	return t
}
func save(id int, firstname string, lastname string) {
	fileName := "persons.txt"
	var fout *os.File
	var err error
	if checkFileIsExist(fileName) {
		fmt.Println("文件不村子啊")
		fout, err = os.Create(fileName)
		if err != nil {
			fmt.Println(fileName, err)
		}
	} else {
		fout, err = os.OpenFile(fileName, os.O_APPEND, 0666)
		if err != nil {
			fmt.Println(fileName, err)
		}
	}
	defer fout.Close()
	fout.WriteString("id:" + strconv.Itoa(id) + " 姓:" + firstname + " 名:" + lastname + "\n")
}

/**
 * 判断文件是否存在  存在返回 true 不存在返回false
 */
func checkFileIsExist(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil || os.IsExist(err)
}
