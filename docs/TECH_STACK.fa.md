# FreelanceOps — Technology Stack

Status: Accepted
Version: 1.0
Date: 2026-11-05
Related: [MVP Definition](./MVP_DEFINITION.fa.md)

## هدف و اصول طراحی

این سند استک قطعی پیاده سازی MVP پروژه FreelanceOps را مشخص می کند. هدف، ساخت داشبورد تک کاربره برای مانیتورینگ وب سایت، ثبت تاریخ انقضای Domain و SSL، نگهداری تاریخچه خرابی و ارسال هشدار و دریافت تایید تمدید از طریق Telegram است.

اصول:
- سادگی MVP بدون ایجاد بدهی فنی غیرضروری برای توسعه آینده
- یک Repository و یک Backend اصلی، بدون Microservices
- Frontend کامپوننت محور با TypeScript و قرارداد API روشن
- مصرف منابع پایین، استقرار قابل تکرار، تست و Backup
- عدم ورود قابلیت های خارج از MVP مانند Team، CRM و Multi-user

## استک نهایی

| بخش | انتخاب | دلیل |
| --- | --- | --- |
| Backend | Go | مصرف منابع مناسب، اجرای دائمی، Concurrency و خروجی قابل استقرار ساده |
| HTTP API Router | Chi | سبک، سازگار با net/http |
| Frontend | React | کامپوننت های قابل استفاده مجدد و توسعه بدون بازنویسی معماری UI |
| Frontend Language | TypeScript | قراردادهای نوع دار و خطاهای کمتر حین توسعه |
| Development / Build | Vite | Dev Server سریع و تولید Static Assets |
| Styling | Tailwind CSS | طراحی یکدست و سریع |
| UI Components | shadcn/ui | کامپوننت های قابل کپی و شخصی سازی، نه سرویس وابسته |
| Client Routing | React Router | مدیریت مسیرها؛ فعلا صفحه اصلی و فرم ها |
| Server State | TanStack Query | Fetch، Cache، Refetch و مدیریت وضعیت درخواست |
| Database | SQLite با WAL | بدون سرویس جداگانه، مناسب نسخه تک کاربره |
| DB Driver / Queries | database/sql با SQLite driver سازگار | SQL شفاف و کنترل شده؛ sqlc اختیاری در آینده |
| DB Migrations | golang-migrate | تغییر نسخه بندی شده Schema |
| Website Monitoring | Go net/http | Checkهای HTTP با Timeout و محدودیت Concurrency |
| TLS / SSL | Go crypto/tls و x509 | بررسی Certificate، اعتبار و Expiration |
| Domain Lookup | RDAP اختیاری + Manual Entry | پوشش مواردی که RDAP اطلاعات کامل نمی دهد |
| Periodic Monitoring | Go time.Ticker | اجرای Checkهای دوره ای با Worker Pool |
| Daily Tasks | robfig/cron | Reminder و Maintenance |
| Notification / Actions | Telegram Bot API | Down، Recovery، Renewal Reminder و Confirmation |
| Reverse Proxy / Web Server | Nginx | ارائه Frontend، TLS و Proxy مسیر API |
| Packaging | Docker multi-stage builds | Imageهای قابل تکرار و کوچک |
| Deployment | Docker Compose | اجرای Nginx و Go با volume دیتابیس |
| CI/CD | GitHub Actions | Lint، Test، Build، Deploy و Health Verification |
| Hosting | Linux VPS | پردازش و مانیتورینگ مستقل و همیشه روشن |

## تصمیم Frontend: React + TypeScript + Vite

Vite فریم ورک UI نیست؛ Dev Server و Build Tool است. React لایه UI و TypeScript زبان توسعه Frontend است. کامپوننت ها و Featureها از ابتدا تفکیک شوند تا قابلیت های بعدی بدون بازنویسی کل Frontend اضافه شوند.

- React Router برای Navigation داخلی
- TanStack Query برای دریافت داده و Polling وضعیت، بدون نیاز فعلی به WebSocket
- Tailwind CSS و shadcn/ui برای ظاهر و کامپوننت های قابل تغییر
- از Redux، Next.js، SSR و State Management اضافه در MVP استفاده نمی شود
- یک API قراردادمحور JSON بین React و Go برقرار می شود؛ در آینده امکان تولید Typeها از OpenAPI وجود دارد

آلترناتیوها: Vue + Vite، SvelteKit، Next.js، HTMX. برای داشبورد تعاملی و توسعه تدریجی، React + Vite انتخاب شد. این انتخاب تضمین نمی کند که هیچ Refactor در آینده لازم نباشد، اما جلوی وابستگی Frontend به HTML تولیدشده توسط Backend را می گیرد.

## تصمیم Reverse Proxy: Nginx

Nginx در Production تنها Entry Point روی 80/443 است:
- مدیریت HTTPS و Redirect به HTTPS
- ارائه فایل های Buildشده Vite از مسیر Static
- ارسال /api/* به Go Backend از شبکه داخلی Compose
- ارسال مسیرهای SPA به index.html برای پشتیبانی از Client-side Routing
- ارسال درخواست Telegram Webhook به Backend، در صورت استفاده از Webhook
- تنظیم Headerهای امنیتی و Timeout مناسب

فایل های Static React در زمان Build تولید می شوند؛ Vite Dev Server روی Production اجرا نمی شود. Go و SQLite نباید از اینترنت مستقیم در دسترس باشند.

## ساختار Repository

```text
www-freelanceops/
├── backend/
│   ├── cmd/server/
│   ├── internal/
│   │   ├── api/
│   │   ├── auth/
│   │   ├── websites/
│   │   ├── monitoring/
│   │   ├── domains/
│   │   ├── certificates/
│   │   ├── notifications/
│   │   └── scheduler/
│   ├── migrations/
│   ├── go.mod
│   └── Dockerfile
├── frontend/
│   ├── src/
│   │   ├── app/
│   │   ├── components/
│   │   ├── features/
│   │   ├── pages/
│   │   └── lib/api/
│   ├── package.json
│   ├── tsconfig.json
│   └── vite.config.ts
├── deploy/
│   └── nginx/
│       └── default.conf
├── docs/
│   ├── MVP_DEFINITION.fa.md
│   └── TECH_STACK.fa.md
├── .github/workflows/
└── compose.yaml
```

## ارتباط بخش ها

```text
Browser
   |
 HTTPS
   v
 Nginx --------> React Static Files (Vite Build)
   |
 /api/*
   v
 Go API + Scheduler + Monitoring + Telegram
   |
 SQLite (persistent volume)
```

یک Process از Go مسئول Scheduler و Monitoring در MVP است؛ اگر بعدا چند Replica اضافه شود، باید اجرای Jobها با Leader Election یا Worker مستقل کنترل شود تا هشدار تکراری تولید نشود.

## رفتار فنی MVP

- Website Check پیش فرض هر 60 ثانیه، HTTP Timeout حدود 10 ثانیه
- محدود کردن Checkهای همزمان و توزیع زمان شروع آن ها؛ جلوگیری از ایجاد هزاران اتصال ناگهانی
- Down بعد از دو Check ناموفق متوالی، Recovery بعد از Check موفق
- Health Policy قابل تعریف: HTTP 4xx همیشه به معنی قطع بودن سایت نیست
- ثبت Check/Incident و Recovery به همراه Timestamp و مدت تقریبی قطعی
- پیشگیری از Notification تکراری با Event State و Idempotency
- Domain Reminder: 14 روز قبل از Expiration
- SSL Reminder: 7 روز قبل از Expiration
- Renewal Confirmation از طریق Telegram با ثبت تاریخ واقعی جدید
- RDAP فقط در صورت دسترس پذیری و قابلیت اتکا؛ ورود دستی تاریخ دامنه همیشه ممکن باشد
- UI Status Refresh با TanStack Query، پیش فرض هر 30 ثانیه
- ذخیره Checkهای خام تا 90 روز به صورت پیش فرض؛ نگهداری بلندمدت Incident Summary

## امنیت و عملیات

- یک حساب Admin با Password Hash امن و Session مبتنی بر Cookieهای HttpOnly، Secure و SameSite؛ بدون Multi-user در MVP
- HTTPS اجباری و نگهداری Secretها در Environment/Secret Storage، نه Repository
- Telegram Bot Token محرمانه، اعتبارسنجی فرستنده / Chat ID و Webhook Secret (در صورت Webhook)
- حفاظت از Monitoring در برابر SSRF: کنترل URL، DNS و IP مقصد، جلوگیری از دسترسی به Loopback، Private، Link-local و Metadata Addresses، و اعتبارسنجی Redirectها
- محدودسازی Request Rate و Timeout روی API و Monitoring
- SQLite در Volume پایدار؛ Backup سازگار با WAL و تمرین Restore
- Migration نسخه بندی شده و Startup/Readiness Health Checks
- لاگ های ساخت یافته، بدون Token و داده حساس
- اجرای Containerها با حداقل Permission و بدون Public Port برای Backend/SQLite

## CI/CD و Deployment

1. Validate و Lint برای Go، TypeScript و Docker Compose
2. Test برای Backend و Frontend
3. Build فایل های Static React و Image مربوط به Go و Nginx
4. Deploy نسخه مشخص به VPS از طریق GitHub Actions و Docker Compose
5. Health Check و Smoke Test از مسیر HTTPS
6. در صورت شکست Verification، Rollback به Image قبلی و گزارش خطا

در محیط Production، فقط Nginx پورت عمومی دارد و سرویس Go روی شبکه داخلی Compose است. Certificateها از طریق Nginx و روش انتخاب شده برای ACME مدیریت و تمدید می شوند؛ Nginx به تنهایی Certificate صادر نمی کند.

## آلترناتیوها و دلایل رد فعلی

| بخش | جایگزین ها | علت انتخاب فعلی |
| --- | --- | --- |
| Backend | Node.js/TypeScript، Python/FastAPI، .NET، Rust | Go ساده و کم مصرف برای Workerهای دائم |
| Frontend | Vue + Vite، Svelte، Next.js، HTMX | React + TypeScript قابلیت توسعه و اکوسیستم بالغ دارد |
| Database | PostgreSQL، MySQL | SQLite برای تک کاربره و نوشتن محدود ساده تر است |
| Proxy | Caddy، Traefik، Apache | تصمیم محصول: Nginx |
| UI | Bootstrap، MUI، Chakra UI | Tailwind + shadcn/ui منعطف و بدون وابستگی بصری سخت |
| Scheduler | Linux cron، Asynq، Temporal | Ticker/Cron داخلی برای MVP کافی است |
| CI/CD | GitLab CI، Jenkins، Woodpecker | GitHub Actions یکپارچه با Repository |
| Deployment | systemd، Kubernetes، PaaS | Compose برای یک VPS و تعداد سرویس کم مناسب است |

## خارج از محدوده

در این مرحله Redis، RabbitMQ، Kubernetes، Microservices، PostgreSQL Cluster، GraphQL، SSR، WebSockets، Team، Client Portal، CRM و چند کانال Notification اضافه نمی شوند. افزودن هرکدام باید با نیاز واقعی و سنجش اثر آن توجیه شود.

## نتیجه

**استک نهایی MVP: Go + Chi + SQLite + React + TypeScript + Vite + Tailwind CSS + shadcn/ui + React Router + TanStack Query + Nginx + Docker Compose + GitHub Actions + Telegram Bot API.**

ساختار از روز اول به صورت Frontend جداگانه و Backend REST API در یک Repository طراحی می شود؛ هدف کاهش احتمال بازنویسی در توسعه های بعدی است، نه حذف کامل نیاز به Refactor.
