package main

import (
    "fmt"
    "time"
    "github.com/zenazn/goji/web"
    "net/http"
    "html/template"
    "./entity"
)

type EmployeesForm struct {
    Title string
    Now string
    EmployeeList []entity.Employee
}

func employees(c web.C, w http.ResponseWriter, r *http.Request) {

    tpl, err1 := template.ParseFiles("view/employees.html")
    if err1 != nil {
        fmt.Fprintf(w, "%#v", err1)
        panic(err1.Error)
    }


    employeeList := []entity.Employee{

        entity.Employee{
            EmployeeId: 1,
            LoginId: "hideki.matsui",
            Ename: "Matsui, Hideki",
            Jname: "松居　英紀",
            Email: "hideki.matsui@hogehoge.dom",
        },
        entity.Employee{
            EmployeeId: 2,
            LoginId: "ichiro.suzuki",
            Ename: "Suzuki, Ichirou",
            Jname: "寿々木　市郎",
            Email: "ichirou.suzuki@hogehoge.dom",
        },
    }

    param := EmployeesForm{
       Title: "社員一覧",
       Now: time.Now().Format("2006/01/02 03:04:05"),
       EmployeeList: employeeList,
    }

    fmt.Printf("### DEBUG ###")
    fmt.Printf("%#v", param)


    err2 := tpl.Execute(w, param)

    if err2 != nil {
        fmt.Fprintf(w, "%#v", err2)
        panic(err2.Error)
    }
}
