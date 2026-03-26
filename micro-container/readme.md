## Directory structure

```
- micro-container/

    - cmd/

        - container/

            - main.go # 프로그램 진입점

    - internal/

        - cgroups/

            - cgroups.go # cgroups 파일 생성 및 삭제

        - container/

            - container.go # 새로운 container만드는 로직

            - config.go # 새로운 container만들 때 쓰는 struct 모음

            - run.go # 컨테이너 프로세스 시작 로직

        - network/

            - bridge.go # 스위치 만들기

            - config.go # 컨테이너 내부 랜선 설정

            - veth.go # 랜선 만들고, 호스트 연결하기
```

### main.go
~~~
func main() {
    switch os.Args[1] {
    case "run":
        // 1단계: 브릿지/Cgroup 준비
        // 2단계: cmd.Start()로 격리된 'child' 실행
        // 3단계: PID 나왔으니 veth 연결 (SetupVeth)
        // 4단계: cmd.Wait()로 자식 종료 대기
        
    case "child":
        // 4단계: 방 안에서 할 일 (Hostname 설정, 네트워크 Configure)
        // 마지막: syscall.Exec("/bin/bash", ...)
    }
}
~~~

