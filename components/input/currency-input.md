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

[Ant Design's Input with Pre-fix](https://ant.design/components/input/)
{% endtab %}

{% tab title="Requirements" %}
**Requirements**

* The input field only accepts numeric values
*  A zero or a single decimal point can be accepted 
* Only two decimal points are allowed
* A dollar sign should appear in the pre-tab of the input to indicate that it is a monetary value.
* The field should not respond to arrow-key or scroll wheel interactions.

**Placeholder**

* none
{% endtab %}

{% tab title="Validation" %}
* Amounts that exceed the maximum and minimum values for a field should show an error or caution on screen until the value is corrected.
* The field should use a currency-styled formatting, with the following behaviors applied when the field is left:
  * Two decimals are always shown, when entered.
  * Commas are added every three characters.
  * Leading Zeros are to be removed.
{% endtab %}

{% tab title="Attributes Format" %}
**Output Format**

Amounts are in decimal format with 2 places after the decimal point. For example, 32.45 is thirty two dollars and forty five cents.
{% endtab %}

{% tab title="Error Handling" %}
**Error Handling**

* Please enter an amount
{% endtab %}

{% tab title="Code" %}
**Code**
{% endtab %}
{% endtabs %}







\*\*\*\*

