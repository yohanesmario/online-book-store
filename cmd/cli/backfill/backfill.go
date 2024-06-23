package main

import (
	"encoding/json"
	"math/rand"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	fxlogrus "github.com/takt-corp/fx-logrus"
	"github.com/yohanesmario/online-book-store/conf"
	"github.com/yohanesmario/online-book-store/internal/domain/model"
	"github.com/yohanesmario/online-book-store/internal/impl/gorm/gormdb"
	"github.com/yohanesmario/online-book-store/internal/impl/gorm/repository"
	"github.com/yohanesmario/online-book-store/internal/service"
	"github.com/yohanesmario/online-book-store/internal/service/book"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
)

func main() {
	fx.New(
		fx.WithLogger(func() fxevent.Logger {
			logrus.SetOutput(os.Stdout)
			logrus.SetLevel(conf.Configuration.LogLevel)
			logrus.SetFormatter(conf.Configuration.LogFormatter)
			return &fxlogrus.LogrusLogger{Logger: logrus.StandardLogger()}
		}),
		gormdb.GormDBModule,
		repository.RepositoryImplModule,
		service.ServiceModule,

		fx.Invoke(backfill),
	)
}

func backfill(bookService book.BookService) {
	total := 1000
	minPrice := 100_000
	maxPrice := 1_000_000 - minPrice
	numberRunes := []rune("0123456789")
	books := make([]*model.Book, total)
	modifier := "BACKFILL_CLI"
	now := time.Now()
	for i := 0; i < total; i++ {
		isbnRunes := make([]rune, 13)
		for j := range isbnRunes {
			isbnRunes[j] = numberRunes[rand.Intn(len(numberRunes))]
		}
		books[i] = &model.Book{
			Title:     "Book Title " + uuid.New().String(),
			Author:    "Author " + uuid.New().String(),
			Isbn:      formatIsbn(isbnRunes),
			Price:     float64(rand.Intn(maxPrice) + minPrice),
			CreatedBy: &modifier,
			CreatedAt: &now,
			UpdatedBy: &modifier,
			UpdatedAt: &now,
		}
	}
	booksJson, err := json.MarshalIndent(books, "", "  ")
	if err != nil {
		logrus.Errorf("Failed to marshal books: %v", err)
		os.Exit(1)
	}
	logrus.Infof("Books: %s", booksJson)
	logrus.Infof("Creating %d books...", total)
	_, err = bookService.CreateBooks(books)
	if err != nil {
		logrus.Errorf("Failed to create books: %v", err)
		os.Exit(1)
	}
	logrus.Infof("Books created.")
}

func formatIsbn(runes []rune) string {
	return string(runes[:3]) + "-" + string(runes[3:5]) + "-" + string(runes[5:10]) + "-" + string(runes[10:])
}
