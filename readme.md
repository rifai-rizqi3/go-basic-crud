- main.go
  - func main
  ```
  func main() {
	config.ConnectDB()

	// 1. Homepage
	http.HandleFunc("/", homecontroller.Welcome)

	log.Println("Starting server on :9003")
	http.ListenAndServe(":9003", nil)
  }
  ```

- Folder
  - models (komunikasi db)
  - views (untuk menampilkan web)
  - controllers
  - config
  - entities (file struck)
- Folder
  - install go-sql-driver
  - import code
  - setting username & password
    - config
      - database.go
    
      ```
      var DB *sql.DB

      func ConnectDB() {
          db, err := sql.Open("mysql", "rnrifai:RNRif@i1212@/go_products")
          if err != nil {
          panic(err)
      }

      log.Println("Connected to database")
      DB = db
      }
      ```
- controllers
  - homecontroller
    - homecontroller.go
    ```
    func Welcome(w http.ResponseWriter, r *http.Request) {
	    temp, err := template.ParseFiles("views/home/index.html")
	    if err != nil {
		    panic(err)
	    }

	    temp.Execute(w, nil)
    }

    ```
- views
  - home
    - index.html
      - masukan template bootsrap
  - main
  ```
    http.HandleFunc("/categories", categorycontroller.Index)
	http.HandleFunc("/categories/add", categorycontroller.Add)
	http.HandleFunc("/categories/edit", categorycontroller.Edit)
	http.HandleFunc("/categories/delete", categorycontroller.Delete)
  ```
- contollers
  - categorycontroller
    - categorycontroller.go
    ```
    func Index(w http.ResponseWriter, r *http.Request) {

    }

    func Add(w http.ResponseWriter, r *http.Request) {

    }

    func Edit(w http.ResponseWriter, r *http.Request) {

    }

    func Delete(w http.ResponseWriter, r *http.Request) {

    }
    ```
    
- entities
  - category.go
