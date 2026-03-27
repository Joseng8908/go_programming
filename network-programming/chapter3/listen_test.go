package chapter3

import(
	"net"
	"testing"
)

func TestListener(t *testing.T){
	// net.Listen 함수는 네트워크의 종류(tcp)와 콜론으로 구분된 IP주소와
	// 포트 문자열을 매개변수로 받음
	// 반환값으로는 net.Listener 인터페이스와 에러 인터페이스를 반환
	// 함수가 성공적으로 반환되면 리스너는 특정 ip주소와 포트번호에 바인딩
	// binding이란, 운영체제가 지정된 IP주소의 포트를 해당 리스너에게 단독으로 할당!
	// 포트를 0으로 설정하면 무작위 포트번호 할당
	// IPv4만 사용하려면 tcp4 사용
	// IPv6까지 사용하려면 tcp6 사용
	listener, err := net.Listen("tcp", "127.0.0.1:0")

	
	if err != nil {
		t.Fatal(err)
	}

	// defer 구문과 Close() 메소드를 사용하여 항상 리스너를 종료하는 것을 확정
	// deadlock 가능성 없애기
	defer func() {_ = listener.Close()} ()

	t.Logf("bound to %q", listener.Addr())
//	// tcp 수신 연결 응답 및 요청 처리하기
//	for {
//		conn, err := listener.Accept()
//		if err != nil {
//			return err
//		}
//		go func(c net.Conn) {
//			defer c.Close() 
//
//			// TCP연결을 사용하여 비지니스 로직 작성
//		} (conn)
//	}
}
