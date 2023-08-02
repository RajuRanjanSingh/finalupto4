package main

import (
	"log"
	"github.com/go-pg/pg/v9"
	"github.com/gin-gonic/gin"
	"net/http"
	orm "github.com/go-pg/pg/v9/orm"
	//"github.com/jackc/pgtype"
	"github.com/gin-contrib/cors"
	"os"
	"time"
	

)
var dbUser = os.Getenv("DB_USER")
var dbPass = os.Getenv("DB_PASS")
var dbaddr = os.Getenv("DB_ADDR")
var dbDatabase = os.Getenv("DB_DATABASE")


// Connecting to db
func connect() *pg.DB {
	opts := &pg.Options{
		User: dbUser,
		Password: dbPass,	
		Addr: dbaddr,
		Database: dbDatabase,
	}
	var db *pg.DB = pg.Connect(opts)
	if db == nil {
		log.Printf("Failed to connect")
	}
	log.Printf("Connected to db")
	createTable(db)
	initiateDB(db)
	return db
}

type Employee struct {
	Id int `json:"id" pg:"id"`
	Name     string  `json:"name"`
	LeaveType   string  `json:"leave_type"`
    Fromdate time.Time  `json:"fromdate"`
    Todate string `json:"todate"`
	Team_Name   string    `json:"team_name" `
	File_upload string `json:"file_upload"`
	Reporter    string `json:"reporter" `
}


// Create User Table
func createTable(db *pg.DB) {
	opts := &orm.CreateTableOptions{
	IfNotExists: true,
	}
	createError := db.CreateTable(&Employee{}, opts)
	if createError != nil {
		log.Printf("Error while creating employee table")
	}
	log.Printf("table created")
}


// INITIALIZE DB CONNECTION 
var dbConnect *pg.DB
func initiateDB(db *pg.DB) {
	dbConnect = db
}

func getTable(c *gin.Context){


	var Employees []Employee
	err := dbConnect.Model(&Employees).Select()
    if err != nil {
	log.Printf("Error while getting all leave form")
	c.JSON(http.StatusInternalServerError, gin.H{
		"status":  http.StatusInternalServerError,
		"message": "Something went wrong",
		})
	return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "All employee table",
		"data": Employees,
	})
	return
}

type View struct {
    NAME       string  `json:"name"`
    TotalLeave float64 `json:"total_leave"`
}
func getView(c *gin.Context) {
    var views []View
    query := `
	SELECT name, COUNT(*) AS total_leave
FROM employees
WHERE DATE_PART('year', fromdate) = 2023
GROUP BY name
ORDER BY total_leave DESC
LIMIT 5;
    `
    _, err := dbConnect.Query(&views, query)
    if err != nil {
        log.Printf("Error while executing the view query: %v", err)
        c.JSON(http.StatusInternalServerError, gin.H{
            "status":  http.StatusInternalServerError,
            "message": "Something went wrong",
        })
        return
    }
    c.JSON(http.StatusOK, gin.H{
        "status":  http.StatusOK,
        "message": "Data from the view",
        "data":    views,
    })
}


type Manager struct {
	Manager    string  `json:"manager"`
    TotalLeave float64 `json:"total_leave"`
}
func getKPI4(c *gin.Context) {
    var Managers []Manager
    query := `
	SELECT reporter as Manager, COUNT(*) AS total_leave
FROM employees
WHERE fromdate >= '2023-01-01'
  AND fromdate <= '2023-03-31'
GROUP BY reporter
ORDER BY total_leave;
    `
    _, err := dbConnect.Query(&Managers, query)
    if err != nil {
        log.Printf("Error while executing the view query: %v", err)
        c.JSON(http.StatusInternalServerError, gin.H{
            "status":  http.StatusInternalServerError,
            "message": "Something went wrong",
        })
        return
    }
    c.JSON(http.StatusOK, gin.H{
        "status":  http.StatusOK,
        "message": "Data from the view",
        "data":    Managers,
    })
} 



func getKPI6(c *gin.Context) {
	var Team []struct {
		Team_Name string  
		Leave_type string 
		Leave_count int 
	}
    //var Teams []Team
    query := `
SELECT team__name, leave_type, COUNT(*) AS leave_count
FROM employees
WHERE DATE_PART('year', fromdate) = 2022
  AND team__name IN (
    SELECT team__name
    FROM employees
    WHERE DATE_PART('year', fromdate) = 2022
    GROUP BY team__name
    ORDER BY COUNT(*) DESC
    LIMIT 2
  )
GROUP BY team__name, leave_type;
    `
    _, err := dbConnect.Query(&Team, query)
    if err != nil {
        log.Printf("Error while executing the view query: %v", err)
        c.JSON(http.StatusInternalServerError, gin.H{
            "status":  http.StatusInternalServerError,
            "message": "Something went wrong",
        })
        return
    }
    c.JSON(http.StatusOK, gin.H{
        "status":  http.StatusOK,
        "message": "Data from the view",
        "data":    Team,
    })
} 

func postLeave(c *gin.Context) {
    var employees Employee
	c.BindJSON(&employees)
    insertError := dbConnect.Insert(&employees)
    if insertError != nil {
	log.Printf("Error while inserting new employee into db, Reason: %v\n", insertError)
	c.JSON(http.StatusInternalServerError, gin.H{
		"status":  http.StatusInternalServerError,
		"message": "Something went wrong",
	})
	return
}
	c.JSON(http.StatusCreated, gin.H{
		"status":  http.StatusCreated,
		"message": "Table created Successfully",	
	})
	return
}

func routes(router *gin.Engine) {
	router.GET("/", welcome)
	router.GET("/get", getTable)
	router.POST("/post",postLeave)
	router.GET("/getView", getView)
	router.GET("/getkpi4", getKPI4)
	router.GET("/getkpi6", getKPI6)
}

func welcome(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "Welcome To API",
	})
	return
}


func main() {
	connect()                     // Connect DB
	router := gin.Default()   
	router.Use(cors.Default())
	routes(router)                // Route Handlers 
	log.Fatal(router.Run(":4747"))
}


/*  {
      "name": "Raju Ranjan Singh",
      "leave_type": "sickleave",
      "fromdate": "2023-06-21",
      "todate": "2023-06-25",
      "team_name": "Designops",
      "file_upload": "medical.png",
      "reporter": "Sahil sir"
    }
*/