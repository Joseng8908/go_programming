# Day 2 - goroutine leak
 
## Experiment

~~~go
func main() {
    // 1. Unbuffered 채널 생성
    dataChan := make(chan int)
    ctx, cancel := context.WithCancel(context.Background())

    // 2. 워커 고루틴 실행
    go func() {
        // 이 defer는 과연 실행될 수 있을까?
        defer fmt.Println("Worker finished") 
        for {
            select {
            case <-ctx.Done():
                return
            case d := <-dataChan:
                fmt.Println("Received:", d)
            }
        }
    }()

    dataChan <- 1
    // cancel()을 호출하지 않고, dataChan도 닫지 않은 채 메인이 끝난다면?
    // 혹은 dataChan을 닫았는데 워커가 모르고 보내려고 한다면?
}
~~~

## Observation
- 채널을 닫지 않으면 range나 select 대기 중인 고루틴은 영원히 메모리에 상주함.
- 이미 닫힌 채널에 데이터를 보내려고 시도하는 순간 프로세스는 즉시 종료, 즉 패닉
- select는 케이스가 동시에 준비되면 무작위로 선택. 취소 신호가 와도 데이터를 하나 더 처리할 확률이 있음

## Why?
1. Leak: Go 런타임은 채널에 누군가 데이터를 보낼 가능성이 아주 조금이라도 있으면 고루틴을 수거하지 않음
2. Panic: 닫힌 채널에 보내는 것은 절대 금기, 즉 설계 실수
3. Select: Go의 철학의 공정함을 잘 나타내는 경우. 특정 케이스에 가중치를 두지 않음로써 기아 현상을 방지하려 하기 때문에

## Key Insights
1. Ownership: 채널을 만든 놈이 닫을 책임도 있다. (보통 Sender)
2. Graceful Shutdown: context.Context는 고루틴을 죽이는 놈이다. 이건 취소를 위해 만들어진 데이터셋이다. defer cancel()을 필수
3. Wait Before Close: 여러 Sender가 있을 때는 sync.WaitGroup으로 모두 끝난지 확인 한 후 채널을 닫아야한다

## Mental Model
고루틴은 시한 폭탄
모든 go키워드를 사용해 고루틴을 만들 때 마다 이놈은 언제, 어떻게 종료되는지를 선언하지 않으면 그놈은 결국 서버 메모리를 다 써버리거나, 
패닉이 일어나 모두를 죽일 수 있다
