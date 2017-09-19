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
	"os"
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
	fmt.Println("")
}

//
//
//
func fmem(u string) {
	//
	//
	//
	Info("WARNING - DON'T TRY AGAINST HUGE REPOSITORY. Maybe OOM will kill you.")
	//
	//
	//
	Info("git clone %s", u)
	r, err := git.Clone(memory.NewStorage(), nil, &git.CloneOptions{
		URL:      u,
		Progress: os.Stdout,
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
	rer, err := r.Log(&git.LogOptions{From: ref.Hash()})
	CheckIfError(err)
	//
	//
	//
	err = rer.ForEach(func(c *object.Commit) error {
		fmt.Println(c.Author.Name + ";" + strings.Replace(c.Author.Email, " ", "", -1))
		return nil
	})
	CheckIfError(err)
}

//
//
//
func lstor(url string, d string) {

	Info("git clone %s %s", url, d)

	r, err := git.PlainClone(d, false, &git.CloneOptions{
		URL:      url,
		Progress: os.Stdout,
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
	rer, err := r.Log(&git.LogOptions{From: ref.Hash()})
	CheckIfError(err)
	//
	//
	//
	err = rer.ForEach(func(c *object.Commit) error {
		fmt.Println(c.Author.Name + ";" + strings.Replace(c.Author.Email, " ", "", -1))
		return nil
	})
	CheckIfError(err)
	//
	//
	//
	os.RemoveAll(d)
	fmt.Println("------------------")
	Info("%s deleted", d)

}

//
//
//
func main() {
	re := flag.String("r", "repository not found - try ./githarvesting.go -r https://github.com/author/repository", "Repository: -r=https://repository")
	me := flag.String("m", "no", "Storage in-memory")
	//pa := flag.String("d", "/tmp/_git/", "Local Storage")
	flag.Parse()
	repo := *re
	memo := *me
	//dst := *pa
	dst := "/tmp/_git/"
	banner()
	if strings.Contains(memo, "yes") {
		fmem(repo)
	} else {
		lstor(repo, dst)
	}
}
