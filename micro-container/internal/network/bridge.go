package network

import (
	"fmt"
	"net"

	"github.com/vishvananda/netlink"
)

func GetOrCreateBridge(bridgeName string) error {
	
	// 지정하려는 브릿지의 이름이 이미 있는지 확인하는 로직
	// 호스트의 모든 네트워크 장치 가져오기
	// 이름으로 장치 찾아보기
	link, err := netlink.LinkByName(bridgeName)

	// 만약 err가 안뜨면 장치가 존재한다는 뜻이니까 그 장치가 bridge인지 확인하고
	// 그대로 재사용 하기
	if err == nil {
		if link.Type() == "bridge" {
			fmt.Println("log: 이미 존재하는 브릿지입니다. 그대로 재사용하겠습니다.")
			// 브릿지 활성화되어 있는지 확인하기
			if IsBridgeUp(bridgeName) {
				return nil
			}
			fmt.Println("log: 브릿지가 꺼져 있어 활성화 합니다")
			return netlink.LinkSetUp(link)
		}
		return fmt.Errorf("log: 이름은 같지만 브릿지 타입이 아닙니다.") 
	}

	// 만약 err가 뜨면 장치가 없다는 뜻이니까 정상
	// 새로 브릿지 만들기
	// 브릿지 상세 설정 객체 생성하기
	// 객체 이름: linkattributes
	la := netlink.NewLinkAttrs()
	la.Name = bridgeName

	// 브릿지 전용 객체 만들기
	br := &netlink.Bridge{LinkAttrs: la}
	
	// 커널에 브릿지 추가 요청(즉 스위치 만들어달라고 요청하는거임)
	if err := netlink.LinkAdd(br); err != nil{
		return err
	}

	// 브릿지에 ip주소 부여하기, 즉 스위치 주소 지정해주는 것임
	addr, _ := netlink.ParseAddr("172.17.0.1/24") //서브넷 주소 1로

	// 생성한 브릿지에 ip주소 할당하기ㅣ
	if err := netlink.AddrAdd(br, addr); err != nil {
		return err
	}

	return netlink.LinkSetUp(br) // 브릿지 활성화 시키기! 
}

// 브릿지가 활성화되어 있는지 확인하는 메소드
func IsBridgeUp(bridgeName string) bool{
	link, err := netlink.LinkByName(bridgeName)
	if err != nil {
		return false
	}

	isUp := link.Attrs().Flags & net.FlagUp != 0

	return isUp
}
