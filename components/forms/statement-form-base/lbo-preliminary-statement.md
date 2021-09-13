# Preliminary Statement

**Description**

The Preliminary Statement Form component has 3 active states and 1 Skeleton loading state. **Only the required fields** must be complete in order to save.

**Child Elements**

Generic Text Field, Currency input, Header, Primary button, Secondary button, Form.

### 1. On-First Load

{% embed url="https://www.figma.com/file/w78ZiMR2USgl1CwXVrcxXv/?node-id=2%3A4766" %}

### 2. Saved State 

In this state, values are submitted and calculated. For instance, taxes are generated and sent to the Right Sidebar

 This state can only be edited by pressing the Edit button.

{% embed url="https://www.figma.com/file/w78ZiMR2USgl1CwXVrcxXv/?node-id=2%3A10825" caption="" %}

### 3. Edit State

This state allows for MSA's to edit their information and resubmit it. Discard will keep values from the previous save.

{% embed url="https://www.figma.com/file/w78ZiMR2USgl1CwXVrcxXv/?node-id=2%3A11583" %}

### 4. Generate Preliminary Statement

Creates a PDF report and a valuation calculation that is needed to Create a Final Valuation. This PDF is able to be downloaded on the members side and will be an object in the documents section.

### 5. Branches in MSA Selection

Based on the input choice of the direct or retail lease options. The flow will change based on the selection of either of these required options.

### 6. Required Fields

| Name | Type | Required |
| :--- | :--- | :--- |
| Address 1 | [Generic](../../../input/text-field.md) |  |
| Address 2 | [Generic](../../../input/text-field.md) |  |
| City | [Generic](../../../input/text-field.md) |  |
| State | Dropdown |  |
|  Phone Number | [Phone](../../../input/phone-number.md) | **Yes** |
|  Lease Type | Dropdown | **Yes** |
| Estimated Taxes Fields \(all\) | [Amount](../../../input/currency-input.md) |  |
| Other Fee's | [Generic](../../../input/text-field.md) |  |
| Amount | [Amount](../../../input/currency-input.md) |  |

### 6. Error Messages

Form validation errors occur when "Save" is pressed and the input displays an Inline Error. 

* Lease type -&gt; Please select a lease type
* Phone number -&gt; Please input a valid phone number

### 7. Code

{% embed url="https://iframe.html?id=forms-prelim-form--primary&args=" %}

{% embed url="https://storybook.carputty.com/axle/iframe.html?id=forms-prelim-form--primary&args=&globals=backgrounds.grid:false&viewMode=story" %}

### 8. Backend service

* [ ] Fee's Calcuation Service needed
* [ ] Form output 



