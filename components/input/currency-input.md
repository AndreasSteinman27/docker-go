---
description: Used for currency and monetary values.
---

# Currency

{% embed url="https://www.figma.com/file/w78ZiMR2USgl1CwXVrcxXv/?node-id=2%3A4522" %}

{% tabs %}
{% tab title="Usage" %}
**Usage**

This field is intended to be used for currency values
{% endtab %}

{% tab title="Parent Object" %}
**Parent Object**
{% endtab %}

{% tab title="Requirements" %}
**Requirements**
{% endtab %}

{% tab title="Error Handling" %}
**Error Handling**
{% endtab %}

{% tab title="Code" %}
**Code**
{% endtab %}
{% endtabs %}

**Intended use**  
This field is intended to be used for currency values.

**Parent Object**  
Ant Design's [**Input with Pre-fix**](https://ant.design/components/input/)

**Functional Requirements**

* The input field only accepts numeric values
*  A single decimal point can be accepted 
* Only two decimal points are allowed
* A dollar sign should appear in the pre-tab of the input to indicate that it is a monetary value.
* Amounts that exceed the maximum and minimum values for a field should show an error or caution on screen until the value is corrected.
*  Forms containing this component should treat the component as unfilled in this state.
* The field should use a currency-styled formatting, with the following behaviors applied when the field is left:
  * Two decimals are always shown, even if no fractional numbers are present.
  * Commas are added every three characters.
  * Leading Zeros are to be removed.
* The field should not respond to arrow-key or scroll wheel interactions.

**Error Messages**

* Please enter an amount

\*\*\*\*

