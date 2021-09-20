---
description: Variant
---

# Preliminary Statement

{% embed url="https://www.figma.com/file/w78ZiMR2USgl1CwXVrcxXv/?node-id=308%3A6033" %}

{% tabs %}
{% tab title="Usage" %}
### **Usage**

The Preliminary Advance Statement is a document created within Axle that serves as a preliminary estimate for a given asset/transaction. The preliminary statement shows estimated transactional amounts, including the vehicle price/outstanding balance, the minimum down payment as calculated against our valuation for the vehicle, and any estimated taxes or fees. Its counterpart is the Final Advance Statement, which shows finalized values once the MSA is able to verify everything.
{% endtab %}

{% tab title="Parent Object" %}
### **Parent Object**

[Input Form](../)

### Child Elements

[Currency input](../../input/currency-input.md), [Header,](../../header/) [Primary button](../../button/), [Secondary button](../../button/secondary-button.md)
{% endtab %}

{% tab title="Requirements" %}
### Requirements

In order to save the form, you must complete the following fields. The others are optional.

* Calculation of Tax's and Fee's are sent to Right Bar
* Performs vehicle valuation
* Provides a different flow based on retail or direct lease dropdown

#### Required fields

In order to save the form, you must complete the following fields. The others are optional.

| Name | Type | Required | Backend |
| :--- | :--- | :--- | :--- |
| Address 1 | Generic |  |  |
| Address 2 | Generic |  |  |
| City | Generic |  |  |
| State | [Dropdown](../../dropdown.md) |  |  |
|  Phone Number | [Phone](../../input/phone-number.md) | **Yes** |  |
|  Lease Type | [Dropdown](../../dropdown.md) | **Yes** |  |
| Estimated Taxes Fields \(all\) | [Currency](../../input/currency-input.md) |  |  |
| Other Fee's | Generic |  |  |
{% endtab %}

{% tab title="States" %}
### On-First Load

In this state, values are submitted and calculated. For instance, taxes are generated and sent to the Right Sidebar

{% embed url="https://www.figma.com/file/w78ZiMR2USgl1CwXVrcxXv/?node-id=2%3A4766" %}

### Saved-State

This state can only be edited by pressing the Edit button.

{% embed url="https://www.figma.com/file/w78ZiMR2USgl1CwXVrcxXv/?node-id=2%3A10825" caption="" %}

### Edit

This state allows for MSA's to edit their information and resubmit it. Discard will keep values from the previous save.

{% embed url="https://www.figma.com/file/w78ZiMR2USgl1CwXVrcxXv/?node-id=2%3A11583" %}
{% endtab %}

{% tab title="Actions" %}
### Actions

{% embed url="https://www.figma.com/file/w78ZiMR2USgl1CwXVrcxXv/?node-id=1306%3A26568" %}

### Generate Preliminary Form

When pressed:

* Creates a PDF report This PDF is able to be downloaded on the member's side and will be an object in the documents section \([Object Table](../../task-tables/object-table/)\) and name "Final Statement of Advance". 
* Checks for required fields
* Provides appropriate error handling
* Performs a Prelim Valuation Calculation

Creates a PDF report and a valuation calculation that is needed to Create a preliminary and Final Valuation. This PDF is able to be downloaded on the member's side and will be an object in the documents section \([Object Table](../../task-tables/object-table/)\).
{% endtab %}

{% tab title="Flow" %}
### Flow

{% embed url="https://www.figma.com/file/w78ZiMR2USgl1CwXVrcxXv/?node-id=308%3A6033" %}
{% endtab %}

{% tab title="Error Handeling" %}
### Error Handeling

Form validation errors occur when "Save" is pressed and the input displays an Inline Error. 

* Lease type -&gt; Please select a lease type
* Phone number -&gt; Please input a valid phone number
{% endtab %}

{% tab title="Code" %}
### Code

{% embed url="https://iframe.html?id=forms-prelim-form--primary&args=" %}
{% endtab %}
{% endtabs %}





