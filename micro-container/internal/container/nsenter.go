package container

/*
#define _GNU_SOURCE
#include <fcntl.h>
#include <sched.h>
#include <unistd.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

__attribute__((constructor)) void enter_namespace() {
    // 1. 환경변수에서 대상 PID와 실행할 명령어를 가져옴
    char *pid = getenv("target_pid");
    char *cmd = getenv("target_cmd")
    
    // exec 명령어가 아닐 때는 그냥 통과 (run, ps 등)
    if (!pid || !cmd) return;

    // 2. 네임스페이스 5종 세트 진입
    char *namespaces[] = {"ipc", "uts", "net", "pid", "mnt"};
    for(int i = 0; i < 5; i++) {
        char path[100];
        sprintf(path, "/proc/%s/ns/%s", pid, namespaces[i]);
        int fd = open(path, O_RDONLY);
        if (fd != -1) {
            setns(fd, 0);
            close(fd);
        }
    }

    // 3. [마법의 구간] 컨테이너의 루트 파일시스템으로 시야 고정
    // setns(mnt) 후에 실행해야 효과가 있음
    char root_path[100];
    sprintf(root_path, "/proc/%s/root", pid);
    chdir(root_path);
    chroot(".");
    chdir("/");

    // 4. 환경변수 정리 (무한 루프 방지 및 보안)
    unsetenv("target_pid");
    unsetenv("target_cmd");

    // 5. 프로세스 교체: 이제 이 프로세스는 /bin/sh가 됨
    char *args[] = {cmd, NULL};
    execvp(cmd, args);

    // 여기까지 왔다면 에러 발생
    perror("execvp 실패");
    exit(1);
}
*/
import "C"

// Go 쪽에서는 호출만 할 뿐, 실제 로직은 위 C 생성자에서 다 끝남
func NsEnter() {}
