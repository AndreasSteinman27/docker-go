---
description: Variant
---

# Final Statement

{% embed url="https://www.figma.com/file/w78ZiMR2USgl1CwXVrcxXv/?node-id=1418%3A26955" %}

{% tabs %}
{% tab title="Usage" %}
### **Usage**

The Final Advance Statement is a document created within Axle that serves as a preliminary estimate for a given asset/transaction. 

The final statement shows: 

* transactional amounts 
* vehicle price/outstanding balance
* The minimum down payment as calculated against our valuation for the vehicle
* Final taxes or fees. 

Its counterpart is the Preliminary Advance Statement, which shows estimated values.
{% endtab %}

{% tab title="Parent Object" %}
### Parent Object
{% endtab %}

{% tab title="Requirements" %}
### Requirements

In order to save the form, you must complete the following fields. The others are optional.

* Calculation of Tax's and Fee's are sent to Right Bar
* Performs vehicle valuation
* Provides a different flow based on retail or direct lease dropdown

#### Required fields

In order to save the form, you must complete the following fields. The others are optional.

### Final Statement Fields

| Name | Type | Required | Backend |
| :--- | :--- | :--- | :---: |
| Address 1 | Generic |  |  |
| Address 2 | Generic |  |  |
| City | Generic |  |  |
| State | [Dropdown](../../dropdown.md) |  |  |
|  Phone Number | [Phone](../../input/phone-number.md) | **Yes** |  |
|  Lease Type | [Dropdown](../../dropdown.md) | **Yes** |  |
| Estimated Taxes Fields \(all\) | [Currency](../../input/currency-input.md) |  |  |
| Other Fee's | Generic |  |  |

### Required Fields Not on Page

| Name | Type |  |
| :--- | :--- | :--- |
|  |  |  |
{% endtab %}

{% tab title="States" %}
### States

### On-First Load

In this state, values are submitted and calculated. For instance, taxes are generated and sent to the Right Sidebar.

{% embed url="https://www.figma.com/file/w78ZiMR2USgl1CwXVrcxXv/?node-id=1418%3A26955" %}

### Saved-State

This state can only be edited by pressing the Edit button.

{% embed url="https://www.figma.com/file/w78ZiMR2USgl1CwXVrcxXv/?node-id=1418%3A27344" %}

### Edit

This state allows for MSA's to edit their information and resubmit it. Discard will keep values from the previous save.

{% embed url="https://www.figma.com/file/w78ZiMR2USgl1CwXVrcxXv/?node-id=1418%3A27148" %}
{% endtab %}

{% tab title="Actions" %}
### Actions

### Generate Final Statement Form

When pressed:

* Creates a PDF report This PDF is able to be downloaded on the member's side and will be an object in the documents section \([Object Table](../../task-tables/object-table/)\) and name "Final Statement of Advance". 
* Checks for required fields
* Provides approiate error handling
* Performs a Final Valuation Calculation
{% endtab %}

{% tab title="Error Handeling" %}
Form validation errors occur when "Save" is pressed and the input displays an Inline Error. 

* Lease type -&gt; Please select a lease type
* Phone number -&gt; Please input a valid phone number

Alert Badge

* Documents - &gt; Please provide  "X" document
* 
{% endtab %}

{% tab title="Code" %}

{% endtab %}
{% endtabs %}

