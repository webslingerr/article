package main

import (
	"fmt"
	"strings"
)

type Author struct {
  Id uint8
  Name string
  Age uint8
}

type Article struct {
  Id uint8
  Title string
  Desc string
  Author_id uint8
}

type InMemoryDB struct {
  authors []Author
  articles []Article
}

func (i *InMemoryDB) CreateAuthor(author Author) {
  for _, val := range i.authors {
    if val.Id == author.Id {
      fmt.Println("WARNING: Author with this id already exists")
      return
    }
  }
  i.authors = append(i.authors, author)
  fmt.Println("SUCCESS: Created a new author")
}

func (i *InMemoryDB) UpdateAuthor(author Author) {
  for in, val := range i.authors {
    if val.Id == author.Id {
      i.authors[in].Name = author.Name
      i.authors[in].Age = author.Age
      fmt.Println("SUCCESS: Author has been updated")
      return
    }
  }
  fmt.Println("WARNING: Author with this id doesn't exist")
}

func (i *InMemoryDB) DeleteAuthor(author Author) {
  for in, val := range i.authors {
    if val.Id == author.Id {
      i.authors = append(i.authors[:in], i.authors[in+1:]...)
      fmt.Println("SUCCESS: Deleted the user")
      return
    }
  }
}

func (i *InMemoryDB) GetByIdAuthor(Id uint8) {
  for _, val := range i.authors {
    if val.Id == Id {
      fmt.Println("SUCCESS: Found author")
      fmt.Printf("\tAuthor id: %d\n", val.Id)
      fmt.Printf("\tAuthor age: %d\n", val.Age)
      fmt.Printf("\tAuthor name: %v\n", val.Name)
      return
    }
  }
  fmt.Println("WARNING: There is no author with this id")
}

func (i *InMemoryDB) GetAllAuthor(limit, offset int, s string) []Author {
  authors := []Author{}
  if limit+offset>len(i.authors) {
    return authors
  }
  for _, val := range i.authors[offset:] {
    if limit == 0 {
      break
    }
    if strings.HasPrefix(strings.ToLower(val.Name), s) {
      authors = append(authors, val)
      limit--
    }
  }
  return authors
}

func (i *InMemoryDB) CreateArticle(article Article) {
  for _, val := range i.articles {
    if val.Id == article.Id {
      fmt.Println("WARNING: Article with this id already exists")
      return
    }
  }
  i.articles = append(i.articles, article)
  fmt.Println("SUCCESS: Created a new article")
}

func (i *InMemoryDB) UpdateArticle(article Article) {
  for in, val := range i.articles {
    if val.Id == article.Id {
      i.articles[in].Desc = article.Desc
      i.articles[in].Title = article.Title
      fmt.Println("SUCCESS: Updated article")
      return
    }
  }
  fmt.Println("WARNING: There is no article with this id")
}

func (i *InMemoryDB) DeleteArticle(article Article) {
  for in, val := range i.articles {
    if val.Id == article.Id {
      i.articles = append(i.articles[:in], i.articles[in+1:]...)
      fmt.Println("SUCCESS: Deleted article")
      return
    }
  }
  fmt.Println("There is no article with this id")
}

func (i *InMemoryDB) GetByIdArticle(id uint8) {
  for _, val := range i.articles {
    if val.Id == id {
      fmt.Println("SUCCESS: Got article with this id")
      fmt.Println("Id:", val.Id)
      fmt.Println("Title:", val.Title)
      fmt.Println("Description:", val.Desc)
      i.GetByIdAuthor(val.Author_id)
      return
    }
  }
  fmt.Println("WARNING: There is no article with this id")
}

func (i *InMemoryDB) GetAllArticle(limit, offset int, search string, author_id uint8) []Article {
  articles := []Article{}
  if limit+offset>len(i.articles) {
    return articles
  }
  for _, val := range i.articles[offset:] {
    if limit == 0 {
      break
    } 
    if val.Author_id == author_id {
      if strings.HasPrefix(strings.ToLower(val.Title), search) {
        articles = append(articles, val)
        limit--
      }
    }
  }
  return articles
}

func main() {
  db := InMemoryDB{}
  author1 := Author {
    Id: 1,
    Name: "Mark Twen",
    Age: 60,
  }
  article := Article {
    Id: 1,
    Title: "Title",
    Desc: "This is the desctiprion",
    Author_id: 1,
  }

  db.CreateAuthor(author1)
  db.CreateArticle(article)
  db.GetByIdArticle(article.Id)
}