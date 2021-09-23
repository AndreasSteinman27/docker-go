---
description: Variant
---

# Final Statement

{% embed url="https://www.figma.com/file/w78ZiMR2USgl1CwXVrcxXv/?node-id=1632%3A27548" %}



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

### Final Statement Required Fields

| Name | Type | Required | Backend |
| :--- | :--- | :--- | :---: |
| Lease Buyout Amount | [Currency](../../input/currency-input.md) | **Yes** | \*\*\*\* |
| All Tax Fields  | [Currency](../../input/currency-input.md) |  |  |
| Other Fee's | Generic |  |  |
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

### Generate Final Statement Button

When pressed:

* Creates a PDF report. This PDF is able to be downloaded on the member's side and will be an object in the documents section \([Object Table](../../task-tables/object-table/)\) named "Final Advance Statement". 
* Sets the Asset status to final-advance-review.

### Edit Button

* Appears in the Saved state, which is when information was previously entered, but fields were grayed out.
* Clicking it makes the fields editable, make the button disappear, and make the Discard and Save buttons appear

### Save Button

* Appears in the Edit state and in the on first load states.
* When clicked, runs form validation
  * Any errors should cause the form to display the error state for the affected field, and for processing to stop.
* If validation proceeds without issue, then the data should be saved to the database, and the fields should be grayed out and made uneditable.
* The sidebar should be updated with the final fee total.
* The sidebar should be updated with the Lease Buy Out Amount, the Maximum Allowed, the Advance Amount, the Transaction Total, and the Minimum Down Payment. 
  * These values must be calculated as well.
  * No valuations should be run, the final vehicle valuation should have been run when the odometer image was verified in the documents upload stage \([Odometer Image Drawer](../../drawer/drawer-odometer.md)\).
* The save button should then disappear and be replaced by the Edit button and the Generate Preliminary Statement button.

### Discard Button

* Appears in the Edit state
* When clicked, discards any changes made and returns all fields back to their previous values. The fields should then be grayed out and uneditable
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

