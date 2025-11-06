# Use Case
## System
	- I, as an Administrator, want to configure the types of Associate Availability (Active, Time0ff, License, Inactive) offered the facility;
	- I, as an Administrator, want to configure the Areas of work (Mercantile, Office, Yard, Well) within the facility;
	- I, as an Administrator, want to configure the Locals within Areas of work ( Mercantile / Canteen, Yard / Mill) within the facility;
	- I, as an Administrator, want to configure the Tasks (Cook, Inventory, Boroqueiro, Tratorista) performed at the facility;
	- I, as an Administrator, want to configure the Payment Method (Daily Wages, Salary, Commission) offered at the facility;
	- I, as the Current Account System, want to retrieve and store the daily price of gold, in Real/grams;
	- I, as the Current Account System, want to retrieve and store the daily exchange between USD and BRL;
	- I, as the Mercantile Manager, want to configure the Canteen Price List;
## Onboarding
	- I, as an Administrator, want to CRUD the Associate record, for a new or returning Associate; 
	- I, as an Associate, while on a Journey, want to update my Associate record;
	- I, as an Administrator, want to CRUD the Journey record for an Associate's;
	- I, as an Administrator, when setting up an Associate's Journey, want set their Availability, Payment Method, Area, Local, Task;
	- I, as an Administrator, when setting up an Associate's Journey, want set their Remuneration, based on their Payment Method;
	- I, as an Administrator, when setting up an Associate's Journey, want the ability to set up two currency accounts, one in Grams of Gold and the other in Reais.
	- I, as a Planer, want to update an Associate's Journey record, when the  Associate's Area, Local, or Task changes;
	- I, as an Administrator, want the system to send me a message informing o the ability to assess the expiration time of an Associate's Journeys and, as they approach expiration, to trigger a replacement plan;
	- I, as an Administrator / Planner / Associate, want the ability to assess the expiration time of an Associate's Journeys and, as they approach expiration, to trigger a replacement plan;
## Expenses
	- I, as an Associate, want to create a cart with Mercantile Items to purchase, including the currency I wish to use to pay for the items, with the total restricted by the limits of my Current Account;
	- I, as a Mercantile Operator, want the system to send me a message to fulfill an Associate Cart purchase;
	- I, as a Mercantile Operator, want to collect and deliver the items in an Associate cart, triggering one charge per item delivered to the Associate's Current Account;
	- I, as an Associate, want to request to send a PIX, including the currency I wish to use to pay for the transaction, with the amount restricted to the limits of my Current Account;
	- I, as a PIX Operator--it could be the system--, want send a PIX requested by an Associate, triggering a charge to the Associate's Current Account;
	- I, as an Operations Manager, want the ability to print a PIX request and ask the requesting associate to sign it;
	- I, as an Associate, want to request to flight, including the currency I wish to use to pay for the transaction, based on the limits of my Current Account;
	- I, as a Flight Coordinator, want to schedule a flight requested by an Associate, triggering a charge to the Associate's Current Account; 
	- I, as an Associate want the ability to request permission to incur on an expense beyond the limits of my Current Account; 
	- I, as an Operations Manager, want the privilege to allow an Associate to incur on an expense beyond the limits of their Current Account; 
	- I, as an Associate want the ability to request to incur on a generic expense,  with the total restricted by the limits of my Current Account;;
	- I, as an Operations Manager, want the privilege to use a Generic Expense form to enable an Associate to incur on a generic expense; 
## Operations
	- I, as an Operations Manager, when planning the  chronogram for the next period of work, want to start planning using the most recent period's.
	- I, as a commissioned Associate working on a well, want to take a day off from work, be replaced by a non-commissioned Associate, receive my daily commission, and pay my replacement daily wage;
	- I, as a commissioned Associate working on a well, want to take a license from work, be replaced by a non-commissioned Associate, and split my commission half and half  with them until my return or the end of my Journey:
	- I, as an Associate, want to receive a message informing me of my next period work assignment, 12 hours prior;
	- I, as an Operations Manager, want the ability to change a chronogram prior to its scheduled start, as necessary, and send a message to the affected Associates' smartphones;
	- I, as an Operations Manager, want the ability to change a chronogram after its completion, as necessary, and send a message to the affected Associates' smartphones;
	- I, as the Operations Manager, want to register a well production, in grams, for a period, at the end of the period;
## Earnings 
	- I, as an Operations Manager, want the ability to trigger the Current Account transactions to pay each Associate, in the currency linked to their task, for their work at the conclusion of a work period;
	- I, as the Current Accounts system, want the ability to postpone registering a commissioned Associate's period earning until the Operations Manager registers the production of the well where the Associated worked.
## Current Account
	- I, as an Associate, want to receive a text message informing me of any and all transactions in my Current Account;
	- I, as an Associate, want to the ability to use a smart phone to review  my current Journeys' Current Account;
	- I, as an Associate, want to the ability to use a smart phone and a Mercantile Printer to print my current Journeys' Current Account summary;
	- I, as an Associate, want the ability to exchange my credit in grams of gold for credit Reais in my Current Account, based on the daily gold exchange rate;
	- I, as an Administrator, want the ability to inspect an Associate' current account depicting their earnings and expenses during their current Journey, including potential earning thru the end of their Journey.
	- I, as an Administrator / Associate, want the ability to view a screen showing the transactions required to zero an Associate's current account, as they end their current journey .
	- I, as an Administrator, want the ability to trigger the transitions required to zero an Associate's current account once both parties pay their dues;
