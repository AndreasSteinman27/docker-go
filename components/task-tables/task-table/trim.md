# Trim

{% embed url="https://www.figma.com/file/w78ZiMR2USgl1CwXVrcxXv/?node-id=851%3A25917" %}

{% tabs %}
{% tab title="Usage" %}
### Usage

This form is used to manage uploaded documents for Autocheck and Monroney Labels. 
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
| [Autocheck Side Drawer](../../drawer/untitled.md) |
| [Maroney Label Drawer](../../../templates/drawer-templates/drawer-trim.md) |
{% endtab %}

{% tab title="Actions" %}
### Autocheck

* Upload opens the [Upload Drawer](../../drawer/upload.md). If the upload was successful, the Upload link should be replaced with Verify.
* The Verify link opens the [Upload/Verify Drawer](../../../templates/drawer-templates/upload-verify-drawer.md). If upload is marked as verified, the Autocheck status should be updated to Complete. If upload is marked as failed, the Autocheck status should be updated to Failed, and the Upload link should appear to the right of the Verify button.

### Monroney Label

* Upload opens the [Upload Drawer](../../drawer/upload.md). If the upload was successful, the Verify link should appear to the left of the Upload link.
* The Verify link opens the [Monroney Label Drawer](../../../templates/drawer-templates/drawer-trim.md). Once the trim has been verified, the status should change from Incomplete to Complete
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



