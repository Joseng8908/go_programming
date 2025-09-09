package main

import (
	"fmt"
	"gopl/pkg/github"
	"log"
	"os"
	"os/exec"
	"strings"
)


func issue() {
	if len(os.Args) < 2 {
		fmt.Println("Write invalid command")
		os.Exit(1)
	}

	subcommand := os.Args[1]

	switch subcommand {
	case "create": createIssue()

	case "read": readIssue()

	case "modify": modifyIssue()

	case "close": readIssue()

	default: 
		fmt.Println("unknown command")
		os.Exit(1)
	}
}

func createIssue() {
	// 1. 임시 파일 생성
	tmpFile, err := os.CreateTemp("", "github-issue-")
	if err != nil {
		log.Fatal(err)
	}
	defer os.Remove(tmpFile.Name()) // 함수 종료 시 파일 삭제

	// 2. 텍스트 편집기 실행
	editor := os.Getenv("EDITOR")
	if editor == "" {
		editor = "vim" // 기본 편집기
	}
	
	cmd := exec.Command(editor, tmpFile.Name())
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	
	if err := cmd.Run(); err != nil {
		log.Fatalf("편집기 실행 오류: %v", err)
	}

	// 3. 임시 파일 내용 읽기
	content, err := os.ReadFile(tmpFile.Name())
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(content), "\n")
	title := lines[0]
	body := strings.Join(lines[1:], "\n")

	// 4. GitHub API 요청 생성
	// 실제 API 요청은 go-github 클라이언트 객체와 함께 호출됩니다.
	// 예를 들어, github.Issues.Create(context.Background(), "owner", "repo", &github.IssueRequest{...})
	github.Issues.Create()
	fmt.Printf("제목: %s\n", title)
	fmt.Printf("본문: \n%s\n", body)
	fmt.Println("이슈 생성을 위한 API 요청을 보냅니다.")
}

func readIssue() {

}

func modifyIssue() {

}

func closeIssue() {

}
