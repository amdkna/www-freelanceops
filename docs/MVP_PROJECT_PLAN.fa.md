# FreelanceOps — MVP Project Management Plan

Status: Ready for Implementation
Version: 1.0
Date: 2026-11-05
Repository: [amdkna/www-freelanceops](https://github.com/amdkna/www-freelanceops)

## ۱. هدف و منبع تصمیم ها

این سند Backlog اجرایی نسخه اول FreelanceOps است و به دو سند پذیرفته شده [MVP Definition](./MVP_DEFINITION.fa.md) و [Technology Stack](./TECH_STACK.fa.md) متکی است. هدف محصول: فریلنسر بتواند سایت ها، وضعیت دسترسی، تاریخ انقضای دامنه و SSL، تاریخچه خرابی و اعلان های تلگرام را در یک ابزار مدیریت کند.

قابلیت های خارج از MVP مثل Multi-user، CRM، Client Portal، حسابداری، Monitoring سرور و Notification چند کاناله وارد فاز انتشار نمی شوند. معماری اصلی: Go + Chi + SQLite، React + TypeScript + Vite، Nginx، Docker Compose و GitHub Actions.

## ۲. مدل مدیریت پروژه

**Epic → Feature → User Story → Task (GitHub Issue).** هر Epic یک خروجی تجاری بزرگ و هر Feature یک قابلیت مشخص است. هر Task یک Issue مستقل با شناسه، اولویت، مراحل پیاده سازی، Test Plan، Acceptance Criteria و Definition of Done دارد. در هر Issue یک User Story مرتبط ثبت شده است.

- **P0 — Release blocker:** نبود یا خرابی آن، MVP عملیاتی یا امن را غیرممکن می کند.
- **P1 — Important:** باید ترجیحا پیش از انتشار عمومی کامل شود، ولی توسعه P0 را متوقف نمی کند.
- **P2 — Enhancement:** بعد از کارکرد مسیر اصلی یا در Release بعدی انجام می شود.
- **P3 — Optional:** فقط بعد از تثبیت محصول و در صورت وجود ظرفیت.

برچسب اولویت در عنوان همه Issueها درج شده، مثلا `[P0][E03/F06] T11 — ...`. شماره Issue همان Task ID در این Backlog است؛ عنوان Epic و Feature نیز برای فیلتر و جستجوی سریع درج می شود.

## ۳. خلاصه حجم کار

- **7 Epic**
- **17 Feature**
- **37 User Story / Task / GitHub Issue** (هر Task یک User Story مشخص دارد)
- **P0: 27**، **P1: 6**، **P2: 3**، **P3: 1**

> در این نسخه برای جلوگیری از پیچیدگی مدیریتی، برای هر Task یک User Story مستقیم داخل همان Issue ثبت شده است. در صورت بزرگ شدن تیم، می توان User Storyها را به Issueهای والد و Taskها را به Sub-issue تبدیل کرد.

## ۴. Backlog بر اساس Epic و Feature

### E01 — Foundation & Security

**هدف:** زیرساخت امن و قابل توسعه

#### F01 — Project Setup

**User Storyهای این Feature:** به عنوان یک فریلنسر می خواهم زیرساخت امن و قابل توسعه تا بتوانم وضعیت و تمدید سایت های خود را بدون ابزارهای متعدد مدیریت کنم. هر Task در Issue مربوطه یک User Story عملیاتی دارد.

| Issue | Priority | Task |
| --- | --- | --- |
| [T01](https://github.com/amdkna/www-freelanceops/issues/1) | P0 | راه اندازی Go/Chi |
| [T02](https://github.com/amdkna/www-freelanceops/issues/2) | P0 | راه اندازی React/TS/Vite |
| [T03](https://github.com/amdkna/www-freelanceops/issues/3) | P0 | قرارداد REST API |

#### F02 — Database & Authentication

**User Storyهای این Feature:** به عنوان یک فریلنسر می خواهم زیرساخت امن و قابل توسعه تا بتوانم وضعیت و تمدید سایت های خود را بدون ابزارهای متعدد مدیریت کنم. هر Task در Issue مربوطه یک User Story عملیاتی دارد.

| Issue | Priority | Task |
| --- | --- | --- |
| [T04](https://github.com/amdkna/www-freelanceops/issues/4) | P0 | SQLite و Migration |
| [T05](https://github.com/amdkna/www-freelanceops/issues/5) | P0 | ورود و Session مدیر |

#### F03 — Operations Baseline

**User Storyهای این Feature:** به عنوان یک فریلنسر می خواهم زیرساخت امن و قابل توسعه تا بتوانم وضعیت و تمدید سایت های خود را بدون ابزارهای متعدد مدیریت کنم. هر Task در Issue مربوطه یک User Story عملیاتی دارد.

| Issue | Priority | Task |
| --- | --- | --- |
| [T06](https://github.com/amdkna/www-freelanceops/issues/6) | P1 | Lint، تست و لاگ ساخت یافته |

### E02 — Website Inventory

**هدف:** مدیریت فهرست سایت ها

#### F04 — Website API

**User Storyهای این Feature:** به عنوان یک فریلنسر می خواهم مدیریت فهرست سایت ها تا بتوانم وضعیت و تمدید سایت های خود را بدون ابزارهای متعدد مدیریت کنم. هر Task در Issue مربوطه یک User Story عملیاتی دارد.

| Issue | Priority | Task |
| --- | --- | --- |
| [T07](https://github.com/amdkna/www-freelanceops/issues/7) | P0 | CRUD سایت |
| [T08](https://github.com/amdkna/www-freelanceops/issues/8) | P0 | امنیت URL و SSRF |

#### F05 — Website Interface

**User Storyهای این Feature:** به عنوان یک فریلنسر می خواهم مدیریت فهرست سایت ها تا بتوانم وضعیت و تمدید سایت های خود را بدون ابزارهای متعدد مدیریت کنم. هر Task در Issue مربوطه یک User Story عملیاتی دارد.

| Issue | Priority | Task |
| --- | --- | --- |
| [T09](https://github.com/amdkna/www-freelanceops/issues/9) | P0 | فرم افزودن و ویرایش سایت |
| [T10](https://github.com/amdkna/www-freelanceops/issues/10) | P0 | لیست مرکزی سایت ها |

### E03 — Availability Monitoring

**هدف:** تشخیص خرابی و ثبت تاریخچه

#### F06 — HTTP Probes

**User Storyهای این Feature:** به عنوان یک فریلنسر می خواهم تشخیص خرابی و ثبت تاریخچه تا بتوانم وضعیت و تمدید سایت های خود را بدون ابزارهای متعدد مدیریت کنم. هر Task در Issue مربوطه یک User Story عملیاتی دارد.

| Issue | Priority | Task |
| --- | --- | --- |
| [T11](https://github.com/amdkna/www-freelanceops/issues/11) | P0 | موتور HTTP Check |
| [T12](https://github.com/amdkna/www-freelanceops/issues/12) | P0 | Scheduler و Worker Pool |

#### F07 — Incident Lifecycle

**User Storyهای این Feature:** به عنوان یک فریلنسر می خواهم تشخیص خرابی و ثبت تاریخچه تا بتوانم وضعیت و تمدید سایت های خود را بدون ابزارهای متعدد مدیریت کنم. هر Task در Issue مربوطه یک User Story عملیاتی دارد.

| Issue | Priority | Task |
| --- | --- | --- |
| [T13](https://github.com/amdkna/www-freelanceops/issues/13) | P0 | منطق DOWN/Recovery |
| [T14](https://github.com/amdkna/www-freelanceops/issues/14) | P0 | ذخیره Check و Incident |
| [T15](https://github.com/amdkna/www-freelanceops/issues/15) | P1 | رابط تاریخچه خرابی |
| [T16](https://github.com/amdkna/www-freelanceops/issues/16) | P2 | پاکسازی داده های قدیمی |

### E04 — Domain & SSL Expiration

**هدف:** مدیریت انقضای دامنه و گواهی

#### F08 — Certificate Status

**User Storyهای این Feature:** به عنوان یک فریلنسر می خواهم مدیریت انقضای دامنه و گواهی تا بتوانم وضعیت و تمدید سایت های خود را بدون ابزارهای متعدد مدیریت کنم. هر Task در Issue مربوطه یک User Story عملیاتی دارد.

| Issue | Priority | Task |
| --- | --- | --- |
| [T17](https://github.com/amdkna/www-freelanceops/issues/17) | P0 | بررسی SSL |
| [T18](https://github.com/amdkna/www-freelanceops/issues/18) | P0 | هشدار ۷ روزه SSL |

#### F09 — Domain Renewal

**User Storyهای این Feature:** به عنوان یک فریلنسر می خواهم مدیریت انقضای دامنه و گواهی تا بتوانم وضعیت و تمدید سایت های خود را بدون ابزارهای متعدد مدیریت کنم. هر Task در Issue مربوطه یک User Story عملیاتی دارد.

| Issue | Priority | Task |
| --- | --- | --- |
| [T19](https://github.com/amdkna/www-freelanceops/issues/19) | P0 | ثبت دستی انقضا و تمدید |
| [T20](https://github.com/amdkna/www-freelanceops/issues/20) | P0 | هشدار ۱۴ روزه دامنه |
| [T21](https://github.com/amdkna/www-freelanceops/issues/21) | P2 | تشخیص اختیاری RDAP |

### E05 — Telegram Notifications

**هدف:** ارسال هشدار و تایید تمدید

#### F10 — Bot Connection

**User Storyهای این Feature:** به عنوان یک فریلنسر می خواهم ارسال هشدار و تایید تمدید تا بتوانم وضعیت و تمدید سایت های خود را بدون ابزارهای متعدد مدیریت کنم. هر Task در Issue مربوطه یک User Story عملیاتی دارد.

| Issue | Priority | Task |
| --- | --- | --- |
| [T22](https://github.com/amdkna/www-freelanceops/issues/22) | P0 | ارتباط امن Telegram Bot |
| [T23](https://github.com/amdkna/www-freelanceops/issues/23) | P0 | هشدار DOWN/Recovery |

#### F11 — Expiry Reminders

**User Storyهای این Feature:** به عنوان یک فریلنسر می خواهم ارسال هشدار و تایید تمدید تا بتوانم وضعیت و تمدید سایت های خود را بدون ابزارهای متعدد مدیریت کنم. هر Task در Issue مربوطه یک User Story عملیاتی دارد.

| Issue | Priority | Task |
| --- | --- | --- |
| [T24](https://github.com/amdkna/www-freelanceops/issues/24) | P0 | Reminder دامنه و SSL |

#### F12 — Renewal Flow

**User Storyهای این Feature:** به عنوان یک فریلنسر می خواهم ارسال هشدار و تایید تمدید تا بتوانم وضعیت و تمدید سایت های خود را بدون ابزارهای متعدد مدیریت کنم. هر Task در Issue مربوطه یک User Story عملیاتی دارد.

| Issue | Priority | Task |
| --- | --- | --- |
| [T25](https://github.com/amdkna/www-freelanceops/issues/25) | P0 | تایید تمدید در تلگرام |
| [T26](https://github.com/amdkna/www-freelanceops/issues/26) | P1 | تاریخچه و Retry اعلان |

### E06 — Dashboard & UX

**هدف:** نمای یکپارچه و تجربه کاربری

#### F13 — Overview Dashboard

**User Storyهای این Feature:** به عنوان یک فریلنسر می خواهم نمای یکپارچه و تجربه کاربری تا بتوانم وضعیت و تمدید سایت های خود را بدون ابزارهای متعدد مدیریت کنم. هر Task در Issue مربوطه یک User Story عملیاتی دارد.

| Issue | Priority | Task |
| --- | --- | --- |
| [T27](https://github.com/amdkna/www-freelanceops/issues/27) | P0 | داشبورد وضعیت و Attention |
| [T28](https://github.com/amdkna/www-freelanceops/issues/28) | P1 | به روزرسانی خودکار |

#### F14 — Usability

**User Storyهای این Feature:** به عنوان یک فریلنسر می خواهم نمای یکپارچه و تجربه کاربری تا بتوانم وضعیت و تمدید سایت های خود را بدون ابزارهای متعدد مدیریت کنم. هر Task در Issue مربوطه یک User Story عملیاتی دارد.

| Issue | Priority | Task |
| --- | --- | --- |
| [T29](https://github.com/amdkna/www-freelanceops/issues/29) | P1 | ریسپانسیو و دسترس پذیری |
| [T30](https://github.com/amdkna/www-freelanceops/issues/30) | P2 | حالت های خطا/خالی/بارگذاری |

### E07 — Deployment & Quality

**هدف:** انتشار امن، بازیابی و پذیرش

#### F15 — Container Runtime

**User Storyهای این Feature:** به عنوان یک فریلنسر می خواهم انتشار امن، بازیابی و پذیرش تا بتوانم وضعیت و تمدید سایت های خود را بدون ابزارهای متعدد مدیریت کنم. هر Task در Issue مربوطه یک User Story عملیاتی دارد.

| Issue | Priority | Task |
| --- | --- | --- |
| [T31](https://github.com/amdkna/www-freelanceops/issues/31) | P0 | Docker و Compose |
| [T32](https://github.com/amdkna/www-freelanceops/issues/32) | P0 | Nginx و HTTPS |

#### F16 — CI/CD

**User Storyهای این Feature:** به عنوان یک فریلنسر می خواهم انتشار امن، بازیابی و پذیرش تا بتوانم وضعیت و تمدید سایت های خود را بدون ابزارهای متعدد مدیریت کنم. هر Task در Issue مربوطه یک User Story عملیاتی دارد.

| Issue | Priority | Task |
| --- | --- | --- |
| [T33](https://github.com/amdkna/www-freelanceops/issues/33) | P0 | GitHub Actions |
| [T34](https://github.com/amdkna/www-freelanceops/issues/34) | P0 | Verify و Rollback |

#### F17 — Backup & Launch

**User Storyهای این Feature:** به عنوان یک فریلنسر می خواهم انتشار امن، بازیابی و پذیرش تا بتوانم وضعیت و تمدید سایت های خود را بدون ابزارهای متعدد مدیریت کنم. هر Task در Issue مربوطه یک User Story عملیاتی دارد.

| Issue | Priority | Task |
| --- | --- | --- |
| [T35](https://github.com/amdkna/www-freelanceops/issues/35) | P0 | Backup/Restore SQLite |
| [T36](https://github.com/amdkna/www-freelanceops/issues/36) | P1 | Runbook و تست پذیرش |
| [T37](https://github.com/amdkna/www-freelanceops/issues/37) | P3 | ابزارهای کمک عیب یابی |

## ۵. ترتیب انجام و وابستگی ها

این Backlog **Sprint با تاریخ قطعی نیست**؛ ترتیب زیر براساس وابستگی فنی تنظیم شده و زمان تخمینی بعد از اولین اجرای واقعی و تعیین ظرفیت توسعه مشخص می شود.

| مرحله | خروجی قابل تحویل | Issueهای کلیدی | شرط عبور |
| --- | --- | --- | --- |
| ۱. Foundation | اسکلت Backend/Frontend، قرارداد API، DB و Login | #1–#5 | شروع محلی، Migration، ورود امن و تست پایه |
| ۲. Inventory | اضافه کردن و مدیریت سایت ها | #7–#10 | CRUD و محافظت SSRF تست شده |
| ۳. Monitoring | بررسی دوره ای، Down و Recovery، تاریخچه | #11–#14 | خرابی ساختگی تشخیص و ثبت شود |
| ۴. Expiration | SSL و Domain expiry با هشدارهای درست | #17–#20 | مرزهای ۷ و ۱۴ روز تست شوند |
| ۵. Telegram | هشدار خرابی، Reminder و تایید تمدید | #22–#25 | دریافت اعلان و ثبت تاریخ جدید پایان به پایان |
| ۶. Dashboard | نمای مرکزی کاربر | #27 | همه موارد نیازمند توجه در یک صفحه دیده شوند |
| ۷. Production | Nginx، Docker، CI/CD، Rollback، Backup | #31–#35 | انتشار HTTPS و تمرین شکست و بازیابی |
| ۸. Polish & Acceptance | تاریخچه UI، Refresh، دسترس پذیری و Runbook | #6, #15, #26, #28, #29, #36 | آزمون پذیرش MVP تکمیل شود |
| ۹. Deferred | Retention، RDAP، Error polish، Diagnostics | #16, #21, #30, #37 | بدون وابستگی به انتشار P0 |

### وابستگی های بحرانی

- #1 و #3 پیش نیاز پیاده سازی APIها هستند؛ #2 و #3 پیش نیاز اجزای Frontend.
- #4 و #5 پیش نیاز ذخیره سازی و دسترسی امن به داده ها هستند.
- #7 و #8 قبل از فعال کردن #11 و #12 ضروری هستند؛ SSRF باید پیش از بررسی URLهای ورودی حل شود.
- #11 → #12 → #13 → #14 مسیر اصلی Monitor و ثبت Incident است.
- #17 → #18 و #19 → #20 مسیر تشخیص و Reminder انقضا هستند.
- #22 پیش نیاز #23، #24 و #25 است. #23 به #13/#14 و #24 به #18/#20 وابسته است.
- #27 به API سایت، وضعیت مانیتورینگ و انقضا نیاز دارد.
- #31 و #32 پیش نیاز انتشار Production هستند؛ #33 و #34 مسیر انتشار قابل اطمینان را تکمیل می کنند؛ #35 شرط بازیابی داده است.

## ۶. Definition of Ready

قبل از شروع هر Task باید این موارد روشن باشند: هدف قابل مشاهده، ورودی/خروجی API در صورت وجود، وابستگی ها، محدوده امنیتی، Acceptance Criteria و Test Plan. Issueهایی که به اطلاعات محیط Production نیاز دارند باید مقادیر را با Secretها و نمونه فایل تنظیمات مشخص کنند، نه با رمز داخل Repository.

## ۷. Definition of Done

هر Issue زمانی Done است که:

1. پیاده سازی با استک مصوب انجام شده باشد.
2. تست های مرتبط موفقیت، شکست و حالت مرزی نوشته و اجرا شده باشند.
3. Lint، Typecheck، Test و Build مربوطه در CI موفق شوند.
4. کنترل های امنیتی مربوطه (Authorization، SSRF، Secret handling و...) پاس شوند.
5. مستندات یا نمونه API در صورت نیاز به روز شوند.
6. Pull Request به main ادغام شود و شواهد بررسی در Issue ثبت شود.
7. برای تسک انتشار، Health Check و در صورت نیاز Rollback نیز آزمایش شده باشد.

## ۸. سناریوهای پذیرش کامل MVP (End-to-End)

- [ ] مدیر وارد داشبورد HTTPS می شود و کاربر ناشناس به داده ها دسترسی ندارد.
- [ ] سایت جدید با فرم ساده اضافه و در لیست مرکزی نشان داده می شود.
- [ ] سایت هر حدود ۶۰ ثانیه بررسی می شود؛ دو شکست متوالی باعث DOWN و یک موفقیت باعث Recovery می شود.
- [ ] DOWN و Recovery به صورت Incident ذخیره و Telegram Notification یکبار ارسال می شوند.
- [ ] سابقه زمان شروع، پایان و مدت حدودی قطعی قابل مشاهده است.
- [ ] تاریخ انقضای دامنه به صورت دستی قابل وارد کردن است و Reminder در ۱۴ روز مانده ارسال می شود.
- [ ] تاریخ SSL خوانده می شود، اعتبار Certificate مشخص است و Reminder در ۷ روز مانده ارسال می شود.
- [ ] مدیر از طریق Telegram تمدید را تایید می کند و تاریخ واقعی جدید ثبت می شود.
- [ ] جدول اصلی Healthy/Down، Domain/SSL expiry و Attention را نشان می دهد.
- [ ] برنامه در صورت Restart داده های SQLite را حفظ می کند.
- [ ] نسخه ناسالم به صورت آزمایشی Deploy می شود و Rollback به آخرین نسخه سالم جواب می دهد.
- [ ] Backup سازگار با WAL ایجاد و روی محیط جدا Restore می شود.

## ۹. ریسک ها و کنترل ها

| ریسک | کنترل |
| --- | --- |
| SSRF با URL ورودی کاربر | #8 پیش از فعال کردن Monitor؛ بررسی DNS، IP و Redirect |
| اعلان تکراری | #13، #23، #24 و #26 با Event ID و Idempotency |
| خطای RDAP یا پشتیبانی نشدن TLD | #19 ثبت دستی ضروری و #21 اختیاری |
| خرابی داده SQLite | #4 WAL و Migration، #35 Backup/Restore آزمایش شده |
| انتشار خراب | #34 Health Verification و Rollback |
| افشای Bot Token یا Session | #5، #22 و تنظیمات امن با Secrets |
| رشد نامحدود داده ها | #16 Retention قابل اضافه شدن پس از MVP |

## ۱۰. روش کار پیشنهادی در GitHub

- یک Repository و Branch اصلی `main` حفظ شود؛ هر Task در Branch کوتاه مدت مانند `feat/t11-http-probes` پیاده شود.
- عنوان Commit/PR شامل Task ID باشد، مثل `feat(monitoring): implement probes (T11)`.
- PR به Issue مربوطه با `Closes #11` متصل شود.
- در شروع فقط Issueهای آماده از نظر Dependencies به In Progress منتقل شوند؛ تعداد کارهای همزمان را محدود نگه دارید.
- اولویت عنوان Issue مبنای تصمیم است؛ GitHub Issues به طور پیش فرض Priority داخلی رسمی ندارد.
- وضعیت های پیشنهادی: Backlog → Ready → In Progress → In Review → Done؛ در صورت استفاده از GitHub Projects قابل پیاده سازی است.

## ۱۱. اولین اقدام عملی

از **#1، #2، #3 و #4** آغاز کنید؛ بعد **#5** را ببندید. با این پنج Issue، پایه کدنویسی واقعی محصول حاضر است. همزمان #31 و #32 را فقط در حد الگوی استقرار محلی بررسی کنید، نه Deploy Production قبل از تکمیل امنیت و تست.

---

**خلاصه:** ۷ Epic، ۱۷ Feature و ۳۷ Issue اجرایی با اولویت P0 تا P3. وظایف P0 مسیر ساخت، امنیت، استقرار و بازیابی MVP هستند. جزئیات مراحل اجرا و تست در متن Issueهای GitHub ثبت شده اند.
