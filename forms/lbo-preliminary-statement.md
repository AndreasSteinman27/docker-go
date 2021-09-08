# LBO Preliminary Statement

LBO Preliminary Statement

The Preliminary Statement Form component has 3 active states and 1 Skeleton loading state. **Only the required fields** must be complete in order to save.

### 1. Empty State 

This the state that is on first load.

{% embed url="https://www.figma.com/file/w78ZiMR2USgl1CwXVrcxXv/?node-id=2%3A4766" %}

### 2. Saved State 

In this state, values are submitted and calculated. For instance, taxes are generated and sent to the Right Sidebar

 This state can only be edited through pressing the Edit button.

{% embed url="https://www.figma.com/file/w78ZiMR2USgl1CwXVrcxXv/?node-id=2%3A10825" caption="" %}

### 3. Edit State

This state allows for MSA's to edit their information and resubmit it. Discard will keep values from the previous save.

{% embed url="https://www.figma.com/file/w78ZiMR2USgl1CwXVrcxXv/?node-id=2%3A11583" %}

### 4. Generate Preliminary Statement

Creates a PDF report and a valuation calculation that is need to Create a Final Valuation. This PDF is able to be downloaded on the memebers side and will be an object in the documents section.

### 5. Branches in MSA Selection

Based on the input choice of the direct or retail lease options. The flow will change based on this options

### 6. Branches in MSA Selection

{% tabs %}
{% tab title="Requirements" %}
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



