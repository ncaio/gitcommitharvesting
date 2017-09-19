//
//	NCAIO - caiogore())(gmail))\.\{]com{]
//
//	source{d}, go-git, (2017), GitHub repository, https://github.com/src-d/go-git
//
//
//
package main

//
//
//
import (
	"flag"
	"fmt"
	"gopkg.in/src-d/go-git.v4"
	. "gopkg.in/src-d/go-git.v4/_examples"
	"gopkg.in/src-d/go-git.v4/plumbing/object"
	"gopkg.in/src-d/go-git.v4/storage/memory"
	"strings"
)

//
//
//

func banner() {
	fmt.Print("\033[H\033[2J")
	Info("----------------------------------------------------------------------")
	Info("[- GITCOMMITHARVESTING by ncaio -]")
	Info(">> caiogore [_|_] (g)mail _._ com")
	Info("----------------------------------------------------------------------")
	Info("WARNING - DON'T TRY THIS AGAINST HUGE REPOSITORY. Maybe OOM will kill you.")
	fmt.Println("")
}

//
//
//
func main() {
	//
	//
	//
	re := flag.String("r", "repository not found - try ./githarvesting.go -r https://github.com/author/repository", "Repository: -r=https://repository")
	flag.Parse()
	repo := *re
	//
	//
	//
	banner()
	//
	//
	//
	Info("git clone %s", repo)
	r, err := git.Clone(memory.NewStorage(), nil, &git.CloneOptions{
		URL: repo,
	})
	CheckIfError(err)
	//
	//
	//
	ref, err := r.Head()
	CheckIfError(err)
	//
	//
	//
	cIter, err := r.Log(&git.LogOptions{From: ref.Hash()})
	CheckIfError(err)
	//
	//
	//
	err = cIter.ForEach(func(c *object.Commit) error {
		fmt.Println(c.Author.Name + ";" + strings.Replace(c.Author.Email, " ", "", -1))
		return nil
	})
	CheckIfError(err)
}
