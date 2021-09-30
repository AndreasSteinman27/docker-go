# Task Table

{% embed url="https://www.figma.com/file/w78ZiMR2USgl1CwXVrcxXv/?node-id=718%3A25252" %}

{% tabs %}
{% tab title="Usage" %}
### Usage

The task table provides a central view for related MSA tasks.
{% endtab %}

{% tab title="Parent Object" %}

{% endtab %}

{% tab title="Requirements" %}
* The task name, status, verification button, and upload button should all be present for a given task.
* [Link buttons](../button/link-button.md) should be used for the verification and upload buttons
* Tasks that are not complete and are marked as required should show an error if the user attempts to progress the stage forward.
{% endtab %}

{% tab title="Interaction" %}
### Interaction

* View/Edit brings up the [Upload/Verify Drawer](../../templates/drawer-templates/upload-verify-drawer.md).
* Start brings up [Verification Drawer](../drawer/untitled.md)
  * Once task is complete the button link turns to Edit
{% endtab %}
{% endtabs %}

### Task Table Components

| Name | Flow |
| :--- | :--- |
| [Trim](../../templates/table-templates/task-table-templates/trim.md) | Lease-buyout |
| [Direct Lease Documents](../../templates/table-templates/task-table-templates/direct-lease-documents.md) | Lease-buyout |
| [Retail Lease Documents](../../templates/table-templates/task-table-templates/retail-lease-documents.md) | Lease-buyout |
| [Disbursement](../../templates/table-templates/task-table-templates/disbursement.md) | Lease-buyout |

