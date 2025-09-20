# Introduction 

I will design and program version 2 of the system I wrote for a client. I built V1 using Google Sheets and their apps scripts programming language, a JavaScript ES6 without support for a few capabilities like modules and arrow functions.

## Good Numbers Make Good Friends

With V1, the client wanted to update its process for recording the income and expenses, current accounts, of its associates to ensure a simple and reliable experience in settling their current accounts upon their departure.

V1 also aimed to improve the quality and transparency of MC associates' expenses transactions at it General Store. 

## Faster, Better, More Elegant and Reliable

V2 goal is to provide the same services using a faster, more reliable, elegant, and secure platform. 

I'll build a single user solution that can be deployed on the cloud or a local server. 
It will include a mobile first UI using [TypeScript](https://www.typescriptlang.org/) and [React](https://react.dev/), enabling `credentialized actors` to operate the parts of the system they are `authorized`.

I'll use a `Domain Driven Design` pattern to build the backend using [golang](https://go.dev/), [GIN](https://gin-gonic.com/), [`RuleGo`](https://github.com/rulego/rulego), and [GORM](https://gorm.io/).

I'll use an [SQLite](https://sqlite.org/) database to persist the data, and migrate to [PostGress](https://www.postgresql.org/) to go to production.

I'll use [Docker](https://www.docker.com/) to containerize the `UI`, `backend`, and `database` separately.

# Ubiquitous Language 
This is the term Eric Evans uses in Domain Driven Design for the practice of building up a common, rigorous language between developers and users. This language should be based on the `Domain Model` used in the software - hence the need for it to be rigorous, since software doesn't cope well with ambiguity. 

Following is the  is the Ubiquitous Language I developed for V1, redacted to be generic:

**`area`** - A generic location, within the facility, where operations relevant to the current account system take place (well, village);

**`active`** - the associate is integrated, ready and able to work in at the employer facility;

**`associate`** - an individual who has provided or provides services to the employer at its facilities; it is one of the top level `ENTITIES` in our domain;

**`availability`** - describes the associate's ability to perform their responsibilities (active, time off, leave, inactive);

**`commission`** - a percentage of the daily production, credited to a commissioned associate for their work at the end of the day;

**`daily pay`** - an associate's compensation, credited in their account  for their work at the end of a work day;

**`current account`** - a record of income earned and expenditure on goods and services consumed by the member during a stay; (RS - expand it to be two accounts, one in Reais and the other in Gold).

**`income`** - A generic term for various forms of compensation.

**`inactive`** - The associate has completed their stay and is no longer in contract with the employeer;

**`Leave`** - A period during a `stay` when an associate is absent from work for more than one day, to attend to personal matters. Clients have rules to ensure that critical operational functions are staffed and to compensate the associates that step in; in some cases, employers also split the compensation, half and half, of the associate that took `leave` with the one who replaced them;

**`mercantile`** - also known as the `canteen`, it is place where the employer's associates purchase goods and services; we opted for `mercantile` since is broader, enabling the employer to consolidade services other than providing meals;

`method` - The way in which the employer compensates the associate (Daily, Salary, Percentage, M-Percentage)

**`Stay`** - a period during which an ssociate is working or has worked at the employer's operations; it is common for my clientele contract their associates for a period of time, send them home to rest for a short period, and contract them back again;

**`time Off`** - a short period, `one-day to two days` during a `stay`, when an associate chooses not to work. Clients have rules to ensure that critical operational functions are staffed and to compensate the associates that step in; in some casees, employers also charge the associates who take the `time off` with the compensation of the associate who steped in for them;

**`salary`** - A a form of incoe paid to associates at the end of a month;

`Sste` - A specific area within the employer's facility where operations relevant to the current account system take place (well 3, mercantile, warehouse).

**`s-percentage`** - a portion of a commission derived from daily production, offered to pay an associate who stepped in to replace a commissioned associate who took `leave' and `work` at a critical job;

**`shif`t** - a period of work within a day; often, employers operate two twelve-hour daily shifts, day and night;

**`task`** - The name of the activity performed by an associtate at a location (cook, tire repairman, tractor driver, etc.);

**`transaction`** - captures an atomic income or expense of interest to the current account system;

# Description
V2 will be a fast, secure, elegant, and reliable `current account` system to record a`ssociates'` `income` and `expenses` during their `stay` at the employer's facilities. It will use a sophisticated `planning` feature to plan the the associates' work that leads to them earning `income`, as well a versitile `expense` feature that enables the associate to order and pick up goods and services at the `mercantile`.

## Associates
Individuals working at the employer's facilities. The administrator adds and updates `associates` records as needed. `Associate` records are permanent and are never deleted; associates no longer involved with the employer have their `deleted` set;

### Stay
Associates work at the employer's facilities for mutually agreed periods-- twelve weeks in V1, configuragle in V2--, called `stays`, after which they return to their residences for a mutually agreed period-- three weeks in V1, configuragle in V2. When an associate begins their `stay`, if necessary, the administrator updates their associate record, and then creates a new `stay` record for the associate. The `stay` record establishes the `method`, `area`, `location`, `task`, as well as the `compensation` for the associate's` stay`. The system allows planner to adjust them as needed. The system provides mechanisms to alert planners and managers of associates with stays about to expire.

## Current Accounts
The employer's associates earn `Income` as a result of their work in their facilities. They incur `Expenses` when purchasing goods and services at the `mercantile`. The system supports associates having multiple current accounts, one for each tipe of income (hard corrency, gold, etc.)

## Incomes
The employer's associates derive their income from performing work outlined on theiry daily work plans. 

### Work Plans
The work planning feature assigns associates to their daily functions, record the work they performed, and document their earnings as a result of this work in their current accounts. It guides the associates to perform the necessary work so that they earn their credits in their current accounts. It consists of the following steps: 1. Model, 2. Plan, 3. Inform, 4. Account
#### Model
The first step is to includes only associates with `active stays` (the responsible member updates a member's stay to Inactive at the time of closing their current account).
#### Plan
#### Inform
#### Account

# Lost and Found
## Ubiquotous Language
[RS] I removed these terms since they are specific to the original client, a mining company; we might bring them back in case it becomes a requirement for a new customer.

**`casa`** da queimada (burnt house) - place where members separate the gold collected in the well; today this gold is weighed before being sent to the safe house;

**`dompensation`** - The amount (BRL or gold) that the employ pays for each task/payment;

**`Leave`** - A period during a `stay` when an associate is absent from work for more than one day, to attend to personal matters.  [RS] `I left a more generic definition above` For an associate assigned to well operations, the associate will split their percentage 50/50 with the day laborer who replaced them until the end of their leave.

**`m-percentage`** - half of a predetermined daily percentage (1/26 of 25%), derived from daily production, offered to pay the associate who took leave and the day laborer who replaced them; [RS] `I left a more generic definition above`

**`Office`** - the location where associates perform administrative functions, such as requesting a PIX transfer, requesting time off, applying for a license, reviewing their current knowledge, etc.;

**`Percentage`** - a commission (1/13 of 25%), derived from a well's daily production, offered to pay an associate for their work at the end of the day at said well; [RS] `I left a more generic definition above`

`**Report**` - a document, printed or digital, that presents information in an organized format for a specific audience and purpose.

**`spreadsheet`** - a web-based application, using rows and columns, that allows users to create, update, modify, and share data online in real time.

 **`Stay`** - a period during which an ssociate is working or has worked at the employer's operations; [RS - I removed it from the main definition, to allow it to be more generic] it lasts no more than ninety days and may be interrupted by time off, leave, or termination.

**`Time off`** - a one-day to two days period during a `stay` and when an associate chooses not to work. [RS] `I left a more generic definition above` In the case of an associate working at the well, they pay the daily rate of the associate who replaced them. In the case of a day laborer, they simply do not work, but earn their daily rate; if the day laborer needs more than one day off, their current account is closed, their stay is terminated, and the employer returns them to their city.

**`vault`** - place where associates weigh the gold separated in the burn; the result is divided between the employer (75%) and the associates (1/13 of 25% for each) working in the well where the extracted gold originated; 

**`village`** - an area within the employer facilities where associates live, eat, and where the company's administrative facilities are located.

**`Well`** - an area where associates perform tasks to extract gold; the employer currently operates x week;









