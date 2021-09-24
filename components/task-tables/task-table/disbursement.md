---
description: Variant
---

# Disbursement

{% embed url="https://www.figma.com/file/w78ZiMR2USgl1CwXVrcxXv/?node-id=624%3A24574" %}

{% tabs %}
{% tab title="Usage" %}
### Usage

Disbursment is the final section for the lease-buyout flow and signfies to OPS that the application is done. 
{% endtab %}

{% tab title="Requirements" %}
### Requirements

* Title Application, Bill of Sales, Insurance Loss Payee Verification, Insurance Coverage Verification, Engagment Letter uses [Upload/Verify Side Drawer](../../drawer/upload-verify-drawer.md).
* Verbal Insurance Policy uses
{% endtab %}

{% tab title="Action" %}
* For all tasks except Verbal Insurance Policy Check:
  * When the Upload button is clicked it brings up the [Side Drawer - Upload](../../drawer/upload.md)
    * When a file is uploaded sucessfully the Uploaded status and Icon appears \(see States\)
    * When a task is complete, the badge is marked as complete
* For Verbal Insurance Policy Check:
  * When Verify is clicked, the [Verification Drawer](../../drawer/untitled.md) appears. 
    * If the policy is successfully verified, the badge should be marked as complete.
    * If the policy is not verified, the badge should remain as incomplete.
* Finish Transaction marks the Asset record status as Ready To Fund.
  * An email is sent to Operations indicating that the vehicle is ready to fund.
  * Another email is sent to the member with an update on their loan status.
{% endtab %}

{% tab title="States" %}
### States

### On-loading

{% embed url="https://www.figma.com/file/w78ZiMR2USgl1CwXVrcxXv/?node-id=624%3A24574" %}

### Status Changes

{% embed url="https://www.figma.com/file/w78ZiMR2USgl1CwXVrcxXv/?node-id=2138%3A16338" %}

### Completed 

{% embed url="https://www.figma.com/file/w78ZiMR2USgl1CwXVrcxXv/?node-id=640%3A24741" %}
{% endtab %}
{% endtabs %}



