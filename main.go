package git_questionare

import (
    question "git-questionare/pkg/question"
)

func GitIntermediate() question.Questionare {
    questionare := question.Questionare{
        {
            Question: "How to stage files in git?",
            Options: []string{
                "git add",
                "git include .",
                "git watch .",
            },
            Answer: "git add",
        },
        {
            Question: "How to send files and commits in git to a remote repository?",
            Options: []string{
                "git push",
                "git send",
                "git send origin master",
            },
            Answer: "git push",
        },
        {
            Question: "How to get the latest changes from a remote repository?",
            Options: []string{
                "git pull",
                "git fetch origin master",
                "git retrieve latest",
            },
            Answer: "git pull",
        },
        {
            Question: "How can you see the commit history in a git repository?",
            Options: []string{
                "git log",
                "git history",
                "git show",
            },
            Answer: "git log",
        },
    }
    return questionare
}
func GitBasics() question.Questionare {
    questionare := question.Questionare{
        {
            Question: "How to init a git repository?",
            Options: []string{
                "git init",
                "git create",
                "git start",
            },
            Answer: "git init",
        },
        {
            Question: "How to stage files in git?",
            Options: []string{
                "git add",
                "git include .",
                "git watch .",
            },
            Answer: "git add",
        },
        {
            Question: "How to send files and commits in git to a remote repository?",
            Options: []string{
                "git push",
                "git send",
                "git send origin master",
            },
            Answer: "git push",
        },
        {
            Question: "How to get the latest changes from a remote repository?",
            Options: []string{
                "git pull",
                "git fetch origin master",
                "git retrieve latest",
            },
            Answer: "git pull",
        },
        {
            Question: "How can you see the commit history in a git repository?",
            Options: []string{
                "git log",
                "git history",
                "git show",
            },
            Answer: "git log",
        },
    }
    return questionare
}
