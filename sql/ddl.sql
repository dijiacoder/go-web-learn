-- 创建数据库
CREATE DATABASE IF NOT EXISTS schooldb;
USE schooldb;

-- 学生表
CREATE TABLE IF NOT EXISTS students
(
    id             INT AUTO_INCREMENT PRIMARY KEY,
    name           VARCHAR(100) NOT NULL COMMENT '学生姓名',
    id_card        VARCHAR(18)  NOT NULL UNIQUE COMMENT '身份证号',
    admission_date DATE         NOT NULL COMMENT '入学日期',
    class_id       INT COMMENT '所属班级',
    birthdate      DATE         NOT NULL COMMENT '生日',
    parent_name    VARCHAR(100) NOT NULL COMMENT '家长姓名',
    parent_phone   VARCHAR(20)  NOT NULL COMMENT '家长联系电话',
    created_at     TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    updated_at     TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    is_del         TINYINT   DEFAULT 0 COMMENT '是否删除（0: 未删除, 1: 已删除）'
);

-- 班级表
CREATE TABLE IF NOT EXISTS classes
(
    id           INT AUTO_INCREMENT PRIMARY KEY,
    name         VARCHAR(100) NOT NULL COMMENT '班级名称',
    grade        VARCHAR(50)  NOT NULL COMMENT '年级',
    head_teacher VARCHAR(100) NOT NULL COMMENT '班主任',
    created_at   TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    updated_at   TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    is_del       TINYINT   DEFAULT 0 COMMENT '是否删除（0: 未删除, 1: 已删除）'
);

-- 添加外键约束
ALTER TABLE students
    ADD CONSTRAINT fk_student_class
        FOREIGN KEY (class_id) REFERENCES classes (id);