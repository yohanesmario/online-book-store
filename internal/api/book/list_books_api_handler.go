package book

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/yohanesmario/online-book-store/internal/api/common/auth"
	"github.com/yohanesmario/online-book-store/internal/api/common/data"
	"github.com/yohanesmario/online-book-store/internal/api/common/errorcode"
	"github.com/yohanesmario/online-book-store/internal/api/common/res"
	"github.com/yohanesmario/online-book-store/internal/service/book"
)

type listBooksResponsePayload res.BaseResponse[listBooksResponseData]

type listBooksResponseData struct {
	Books  []data.BookData `json:"books"`
	LastID int32           `json:"lastId"`
}

// List Books godoc
//
// @Summary     List Books
// @Description List all books, paginated by lastId and limit.
// @Tags        Book APIs
// @Produce     json
//
// @Param  lastId  query  int32  false  "lastId of the fetched books. Default to 0"
// @Param   limit   query  int    false  "number of books to fetch. Default to 10"
//
// @Success  200  {object}  listBooksResponsePayload  "List of books returned. Will return lastId=-1 if no more books to fetch"
// @Failure  400  {object}  res.ErrorResponse         "Input can't be parsed"
// @Failure  500  {object}  res.ErrorResponse         "Internal server error"
//
// @Security  BearerAuth
// @Router    /api/v1/books  [get]
func RegisterListBooksAPI(app *fiber.App, bookService book.BookService) {
	app.Get("/api/v1/books", auth.JWTAuth, func(c *fiber.Ctx) error {
		lastID := int32(c.QueryInt("lastId", 0))
		limit := c.QueryInt("limit", 10)

		books, err := bookService.GetBooksByLastID(lastID, limit)
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(res.ErrorResponse{
				Error: res.ErrorData{
					Code:    errorcode.InvalidInput,
					Message: "Internal server error",
				},
			})
		}

		lastId := int32(-1)
		bookList := make([]data.BookData, len(books))
		if len(books) > 0 {
			for i, b := range books {
				if b.ID > lastId {
					lastId = b.ID
				}
				bookList[i] = data.BookData{
					ID:     b.ID,
					Title:  b.Title,
					Author: b.Author,
					Isbn:   b.Isbn,
					Price:  &b.Price,
				}
			}
		}

		return c.Status(http.StatusOK).JSON(listBooksResponsePayload{
			Data: listBooksResponseData{
				Books:  bookList,
				LastID: lastId,
			},
		})
	})
}
