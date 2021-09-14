# Social Security

{% embed url="https://www.figma.com/file/w78ZiMR2USgl1CwXVrcxXv/?node-id=1%3A2" %}

**Intended use**  
To capture a user's Social Security Number in a secure and consistent manner

**Parent Object**  
[Ant Design's Input **-** Password component](https://ant.design/components/input/)

**Functional Requirements**

* The field must only accept numbers
* Leading zeroes must be supported.
* The SSN should be treated as a string and not a number.
* The field must be limited to 10 characters
* Dashes should automatically be inserted after the third and fifth digits
* MSAs should be able to edit individual digits without the auto-formatting getting in the way.

**Error Messages**

* Please enter a valid Social Security Number
  * Shown if non-numerical characters are entered, or if there are fewer than 10 digits.
* This field is required
  * Shown if the field is marked required and is not correctly filled in when the form is submitted

\*\*\*\*

