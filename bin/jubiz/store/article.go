package store

import (
	"path/filepath"
	"sort"
	"sync"

	"github.com/maxence-charriere/jubiz"
	"github.com/murlokswarm/app"
	"github.com/murlokswarm/flux"
)

// Action names related to articles.
var (
	LoadLocalArticles = "art-load-local"
	SaveArticles      = "art-save"
	DownloadArticles  = "art-download"
	UpdateArticle     = "art-update"
	ShareArticle      = "art-share"
)

// Event names related to articles.
var (
	LocalArticlesLoaded = "art-local-loaded"
	ArticlesSaved       = "art-saved"
	ArticlesDownloading = "art-downloading"
	ArticlesDownloaded  = "art-downloaded"
	ArticleUpdated      = "art-updated"
)

var (
	articlesName = filepath.Join(app.Storage().Default(), "articles.json")
)

type articleStore struct {
	flux.Store

	mutex    sync.Mutex
	articles jubiz.ArticleList
}

func (s *articleStore) OnDispatch(a flux.Action) error {
	switch a.Name {
	case LoadLocalArticles:
		s.LoadFromLocalResources()

	case SaveArticles:
		return s.Save()

	case DownloadArticles:
		return s.Download()

	case UpdateArticle:
		s.Update(a.Payload.(jubiz.Article))

	case ShareArticle:
		app.Share().URL(a.Payload.(jubiz.Article).URL)
	}
	return nil
}

func (s *articleStore) LoadFromLocalResources() error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	var articles jubiz.ArticleList
	if err := jubiz.FileUnmarshal(articlesName, &articles); err != nil {
		return err
	}
	sort.Sort(articles)
	s.articles = articles

	s.Emit(flux.Event{
		Name:    LocalArticlesLoaded,
		Payload: articles,
	})
	return nil
}

func (s *articleStore) Save() error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if err := jubiz.FileMarshal(articlesName, s.articles); err != nil {
		s.Emit(flux.Event{
			Name:  ArticlesSaved,
			Error: err,
		})
		return err
	}
	return nil
}

func (s *articleStore) Download() error {

	s.Emit(flux.Event{
		Name: ArticlesDownloading,
	})

	feed, err := jubiz.GetFeed()
	if err != nil {
		s.Emit(flux.Event{
			Name:  ArticlesDownloaded,
			Error: err,
		})
		return err
	}

	s.mutex.Lock()
	defer s.mutex.Unlock()

	articles := jubiz.MakeArticlesFromFeed(feed)
	newCount := 0

	articles, newCount = jubiz.MergeArticleLists(s.articles, articles)
	if newCount != 0 {
		app.Dock().SetBadge(newCount)
	}

	sort.Sort(articles)
	s.articles = articles

	s.Emit(flux.Event{
		Name:    ArticlesDownloaded,
		Payload: articles,
	})
	return nil
}

func (s *articleStore) Update(a jubiz.Article) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	for i, art := range s.articles {
		if art.ID == a.ID {
			s.articles[i] = a
			break
		}
	}

	s.Emit(flux.Event{
		Name:    ArticleUpdated,
		Payload: a,
	})
}
