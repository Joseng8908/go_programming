package network

import (
	"testing"

	"github.com/vishvananda/netlink"
)

func TestGetOrCreateBridge(t *testing.T) {
	bridgeName := "test-bridge0"

	// 1. 테스트 전 청소: 혹시 남아있을지 모를 브릿지 삭제
	cleanup(bridgeName)
	defer cleanup(bridgeName)

	// 2. 브릿지 생성 테스트
	t.Run("CreateNewBridge", func(t *testing.T) {
		err := GetOrCreateBridge(bridgeName)
		if err != nil {
			t.Fatalf("브릿지 생성 실패: %v", err)
		}

		// 실제로 생성되었는지 커널에서 확인
		link, err := netlink.LinkByName(bridgeName)
		if err != nil {
			t.Errorf("생성된 브릿지를 찾을 수 없음: %v", err)
		}

		if link.Type() != "bridge" {
			t.Errorf("장치 타입이 bridge가 아님: %s", link.Type())
		}

		// IP 주소가 제대로 할당되었는지 확인
		addrs, _ := netlink.AddrList(link, netlink.FAMILY_V4)
		foundIP := false
		for _, addr := range addrs {
			if addr.IP.String() == "172.17.0.1" {
				foundIP = true
				break
			}
		}
		if !foundIP {
			t.Error("IP 주소(172.17.0.1)가 브릿지에 할당되지 않았습니다.")
		}
	})

	// 3. 재사용 테스트 (이미 존재하는 경우)
	t.Run("ReuseExistingBridge", func(t *testing.T) {
		err := GetOrCreateBridge(bridgeName)
		if err != nil {
			t.Errorf("기존 브릿지 재사용 시 에러 발생: %v", err)
		}
		
		if !IsBridgeUp(bridgeName) {
			t.Error("브릿지가 Up 상태여야 하는데 Down 상태입니다.")
		}
	})
}

// 테스트용 헬퍼 함수: 생성된 네트워크 장치 삭제
func cleanup(name string) {
	link, err := netlink.LinkByName(name)
	if err == nil {
		netlink.LinkDel(link)
	}
}
