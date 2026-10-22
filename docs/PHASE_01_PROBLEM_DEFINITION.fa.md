# FreelanceOps - Phase 1: Problem Definition

Status: In Progress

## هدف

این Document مسئله اصلی FreelanceOps را از دید تجربه واقعی User تعریف می‌کند.

هدف Phase 1 این است که قبل از تعریف Feature، Architecture یا Technology Stack، مشخص شود Product دقیقا قرار است چه مشکلی را حل کند.

---

## مسئله اصلی

Freelancerهایی که همزمان چند Client و Service را مدیریت می‌کنند، اطلاعات عملیاتی و مسئولیت‌های جاری خود را در چند محل مختلف نگهداری می‌کنند.

در وضعیت فعلی، اطلاعات ممکن است بین موارد زیر پراکنده باشند:

- Folderهای Local روی Computer
- Excel
- Google Sheets
- Notes
- Telegram
- WhatsApp
- Mizito
- Jira
- GitHub
- Password Manager

در نتیجه، یک نمای مرکزی و قابل اعتماد برای پاسخ به این سوال وجود ندارد:

> الان برای کدام Client و کدام Service باید اقدامی انجام شود؟

---

## مشکلات مشاهده‌شده

### 1. Renewal و Expiration

User دید مرکزی و قابل اعتمادی از موارد نزدیک به Expiration ندارد.

نمونه موارد:

- Domain Renewal
- SSL Certificate Expiration
- Hosting Renewal
- Service Subscription Renewal

این مشکل می‌تواند باعث شود Expiration فقط زمانی دیده شود که Service وارد وضعیت بحرانی شده یا از دسترس خارج شده است.

### 2. پراکندگی اطلاعات

اطلاعات Clientها در Fileها، Spreadsheetها، Noteها و Toolهای مختلف پخش شده‌اند.

در نتیجه:

- پیدا کردن اطلاعات زمان‌بر است.
- Context مربوط به هر Client یکپارچه نیست.
- بخشی از اطلاعات ممکن است قدیمی یا فراموش شود.
- یک Source of Truth مشخص وجود ندارد.

### 3. نبود Daily Operational View

User همیشه نمی‌داند امروز باید روی چه چیزی تمرکز کند.

Taskها، Renewalها، Maintenanceها و Follow-upها در چند محل مختلف قرار دارند و سیستم واحدی برای اولویت‌بندی Attention وجود ندارد.

### 4. فراموش شدن Maintenance

Maintenanceهای دوره‌ای ممکن است فراموش شوند.

حتی زمانی که برنامه Maintenance مشخص است، نبود Reminder یا Tracking منظم باعث می‌شود اجرای واقعی با برنامه هماهنگ نباشد.

### 5. Backup Verification

برای بررسی Backupها برنامه وجود دارد، اما اجرای آن همیشه منظم نیست.

مسئله اصلی نبود Reminder و Follow-up قابل اعتماد است.

بنابراین مشکل فقط تعریف Backup Policy نیست؛ اطمینان از انجام شدن Checkهای دوره‌ای نیز بخشی از مسئله است.

### 6. پراکندگی مکالمات Client

Communication با Clientها بین چند Platform مختلف پخش شده است.

نمونه:

- WhatsApp
- Telegram
- Mizito
- Jira

در نتیجه ممکن است:

- Request یا Promise مربوط به Client فراموش شود.
- Context یک Conversation در محل دیگری باقی بماند.
- Task حاصل از یک Conversation وارد Workflow اصلی نشود.

---

## Incidentهای واقعی

مشکل تعریف‌شده صرفا فرضی نیست و در تجربه واقعی رخ داده است.

موارد ثبت‌شده:

- SSL Certificate Expire شده است.
- انجام یک Task مربوط به Client فراموش شده است.
- Renewal / Expiration بعضی Serviceها در زمان مناسب دیده نشده است.

این Incidentها نشان می‌دهند که مسئله اصلی تنها Storage اطلاعات نیست؛ مشکل در تبدیل اطلاعات به Action به‌موقع است.

---

## روش فعلی مدیریت

در حال حاضر User از ترکیبی از Toolهای زیر استفاده می‌کند:

- Telegram
- WhatsApp
- Excel
- Google Sheets
- Jira
- GitHub
- Notes
- Password Manager
- Mizito
- Local File / Folder

هیچ‌کدام به تنهایی Operational View کامل Client را ارائه نمی‌دهند.

---

## مهم‌ترین Problem Area

در این مرحله مهم‌ترین مسئله Product به صورت زیر اولویت‌بندی شده است:

> نگهداری و یادآوری Renewal و Expiration مربوط به Serviceهای Client.

این موضوع باید یکی از اولین Valueهای قابل لمس FreelanceOps باشد.

سیستم باید بتواند:

1. تاریخ‌های مهم را نگهداری کند.
2. موارد نزدیک به Expiration را تشخیص دهد.
3. قبل از رسیدن به وضعیت بحرانی Reminder ایجاد کند.
4. موارد نیازمند Action را در یک View مرکزی نمایش دهد.
5. وضعیت Action مربوط به Renewal را Track کند.

---

## Insight اصلی

اطلاعات پراکنده بخشی از مشکل است، اما Problem اصلی صرفا Centralized Storage نیست.

اگر تمام اطلاعات در یک Database جمع شوند ولی User همچنان مجبور باشد خودش دائما آن‌ها را Check کند، مسئله اصلی حل نشده است.

FreelanceOps باید از مدل زیر عبور کند:

```text
Store Information
```

و به مدل زیر برسد:

```text
Store → Understand → Remind → Prioritize → Track Action
```

---

## Product Opportunity

فرصت اصلی FreelanceOps ساخت یک Operational Memory برای Freelancer است.

Product باید بتواند از اطلاعات Clientها نتیجه عملیاتی استخراج کند.

برای مثال، به جای اینکه فقط این Data را ذخیره کند:

```text
example.com
Expiration Date: 2027-01-20
```

باید بتواند به User بگوید:

```text
Action Required

example.com
Domain expires in 14 days.
Renewal has not been confirmed.
```

بنابراین Value اصلی Product از Data Entry ایجاد نمی‌شود؛ از تبدیل Data به Attention و Action ایجاد می‌شود.

---

## Problem Statement - کوتاه

> Freelancerهایی که چند Client را مدیریت می‌کنند، به دلیل پراکندگی اطلاعات و نبود Reminder و Operational View مرکزی، ممکن است Renewal، Maintenance و Taskهای مهم Client را فراموش کنند.

---

## Problem Statement - کامل

Freelancerهایی که همزمان مسئول چند Client و تعداد زیادی Service، Asset و عملیات دوره‌ای هستند، اطلاعات خود را بین Fileها، Spreadsheetها، Noteها، Communication Toolها و سیستم‌های مختلف نگهداری می‌کنند.

به دلیل نبود یک Source of Truth عملیاتی که بتواند تاریخ‌های مهم، Renewalها، Maintenanceها و Taskهای نیازمند Attention را تشخیص داده و به موقع Surface کند، User مجبور است بخش زیادی از وضعیت Clientها را در حافظه خود نگه دارد یا چند Tool را به صورت دستی Check کند.

این وضعیت باعث می‌شود برخی Renewalها یا Maintenanceها دیر دیده شوند، Taskهای Client فراموش شوند و مشکلاتی مانند Expiration یک SSL Certificate قبل از اقدام User رخ دهند.

FreelanceOps باید این Operational State را به صورت مرکزی نگهداری کند و قبل از بحرانی شدن وضعیت، موارد نیازمند Action را به User یادآوری و اولویت‌بندی کند.

---

## فرضیه اولیه ارزش محصول

اگر FreelanceOps بتواند اطلاعات عملیاتی Client را در یک محل نگهداری کند و Renewal، Maintenance و Taskهای مهم را قبل از موعد به User Surface کند، Freelancer وابستگی کمتری به حافظه شخصی و Check دستی Toolهای مختلف خواهد داشت.

نتیجه مورد انتظار:

- Renewal فراموش‌شده کمتر
- Maintenance فراموش‌شده کمتر
- Incident قابل پیشگیری کمتر
- زمان کمتر برای پیدا کردن Context
- Mental Load کمتر
- پاسخ‌گویی منظم‌تر به Client

---

## موارد باز Phase 1

برای تکمیل Phase 1 هنوز باید موارد زیر مشخص شوند:

- [ ] هزینه واقعی فراموش شدن یک Renewal یا Maintenance
- [ ] Frequency رخ دادن این مشکلات
- [ ] Severity هر نوع مشکل
- [ ] مشخص کردن بخش‌هایی که بیشترین Mental Load را ایجاد می‌کنند
- [ ] مشخص کردن Eventهایی که معمولا دیر تشخیص داده می‌شوند
- [ ] نهایی کردن Problem Statement بعد از بررسی موارد بالا

---

## نتیجه فعلی

تا این مرحله، مشکل اصلی FreelanceOps به شکل زیر قابل خلاصه‌سازی است:

> اطلاعات Clientها فقط نیاز به نگهداری ندارند؛ آن‌ها باید به Action قابل پیگیری تبدیل شوند.

و اولین Use Case با Priority بالا:

> Renewal و Expiration Tracking همراه با Reminder قبل از ایجاد مشکل.
