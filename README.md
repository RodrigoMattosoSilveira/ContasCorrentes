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

**`cronogram`** - a `work plan` for a `period`, informing `associates` of the `sites` and `tasks` they are to perform for the `period`; its workflow is: 1. Model, 2. Plan, 3. Inform, 4. Account

**`currency`** - a a medium of exchange to record an `associate's` `income` and `expenditure` during a `stay`; mining companies have two currencies, `Real` and `Gold`;

**`current account`** - a running record of an associate's `earnings` and `expenses` for all of their `currency acccouts`;

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

**`time Off`** - a short period, `one-day to two days` during a `stay`, when an associate chooses not to work. Clients have rules to ensure that critical operational functions are staffed and to compensate the associates that step in; in some casees, employers also charge the associates who take the `time off` with the compensation of the associate who steped in for them;

**`real account`** - a running record of an associate's `income` and `expenditures`, in `Real`, during a `stay`;

**`salary`** - A a form of incoe paid to associates at the end of a month;

**`site`** - A `specific area` within the employer's facility where operations relevant to the `current account` system take place; some of my client sites were canteen, cofre, well 1;

**`s-percentage`** - a portion of a commission derived from daily production, my client paid an offered to pay an associate who stepped in to replace a commissioned associate who took `leave' and `work` at a critical job;

**`shift`** - a period of work within a day; often, employers operate two twelve-hour daily shifts, day and night;

**`stay`** - a `period` during which an `associate` is working or has worked at the employer's operations; the default is 12 weeks, followed by at least 3 weeks off; associates in good stand are invited back for new `stays`;

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

The business supports the folling kidns of expenses:
- **`mercantile`** - The `associate` can use the same `currency account` to `purchase many items` in a `single mercantile transaction` the application ensures the associate has the required credit, the `mercantile associate` delivers the purchased goods;
- **`PIX`** - This is a mechanism for an `associate` to `remit monetary currency` to a named bank; the `default bank` is on their associate record; the application ensures the associate has the required credit, the `mercantile manager` assists with the remitance; perhaps the system can accomplish this automatically;
- **`flight`** - The business charges the associate for the flights to and from its operating installations. The business charges the `to flight` on the first day of their `stay`. The business charges the `from flight` on the last day of their `stay`; the `personnel manager` performs both operations; 
- **`timeoff`** - The `schedule planner` charges the commissioned associate for the wages of the non-commussioned associate that replaced them;
- **`exchange`** - The business enables `associates` with `multiple currency accounts`, to exchange `from` the `non-monetary to monetary acconts`, as long as their `non-moneraty accont` has sufficient credit; the application ensures the associate has the required credit, the `mercantile manager` assists with the operation; 
- **` close stay`** - The `personnel manager` summarizes the `credits` or `debits` of all the `associate's account`, and determines  and who has to pay who to close the each account. Again, at the `close stay` time, the associate will sign a document expressing their consent with the results of their stay.

In all cases, associates with `accounts in multiple currencies` will have the option to choose the currency they wish to pay for their expense, as long as the account has enough credit to cover the expenses. In all cases the associate will have an electronic or paper copy describing the transaction.

### Cronograms
The work planning feature assigns associates to their daily functions, records the work they performed, and documents their earnings as a result of this work in their current accounts. It guides the associates to perform the necessary work so that they earn their credits in their current accounts. It consists of the following steps: 1. Model, 2. Plan, 3. Inform, 4. Account

#### Model
The first step is to includes only associates with `active stays` (the responsible member updates a member's stay to Inactive at the time of closing their current account).

The `schedule planner` uses the `active stay` records to plan the daily work periods; the default is a `daily` and a `nighlty` periods. The `schedule planner` uses the `most recent published work plan` for the period they are planning to simplify their task; they can add or remove associates from the work plan, as well as change the locations and tasks of associates.

#### Plan
The `schedule planner` refines the plan to account for associates who took timeoff or licenses, or any other operation requirement for the period they are planning.

#### Inform
The `schedule planner` informs the associates of their next period assignments.

#### Account
The `schedule planner` refines the record of the work performed for the period and triggers the addition of the associates' `earnings transactions`.

# System Design
I'll describe the software design for `V2`.

## Good Numbers Make Good Friends
With `V1`, my client wanted to update its process for recording the income and expenses, `current accounts`, transactions of its associates to ensure a simple and reliable experience in settling their current accounts upon their departure. Furthermore, `V1` also aimed to improve the quality and transparency of associates' expenses at the canteen. 

## Faster, Better, More Elegant and Reliable
V2 goal is to provide a superset of services using a faster, more reliable, elegant, and secure platform. 

I'll build a multi user solution that can be deployed on the cloud or a local server. 

It will use a mobile first UI based on the [the golang template engine](https://golangdocs.com/templates-in-golang) driving [HTMX templates](https://htmx.org/), supported by [GIN routes](https://github.com/gin-gonic/gin) enabling using [Golang Gin JWT](https://github.com/appleboy/gin-jwt) to allow `credentialized actors` to operate the parts of the system they are `authorized` .


I'll use a `Domain Driven Design` pattern to build the backend using [golang](https://go.dev/), [GIN](https://gin-gonic.com/), [`RuleGo`](https://github.com/rulego/rulego), and [GORM](https://gorm.io/).

I'll use an [SQLite](https://sqlite.org/) database to persist the data during constrution. I might migrate to [PostGress](https://www.postgresql.org/) to go to production.

I'll use [Docker](https://www.docker.com/) to containerize the `UI`, `backend`, and `database` separately.

## Actors
These actors identify the roles and authorizarions that the system will support:
- **`system administrator`** - Individual with credentials to perform all system operations;
- **`associate`** - Contracted by the business to work at its facilities, can view their current accounts, and initiate `expense` activities;
- **`mercantile associate`** - Receives notifications of `action required` about `expense requests`, assists in delivering `mercantile` items for `approved expenses`;
- **`mercantile manager`** - Receives notifications of `action required` about `PIX, exxchange requests`, assists in completing the approved expenses`;
- **`personnel manager`** - `Onboards associates`, ensures their `associate and stay records` are correct, `closes the associate stay`;
- **`schedule planer`** - Prepares the `daily work assignments`, inform associates of their `assignments`, and `records the associates earnings` in the `current accounts`;

## DDD
### Entities
- Associate
- Account
- Stay
- Price List
- Mercantile Item
- Mercantile Item Price

### Value Objects
- Account Type (derived from currencies)
- Area (configurable: Mercantile, Village, Office)
- Affilication (afiliated, terminated)
- Availability (active, timeoff, license, complete)
- Address (Residence, Emergency)
- Bank
- Cell Phone (Personal, Emergency)
- Currency (configurable: Real, Gold)
- Earning (transaction)
- Email  (Personal, Emergency)
- Expense (transaction)
- Flight (Date, Airline Flight, From, To, time departure, time arrival)
- Location (configurable)
- Payment Methodo (configurable: daily pay, weekly pay, salary, commission)

### Services
- **`Associate`**:
  - The `personnel manager` triggers the creation of the` associate`; 
  - Authorized actors read the associate records;
  - The `personnel manager` can update the associate records;; 
  - The system never deletes `associates`; it updates its state (inactive, terminated), and it sets a delete date which can be revoked in special circunstances;
- **`Account`**:
  - The `personnel manager` triggers the creation of the `account`; `an associate must have one monetary account`; 
  - authorized actors read the associates accounts, as for instance the `mercantile manager`; 
  - The system updates accoungs via `earnings and expenses transactions`; 
  - The system never deletes `accounts`; it updates their state (inactive, terminated), and it sets a delete date which can be revoked in special circunstances;
- **`Stay`**:
  - The `personnel manager` triggers the creation of the `stay`;
  - Authorized personnel read them;
  - The `personnel manager` updates them;
  - The system never deletes a `stay`; it sets its `state` as `complete`, and the `delete` date which can be revoked in special circunstances;
- **`Price List`**
  -  The `mercantile manager` triggers the creation of the `price list`;
  -  Anyone can read the `price list`
  -  The `mercantile manager` updates the `price list`;
  -  The system never deletes a `price list`; it sets its `state` as `discontinued`, and the `delete` date which can be revoked in special circunstances;
-  **`Mercantile Item`**
  -  The `mercantile manager` triggers the creation of the `mercantile item`;
  -  The `system` reads the `mercantile item` when showing the `price list`; the ;
  -  The `mercantile manager` updates the `mercantile item`;
  -  The system never deletes a`mercantile item`; it sets its `state` as `discontinued`, and the `delete` date which can be revoked in special circunstances;
-  **`Mercantile Item Price`**
  -  The `mercantile manager` triggers the creation of the `mercantile item`;
  -  The `system` reads the `mercantile item price` when showing the `mercantile item `; the ;
  -  The `mercantile manager` updates the `mercantile item price`;
  -  The system never deletes a`mercantile item price`; it sets its `state` as `discontinued`, and the `delete` date which can be revoked in special circunstances;
- **`Area`** (configurable: Mercantile, Village, Office)
  - The `system administrator` configures the `areas` of interest in the facility;
  - The `system` reads and attaches to records that require them as a characteristic; 
  - The `system administrator` does not update them;
  - The `system administrator` marks them as read only for existing assignments, and to be ignored in any future assignments;
- **`Affilication`** (afiliated, terminated) - These are defined at system configuration time and never changed;
- **`Availability`** (active, timeoff, license, complete) - These are defined at system configuration time and never changed;
- **`Address`** (Residence, Emergency)
  - The `personnel manager` configures new `address`;
  - An `associate` can read their own `email`; the the `mercantile manager`  and `personnel manager` can read all associates' `emails`;
  - The `associate` triggers the update of an `address`;
  - The system never deletes a`address`, although an `associate` might remove it leaving the address without any references;
- **`Bank`**
  - The `personnel manager` configures new `bank`;
  - The `system` reads and shows `bank`;
  - The `associate` triggers the update of a `bank`;
  - The system never deletes a`bank`, although an `bank` might remove it leaving the address without any references;
- **`Cell Phone`** (Personal, Emergency)
  - The `personnel manager` configures new `Cell Phone`;
  - An `associate` can read their own `email`; the the `mercantile manager`  and `personnel manager` can read all associates' `emails`;
  - The `associate` triggers the update of a `Cell Phone`;
  - The system never deletes a`Cell Phone`, although an `Cell Phone` might remove it leaving the address without any references;
- **`Currency`** (configurable: Real, Gold)
  - The `personnel manager` configures new `Cell Phone`;
  - The `system` reads and shows `currency`;
  - The `personnel manager` triggers the update of a `personnel manager`;
  - The system never deletes a`personnel manager`, although an `personnel manager` might remove it leaving the address without any references;
- **`Earning`** (transaction)
  - The `schedule planner` trigger the generation of a new `earning` transaction;
  - An `associate` can read their own `earning` transactions; the `mercantile manager` and `personnel manager` can read all associates' `earning` transactions;
  - No one can update an `earning`transaction;
  - No one can delete an `earning` transaction;
- **`Email`**  (Personal, Emergency)
  - The `personnel manager` configures new `email`;
  - An `associate` can read their own `email`; the the `mercantile manager`  and `personnel manager` can read all associates' `emails`;
  - The `associate` triggers the update of an `email`;
  - The system never deletes an `email`, although an `email` might remove it leaving the address without any references;
- **`Expense`** (transaction)
  - The `associate` triggers the generation of a new `expense` transaction;
  - An `associate` can read their own `expense` transactions; the `mercantile manager`  and `personnel manager` can read all associates' `expense` transactions;
  - No one can update an `expense`transaction;
  - No one can delete an `expense` transaction;
- **`Flight`** (Date, Airline Flight, From, To, time departure, time arrival)
  - The `personnel manager` configures new `flight`;
  - An `associate` can read their own `flights`; the the `mercantile manager`  and `personnel manager` can read all associates' `flights`;
  - The `personnel manager` triggers the update of an `flights`;
  - The system never deletes an `flights`, although an `flights` might remove it leaving the address without any references;
- **`Location`** (configurable)
  - The `system administrator` configures the `locations` of interest in the facility;
  - The `system` reads and attaches to records that require them as a characteristic; 
  - The `system administrator` does not update them;
  - The `system administrator` marks them as read only for existing assignments, and to be ignored in any future assignments;
- **`Payment`** Methodo (configurable: daily pay, weekly pay, salary, commission)
  - The `personnel manager` configures new `payment` method;
  - An `associate` can read their own `payment` method; the the `mercantile manager`  and `personnel manager` can read all associates' `payment` methods;
  - The  `personnel manager` or the `schedule planner` triggers the update of a `payment` method;
  - The system never deletes a`address`, although an `associate` might remove it leaving the address without any references;


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

# Links
[Client-side form validation](https://developer.mozilla.org/en-US/docs/Learn_web_development/Extensions/Forms/Form_validation)
[HTMX](https://htmx.org/docs/)
[Bootstrap forms](https://getbootstrap.com/docs/5.3/forms/layout/)
[Bootstrap Icons](https://icons.getbootstrap.com/)
[Form Controls for Mobiles](https://phoenix-react-alt.prium.me/modules/forms/form-control)
[Bootstrap Validation](https://getbootstrap.com/docs/5.0/forms/validation/)
[How to Create a Full-Stack Boilerplate with Go, Fiber, and HTMX](https://medium.com/@ahsanmubariz/how-to-create-a-full-stack-boilerplate-with-go-fiber-and-htmx-6012452aaa96), [github](https://github.com/ahsanmubariz/go_htmx_fiber_boilerplate)
[Understanding Web Project Folder Structure](https://medium.com/@mohitkhubchandani88/understanding-web-project-folder-structure-6cc7ae62fb40)
[HTMX Error Handling](https://github.com/bigskysoftware/htmx/blob/master/www/content/extensions/response-targets.md)