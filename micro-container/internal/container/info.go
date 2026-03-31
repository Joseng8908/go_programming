package container

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

// 정보 반환용 객체입니다
type ContainerInfo struct {
	ID         string     `json:"id"`
	Pid        int        `json:"pid"`
	Command    string     `json:"command"`
	CreateTime time.Time  `json:"CreateTime"`
	Status     string     `json:"status"`

	// 추후 추가할 것: IP주소, 메모리 제한 정도 등등
}

// 고정적으로 info를 저장할 주소의 위치
const InfoLocation = "/var/lib/micro-container/containers"

func WriteContainerInfo(info *ContainerInfo) error {
	// 저장할 폴더 경로 생성하기
	dirPath := filepath.Join(InfoLocation, info.ID)
	if err := os.MkdirAll(dirPath, 0755); err != nil {
		return fmt.Errorf("정보 저장 폴더 생성 실패...: %v", err)
	}

	// 파일 생성
	configPath := filepath.Join(dirPath, "config.json")
	file, err := os.Create(configPath)
	if err != nil {
		return fmt.Errorf("정보 파일 생성 실패...: %v", err)
	}
	defer file.Close()

	// json으로 마샬링해서 파일에 쓰기
	jsonBytes, err := json.MarshalIndent(info, "", "  ")
	if err != nil {
		return fmt.Errorf("JSON 변환 실패... : %v", err)
	}

	if _, err := file.Write(jsonBytes); err != nil {
		return fmt.Errorf("파일 쓰기 실패")
	}

	return nil
}

func ReadContainerInfo(containerID string) (*ContainerInfo, error) {
	configPath := filepath.Join(InfoLocation, containerID, "config.json")

	content, err := os.ReadFile(configPath)
	if err != nil {
		return nil, err
	}

	var info ContainerInfo
	if err := json.Unmarshal(content, &info); err != nil {
		return nil, err
	}

	return &info, nil
}

 
// 모든 컨테이너 정보를 가져오는 ps 명령어 구현 핵심
func GetInfoList() ([]*ContainerInfo, error) {
	var containers []*ContainerInfo

	// 컨테이너가 저장된 디렉토리의 모든 폴더 읽어오기
	files, err := os.ReadDir(InfoLocation)
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		if file.IsDir() {
			// 여기서 file.Name()넣어도 되는데 containerID를 넣는게 좋아보이기도 함
			info, err := ReadContainerInfo(file.Name())
			if err == nil {
				containers = append(containers, info)
			}
		}
	}

	return containers, nil
}
