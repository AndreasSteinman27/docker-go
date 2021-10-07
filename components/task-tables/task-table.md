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

### Upload/Edit Button

* This button should always be present.
* For tasks with files, the button should bring up the View/Upload Drawer, for tasks without files \(verification only\), the button should bring up the Verification Drawer.
  * Certain tasks, such as odometer and Monroney, have exceptions, refer to the template for these.
* If the task does have an associated file, and the file is not present, the button should say Upload.
  * Non-file tasks should start at Verify.
* If the task's verification has not been set, the button should say Verify.
* If the task has been marked as Verified, the button should say Edit.

### View Button

* The View button should appear for tasks that are associated with a file, that also have a file uploaded
* The button should link to the file in question.
* If no file is present for a task, the button should not appear.
{% endtab %}
{% endtabs %}

### Task Table Components

| Name | Flow |
| :--- | :--- |
| [Trim](../../templates/table-templates/task-table-templates/trim.md) | Lease-buyout |
| [Direct Lease Documents](../../templates/table-templates/task-table-templates/direct-lease-documents.md) | Lease-buyout |
| [Retail Lease Documents](../../templates/table-templates/task-table-templates/retail-lease-documents.md) | Lease-buyout |
| [Disbursement](../../templates/table-templates/task-table-templates/disbursement.md) | Lease-buyout |

