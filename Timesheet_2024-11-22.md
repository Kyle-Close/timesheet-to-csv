---
tags: timesheet
---

**2024-11-22: Timesheet**

```dataviewjs 
let numbers = dv.current().hours;
let sum = 0;

for (const num of numbers) {
  sum += num
}

dv.el("b", "Total Hours: " + sum);
```
---

| **Hours**      | **Name**                      | **Link**                                               | **Notes**                                                                                                                                                                                                                         |
|:-------------- |:----------------------------- |:------------------------------------------------------ |:--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| (hours:: 0.5)  | Digi-Key Corporation – PO     | [Link](https://conexiom.atlassian.net/browse/IQ-34300) |                                                                                                                                                                                                                                   |
| (hours:: 0.5)  | RUBIX Group - FR - Orexad     | [Link](https://conexiom.atlassian.net/browse/IQ-34188) |                                                                                                                                                                                                                                   |
| (hours:: 0.25) | Graybar – PO                  | [Link](https://conexiom.atlassian.net/browse/IQ-34331) |                                                                                                                                                                                                                                   |
| (hours:: 0.25) | Digi-Key Corporation – PO     | [Link](https://conexiom.atlassian.net/browse/IQ-34340) |                                                                                                                                                                                                                                   |
| (hours:: 1.75) | Graybar - Pricing             | [Link](https://conexiom.atlassian.net/browse/IQ-34100) | Complex graybar pricing request. Logic was all being done in pre-proc. Customer sends in a huge text file that needs to be parsed and added to an xref table. The format of their input text file changed needing update to logic |
| (hours:: 0.5)  | Parker EMEA - Benelux         | [Link](https://conexiom.atlassian.net/browse/IQ-34207) |                                                                                                                                                                                                                                   |
| (hours:: 0.5)  | Sherwin-Williams - EMEA - INV | [Link](https://conexiom.atlassian.net/browse/IQ-34256) |                                                                                                                                                                                                                                   |
| (hours:: 0.5)  | Digi-Key Corporation – PO     | [Link](https://conexiom.atlassian.net/browse/IQ-34290) |                                                                                                                                                                                                                                   |
| (hours:: 0.5)  | Digi-Key Corporation – PO     | [Link](https://conexiom.atlassian.net/browse/IQ-34301) |                                                                                                                                                                                                                                   |
| (hours:: 0.5)  | Digi-Key Corporation – PO     | [Link](https://conexiom.atlassian.net/browse/IQ-34313) |                                                                                                                                                                                                                                   |
| (hours:: 0.75) | Sherwin-Williams - EMEA - INV | [Link](https://conexiom.atlassian.net/browse/IQ-34206) |                                                                                                                                                                                                                                   |
| (hours:: 1.5)  | Sherwin-Williams - EMEA - INV | [Link](https://conexiom.atlassian.net/browse/IQ-34366) | Customer wanted us to use custom logic for numeric group/decimal separator variation. Seemed doable since there was specific criteria for each format that would be consistent. Did it but had some trouble along the way.        |
|                |                               |                                                        |                                                                                                                                                                                                                                   |


```
Customer wanted us to use custom logic for numeric group/decimal separator variation. Seemed doable since there was specific criteria for each format that would be consistent. Did it but had some trouble along the way.
```
