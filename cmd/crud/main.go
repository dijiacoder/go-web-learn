package main

import (
	"context"
	"fmt"
	"github.com/dijiacoder/go-web-learn/internal/query"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// 初始化数据库连接
	dsn := "root:123456@tcp(127.0.0.1:3306)/schooldb?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	ctx := context.Background()
	q := query.Use(db).Student

	// Create
	//newStudent := &query.Student{Name: "John", Age: 20}
	//err = q.WithContext(ctx).Create(newStudent)
	//if err != nil {
	//	fmt.Println("Create error:", err)
	//}

	// Read
	students, err := q.WithContext(ctx).Where(q.ID.Eq(1)).Find()
	if err != nil {
		fmt.Println("Read error:", err)
	} else {
		for _, student := range students {
			fmt.Printf("ID: %d, Name: %s, Birthdate: %s, ClassID: %d\n",
				student.ID, student.Name, student.Birthdate, student.ClassID)

			fmt.Println(students)
		}
	}

	// Update
	_, err = q.WithContext(ctx).Where(q.ID.Eq(1)).Update(q.ClassID, 1)
	if err != nil {
		fmt.Println("Update error:", err)
	}
}
