package chapter3

import (
	"io"
	"net"
	"testing"
) 

func TestDial(t *testing.T) {
	// 랜덤 포트에 리스너 생성하기
	listener, err := net.Listen("tcp", "127.0.0.1:")
	if err != nil {
		t.Fatal(err)
	} 

	// 채널링
	done := make(chan struct{})

	go func() {
		defer func() {done <- struct{}{}} () 

		// 여기서부터 반복문을 돌며 연결을 기다림
		for {
			conn, err := listener.Accept()
			if err != nil {
				t.Log(err)
				return
			}
		
			// 연결이 되면 연결 처리 로직을 담당하는 핸들러 구현
			go func(c net.Conn) {
				// 고루틴 종료 후 채널에 context통신 보장
				defer func() {
					c.Close()
					done <- struct{}{}
				} ()

				// 1024바이트 버퍼 만들기
				buf := make([]byte, 1024)
				// 버퍼 읽고 로그 쓰기
				for {
					// EOF 가 뜨면 고루틴 종료
					n, err := c.Read(buf)
					if err != nil {
						if err != io.EOF{
							t.Error(err)
						}
						return
					}
					t.Logf("received: %q", buf[:n])
				}
			}(conn)
		}
	}()

	// tcp로 같은 네트워크의 connect객체 생성하기
	conn, err := net.Dial("tcp", listener.Addr().String())
	if err != nil {
		t.Fatal(err)
	}

	// 고루틴 끝날때까지 기다릴게요!
	conn.Close() 
	<- done
	listener.Close()
	<- done
}
