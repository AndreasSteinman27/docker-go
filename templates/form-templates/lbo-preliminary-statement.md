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

[Input Form](../../components/form/)

### Child Elements

[Currency input](../../components/input/currency-input.md), [Header,](../../components/headers/header/) [Primary button](../../components/button/), [Secondary button](../../components/button/secondary-button.md)
{% endtab %}

{% tab title="Requirements" %}
### Requirements

* Calculation of Taxes and Fees are sent to Right Bar
* Performs vehicle valuation upon submission using the data provided in Trim Verification
* Provides a different flow based on retail or direct lease dropdown

#### Fields

Fields marked as Required must be completed and validated to submit this form. All others are optional.

| Name | Type | Required | Backend |
| :--- | :--- | :--- | :--- |
| Lessor Details - Address 1 | Generic |  |  |
| Lessor Details - Address 2 | Generic |  |  |
| Lessor Details - City | Generic |  |  |
| Lessor Details - State | [Dropdown](../../components/dropdown.md) |  |  |
| Lessor Details - Phone Number | [Phone](../../components/input/phone-number.md) | **Yes** |  |
| Lease Type | [Dropdown](../../components/dropdown.md) | **Yes** |  |
| Estimated Taxes Fields \(all\) | [Currency](../../components/input/currency-input.md) |  |  |
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

### Generate Preliminary Statement Button

When pressed:

* Creates a PDF report. This PDF is able to be downloaded on the member's side and will be an object in the documents section \([Object Table](../../components/task-tables/object-table/)\) named "Preliminary Advance Statement". 
* Sets the Asset status to prelim-advance-review.

### Edit Button

* Appears in the Saved state, which is when information was previously entered, but fields were grayed out.
* Clicking it makes the fields editable, make the button disappear, and make the Discard and Save buttons appear

### Save Button

* Appears in the Edit state and in the on first load states.
* When clicked, runs form validation
  * Any errors should cause the form to display the error state for the affected field, and for processing to stop.
* If validation proceeds without issue, then the data should be saved to the database, and the fields should be grayed out and made uneditable.
* The sidebar should be updated with the estimated fees, and the other transaction fields should be recalculated \(Advance Amount, Transaction Total, and Minimum Down Payment\).
  * No valuations should be run, the prelim valuation is performed when the trim verification occurs \([Trim](../../components/task-tables/task-table/trim.md)\).
* The sidebar should be updated with the Lessor Details entered.
* The save button should then disappear and be replaced by the Edit button and the Generate Preliminary Statement button.

### Discard Button

* Appears in the Edit state
* When clicked, discards any changes made and returns all fields back to their previous values. The fields should then be grayed out and uneditable.
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





