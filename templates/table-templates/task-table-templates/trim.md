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

{% page-ref page="../../../components/task-tables/task-table.md" %}
{% endtab %}

{% tab title="Requirements" %}
#### Fields

Fields marked as Required must be completed and validated to submit the form attached to this table.

| Name | Image Required | Verification Required |
| :--- | :--- | :--- |
| Autocheck | **Yes** | Yes |
| Monroney Label | **Yes** | Yes |

### Related Components

| Component |
| :--- |
| [Autocheck Side Drawer](../../../components/drawer/verification.md) |
| [Maroney Label Drawer](../../drawer-templates/drawer-trim.md) |
{% endtab %}

{% tab title="Actions" %}
### Autocheck

* The View/Edit link opens the [Upload/Verify Drawer](../../drawer-templates/upload-verify-drawer.md). If upload is marked as verified, the Autocheck status should be updated to Complete. If upload is marked as failed, the Autocheck status should be updated to Failed.

### Monroney Label

* The Verify link opens the [Monroney Label Drawer](../../drawer-templates/drawer-trim.md). Once the trim has been verified, the status should change from Incomplete to Complete
  * It is impossible to "fail" a Monroney Label.
{% endtab %}

{% tab title="States" %}
### **States**

### **On Load**

{% embed url="https://www.figma.com/file/w78ZiMR2USgl1CwXVrcxXv/?node-id=851%3A25917" %}

### **Image Uploaded & Needs to be Verified**

{% embed url="https://www.figma.com/file/w78ZiMR2USgl1CwXVrcxXv/?node-id=2138%3A16035" %}

### Complete Status

{% embed url="https://www.figma.com/file/w78ZiMR2USgl1CwXVrcxXv/?node-id=876%3A25901" %}
{% endtab %}

{% tab title="Error Handling" %}

{% endtab %}

{% tab title="Code" %}
\*\*\*\*
{% endtab %}
{% endtabs %}





\*\*\*\*



