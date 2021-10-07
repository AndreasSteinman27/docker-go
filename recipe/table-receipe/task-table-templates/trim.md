---
description: Variant of Task Table
---

# Trim Table

![](../../../.gitbook/assets/trim-verification-onload.png)

{% tabs %}
{% tab title="Usage" %}
### Usage

This form is used to manage uploaded documents for Autocheck and Monroney Labels. 

### Primary Component

{% page-ref page="../../../ingredients/task-tables/task-table.md" %}
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

See [Task Table](../../../ingredients/task-tables/task-table.md) for additional interactions.

### Related Components

| Component |
| :--- |
| [Autocheck Side Drawer](../../drawer/verification.md) |
| [Maroney Label Drawer](../../drawer/drawer-trim.md) |
{% endtab %}

{% tab title="Validation/Error Handling" %}
See [Task Table](../../../ingredients/task-tables/task-table.md)
{% endtab %}

{% tab title="States" %}
### **States**

### **On Load**

![](../../../.gitbook/assets/trim-verification-onload%20%281%29.png)

### **Image Uploaded & Needs to be Verified**

![OnLoad](../../../.gitbook/assets/side-drawer-upload-maroney-label-drawer%20%281%29.png)

![Maroney Label Uploaded](../../../.gitbook/assets/side-drawer-upload-maroney-label-drawer-complete%20%281%29.png)

### Complete Status

![](../../../.gitbook/assets/trim-verifcation-status-change.png)
{% endtab %}

{% tab title="Code" %}
\*\*\*\*
{% endtab %}
{% endtabs %}







\*\*\*\*



