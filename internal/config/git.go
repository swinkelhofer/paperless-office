package config

import (
	"time"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
	log "github.com/sirupsen/logrus"
)

func (c *Configuration) InitGitWorker() {
	c.GitPushChan = make(chan CommitMessage, 100)
	go c.GitPushWorker(c.GitPushChan)
}

func (c *Configuration) GitPushWorker(gitPushChan chan CommitMessage) {
	var (
		commit CommitMessage
		err    error
	)

	for {
		select {
		case commit = <-gitPushChan:
			if err = c.GitStageAndPush(string(commit)); err != nil {
				log.Errorf("[gitWorker] %s", err.Error())
			}
		}
	}
}

func (c *Configuration) GitStageAndPush(commitMessage string) error {
	var (
		workTree *git.Worktree
		err      error
		remote   *git.Remote
	)

	if c.Git.gitRepo == nil {
		log.Debugf("Target directory for processed PDFs is not a valid git repository")
		return nil
	}

	if workTree, err = c.Git.gitRepo.Worktree(); err != nil {
		log.Errorf("Get worktree: %s", err.Error())
		return err
	}

	if err = workTree.AddWithOptions(&git.AddOptions{All: true}); err != nil {
		log.Errorf("Stage files: %s", err.Error())
		return err
	}

	if _, err = workTree.Commit(commitMessage, &git.CommitOptions{
		All: true,
		Author: &object.Signature{
			Name:  c.Git.Author.Name,
			Email: c.Git.Author.Email,
			When:  time.Now(),
		},
	}); err != nil {
		log.Errorf("Commit: %s", err.Error())
		return err
	}

	if remote, err = c.Git.gitRepo.Remote("origin"); err != nil {
		log.Errorf("Get Remote: %s", err.Error())
		return err
	}

	if err = remote.Push(&git.PushOptions{}); err != nil {
		log.Errorf("Push: %s", err.Error())
		return err
	}
	return nil
}
