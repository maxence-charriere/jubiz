package main

import "testing"
import "html"

const (
	basicItemContent = `
<p>Le photographe russe <a target="_blank" href="http://www.greatdane.photography/">Andy Seliverstoff</a> immortalise d&rsquo;adorables portraits mettant en scène des enfants accompagnés de leur chien protecteur et de grande taille. De jolies pauses  remplies de tendresse et d&rsquo;humour tant la complicité entre les animaux et leurs petits maîtres semble grande. Des scènes dignes des films et dessins animés de notre enfance.</p>
<p><a href="http://www.fubiz.net/2016/12/12/adorable-pictures-of-kids-with-big-dogs/bigdogskids0/" rel="attachment wp-att-824268"><img class="alignnone size-medium wp-image-824268" src="http://www.fubiz.net/wp-content/uploads/2016/12/bigdogskids0-900x529.jpg" alt="" width="900" height="529" /></a></p>
<p><a href="http://www.fubiz.net/2016/12/12/adorable-pictures-of-kids-with-big-dogs/bigdogskids1/" rel="attachment wp-att-824269"><img class="alignnone size-medium wp-image-824269" src="http://www.fubiz.net/wp-content/uploads/2016/12/bigdogskids1-900x600.jpg" alt="" width="900" height="600" /></a></p>
<p><a href="http://www.fubiz.net/2016/12/12/adorable-pictures-of-kids-with-big-dogs/bigdogskids2/" rel="attachment wp-att-824270"><img class="alignnone size-medium wp-image-824270" src="http://www.fubiz.net/wp-content/uploads/2016/12/bigdogskids2-900x639.jpg" alt="" width="900" height="639" /></a></p>
<p><a href="http://www.fubiz.net/2016/12/12/adorable-pictures-of-kids-with-big-dogs/bigdogskids3/" rel="attachment wp-att-824271"><img class="alignnone size-medium wp-image-824271" src="http://www.fubiz.net/wp-content/uploads/2016/12/bigdogskids3-900x675.jpg" alt="" width="900" height="675" /></a></p>
<p><a href="http://www.fubiz.net/2016/12/12/adorable-pictures-of-kids-with-big-dogs/bigdogskids4/" rel="attachment wp-att-824272"><img class="alignnone size-medium wp-image-824272" src="http://www.fubiz.net/wp-content/uploads/2016/12/bigdogskids4-900x608.jpg" alt="" width="900" height="608" /></a></p>
<p><a href="http://www.fubiz.net/2016/12/12/adorable-pictures-of-kids-with-big-dogs/bigdogskids5/" rel="attachment wp-att-824273"><img class="alignnone size-medium wp-image-824273" src="http://www.fubiz.net/wp-content/uploads/2016/12/bigdogskids5-900x600.jpg" alt="" width="900" height="600" /></a></p>
<p><a href="http://www.fubiz.net/2016/12/12/adorable-pictures-of-kids-with-big-dogs/bigdogskids6/" rel="attachment wp-att-824274"><img class="alignnone size-medium wp-image-824274" src="http://www.fubiz.net/wp-content/uploads/2016/12/bigdogskids6-900x600.jpg" alt="" width="900" height="600" /></a></p>
<p><a href="http://www.fubiz.net/2016/12/12/adorable-pictures-of-kids-with-big-dogs/bigdogskids7/" rel="attachment wp-att-824275"><img class="alignnone size-medium wp-image-824275" src="http://www.fubiz.net/wp-content/uploads/2016/12/bigdogskids7-900x600.jpg" alt="" width="900" height="600" /></a></p>
<p><a href="http://www.fubiz.net/2016/12/12/adorable-pictures-of-kids-with-big-dogs/bigdogskids8/" rel="attachment wp-att-824276"><img class="alignnone size-medium wp-image-824276" src="http://www.fubiz.net/wp-content/uploads/2016/12/bigdogskids8-900x600.jpg" alt="" width="900" height="600" /></a></p>
<p><a href="http://www.fubiz.net/2016/12/12/adorable-pictures-of-kids-with-big-dogs/bigdogskids9/" rel="attachment wp-att-824277"><img class="alignnone size-medium wp-image-824277" src="http://www.fubiz.net/wp-content/uploads/2016/12/bigdogskids9-900x600.jpg" alt="" width="900" height="600" /></a></p>
<p><a href="http://www.fubiz.net/2016/12/12/adorable-pictures-of-kids-with-big-dogs/bigdogskids10/" rel="attachment wp-att-824278"><img class="alignnone size-medium wp-image-824278" src="http://www.fubiz.net/wp-content/uploads/2016/12/bigdogskids10-900x676.jpg" alt="" width="900" height="676" /></a></p>
<p><a href="http://www.fubiz.net/2016/12/12/adorable-pictures-of-kids-with-big-dogs/bigdogskids11/" rel="attachment wp-att-824279"><img class="alignnone size-medium wp-image-824279" src="http://www.fubiz.net/wp-content/uploads/2016/12/bigdogskids11-900x600.jpg" alt="" width="900" height="600" /></a></p>
<p><a href="http://www.fubiz.net/2016/12/12/adorable-pictures-of-kids-with-big-dogs/bigdogskids12/" rel="attachment wp-att-824280"><img class="alignnone size-medium wp-image-824280" src="http://www.fubiz.net/wp-content/uploads/2016/12/bigdogskids12-900x600.jpg" alt="" width="900" height="600" /></a></p>
<p><a href="http://www.fubiz.net/2016/12/12/adorable-pictures-of-kids-with-big-dogs/bigdogskids13/" rel="attachment wp-att-824281"><img class="alignnone size-medium wp-image-824281" src="http://www.fubiz.net/wp-content/uploads/2016/12/bigdogskids13-900x600.jpg" alt="" width="900" height="600" /></a></p>
<p><a href="http://www.fubiz.net/2016/12/12/adorable-pictures-of-kids-with-big-dogs/bigdogskids14/" rel="attachment wp-att-824282"><img class="alignnone size-medium wp-image-824282" src="http://www.fubiz.net/wp-content/uploads/2016/12/bigdogskids14-900x600.jpg" alt="" width="900" height="600" /></a></p>

<a href='http://www.fubiz.net/2016/12/12/adorable-pictures-of-kids-with-big-dogs/bigdogskids2/'><img width="150" height="150" src="http://www.fubiz.net/wp-content/uploads/2016/12/bigdogskids2-150x150.jpg" class="attachment-thumbnail" alt="bigdogskids2" /></a>
<a href='http://www.fubiz.net/2016/12/12/adorable-pictures-of-kids-with-big-dogs/bigdogskids7/'><img width="150" height="150" src="http://www.fubiz.net/wp-content/uploads/2016/12/bigdogskids7-150x150.jpg" class="attachment-thumbnail" alt="bigdogskids7" /></a>
<a href='http://www.fubiz.net/2016/12/12/adorable-pictures-of-kids-with-big-dogs/bigdogskids8/'><img width="150" height="150" src="http://www.fubiz.net/wp-content/uploads/2016/12/bigdogskids8-150x150.jpg" class="attachment-thumbnail" alt="bigdogskids8" /></a>
<a href='http://www.fubiz.net/2016/12/12/adorable-pictures-of-kids-with-big-dogs/bigdogskids6/'><img width="150" height="150" src="http://www.fubiz.net/wp-content/uploads/2016/12/bigdogskids6-150x150.jpg" class="attachment-thumbnail" alt="bigdogskids6" /></a>
<a href='http://www.fubiz.net/2016/12/12/adorable-pictures-of-kids-with-big-dogs/bigdogskids9/'><img width="150" height="150" src="http://www.fubiz.net/wp-content/uploads/2016/12/bigdogskids9-150x150.jpg" class="attachment-thumbnail" alt="bigdogskids9" /></a>
<a href='http://www.fubiz.net/2016/12/12/adorable-pictures-of-kids-with-big-dogs/bigdogskids5/'><img width="150" height="150" src="http://www.fubiz.net/wp-content/uploads/2016/12/bigdogskids5-150x150.jpg" class="attachment-thumbnail" alt="bigdogskids5" /></a>
<a href='http://www.fubiz.net/2016/12/12/adorable-pictures-of-kids-with-big-dogs/bigdogskids10/'><img width="150" height="150" src="http://www.fubiz.net/wp-content/uploads/2016/12/bigdogskids10-150x150.jpg" class="attachment-thumbnail" alt="bigdogskids10" /></a>
<a href='http://www.fubiz.net/2016/12/12/adorable-pictures-of-kids-with-big-dogs/bigdogskids4/'><img width="150" height="150" src="http://www.fubiz.net/wp-content/uploads/2016/12/bigdogskids4-150x150.jpg" class="attachment-thumbnail" alt="bigdogskids4" /></a>
<a href='http://www.fubiz.net/2016/12/12/adorable-pictures-of-kids-with-big-dogs/bigdogskids12/'><img width="150" height="150" src="http://www.fubiz.net/wp-content/uploads/2016/12/bigdogskids12-150x150.jpg" class="attachment-thumbnail" alt="bigdogskids12" /></a>
<a href='http://www.fubiz.net/2016/12/12/adorable-pictures-of-kids-with-big-dogs/bigdogskids3/'><img width="150" height="150" src="http://www.fubiz.net/wp-content/uploads/2016/12/bigdogskids3-150x150.jpg" class="attachment-thumbnail" alt="bigdogskids3" /></a>
<a href='http://www.fubiz.net/2016/12/12/adorable-pictures-of-kids-with-big-dogs/bigdogskids13/'><img width="150" height="150" src="http://www.fubiz.net/wp-content/uploads/2016/12/bigdogskids13-150x150.jpg" class="attachment-thumbnail" alt="bigdogskids13" /></a>
<a href='http://www.fubiz.net/2016/12/12/adorable-pictures-of-kids-with-big-dogs/bigdogskids14/'><img width="150" height="150" src="http://www.fubiz.net/wp-content/uploads/2016/12/bigdogskids14-150x150.jpg" class="attachment-thumbnail" alt="bigdogskids14" /></a>
<a href='http://www.fubiz.net/2016/12/12/adorable-pictures-of-kids-with-big-dogs/bigdogskids1/'><img width="150" height="150" src="http://www.fubiz.net/wp-content/uploads/2016/12/bigdogskids1-150x150.jpg" class="attachment-thumbnail" alt="bigdogskids1" /></a>
<a href='http://www.fubiz.net/2016/12/12/adorable-pictures-of-kids-with-big-dogs/bigdogskids0/'><img width="150" height="150" src="http://www.fubiz.net/wp-content/uploads/2016/12/bigdogskids0-150x150.jpg" class="attachment-thumbnail" alt="bigdogskids0" /></a>
<a href='http://www.fubiz.net/2016/12/12/adorable-pictures-of-kids-with-big-dogs/bigdogskids11/'><img width="150" height="150" src="http://www.fubiz.net/wp-content/uploads/2016/12/bigdogskids11-150x150.jpg" class="attachment-thumbnail" alt="bigdogskids11" /></a>

<img src="http://feeds.feedburner.com/~r/fubiz/~4/uclcutSERv8" height="1" width="1" alt=""/>
    `
)

func TestMakeArticleTest(t *testing.T) {
	t.Log(makeArticleText(basicItemContent))
}

func TestMakeArticleImages(t *testing.T) {
	t.Log(makeArticleImages(basicItemContent))
}

func TestMakeArticlesFromFeed(t *testing.T) {
	feed, _ := getFeed()
	t.Logf("%+v", makeArticlesFromFeed(feed))
}

func TestUnescape(t *testing.T) {
	txt := "\u003cp\u003e\u003ca target=\"_blank\" href=\"https://www.flickr.com/photos/neon_tambourine/\"\u003eLukasz Wierzbowski\u003c/a\u003e est un photographe autodidacte vivant à Wroclaw. Il imagine des clichés capturés dans la vie quotidienne. Ses modèles posent dans des poses inhabituelles, ce qui offrent des photographies remplies d’humour, avec une touche vintage. Il semble saisir ses sujets sur le vif, par surprise.\u003c/p\u003e\n\n\u003c/p\u003e\n\n\u003c/p\u003e\n\n\u003c/p\u003e\n\n\u003c/p\u003e\n\n\u003c/p\u003e\n\n\u003c/p\u003e\n\n\u003c/p\u003e"
	txt = html.UnescapeString(txt)
	t.Log(txt)
}
