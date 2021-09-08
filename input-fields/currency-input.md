# Amount

**Intended use**  
This field is intended to be used for currency values.

**Parent Object**  
Ant Design's InputNumber: [https://ant.design/components/input-number/](https://ant.design/components/input-number/)

**Functional Requirements**

* The input field should only accept numeric values, including a single decimal point.
* Values with greater precision than two decimal points are not allowed.
* A dollar sign should appear in the pre-tab of the input to indicate that it is a monetary value.
* Amounts that exceed the maximum and minimum values for a field should show an error on screen until the value is corrected. Forms containing this component should treat the component as unfilled in this state.
* The field should use a currency-styled formatting, with the following behaviors applied when the field is left:
  * Two decimals are always shown, even if no fractional numbers are present.
  * Commas are added every three characters.
  * Trailing Zeros are to be removed.
* The field should not respond to arrow-key or scroll wheel interactions.

   


