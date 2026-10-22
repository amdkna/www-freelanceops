# FreelanceOps - MVP Definition

Status: Accepted

## هدف

این Document محدوده نسخه اول FreelanceOps را مشخص می‌کند.

هدف MVP این است که یک Freelancer بتواند وضعیت Websiteها، Domainها و SSL Certificateهای خود را در یک محل ببیند و در صورت خرابی یا نزدیک شدن به Expiration از طریق Telegram اطلاع دریافت کند.

---

## کاربر نسخه اول

نسخه اول فقط برای یک User ساخته می‌شود:

> Freelancer

در حال حاضر نیازی به Team، Client Portal، Workspace یا Multi-user وجود ندارد.

---

## ارزش اصلی محصول

MVP باید سه مشکل را همزمان حل کند:

1. Renewalها فراموش نشوند.
2. خرابی Websiteها دیر دیده نشود.
3. وضعیت همه موارد در یک صفحه و یک محل مشخص باشد.

تعریف کوتاه:

> FreelanceOps باید Websiteها را Monitor کند، تاریخ‌های مهم را نگه دارد و در زمان مناسب گزارش و Reminder ارسال کند.

---

## صفحه اصلی

User باید بتواند همه Websiteها را در یک صفحه ببیند.

برای هر Website حداقل این اطلاعات نمایش داده شود:

- نام Website
- Domain
- وضعیت فعلی Website
- تاریخ Expiration مربوط به Domain
- تاریخ Expiration مربوط به SSL Certificate
- وضعیت نیاز به Attention

هدف صفحه اصلی این است که User بدون ورود به چند صفحه بتواند سریع متوجه شود:

- چه Websiteهایی سالم هستند
- چه Websiteهایی مشکل دارند
- چه Domainهایی نزدیک Expiration هستند
- چه SSLهایی نزدیک Expiration هستند

---

## Website Monitoring

Websiteها باید به صورت منظم Check شوند.

اگر Website از دسترس خارج شد:

- User باید سریع مطلع شود.
- Notification از طریق Telegram ارسال شود.
- Event مربوط به خرابی باید ذخیره شود.

وقتی Website دوباره در دسترس قرار گرفت:

- بهتر است وضعیت Healthy شود.
- Event مربوط به Recovery نیز ذخیره شود.

---

## تاریخچه خرابی

History مربوط به Website Availability باید نگهداری شود.

هدف این است که در آینده بتوان مشاهده کرد:

- یک Website چند بار Down شده است
- Down شدن چه زمانی اتفاق افتاده است
- چه زمانی Recover شده است
- مدت تقریبی مشکل چقدر بوده است

برای MVP، Storage می‌تواند در Database یا JSON باشد.

انتخاب نهایی روش Storage در مرحله Stack Selection انجام خواهد شد.

---

## Domain Renewal

برای Domainها باید Expiration Date نگهداری شود.

Notification مربوط به Domain باید:

- 14 روز قبل از Expiration ارسال شود.

در صورت نیاز می‌توان در آینده Reminderهای بیشتری اضافه کرد.

---

## SSL Renewal

برای SSL Certificate باید Expiration Date نگهداری شود.

Notification مربوط به SSL باید:

- 7 روز قبل از Expiration ارسال شود.

---

## Telegram Notification

Telegram Channel اصلی Notification در MVP است.

Telegram Bot باید برای موارد زیر پیام ارسال کند:

- Website Down
- Website Recovery
- Domain نزدیک Expiration
- SSL نزدیک Expiration
- Confirmation مربوط به Renewal

---

## Renewal Confirmation

پس از Renewal، User باید بتواند Renewal را Confirm کند.

روش مورد انتظار:

- User از طریق Telegram اعلام کند که Renewal انجام شده است.
- سیستم وضعیت مربوطه را Update کند.
- Expiration Date جدید ثبت شود.

جزئیات دقیق Interaction در Telegram در مرحله Product Flow مشخص می‌شود.

---

## اضافه کردن Website

MVP باید یک Form ساده برای Add کردن Website داشته باشد.

هدف این Form سادگی است.

User نباید برای اضافه کردن یک Website مجبور باشد اطلاعات پیچیده یا فنی وارد کند.

اطلاعات مورد نیاز باید تا حد امکان محدود و واضح باشند.

Candidate اولیه:

- نام Website
- Domain / URL
- Domain Expiration Date
- SSL Expiration Date

در صورت امکان، بعضی از اطلاعات می‌توانند در آینده Auto-detect شوند.

---

## Client Structure

در MVP نیازی به ساختار جداگانه برای Clientها وجود ندارد.

همه Websiteها در یک List مرکزی نمایش داده می‌شوند.

Client Management به نسخه‌های بعدی منتقل می‌شود.

---

## Scope نسخه اول

### In Scope

- Website List
- Website Status
- Website Monitoring
- Website Down Notification
- Website Recovery Notification
- Domain Expiration Date
- Domain Reminder
- SSL Expiration Date
- SSL Reminder
- Telegram Bot
- Renewal Confirmation
- Website History
- Add Website Form
- Attention Status

### Out of Scope

در MVP موارد زیر وجود ندارند:

- Multi-user
- Team Support
- Client Portal
- CRM
- Project Management
- Accounting
- Password Management
- Backup Monitoring
- Server Monitoring
- Complex Reporting
- Mobile App
- Multiple Notification Channels
- Advanced Client Management
- Enterprise Monitoring

---

## رفتار مورد انتظار MVP

نمونه Scenario:

1. User یک Website اضافه می‌کند.
2. Website در List اصلی نمایش داده می‌شود.
3. سیستم وضعیت Website را Check می‌کند.
4. اگر Website Down شود، Telegram Notification ارسال می‌شود.
5. Event در History ذخیره می‌شود.
6. وقتی Website Recover شود، Telegram Notification ارسال می‌شود.
7. اگر Domain تا 14 روز آینده Expire شود، Reminder ارسال می‌شود.
8. اگر SSL تا 7 روز آینده Expire شود، Reminder ارسال می‌شود.
9. User پس از Renewal از طریق Telegram اعلام می‌کند که Renewal انجام شده است.
10. سیستم وضعیت و Expiration Date را Update می‌کند.

---

## معیار موفقیت MVP

MVP زمانی موفق محسوب می‌شود که User بتواند:

- همه Websiteهای خود را در یک صفحه ببیند.
- وضعیت Healthy یا Down بودن آن‌ها را ببیند.
- قبل از Expire شدن Domain هشدار بگیرد.
- قبل از Expire شدن SSL هشدار بگیرد.
- Down شدن Website را سریع در Telegram بفهمد.
- Renewal را Confirm کند.
- History خرابی‌ها را بعدا مشاهده کند.

---

## تصمیم‌های Product نهایی‌شده

- User: Single Freelancer
- UI: یک صفحه مرکزی برای همه Websiteها
- Notification Channel: Telegram
- Domain Reminder: 14 روز قبل
- SSL Reminder: 7 روز قبل
- Website Failure Alert: سریع
- Website History: لازم است
- Data Persistence: لازم است
- Add Website Form: لازم است
- Client Separation: فعلا لازم نیست

---

## مرحله بعد

بعد از این Document، پروژه آماده ورود به مرحله قبل از Stack Selection است.

در مرحله بعد فقط باید چند Requirement ساده مشخص شوند:

- Website Monitoring هر چند وقت یک بار انجام شود
- Telegram Bot دقیقا چه Actionهایی داشته باشد
- چه Dataهایی حتما باید Persist شوند
- Deployment برای MVP کجا انجام شود
- چه محدودیت Budget یا Hosting وجود دارد

پس از مشخص شدن این موارد، می‌توان Technology Stack را انتخاب کرد.
