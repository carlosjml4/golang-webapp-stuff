package gopher

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jimmykuu/wtforms"
	"labix.org/v2/mgo/bson"
)

// URL: /books
// 图书列表
func booksHandler(w http.ResponseWriter, r *http.Request) {
	c := DB.C("books")
	var chineseBooks []Book
	c.Find(bson.M{"language": "中文"}).All(&chineseBooks)

	var englishBooks []Book
	c.Find(bson.M{"language": "英文"}).All(&englishBooks)
	renderTemplate(w, r, "book/index.html", map[string]interface{}{
		"chineseBooks": chineseBooks,
		"englishBooks": englishBooks,
		"active":       "books",
	})
}

// URL: /book/{id}
// 显示图书详情
func showBookHandler(w http.ResponseWriter, r *http.Request) {
	bookId := mux.Vars(r)["id"]

	c := DB.C("books")
	var book Book
	c.Find(bson.M{"_id": bson.ObjectIdHex(bookId)}).One(&book)

	renderTemplate(w, r, "book/show.html", map[string]interface{}{
		"book":   book,
		"active": "books",
	})
}

// URL: /admin/book/{id}/edit
// 编辑图书
func editBookHandler(w http.ResponseWriter, r *http.Request) {
	bookId := mux.Vars(r)["id"]

	c := DB.C("books")
	var book Book
	c.Find(bson.M{"_id": bson.ObjectIdHex(bookId)}).One(&book)

	form := wtforms.NewForm(
		wtforms.NewTextField("title", "书名", book.Title, wtforms.Required{}),
		wtforms.NewTextField("cover", "封面", book.Cover, wtforms.Required{}),
		wtforms.NewTextField("author", "作者", book.Author, wtforms.Required{}),
		wtforms.NewTextField("translator", "译者", book.Translator),
		wtforms.NewTextArea("introduction", "简介", book.Introduction),
		wtforms.NewTextField("pages", "页数", strconv.Itoa(book.Pages), wtforms.Required{}),
		wtforms.NewTextField("language", "语言", book.Language, wtforms.Required{}),
		wtforms.NewTextField("publisher", "出版社", book.Publisher),
		wtforms.NewTextField("publication_date", "出版年月日", book.PublicationDate),
		wtforms.NewTextField("isbn", "ISBN", book.ISBN),
	)

	if r.Method == "POST" {
		if form.Validate(r) {
			pages, _ := strconv.Atoi(form.Value("pages"))

			err := c.Update(bson.M{"_id": book.Id_}, bson.M{"$set": bson.M{
				"title":            form.Value("title"),
				"cover":            form.Value("cover"),
				"author":           form.Value("author"),
				"translator":       form.Value("translator"),
				"introduction":     form.Value("introduction"),
				"pages":            pages,
				"language":         form.Value("language"),
				"publisher":        form.Value("publisher"),
				"publication_date": form.Value("publication_date"),
				"isbn":             form.Value("isbn"),
			}})

			if err != nil {
				panic(err)
			}

			http.Redirect(w, r, "/admin/books", http.StatusFound)
			return
		}
	}

	renderTemplate(w, r, "book/form.html", map[string]interface{}{
		"adminNav": ADMIN_NAV,
		"book":     book,
		"form":     form,
		"isNew":    false,
	})
}

// URL: /book/{id}/delete
// 删除图书
func deleteBookHandler(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	c := DB.C("books")
	c.RemoveId(bson.ObjectIdHex(id))

	w.Write([]byte("true"))
}

func listBooksHandler(w http.ResponseWriter, r *http.Request) {
	c := DB.C("books")
	var books []Book
	c.Find(nil).All(&books)

	renderTemplate(w, r, "book/list.html", map[string]interface{}{
		"adminNav": ADMIN_NAV,
		"books":    books,
	})
}

func newBookHandler(w http.ResponseWriter, r *http.Request) {
	form := wtforms.NewForm(
		wtforms.NewTextField("title", "书名", "", wtforms.Required{}),
		wtforms.NewTextField("cover", "封面", "", wtforms.Required{}),
		wtforms.NewTextField("author", "作者", "", wtforms.Required{}),
		wtforms.NewTextField("translator", "译者", ""),
		wtforms.NewTextArea("introduction", "简介", ""),
		wtforms.NewTextField("pages", "页数", "", wtforms.Required{}),
		wtforms.NewTextField("language", "语言", "", wtforms.Required{}),
		wtforms.NewTextField("publisher", "出版社", ""),
		wtforms.NewTextField("publication_date", "出版年月日", ""),
		wtforms.NewTextField("isbn", "ISBN", ""),
	)

	if r.Method == "POST" {
		if form.Validate(r) {
			pages, _ := strconv.Atoi(form.Value("pages"))
			c := DB.C("books")
			err := c.Insert(&Book{
				Id_:             bson.NewObjectId(),
				Title:           form.Value("title"),
				Cover:           form.Value("cover"),
				Author:          form.Value("author"),
				Translator:      form.Value("translator"),
				Pages:           pages,
				Language:        form.Value("language"),
				Publisher:       form.Value("publisher"),
				PublicationDate: form.Value("publication_date"),
				Introduction:    form.Value("introduction"),
				ISBN:            form.Value("isbn"),
			})

			if err != nil {
				panic(err)
			}
			http.Redirect(w, r, "/admin/books", http.StatusFound)
			return
		}
	}

	renderTemplate(w, r, "book/form.html", map[string]interface{}{
		"adminNav": ADMIN_NAV,
		"form":     form,
		"isNew":    true,
	})
}
