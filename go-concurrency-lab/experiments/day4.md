# Day 4 - Pipline & Worker Pool 
 
## Experiment
- 워커 3개를 띄우고 10개의 일을 처리하는 워커 풀을 만든다
- context.Cancel을 불렀을 때 워커들이 순차적으로 return하는지 확인한다
~~~go
func WorkerPool(ctx context.Context, jobs <- chan int) {
	var wg sync.WaitGroup

	// 1. 워커 3개 가동 (Fan - out: 한 단계에서 여러 워커를 띄워 병렬로 처리하는 것)
	for i := 1; i <= 3; i++ {
		go func(id int) {
			defer wg.Done()
			for {
				select {
				case <- ctx.Done():
					fmt.Printf("워커 %d 퇴근 \n", id)
					return
				case job, ok := <- jobs:
					if !ok {return}
					fmt.Printf("워커 %d 가 일하는 중: %d\n", id, job)
				}
			}
		}(i)
	}
	wg.Wait()
}
~~~

## Observation
- cancel() 이후 Wait()를 하지 않으면 메인이 먼저 종료되어 워커의 마지막 로그가 안 보일 수 있음.
- 채널 버퍼가 가득 차면 송신측이 멈추는 BackPressure 현상 확인

## Why?
- Wait 전에 Cancel을 해야 종료 트리거를 받을 수 있음
- 고루틴 생성 비용을 아끼고, 시스템 자원(CPU/MEM)을 예측 가능하게 관리하기 위해서

## Key Insights
- Backpressure: 시스템 부하를 조절하는 천연 브레이크.
- Shutdown Sequence: 신호 전파 -> 종료 대기 -> 자원 정리.

## Mental Model: 교통 시스템?
"파이프라인은 꼬리물기가 금지된 교차로"
1. Backpressure: 앞 차가 나가지 않으면 뒤 차는 교차로에 진입하지 못한다. 버퍼는
교차로 사이의 여유 공간일 뿐, 근본적인 정체 해결 못함

2. 감시 고루틴: 신호등(Wait)이 고장나서 차선 하나를 다 막고있으면, 그 뒤의 모든 차는 영원히 움직이지 못함(deadlock)
그래서 언제 신호를 바꿀지를 고민하는건 운전자나, 회전교차로가 아닌, 제 3자가 해야함

3. tutekdns: 사고가 난다면, 견인차를 부르는 것(wait)보다, 일단 뒤에 오는 차들을
우회하라고 알리는(Cancel) 것이 먼저이다

## Supplement
Buffered 채널이면 별도 고루틴을 쓰지 않아도 될까?
-> 워커의 개수와 Buffer의 크기에 따라 달라짐
이론적으로는 워커가 보낼 데이터가 10개고, 버퍼가 100개라면 데드락 걸리지 않음
하지만 보낼 데이터가 100개보다 커지는 순간 바로 끝, 즉 실무에서는 보통 Buffered 채널이어도
별도 고루틴을 사용(감시 고루틴)
