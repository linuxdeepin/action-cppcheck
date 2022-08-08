package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strings"
	"time"

	"github.com/google/go-github/v42/github"
	"github.com/sourcegraph/go-diff/diff"
	"golang.org/x/sync/errgroup"
)

func init() {
	rand.Seed(time.Now().Unix())
}

func main() {
	//*** init ***//
	var file, owner, repo string
	var pullID int
	flag.StringVar(&repo, "repo", "peeweep-test/dde-dock", "owner and repo name")
	flag.StringVar(&file, "f", "/dev/stdin", "cppcheck result in xml format")
	flag.IntVar(&pullID, "pr", 0, "pull request id")
	flag.Parse()
	arr := strings.SplitN(repo, "/", 2)
	owner = arr[0]
	repo = arr[1]

	client := github.NewClient(&http.Client{})

	//*** init end ***/

	//*** check ***//
	var diffs []*diff.FileDiff
	var checkErrs []CppCheckError
	eg, ctx := errgroup.WithContext(context.Background())
	// get pull request diff
	eg.Go(func() error {
		diffRaw, _, err := client.PullRequests.GetRaw(ctx, owner, repo, pullID, github.RawOptions{Type: github.Diff})
		if err != nil {
			return fmt.Errorf("get diff: %w", err)
		}
		diffs, err = diff.ParseMultiFileDiff([]byte(diffRaw))
		if err != nil {
			return fmt.Errorf("parse diff: %w", err)
		}
		return nil
	})
	// get cppcheck result
	eg.Go(func() error {
		errors, err := decodeErrors(file)
		if err != nil {
			return err
		}
		checkErrs = errors
		return nil
	})
	err := eg.Wait()
	if err != nil {
		log.Fatal(err)
	}
	var comments []*github.DraftReviewComment
	for i := range diffs {
		filename := strings.TrimPrefix(diffs[i].NewName, "b/")
		for j := range diffs[i].Hunks {
			startline := int(diffs[i].Hunks[j].NewStartLine)
			endline := startline + int(diffs[i].Hunks[j].NewLines) - 1
			for k := range checkErrs {
				if checkErrs[k].Location == nil {
					continue
				}
				if checkErrs[k].Location.File != filename {
					continue
				}
				if checkErrs[k].Location.Line < startline || checkErrs[k].Location.Line > endline {
					continue
				}
				line, body := checkErrs[k].Location.Line, checkErrs[k].Verbose
				comments = append(comments, &github.DraftReviewComment{
					Path: &filename,
					Line: &line,
					Body: &body,
				})
				fmt.Printf("::warning file=%s,line=%d::%s\n", filename, checkErrs[k].Location.Line, checkErrs[k].Verbose)
			}
		}
	}
}
