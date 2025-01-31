# MyStats 서버

Go, Fiber, PostgreSQL을 사용한 야구 기록 관리 서버

## 프로젝트 구조

```
mystats-server/
├── cmd/
│   └── server/
│       └── main.go           # 메인 애플리케이션 진입점
├── internal/
│   ├── config/
│   │   └── database.go       # 데이터베이스 설정
│   └── ent/
│       └── schema/           # 데이터베이스 스키마
│           ├── user.go
│           ├── game.go
│           ├── batter_record.go
│           └── pitcher_record.go
└── docker-compose.yml        # PostgreSQL용 Docker 설정
```

## 기술 스택

- Go
- Fiber (웹 프레임워크)
- PostgreSQL (데이터베이스)
- Ent (ORM)

## 데이터베이스 스키마

- **User**: 사용자 정보 저장
- **Game**: 경기 기록 관리
- **BatterRecord**: 타격 기록 저장
- **PitcherRecord**: 투구 기록 저장

## 개발 환경 설정

1. PostgreSQL 실행:
```bash
docker compose up -d
```

2. 서버 실행:
```bash
go run ./cmd/server/main.go
```

서버는 `http://localhost:8080`에서 실행됩니다.

## 다음 단계

- 사용자 인증 구현
- CRUD 작업을 위한 REST API 엔드포인트 생성
- 인증 및 로깅을 위한 미들웨어 추가
- Flutter 프론트엔드와 통합
