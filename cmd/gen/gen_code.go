package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
	"gorm.io/plugin/dbresolver"
)

func main() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/schooldb?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("数据库连接失败: " + err.Error())
	}

	err = db.Use(dbresolver.Register(dbresolver.Config{
		Sources:  []gorm.Dialector{mysql.Open(dsn)},
		Replicas: []gorm.Dialector{mysql.Open(dsn)},
		Policy:   dbresolver.RandomPolicy{},
	}))
	if err != nil {
		panic("dbresolver 插件初始化失败: " + err.Error())
	}

	// 修改后的生成器配置
	g := gen.NewGenerator(gen.Config{
		OutPath:       "D:\\Web3\\go-web-learn\\internal\\models\\query",
		ModelPkgPath:  "D:\\Web3\\go-web-learn\\internal\\models\\entity",
		Mode:          gen.WithDefaultQuery | gen.WithQueryInterface | gen.WithoutContext,
		FieldNullable: true,
	})

	g.UseDB(db)

	// 明确指定要生成的模型
	g.ApplyBasic(g.GenerateModel("students"),
		g.GenerateModel("classes"))

	g.Execute()

	fmt.Println("模型和查询代码生成成功")
}
