---
description: Variant 1 of a base form.
---

# Preliminary & Final Statements

**Description**

The Preliminary Statement Form component has 3 active states and 1 Skeleton loading state. **Only the required fields** must be complete in order to save.

**Child Elements**

[Generic Text Field](), [Currency input](../input/currency-input.md), Header, [Primary button](../button/), [Secondary button](../button/secondary-button.md)

## 1. States

#### On-First Load

{% embed url="https://www.figma.com/file/w78ZiMR2USgl1CwXVrcxXv/?node-id=2%3A4766" %}

#### Saved State 

In this state, values are submitted and calculated. For instance, taxes are generated and sent to the Right Sidebar

 This state can only be edited by pressing the Edit button.

{% embed url="https://www.figma.com/file/w78ZiMR2USgl1CwXVrcxXv/?node-id=2%3A10825" caption="" %}

#### Edit

This state allows for MSA's to edit their information and resubmit it. Discard will keep values from the previous save.

{% embed url="https://www.figma.com/file/w78ZiMR2USgl1CwXVrcxXv/?node-id=2%3A11583" %}

## 2. Generate Statements

Creates a PDF report and a valuation calculation that is needed to Create a Final Valuation. This PDF is able to be downloaded on the member's side and will be an object in the documents section.

## 3. Flow

Based on the input choice of the direct or retail lease options. The flow will change based on the selection of either of these required options.

## 4. Required Fields

In order to save the form, you must complete the following fields. The others are optional.

| Name | Type | Required |
| :--- | :--- | :--- |
| Address 1 | [Generic]() |  |
| Address 2 | [Generic]() |  |
| City | [Generic]() |  |
| State | [Dropdown](../dropdown.md) |  |
|  Phone Number | [Phone](../input/phone-number.md) | **Yes** |
|  Lease Type | [Dropdown](../dropdown.md) | **Yes** |
| Estimated Taxes Fields \(all\) | [Currency](../input/currency-input.md) |  |
| Other Fee's | [Generic]() |  |

## 5. Error Messages

Form validation errors occur when "Save" is pressed and the input displays an Inline Error. 

* Lease type -&gt; Please select a lease type
* Phone number -&gt; Please input a valid phone number

## 6. Code

{% embed url="https://iframe.html?id=forms-prelim-form--primary&args=" %}

{% embed url="https://storybook.carputty.com/axle/iframe.html?id=forms-prelim-form--primary&args=&globals=backgrounds.grid:false&viewMode=story" %}

## 7. Backend service

```text
Code goes here - test me please
```

**\(needed\)**



