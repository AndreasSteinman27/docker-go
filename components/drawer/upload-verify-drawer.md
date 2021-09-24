# Upload/Verify Drawer

{% embed url="https://www.figma.com/file/w78ZiMR2USgl1CwXVrcxXv/?node-id=1149%3A37201" %}

{% tabs %}
{% tab title="Usage" %}
### Usage

The Upload/Verify Drawer is used for components that require MSA verification of an uploaded image.
{% endtab %}

{% tab title="Requirements" %}
### Requirement

* The MSA must be able to mark the image pass/fail to Save
{% endtab %}

{% tab title="States" %}
### Onload Task Table

{% embed url="https://www.figma.com/file/w78ZiMR2USgl1CwXVrcxXv/?node-id=718%3A25252" %}

#### Upload

![](../../.gitbook/assets/side-drawer-upload-onload.png)

#### Upload Complete, MSA saves Image

![](../../.gitbook/assets/side-drawer-upload-complete.png)

#### Upload/Verify Side Drawer Appears OnLoad

![](../../.gitbook/assets/verfiy-onload.png)

#### Pass

![](../../.gitbook/assets/verfiy-pass-image.png)

#### Fail

![](../../.gitbook/assets/verfiy-fail-image.png)

### Complete

{% embed url="https://www.figma.com/file/w78ZiMR2USgl1CwXVrcxXv/?node-id=2699%3A32068" %}
{% endtab %}

{% tab title="Interactions" %}
### Interactions

### Visual Inspection

* If the MSA selects Pass on the Visual Inspection dropdown, the Save button should be activated.
* If the MSA selects Fail on the Visual Inspection dropdown, the Save button should be activated. 

### Save Button

* The Save button's default state is inactive.
  * If it is clicked while inactive, nothing should happen.
* If the Save Button is clicked while active, the associated document should be updated with the appropriate status, and the drawer should close:
  * Verified if the Visual Inspection is Pass
  * Failed if the Visual Inspection is Fail

### Discard and Close Buttons:

* If Discard or Close is clicked, the drawer should close, and any entered values should be discarded.
{% endtab %}
{% endtabs %}

### Tasks that use this Drawer Component

| Name | Task Table |
| :--- | :--- |
| Driver's License Front | [Direct Lease Documents](../task-tables/task-table/direct-lease-documents.md) |
| Driver's License Back | [Direct Lease Documents](../task-tables/task-table/direct-lease-documents.md) |
| Proof of Insurance | [Direct Lease Documents](../task-tables/task-table/direct-lease-documents.md) |
| Notarized Attorney | [Direct Lease Documents](../task-tables/task-table/direct-lease-documents.md) |
| Bill of Sale | [Disbursement](../task-tables/task-table/disbursement.md) |
| Title Application | [Disbursement](../task-tables/task-table/disbursement.md) |
| Insurance Coverage Verification | [Disbursement](../task-tables/task-table/disbursement.md) |
| Insurance Loss Payee Verification | [Disbursement](../task-tables/task-table/disbursement.md) |
| Engagement Letter | [Disbursement](../task-tables/task-table/disbursement.md) |

