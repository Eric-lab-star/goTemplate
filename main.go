package main

import (
	"html/template"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

type book struct {
	Title      string
	Author     string
	PageCount  int
	Review     string
	Categories []string
	PubDate    string
}

var (
	TheDeepEnd = book{
		Title:     "the deep end",
		Author:    "jeff kinney",
		PageCount: 224,
		Review: `
		The Heffley clan has been stuck living together in Gramma’s basement for two months, waiting for the family home to be repaired, and the constant togetherness has been getting on everybody’s nerves. Luckily Greg’s Uncle Gary has a camper waiting for someone to use it, and so the Heffleys set off on the open road looking for an adventurous vacation, hoping the changing scenery will bring a spark back to the family unit. The winding road leads the Heffleys to a sprawling RV park, a setting teeming with possibilities for Greg to get up to his usual shenanigans. Greg’s snarky asides and misadventures continue to entertain. At this point the Wimpy Kid books run like a well-oiled machine, paced perfectly with witty lines, smart gags, and charming cartoons. Kinney knows just where to put a joke, the precise moment to give a character shading, and exactly how to get the narrative rolling, spinning out the oddest plot developments. The appreciation Kinney has for these characters seeps through the novels, endearing the Heffleys to readers even through this title, the 15th installment in a franchise boasting spinoffs, movies, and merchandise. There may come a time when Greg and his family overstay their welcome, but thankfully that day still seems far off.	
		`,
		Categories: []string{"children's entertainment and sports", "general graphic novels and comics"},
		PubDate:    "2020-10-27",
	}
	WreckingBall = book{
		Author:     "jeff kinney",
		PageCount:  224,
		Review:     "Whekn Great Aunt Reba dies, she leaves some money to the family. Greg’s mom calls a family meeting to determine what to do with their share, proposing home improvements and then overruling the family’s cartoonish wish lists and instead pushing for an addition to the kitchen. Before bringing in the construction crew, the Heffleys attempt to do minor maintenance and repairs themselves—during which Greg fails at the work in various slapstick scenes. Once the professionals are brought in, the problems keep getting worse: angry neighbors, terrifying problems in walls, and—most serious—civil permitting issues that put the kibosh on what work’s been done. Left with only enough inheritance to patch and repair the exterior of the house—and with the school’s dismal standardized test scores as a final straw—Greg’s mom steers the family toward moving, opening up house-hunting and house-selling storylines (and devastating loyal Rowley, who doesn’t want to lose his best friend). While Greg’s positive about the move, he’s not completely uncaring about Rowley’s action. (And of course, Greg himself is not as unaffected as he wishes.) The gags include effectively placed callbacks to seemingly incidental events (the “stress lizard” brought in on testing day is particularly funny) and a lampoon of after-school-special–style problem books. Just when it seems that the Heffleys really will move, a new sequence of chaotic trouble and property destruction heralds a return to the status quo. Whew.",
		Categories: []string{"childrend's entertainment and sports", "general graphic novels and comics"},
		PubDate:    "2019-11-05",
	}
	DogManAndCatKid = book{
		Title:      "dog man and cat kid",
		Author:     "dav pilkey",
		PageCount:  256,
		Review:     "The Steinbeck novel’s Cain/Abel motif gets some play here, as Petey, “world’s evilest cat” and cloned Li’l Petey’s original, tries assiduously to tempt his angelic counterpart over to the dark side only to be met, ultimately at least, by Li’l Petey’s “Thou mayest.” (There are also occasional direct quotes from the novel.) But inner struggles between good and evil assume distinctly subordinate roles to riotous outer ones, as Petey repurposes robots built for a movie about the exploits of Dog Man—“the thinking man’s Rin Tin Tin”—while leading a general rush to the studio’s costume department for appropriate good guy/bad guy outfits in preparation for the climactic battle. During said battle and along the way Pilkey tucks in multiple Flip-O-Rama inserts as well as general gags. He lists no fewer than nine ways to ask “who cut the cheese?” and includes both punny chapter titles (“The Bark Knight Rises”) and nods to Hamiltonand Mary Poppins. The cartoon art, neatly and brightly colored by Garibaldi, is both as easy to read as the snappy dialogue and properly endowed with outsized sound effects, figures displaying a range of skin colors, and glimpses of underwear (even on robots).",
		Categories: []string{"general graphic novels and comics", "mystery and crime"},
		PubDate:    "2017-12-26",
	}
	Ghost = book{
		Title:      "ghost",
		Author:     "raina telgemeir",
		PageCount:  256,
		Review:     "Dad has a new job, but it’s little sister Maya’s lungs that motivate the move: she has had cystic fibrosis since birth—a degenerative breathing condition. Despite her health, Maya loves adventure, even if her lungs suffer for it and even when Cat must follow to keep her safe. When Carlos, a tall, brown, and handsome teen Ghost Tour guide introduces the sisters to the Bahía ghosts—most of whom were Spanish-speaking Mexicans when alive—they fascinate Maya and she them, but the terrified Cat wants only to get herself and Maya back to safety. When the ghost adventure leads to Maya’s hospitalization, Cat blames both herself and Carlos, which makes seeing him at school difficult. As Cat awakens to the meaning of Halloween and Day of the Dead in this strange new home, she comes to understand the importance of the ghosts both to herself and to Maya. Telgemeier neatly balances enough issues that a lesser artist would split them into separate stories and delivers as much delight textually as visually. The backmatter includes snippets from Telgemeier’s sketchbook and a photo of her in Día makeup.",
		Categories: []string{"general graphic novels", "childrens's social theme", "children's family"},
		PubDate:    "2016-09-13",
	}
	LegacyAndTheDouble = book{
		Title:      "legacy and the double",
		Author:     "annie methew",
		PageCount:  224,
		Review:     "In this sequel to Legacy and the Queen (2019), Legacy Petrin and her friends Javi and Pippa have returned to Legacy’s home province and the orphanage run by her father. With her friends’ help, she is in training to defend her championship when they discover that another player, operating under the protection of High Consul Silla, is presenting herself as Legacy. She is so convincing that the real Legacy is accused of being an imitation. False Legacy has become a hero to the masses, further strengthening Silla’s hold, and it becomes imperative to uncover and defeat her. If Legacy is to win again, she must play her imposter while disguised as someone else. Winning at tennis is not just about money and fame, but resisting Silla’s plans to send more young people into brutal mines with little hope of better lives. Legacy will have to overcome her fears and find the magic that allowed her to claim victory in the past. This story, with its elements of sports, fantasy, and social consciousness that highlight tensions between the powerful and those they prey upon, successfully continues the series conceived by late basketball superstar Bryant. As before, the tennis matches are depicted with pace and spirit. Legacy and Javi have brown skin; most other characters default to White.",
		Categories: []string{"children's social theme", "childrens's entertainment", "children's science fiction and fantasy"},
		PubDate:    "2021-08-24",
	}
)

func main() {
	books := []book{LegacyAndTheDouble, Ghost, TheDeepEnd, WreckingBall, DogManAndCatKid}
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", homeHandler(books))
	fileServer(r, "/static")
	http.ListenAndServe("localhost:8000", r)
}

func homeHandler(data []book) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles("blog.gotmpl", "book.gotmpl", "header.gotmpl")
		if err != nil {
			panic(err)
		}
		err = tmpl.Execute(w, data)
		if err != nil {
			panic(err)
		}
	}
}

func fileServer(r chi.Router, path string) {
	// url에 {}*가 있는지 확인후 거르기
	if strings.ContainsAny(path, "{}*") {
		panic("FileServer does not permit any URL parameters.")
	}
	//url의 마지막 글자가 "/"가 아닌경우 "/"를 붙여서 리다이렉트 시킴
	if path != "/" && path[len(path)-1] != '/' {
		r.Get(path, http.RedirectHandler(path+"/", 301).ServeHTTP)
		path += "/"
	}
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	filesDir := http.Dir(filepath.Join(wd, "static"))

	path += "*"
	r.Get(path, func(w http.ResponseWriter, r *http.Request) {
		rctx := chi.RouteContext(r.Context())
		pathPrefix := strings.TrimSuffix(rctx.RoutePattern(), "/*")
		fs := http.StripPrefix(pathPrefix, http.FileServer(filesDir))
		fs.ServeHTTP(w, r)
	})
}
