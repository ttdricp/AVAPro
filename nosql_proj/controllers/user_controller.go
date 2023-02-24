package controllers

import (
	"context"
	"fmt"
	"github.com/DigitalTroller/nosql_proj/configs"
	"github.com/DigitalTroller/nosql_proj/models"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"reflect"
	"strconv"
	"strings"
	"time"

	helper "github.com/DigitalTroller/nosql_proj/helpers"
	"github.com/gin-gonic/gin"
)

var userCollection *mongo.Collection = configs.GetCollection(configs.DB, "users")
var taskCollection *mongo.Collection = configs.GetCollection(configs.DB, "tasks")
var validate = validator.New()

//func CreateUser() http.HandlerFunc {
//	return func(rw http.ResponseWriter, r *http.Request) {
//		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
//		var user models.User
//		defer cancel()
//
//		//validate the request body
//		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
//			rw.WriteHeader(http.StatusBadRequest)
//			response := responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}}
//			json.NewEncoder(rw).Encode(response)
//			return
//		}
//
//		//use the validator library to validate required fields
//		if validationErr := validate.Struct(&user); validationErr != nil {
//			rw.WriteHeader(http.StatusBadRequest)
//			response := responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}}
//			json.NewEncoder(rw).Encode(response)
//			return
//		}
//
//		newUser := models.User{
//			Id:       primitive.NewObjectID(),
//			First_name:    user.First_name,
//			Last_name:     *string  ,
//			Password:      *string  ,
//			Email:         *string  ,
//			Phone         *string  ,
//			Token         *string  ,
//			Refresh_token *string  ,
//			Created_at    time.Time,
//			Updated_at    time.Time,
//			User_id       string   ,
//			Name:     user.Name,
//			Location: user.Location,
//			Title:    user.Title,
//		}
//		result, err := userCollection.InsertOne(ctx, newUser)
//		if err != nil {
//			rw.WriteHeader(http.StatusInternalServerError)
//			response := responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}}
//			json.NewEncoder(rw).Encode(response)
//			return
//		}
//
//		rw.WriteHeader(http.StatusCreated)
//		response := responses.UserResponse{Status: http.StatusCreated, Message: "success", Data: map[string]interface{}{"data": result}}
//		json.NewEncoder(rw).Encode(response)
//	}
//}

func GetAUser(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var user models.UserZero
	defer cancel()
	cookie := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJFbWFpbCI6ImFiYXlAbWFpbC5ydSIsIkZpcnN0X25hbWUiOiJBYmF5IiwiTGFzdF9uYW1lIjoiTXVraHRhcnVseSIsIlVpZCI6IjYzZjZlMTcxNjNkN2MxZTIwZmE4NDNjNSIsImV4cCI6MTY3NzMyMDQ3NX0.z_rep7LeF7gmf1301qkBJ35qmpgYiJNJSVURZtZz84k"
	fmt.Println(cookie)
	err := userCollection.FindOne(ctx, bson.M{"token": cookie}).Decode(&user)
	if err != nil {
		//log.Panic(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error occured while checking for the admin restriction"})
		return
	}
	fmt.Println(&user.Role)

	c.JSON(http.StatusOK, user.User_id)
	return
}

//func EditAUser() http.HandlerFunc {
//	return func(rw http.ResponseWriter, r *http.Request) {
//		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
//		params := mux.Vars(r)
//		userId := params["userId"]
//		var user models.User
//		defer cancel()
//
//		objId, _ := primitive.ObjectIDFromHex(userId)
//
//		//validate the request body
//		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
//			rw.WriteHeader(http.StatusBadRequest)
//			response := responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}}
//			json.NewEncoder(rw).Encode(response)
//			return
//		}
//
//		//use the validator library to validate required fields
//		if validationErr := validate.Struct(&user); validationErr != nil {
//			rw.WriteHeader(http.StatusBadRequest)
//			response := responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}}
//			json.NewEncoder(rw).Encode(response)
//			return
//		}
//
//		update := bson.M{"name": user.Name, "location": user.Location, "title": user.Title}
//
//		result, err := userCollection.UpdateOne(ctx, bson.M{"_id": objId}, bson.M{"$set": update})
//		if err != nil {
//			rw.WriteHeader(http.StatusInternalServerError)
//			response := responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}}
//			json.NewEncoder(rw).Encode(response)
//			return
//		}
//
//		//get updated user details
//		var updatedUser models.User
//		if result.MatchedCount == 1 {
//			err := userCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&updatedUser)
//
//			if err != nil {
//				rw.WriteHeader(http.StatusInternalServerError)
//				response := responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}}
//				json.NewEncoder(rw).Encode(response)
//				return
//			}
//		}
//
//		rw.WriteHeader(http.StatusOK)
//		response := responses.UserResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": updatedUser}}
//		json.NewEncoder(rw).Encode(response)
//	}
//
//}

//func DeleteAUser() http.HandlerFunc {
//	return func(rw http.ResponseWriter, r *http.Request) {
//		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
//		params := mux.Vars(r)
//		userId := params["userId"]
//		defer cancel()
//		objId, _ := primitive.ObjectIDFromHex(userId)
//
//		result, err := userCollection.DeleteOne(ctx, bson.M{"_id": objId})
//
//		if err != nil {
//			rw.WriteHeader(http.StatusInternalServerError)
//			response := responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}}
//			json.NewEncoder(rw).Encode(response)
//			return
//		}
//
//		if result.DeletedCount < 1 {
//			rw.WriteHeader(http.StatusNotFound)
//			response := responses.UserResponse{Status: http.StatusNotFound, Message: "error", Data: map[string]interface{}{"data": "User with specified ID not found!"}}
//			json.NewEncoder(rw).Encode(response)
//			return
//		}
//
//		rw.WriteHeader(http.StatusOK)
//		response := responses.UserResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": "User successfully deleted!"}}
//		json.NewEncoder(rw).Encode(response)
//	}
//
//}

//func GetAllUser() http.HandlerFunc {
//	return func(rw http.ResponseWriter, r *http.Request) {
//		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
//		var users []models.User
//		defer cancel()
//
//		results, err := userCollection.Find(ctx, bson.M{})
//
//		if err != nil {
//			rw.WriteHeader(http.StatusInternalServerError)
//			response := responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}}
//			json.NewEncoder(rw).Encode(response)
//			return
//		}
//
//		//reading from the db in an optimal way
//		defer results.Close(ctx)
//		for results.Next(ctx) {
//			var singleUser models.User
//			if err = results.Decode(&singleUser); err != nil {
//				rw.WriteHeader(http.StatusInternalServerError)
//				response := responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}}
//				json.NewEncoder(rw).Encode(response)
//			}
//
//			users = append(users, singleUser)
//		}
//
//		rw.WriteHeader(http.StatusOK)
//		response := responses.UserResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": users}}
//		json.NewEncoder(rw).Encode(response)
//	}
//
//}

func HashPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		log.Panic(err)
	}

	return string(bytes)
}

// VerifyPassword checks the input password while verifying it with the passward in the DB.
func VerifyPassword(userPassword string, providedPassword string) (bool, string) {
	err := bcrypt.CompareHashAndPassword([]byte(providedPassword), []byte(userPassword))
	check := true
	msg := ""

	if err != nil {
		msg = fmt.Sprintf("login or passowrd is incorrect")
		check = false
	}

	return check, msg
}

func Main() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func SignUp() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var user models.User
		if err := c.Bind(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		validationErr := validate.Struct(user)
		if validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
			return
		}
		count, err := userCollection.CountDocuments(ctx, bson.M{"email": user.Email})
		defer cancel()
		if err != nil {
			log.Panic(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occured while checking for the email"})
			return
		}

		password := HashPassword(*user.Password)
		user.Password = &password
		role := "User"
		user.Role = &role
		count, err = userCollection.CountDocuments(ctx, bson.M{"phone": user.Phone})
		defer cancel()
		if err != nil {
			log.Panic(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occured while checking for the phone number"})
			return
		}

		if count > 0 {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "this email or phone number already exists"})
			return
		}

		user.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		user.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		user.ID = primitive.NewObjectID()
		user.User_id = user.ID.Hex()
		token, refreshToken, _ := helper.GenerateAllTokens(*user.Email, *user.First_name, *user.Last_name, user.User_id)
		user.Token = &token
		user.Refresh_token = &refreshToken
		resultInsertionNumber, insertErr := userCollection.InsertOne(ctx, user)
		if insertErr != nil {
			msg := fmt.Sprintf("User item was not created")
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}
		defer cancel()

		fmt.Println(resultInsertionNumber)

		//c.JSON(http.StatusOK, resultInsertionNumber)
		c.Request.Header.Set("token", token)
		c.Redirect(http.StatusFound, "/login")
	}
}

func signIn(writer http.ResponseWriter, request *http.Request) {
	// your business logic here
	http.SetCookie(writer, &http.Cookie{
		Name:  "signed in cookie",
		Value: "True",
	})
}

// Login is the api used to tget a single user
func Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var user models.UserLogin
		var foundUser models.UserFound
		if err := c.Bind(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}

		err := userCollection.FindOne(ctx, bson.M{"email": user.Email}).Decode(&foundUser)
		defer cancel()
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "login or passowrd is incorrect"})
			return
		}

		passwordIsValid, msg := VerifyPassword(*user.Password, *foundUser.Password)
		defer cancel()
		if passwordIsValid != true {
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}

		token, refreshToken, _ := helper.GenerateAllTokens(*foundUser.Email, *foundUser.First_name, *foundUser.Last_name, foundUser.User_id)

		helper.UpdateAllTokens(token, refreshToken, foundUser.User_id)

		c.SetCookie("token", token, 3600, "/task", "localhost", false, true)
		c.SetCookie("refreshToken", refreshToken, 3600, "/task", "localhost", false, true)
		c.Request.Header.Set("token", token)
		fmt.Println(c.GetHeader("token"))
		cookie, _ := c.Request.Cookie("token")
		refreshcookie, _ := c.Request.Cookie("refreshToken")
		fmt.Println(cookie)
		fmt.Println(refreshcookie)
		return
		//c.Writer.Header().Set("token", token)
		//fmt.Println(c.GetHeader("token"))
		//c.JSON(http.StatusOK, foundUser)
		//c.Next()
		//c.Writer.Header().Set("token", token)
	}
}
func GetDataA(c *gin.Context) time.Time {
	var timme models.StructA
	c.Bind(&timme)
	c.JSON(200, gin.H{
		"TillDate": timme.FieldA,
	})
	date := timme.FieldA
	correctStrDate := strings.Replace(date, "T", " ", 1)
	correctTime, _ := time.Parse("2006-01-02 15:04", correctStrDate)
	fmt.Println(date)
	fmt.Println(correctTime)
	fmt.Println(reflect.TypeOf(correctTime))

	return correctTime
}

func sPtr(s string) *string        { return &s }
func sPtri(s int) *int             { return &s }
func sPtrt(s time.Time) *time.Time { return &s }

func AddNewTask() gin.HandlerFunc {
	return func(c *gin.Context) {
		//cookie := c.GetHeader("token")
		//fmt.Println(cookie)
		GetAUser(c)
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var task models.Task
		Time := GetDataA(c)
		//*task.Till_date = time
		Name := c.PostForm("Task name")
		task.Name = sPtr(Name)
		fmt.Println(task.Name)
		Category := c.PostForm("category")
		task.Category = sPtr(Category)
		fmt.Println(task.Category)
		Description := c.PostForm("Description")
		task.Description = sPtr(Description)
		fmt.Println(task.Description)
		Points, _ := strconv.Atoi(c.PostForm("Points"))
		task.Points = sPtri(Points)
		fmt.Println(task.Points)
		task.Till_date = sPtrt(Time)
		fmt.Println(task.Till_date)

		fmt.Println(task.Task_id)

		task.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		task.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		task.ID = primitive.NewObjectID()
		task.Task_id = task.ID.Hex()

		resultInsertionNumber, insertErr := taskCollection.InsertOne(ctx, task)
		if insertErr != nil {
			msg := fmt.Sprintf("Task item was not created")
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}
		defer cancel()

		fmt.Println(resultInsertionNumber)

		//c.JSON(http.StatusOK, resultInsertionNumber)
		c.Redirect(http.StatusFound, "/task")
	}
}

func GetTasks() gin.HandlerFunc {
	return func(c *gin.Context) {
		var tasks []models.Task
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		cursor, err := taskCollection.Find(ctx, bson.D{})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occured while fetching the tasks"})
			return
		}
		fmt.Println(tasks)

		defer cursor.Close(ctx)

		for cursor.Next(ctx) {
			var task models.Task
			if err := cursor.Decode(&task); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "error occured while decoding the tasks"})
				return
			}
			tasks = append(tasks, task)
		}

		if err := cursor.Err(); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occured while iterating the tasks"})
			return
		}

		c.HTML(http.StatusOK, "tasks.html", gin.H{
			"tasks": tasks,
		})

	}
}

func TaskExec() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the task ID from the URL parameter
		taskID := c.Param("id")

		// Find the task with the specified ID in the database
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		var task models.Task
		err := taskCollection.FindOne(ctx, bson.M{"task_id": taskID}).Decode(&task)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "task not found"})
			return
		}

		// Render the task_exec.html page with the task details
		c.HTML(http.StatusOK, "task_exec.html", gin.H{
			"title": task.Name,
			"task":  task,
		})
	}
}

func FileUpload() gin.HandlerFunc {
	return func(c *gin.Context) {
		file, err := c.FormFile("file")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		// Save the file to disk
		if err := c.SaveUploadedFile(file, "./uploads/"+file.Filename); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		// Return a success message
		c.JSON(http.StatusOK, gin.H{
			"message": "File uploaded successfully",
		})
	}
}
