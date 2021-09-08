# LBO Preliminary Statement

LBO Preliminary Statement

{% tabs %}
{% tab title="Requirements" %}
The **Preliminary Statement Form** component has 3 active states and 1 Skeleton loading state. The required fields must be complete and the documents uploaded to progress to Final Statement in the Navigation Bar.

**Actions/States**

* On First Load / Default

  * an empty state in which the form has no values



* Saved State
  * The Right Sidebar is based on values when "saved" is pressed
  * Fee’s are generated when inputed



* Editing Sate
  * MSA can change text fields
  * Save overrides previous values in the database



* Generate Preliminary
  * Generate valuation is placed in "Archive" as an object
  * Preliminary Valuation is Calculated based on current values imputed

**Direct & Retail Lease Options**

* Direct and Retail lease has 2 types of fields which provide different document and text requirements
* Direct Lease ask for additional documentation
* Retail Lease provides different contact information

**Preliminary LBO Text Fields**

| Text Field / Dropdown | Required |
| :--- | :--- |
| Address 1 |  |
| Address 2 |  |
|  City |  |
| State |  |
|  Zip Code |  |
| [ Phone Number](../input-fields/phone-number.md) | Yes |
|  Lease Type | Yes |
| Estimated Taxes  |  |
{% endtab %}

{% tab title="Design" %}
{% embed url="https://www.figma.com/file/LAf8e0JKG93H5LxR9Asfbu/?node-id=1594%3A156118" caption="Lease Buyout Component" %}

{% embed url="https://www.figma.com/file/LAf8e0JKG93H5LxR9Asfbu/?node-id=1594%3A147559" %}
{% endtab %}

{% tab title="Error Messages" %}
**Inline Error Messages \(onSubmit\)**

* Please input a valid phone number
* Please select a lease type
{% endtab %}

{% tab title="Code" %}

{% endtab %}
{% endtabs %}



