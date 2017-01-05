package main

import (
	"sync"

	"sort"

	"github.com/murlokswarm/app"
	"github.com/murlokswarm/flux"
)

var (
	articlesFilename = app.Resources().Join("articles.json")

	articlesRead   = "articles-read"
	articlesSave   = "articles-save"
	articlesGet    = "articles-get"
	articleSet     = "article-set"
	articleDisplay = "article-display"
)

type ArticlesStore struct {
	flux.Store

	mutex    sync.Mutex
	articles articleList
}

func (s *ArticlesStore) OnDispatch(a flux.Action) error {
	switch a.Name {
	case articlesRead:
		s.Read()

	case articlesSave:
		return s.Save()

	case articlesGet:
		return s.Get()

	case articleDisplay:
		s.Display(a.Payload.(article))

	case articleSet:
		s.Set(a.Payload.(article))
	}
	return nil
}

func (s *ArticlesStore) Read() error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	articles, err := readArticles(articlesFilename)
	if err != nil {
		return err
	}

	sort.Sort(articles)
	s.articles = articles

	s.Emit(flux.Event{
		Name:    articlesRead,
		Payload: articles,
	})
	return nil
}

func (s *ArticlesStore) Save() error {
	e := flux.Event{Name: articlesSave}

	s.mutex.Lock()
	defer s.mutex.Unlock()

	if err := saveArticles(articlesFilename, s.articles); err != nil {
		e.Error = err
		s.Emit(e)
		return err
	}
	return nil
}

func (s *ArticlesStore) Get() error {
	e := flux.Event{Name: articlesGet}

	feed, err := getFeed()
	if err != nil {
		e.Error = err
		s.Emit(e)
		return err
	}

	s.mutex.Lock()
	defer s.mutex.Unlock()

	articles := makeArticlesFromFeed(feed)
	articles = mergeArticleLists(s.articles, articles)
	sort.Sort(articles)
	s.articles = articles

	e.Payload = articles
	s.Emit(e)
	return nil
}

func (s *ArticlesStore) Display(a article) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	s.Emit(flux.Event{
		Name:    articleDisplay,
		Payload: a,
	})
}

func (s *ArticlesStore) Set(a article) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	for i, art := range s.articles {
		if art.URL == a.URL {
			s.articles[i] = a
			break
		}
	}

	s.Emit(flux.Event{
		Name:    articleSet,
		Payload: a,
	})
}
