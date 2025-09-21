# Introduction 
I have a client that operates a mining company in the heart of the Amazon, where the only suitable means of tranportation are with `small river canoes` and `small airplanes`. This geographic restriction requires my client to host its `associates` working on its mining operations for extended periods of time; it also requires my client to operate a `mercantile` with the goods and services required for their associates to have a productive and comfortable `stay`.

One of their major challenges is to keep a reliable record of their associates `current accounts`, including a detailed journal of each associate's `earnings` and `expenses`.

I built a software solution to assist my client with their `current accounts` difficulties; I leveraged their existing Excel solution, to write a `Google Sheets` solution, using its apps script programming language, a `JavaScript ES6 without support for a few capabilities like modules and arrow functions.

I decided to upgrade the solution I built for my client, `V1`, providing the same services using a faster, more reliable, elegant, and secure platform, `V2`. 

This document consists of a description of a business similar to my client's as well of a more sophisticated and general software solution suitable for any prospective customers with similar operations. There are hundreds of propspective clients like this in the area.

# Business Domain
This business domain description is and spice generic, hopefully opening the space for `V2` being suitable not only for mining, but for agricultural operations as well.

## Ubiquitous Language 
This is the term Eric Evans uses in Domain Driven Design for the practice of building up a common, rigorous language between developers and users. This language should reflect the `Business Domain` and be based on the `Domain Model` used in the software - hence the need for it to be rigorous, since software doesn't cope well with ambiguity. 

Following is the the Ubiquitous Language I developed for `V1`, redacted to be generic and apply for `V2`:

**`active`** - a state of the associate's `stay`, the `planner` can assign work taks to them;

**`affiliated`** - A state of an `associate` indicating that they remain connected with the organizarion, even without an active `stay`;

**`area`** - A generic location, within the facility, where operations relevant to the current account system take place (well, village);

**`account`** - a running record of an associate's `income` and `expenditures` during a `stay` for a given `currency`; mining companies's associates have `accoounts` in `Real` and `Gold`;

**`associate`** - an individual who has provided or provides services to the employer at its facilities; it is one of the top level `ENTITIES` in our domain;

**`available credit`** - The difference between an associate's `earnings` and `expenses` for a given account;

**`availability`** - an enumeration of states an associate has relative to their current `stay`; (active, time off, leave, inactive);

**`commission`** - income earned based on productivity; associates hired to work on wells earn commissions, instead of fixed pay; this is based on the daily productivty of the wells they worked on, and credited to their `gold account` at the end of the day;.

**`complete`** - a state of the `associate's stay`, alerting the `planner` that they associate stay ended, and that the`planner` cannot assign work taks to them;

**`cronogram`** - a `work plan` for a `period`, informing `associates` of the `sites` and `tasks` they are to perform for the `period`;

**`currency`** - a a medium of exchange to record an `associate's` `income` and `expenditure` during a `stay`; mining companies have two currencies, `Real` and `Gold`;

**`current account`** - a running record of an associate's `currency account's` `income` and `expenditures` during a `stay;`;

**`daily pay`** - an associate's compensation, credited in their account  for their work at `the end of a work day`;

**`earning`** - A generic term for various forms of compensation.

**`earning`** - A quantity `earned` by an `associate` for work done at their employer's facility; at my client, this amount can be either in the `local currency`, `Real`, of in `grams of gold` based on their `commission` derived from the `well` `yields` they worked on;

**`earning transaction`** - A `transaction` expressing an `associate's` earnings as the result of performing a `cronogram` `task`;

**`earning potential`** - A calculation based on the `amount of time` left on an associate's current `stay` and their `remuneration` stated in their current `stay`;

**`employer`** - The organization, hosting associates to working on it facilities;

**`expense`** - A generic term for various forms of expenditures by an associate at the employer's merchantile.

**`expenses ceiling`** - The sum of an associated `available credit` and `earning potential`, representing the maximum amount of expenses an associate might incur at a given point in time;

**`expense currency`** - The currency the associate uses to pay for an expense;

**`expense transaction`** - A `transaction` expressing an `associate's` `puchase` of `goods and services` at the `mercantile`; at my client, its amount can be either in the `local currency`, `Real`, of in `grams of gold` based on their `commission` derived from the `well` `yields` they worked on;

**`flight`** - My client charges to fly its associates to and from it mine location;

**`gold account`** - a running record of an associate's `income` and `expenditures`, in `gold`, during a `stay`;

**`income`** - A generic term for various forms of compensation.

**`income transaction`** - A `transaction` expressing an `associate's` earnings as the result of performing a `cronogram` `task`;

**`income currency`** - The `currency` the employer used to pay for an `income transaction`;

**`leave`** - A period during a `stay` when an associate is absent from work for more than one day, to attend to personal matters. Clients have rules to ensure that critical operational functions are staffed and to compensate the associates that step in; in some cases, employers also split the compensation, half and half, of the associate that took `leave` with the one who replaced them;

**`license`** - a state of the `associate's stay`, alerting the `planner` that they associate took an extended timne off, and that the`planner` cannot assign work taks to them;

**`mercantile`** - also known as the `canteen`, it is place where the employer's associates purchase goods and services; we opted for `mercantile` since is broader, enabling the employer to consolidade services other than providing meals;

`method` - The way in which the employer compensates the associate (Daily, Weekly. Salary, Percentage, S-Percentage)

**`PIX`** - My client provides a service to enable its associates to remit money (not gold) to a banch account of their choice;

**`stay`** - a `period` during which an `associate` is working or has worked at the employer's operations; my client has stays that last 12 weeks, followed by at least 3 weeks off; associates in good stand with my client are invited back for new stays;

**`time Off`** - a short period, `one-day to two days` during a `stay`, when an associate chooses not to work. Clients have rules to ensure that critical operational functions are staffed and to compensate the associates that step in; in some casees, employers also charge the associates who take the `time off` with the compensation of the associate who steped in for them;

**`real account`** - a running record of an associate's `income` and `expenditures`, in `Real`, during a `stay`;

**`salary`** - A a form of incoe paid to associates at the end of a month;

`site` - A `specific area` within the employer's facility where operations relevant to the `current account` system take place; some of my client sites were canteen, cofre, well 1;

**`s-percentage`** - a portion of a commission derived from daily production, my client paid an offered to pay an associate who stepped in to replace a commissioned associate who took `leave' and `work` at a critical job;

**`shif`t** - a period of work within a day; often, employers operate two twelve-hour daily shifts, day and night;

**`task`** - The name of the activity performed by an associtate at a location (cook, tire repairman, tractor driver, etc.);

**`terminated`** - A state of an `associate` indicating that his relationship with my client has been severed;

**`time off`** - a state of the associate's `stay`, alerting the `planner` that they associate took a few days off, and that the planner cannot assign the associate to any tasks;

**`transaction`** - captures an atomic income or expense of interest to the current account system;

**`weekly pay`** - an associate's compensation, credited in their account  for their work `at the end of a work week`;

## Description
### Earnings
#### Onboarding
When onboarding an individual who is new to the organization, the `personnel manager` adds their `associate` record, setting its state as `affiliated`.  Associates work at the business' facility for a few months; we call these peridos `stays`; the detault stay length is 12 weeks. After a `stay`, an associatte takes a few weeks to rest; the default rest lenght is 3 weeks. Associates in good stead with the business remain `affiliated`, otherwise they are `terminated`.

When onboarding an individual, the `personnel manager` also defines their `stay` record, setting its state as `active`; this `stay` record would be different than those of `previous stays` the associate might have had. This enable the `schedule manager` to assign work for the associate. Among other things, the `stay record` includes `date the stay started`, the `remuneration method and amount`.

#### Scheduling Work
The `schedule planner` uses the `stay` records to plan the daily work periods; the default is a daily and a nighlty periods. The `schedule planner` uses the `most recent published work plan` of the period they are planning to simplify their task; they can add or remove associates from the work plan, as well as change the locations and tasks of associates.

Througout their stay, an associate can take `timeoff`, causing the `personnel manager` to update their `stay` record status to `timeoff`. This enables the `schedule planner` assign a replacement, if necessary. When a `commissioned associate` take `timeoff`, the `schedule planner` assigns a replacement among `non-commissioned` associates. The business pays the commission and the wages of both normanlly, but `charges the commissioned associate` with the `wages of the non-commissioned associate` that replaced them;

#### Exceptions
Througout their stay, a `commissioned associate` can take a `license`, causing the `personnel manager` to update their `stay` record status to `license`. This enables the `schedule planner` to assign a replacement; in this case, the business splits the  `commissioned associate`  commission with the  `non-commissioned associate` that replaced them; the default split is half and half.

#### Earnings Accounting
At the `end of a daily work period`, the `schedule planner` ensures that all `associates executed their tasks as planned`, `adjusts the plan as required`, and `records the assocciates' earning in their current accounts`.

### Expenses
During an `associate's stay`, the business records their `earnings` and `expenses` in their `current accounts`. When an associate wishes to incur an expense, the `responsible mercantile person` checks the member’s current account to ensure that they have sufficient `available credit` or `earning potential`  to do so. This ensures a smooth closing of the member's current account at the end of their `stay`.

The business supports the folling kids of expenses:
- `mercantile` - The `associate` can use the same `currency accont` to `purchase many items` in a `single mercantile transaction`;
- `PIX` - This is a mechanism for an `associate` to `remit monetary currency`; the `default bank` is on his associate record
- `flight` - The business charges the associate for the flights to and from its operating installations. The business charges the `to flight` on the first day of their `stay`. The business charges the `from flight` on the last day of their `stay`. 
- `timeoff` - This option enables the `schedule planner` to charge the commissioned associate for the wages of the non-commussioned associate that replaced them;
- `exchange` - The business enables `associates` with `multiple currency accounts`, to exchange `from` the `non-monetary to monetary acconts`, as long as their `non-moneraty accont` has sufficient credit.
- ` close stay` - At the end of an associate stay, the personal manager uses the `close stay` option to summarize the credits or debits of all the ssociate's accounts, and outlining who has to pay who to close the each account. Again, at the `close stay` time, the associate will sign a document expressing their consent with the results of their stay.

In the case of `mercantile`, b. `PIX`, c. `flight`, d. `exchange`, e. `others`, associates with `accounts in multiple currencies` will have the option to choose the currency they wish to pay for their expense, as long as the account has enough credit to cover the expenses. In all cases the associate will have an electronic or paper copy describing the transaction.


# System Design
I'll describe the software design for `V2`.

## Good Numbers Make Good Friends
With `V1`, my client wanted to update its process for recording the income and expenses, `current accounts`, transactions of its associates to ensure a simple and reliable experience in settling their current accounts upon their departure. Furthermore, `V1` also aimed to improve the quality and transparency of MC associates' expenses transactions at its canteen. 

## Faster, Better, More Elegant and Reliable
V2 goal is to provide the same services using a faster, more reliable, elegant, and secure platform. 

I'll build a multi user solution that can be deployed on the cloud or a local server. 

It will include a mobile first UI using [TypeScript](https://www.typescriptlang.org/) and [React](https://react.dev/), enabling `credentialized actors` to operate the parts of the system they are `authorized`.

I'll use a `Domain Driven Design` pattern to build the backend using [golang](https://go.dev/), [GIN](https://gin-gonic.com/), [`RuleGo`](https://github.com/rulego/rulego), and [GORM](https://gorm.io/).

I'll use an [SQLite](https://sqlite.org/) database to persist the data during constrution, and migrate to [PostGress](https://www.postgresql.org/) to go to production.

I'll use [Docker](https://www.docker.com/) to containerize the `UI`, `backend`, and `database` separately.




# Description
V2 will be a fast, secure, elegant, and reliable `current account` system to record a`ssociates'` `income` and `expenses` during their `stay` at the employer's facilities. It will use a sophisticated `planning` feature to plan the the associates' work that leads to them earning `income`, as well a versitile `expense` feature that enables the associate to order and pick up goods and services at the `mercantile`.

## Associates
Individuals working at the employer's facilities. The administrator adds and updates `associates` records as needed. `Associate` records are permanent and are never deleted; associates no longer involved with the employer have their `deleted` set;

### Stay
Associates work at the employer's facilities for mutually agreed periods-- twelve weeks in V1, configuragle in V2--, called `stays`, after which they return to their residences for a mutually agreed period-- three weeks in V1, configuragle in V2. When an associate begins their `stay`, if necessary, the administrator updates their associate record, and then creates a new `stay` record for the associate. The `stay` record establishes the `method`, `area`, `location`, `task`, as well as the `compensation` for the associate's` stay`. The system allows planner to adjust them as needed. The system provides mechanisms to alert planners and managers of associates with stays about to expire.

## Current Accounts
The employer's associates earn `Income` as a result of their work in their facilities. They incur `Expenses` when purchasing goods and services at the `mercantile`. The system supports associates having multiple current accounts, one for each tipe of income (hard corrency, gold, etc.)

### Incomes
The employer's associates derive their income from performing work outlined on theiry daily `work plans`, see ahead for more details. The proposed system supports three ways for Carará Mining to pay its associates: i. daily, ii. salary, commission.

#### Daily
The system credits the earnings received by the associates earning daily pay on their current account, in the currendy specificed in their `stay`, at the end of their work periods.

#### Salary
The system credits the earnings received by the associates earning salaries on their current account, in the currendy specificed in their `stay`, on the last day of the month, at the end of their work period.

#### Commission
The system credits the earnings received by the associates earning commission on their current account, in the currendy specificed in their `stay`, at the end of their work periods; thisd is predicated on the employer have recorded the basis upon which the system calculates the commission; otherwise, the sytem keeps the transaction in a wainting state;

### Expenses
During an associate's `stay`, the employer records its revenues and expenses in their current account. 
When an associate places and `expense` order, the employers' responsible person checks the associate’s current account to ensure that they have sufficient `available credit` or `earning potential` during this stay to do so; the sum of both defines the `ceiling of expenses` allowed for the associated at any point in timne. This ensures a smooth closing of the member's current account at the end of their stay. 

## Work Plans
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





kdskdkdkdkd



