package main

import "fmt"

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

func (i *InMemoryDB) GetByIdAuthor(Id uint8) Author {
  for _, val := range i.authors {
    if val.Id == Id {
      fmt.Println("SUCCESS: Found author")
      return val
    }
  }
  fmt.Println("WARNING: There is no author with this id")
  return Author{}
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
  fmt.Println("WARNING: There is no article with this id")
}

func (i *InMemoryDB) GetByIdArticle(id uint8) Article {
  for _, val := range i.articles {
    if val.Id == id {
      fmt.Println("SUCCESS: Got article with this id")
      return val
    }
  }
  fmt.Println("WARNING: There is no article with this id")
  return Article{}
}

func main() {
  
}