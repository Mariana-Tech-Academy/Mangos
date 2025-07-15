package handlers

import (
	_"library/services"
	_"net/http"
)

// type BookHandler struct {
// 	Service *services.BookService
// }

// func createBook(w http.ResponseWriter, r *http.Request) 
// var req *models.library
// 	err = json.NewDecoder(r.Body).Decode(&req)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return
// 	}
// 	userID := r.URL.Query().Get("bookID")  
// 	if req.Title == "" || req.Text == "" {
// 		http.Error(w, "title, text required", http.StatusBadRequest)
// 		return
// 	}
	
// 	req.UserID = userID
// 	req.ID = uuid.NewString()
// 	err = database.Db.Create(&req).Error
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}
// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusOK)
// 	json.NewEncoder(w).Encode(req)

// func updateBook(w http.ResponseWriter, r *http.Request) 


// func deleteBook(w http.ResponseWriter, r *http.Request) 







/*
create book - admin
delete book - admin
update book - admin

list All Books - user
get Book By ID - user
list My Books - user
chechout book - user
check in book - user
find book by genre
find book by author
find book by year
find book by title
*/