package handlers

import "library/services"

type BookHandler struct {
	Service *services.BookService
}
func (h *BookHandler) FindByGenre(w http.ResponseWriter, r *http.Request){
genre := r.URL.Query(.Get("genre")

books, err := h.service.FindByGenre(genre)
if err != nil{
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return 
}
json.NewEncoder(w).Encode(books)
}

/*
Group 1
list All Books - user
get Book By ID - user
list My Books - user
chechout book - user
check in book - user

Group 2
create book - admin    
delete book - admin
update book - admin


Group 3
find book by genre
find book by author
find book by year
find book by title


Group 4
register user
login user
user info
*/