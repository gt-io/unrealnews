# unrealnews
>30초마다 언리얼 이슈를 긁어서 슬랙이나 줄립으로 메세지를 보냄.

# 사용법
1. https://github.com/gt-io/unrealnews/releases 에서 실행파일 다운로드, 압축해제.
2. conf.json 파일을 열어서 줄립 이나 슬랙 bot 환경설정. ( 자세한 설명 생략 )
3. .version 파일을 열어서 검색할 unreal version 셋팅.
4. unrealnew.exe 실행. 

# 윈도우 서비스 기능 ( 관리자권한 )
1. 서비스 등록
```sc create unrealnews binPath= %~dp0\unrealnews.exe start= auto```
2. 서비스 실행
```sc start unrealnews```
3. 서비스 중지
```sc stop unrealnews```
3. 서비스 삭제
```sc delete unrealnews```
