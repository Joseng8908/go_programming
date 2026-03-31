# Micro-Container: Go로 구현한 초경량 컨테이너 엔진
- 이 프로젝트는 도커나, 컨테이너디 같은 외부 도구 없이, 리눅스 커널의 핵심 기능인 Namespace와 Cgroups를 이용하여 컨테이너의 생명주기를 직접 관리하는 교육용 컨테이너 엔진입니다.

## Features
- 격리된 실행 (Run)
- 백그라운드 모드 (Detach): -d 옵션을 이용하여 컨테이너를 백그라운드 프로세스로 실행.
- 프로세스 관리 (PS): 실행 중인 컨테이너의 상태와 정보를 JSON 기반으로 관리 및 조회.
- 컨테이너 진입 (Exec): CGO 생성자 트릭을 이용해 실행 중인 컨테이너에 동적으로 접속(이 부분 공부가 좀 필요합니다)
- 자원 회수: (Stop): 프로세스 종료 및 마운트 해제, Cgroup 삭제 등 모든 자원을 역순으로 정리.

## Skills & Architecture
- Language: Go (Golang)
- Core Logic:
    - Storage: OverlayFS를 이용한 계층형 파일 시스템 구현.
    - Network: Veth Pair와 Bridge를 이용한 컨테이너 전용 네트워크 구성.
    - Resource Limit: Cgroups를 통한 CPU/Memory 제한 기반 마련.
    - Namespace Entry: setns 시스템 콜과 CGO를 이용한 컨테이너 내부 진입 로직.

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

            - info.go # 컨테이너 정보 읽고, 쓰고, 가져오는 로직

            - stop.go # 컨테이너 데몬으로 실행시켰을 때 강제종료하는 로직

        - network/

            - bridge.go # 스위치 만들기

            - config.go # 컨테이너 내부 랜선 설정

            - veth.go # 랜선 만들고, 호스트 연결하기
```

## Execution
- build
~~~bash
make build
# or
go build -o micro-container ./cmd/container/main.go
~~~

- 컨테이너 실행 (Background)
~~~bash
sudo ./micro-container run -d /bin/sh -c "while true; do sleep 1; done"
~~~

- 컨테이너 목록 확인
~~~bash
sudo ./micro-container ps
~~~

- 컨테이너 내부 접속
~~~bash
sudo ./micro-container exec test-container /bin/sh
~~~

- 컨테이너 정지 및 삭제
~~~bash
sudo ./micro-container stop test-container
~~~
 
## Todo
- 한번에 많은 container를 띄우기 위해 각각의 번호 매핑하는 로직과 한번에 몇개 띄우는 로직 짜기.
    - network, ip할당, container이름, 생기는 cgroups파일들 등등 관리도 함께.
- setns쪽 CGO 공부하기(pthread_create 이슈?).
- json파일 쪽 로직 손보기..이거 계속 남아있는지, stop했을때 사라지는지 다시 봐야함.
- 로그 파일 정리하기
- container리소스 모니터링
- Unix/Linux 개념 복습

## 생각해볼만한 주제
- container 리소스 모니터링 기반 스케줄링 알고리즘?
