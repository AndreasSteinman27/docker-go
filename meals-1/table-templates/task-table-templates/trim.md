---
description: Variant of Task Table
---

# Trim Table

{% embed url="https://www.figma.com/file/w78ZiMR2USgl1CwXVrcxXv/?node-id=851%3A25917" %}

{% tabs %}
{% tab title="Usage" %}
### Usage

This form is used to manage uploaded documents for Autocheck and Monroney Labels. 

### Primary Component

{% page-ref page="../../../receips/task-tables/task-table.md" %}
{% endtab %}

{% tab title="Functional Requirements" %}
#### Fields

Fields marked as Required must be completed and validated to submit the form attached to this table.

| Name | Drawer Type |  Required |
| :--- | :--- | :--- |
| Autocheck | Upload/Verify | Yes |
| Monroney Label | Monroney Label | Yes |

### Interactions

### Monroney Label

* The Verify link opens the [Monroney Label Drawer](../../drawer/drawer-trim.md). Once the trim has been verified, the status should change from Incomplete to Complete
  * It is impossible to "fail" a Monroney

See [Task Table](../../../receips/task-tables/task-table.md) for additional interactions.

### Related Components

| Component |
| :--- |
| [Autocheck Side Drawer](../../../receips/drawer/verification.md) |
| [Maroney Label Drawer](../../drawer/drawer-trim.md) |
{% endtab %}

{% tab title="Validation/Error Handling" %}
See [Task Table](../../../receips/task-tables/task-table.md)
{% endtab %}

{% tab title="Designs" %}
### **States**

### **On Load**

{% embed url="https://www.figma.com/file/w78ZiMR2USgl1CwXVrcxXv/?node-id=851%3A25917" %}

### **Image Uploaded & Needs to be Verified**

{% embed url="https://www.figma.com/file/w78ZiMR2USgl1CwXVrcxXv/?node-id=2138%3A16035" %}

### Complete Status

{% embed url="https://www.figma.com/file/w78ZiMR2USgl1CwXVrcxXv/?node-id=876%3A25901" %}
{% endtab %}

{% tab title="Code" %}
\*\*\*\*
{% endtab %}
{% endtabs %}





\*\*\*\*



