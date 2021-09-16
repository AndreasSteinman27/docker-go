# Zipcode

{% embed url="https://www.figma.com/file/w78ZiMR2USgl1CwXVrcxXv/?node-id=317%3A23573" %}

{% tabs %}
{% tab title="Usage" %}
**Usage**

This field accepts zipcodes as parts of a users address
{% endtab %}

{% tab title="Parent Object" %}
**Parent Object**

[Ant Design's Input](https://ant.design/components/input/)
{% endtab %}

{% tab title="Requirements" %}
**Requirements**

* Accepts digits only
* Accepts 5 characters
* Accepts additional 4 characters with formatting 12345-1234
* Input can be type OR pasted
* Non-existing 5 digits Zip results in error messaging
  * ZIP Codes must be either 5 or 5+4 digits. If the field is blank, or has either less than 5 or less than 9 digits are entered, an error should appear.
{% endtab %}

{% tab title="Attributes Format" %}
**Attributes Format**

```text
12345=>12345
12345-1234=>12351234
```
{% endtab %}

{% tab title="Error Handeling" %}
**Error Handeling**

* Please enter a valid zipcode
{% endtab %}

{% tab title="Code" %}
**Code**
{% endtab %}
{% endtabs %}

\*\*\*\*

