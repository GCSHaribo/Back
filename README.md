# Back

## 협업 방법

1. Issue에서 자신이 담당하고 있는 컴포넌트를 확인하고 본인을 Assignee로 등록하세요.
   * 본인이 하고싶은 컴포넌트에 본인을 assignee로 등록하세요
2. Projects의 각 카드는 Issue에 등록되어있는 내용과 동일합니다.
3. develop 브랜치에서 새로운 브랜치를 생성해서 작업해주세요.
4. {seq}-{desc}와 같은 식으로 브랜치를 생성합니다. desc는 설명이므로 자유롭게 하셔도 상관없습니다.
5. commit을 완료하면 develop branch에 PR을 날려주세요.
6. 충돌하는 부분이 없는지만 확인하고 merge합니다.
7. **master브랜치에 commit하거나 PR을 날리는 일은 없어야 합니다.**
8. API명세서는 swagger를 사용하고 Wiki에 문서화합니다.
9. Back-End 각 url마다 routing하는 각각의 handler를 만드는 방식입니다

## COMMIT MESSAGE CONVENTION
기본적으로 COMMIT 메시지는 제목/본문/꼬리말로 작성
```none
type: 제목
본문
꼬리말
```
COMMIT TYPE
1. feat:새로운 기능 추가
2. fix: 버그픽스
3. docs: 문서(api 기능서 등) 수정
4. refactor: 코드 리펙토링
5. test: 테스트코드
6. chore:파일 구조를 바꾸거나 패키지를 새로 다운 받는 경우

제목은 필수지만 본문,꼬리말은 필수사항이 아니니 자신이 덧붙이고 싶으 말이나 지침 등을 적어주시면 됩니다.
과거시제 사용 x

커밋작성 예시
```
feat: 로그인 핸들러 추가
로그인 핸들러를 추가했습니다
```
예시에서는 한글로 작성했지만 되도록이면 영어로 작성하도록 
