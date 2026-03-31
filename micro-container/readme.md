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


