package main

import (
	"strconv"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	//数据库连接池
	"github.com/360EntSecGroup-Skylar/excelize"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	url2 := "xdnphb_test:ZK6MleC5X6wR@tcp(122.114.46.81:3306)/information_schema?charset=utf8mb4&parseTime=True"
	db2, _ := gorm.Open(mysql.Open(url2), &gorm.Config{})

	url := "xdnphb_test:ZK6MleC5X6wR@tcp(122.114.46.81:3306)/phb_test_20150820?charset=utf8mb4&parseTime=True"
	db, err := gorm.Open(mysql.Open(url), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.Logger.LogMode(1)
	//修改连接池复用最大时间
	sqlDB, err := db.DB()
	sqlDB.SetConnMaxLifetime(time.Hour * 5)
	rows, err := db.Raw("SHOW TABLES;").Rows()
	if err != nil {
		panic("failed to connect database")
	}
	f := excelize.NewFile()
	index := f.NewSheet("Sheet2")
	i := 1

	r := "select round(sum(DATA_LENGTH/1024/1024),2) as data from TABLES where table_schema='phb_test_20150820' and table_name=?;"
	for rows.Next() {
		var table, size string
		err := rows.Scan(&table)
		if err != nil {
			panic("failed to connect database")
		}
		d := db2.Raw(r, table).Row()
		d.Scan(&size)
		f.SetCellValue("Sheet2", "A"+strconv.Itoa(i), table)
		f.SetCellValue("Sheet2", "B"+strconv.Itoa(i), size)
		i++
	}
	f.SetActiveSheet(index)
	if err := f.SaveAs("d://Book1.xlsx"); err != nil {
		println(err.Error())
	}

}
