# Task Table

{% embed url="https://www.figma.com/file/w78ZiMR2USgl1CwXVrcxXv/?node-id=718%3A25252" %}

{% tabs %}
{% tab title="Usage" %}
### Overview

The Table is used for tasks that require documents to be uploaded and verified.

### Templates that use this component

| Name | Flow |
| :--- | :--- |
| [Trim](../../recipe/table-receipe/task-table-templates/trim.md) | Lease-buyout |
| [Direct Lease Documents](../../recipe/table-receipe/task-table-templates/direct-lease-documents.md) | Lease-buyout |
| [Retail Lease Documents](../../recipe/table-receipe/task-table-templates/retail-lease-documents.md) | Lease-buyout |
| [Disbursement](../../recipe/table-receipe/task-table-templates/disbursement.md) | Lease-buyout |
{% endtab %}

{% tab title="Functional Requirements" %}
### Requirements

* The task name and Status should be present for all tasks
* Tasks associated with a file should show a View button that links to the uploaded file if an uploaded file is present.
* All tasks should have a single button to upload, verify, or edit the file/verification for a task. The button should change name depending on the state of the object.
* [Link buttons](../button/link-button.md) should be used for the verification and upload buttons
* Tasks that are not complete and are marked as required should show an error if the user attempts to progress the stage forward.

### Interactions

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

{% tab title="Error Handling" %}
* This task is required.
{% endtab %}

{% tab title="Designs" %}

{% endtab %}

{% tab title="Code" %}

{% endtab %}
{% endtabs %}



### Task Table Components

| Name | Flow |
| :--- | :--- |
| [Trim](../../recipe/table-receipe/task-table-templates/trim.md) | Lease-buyout |
| [Direct Lease Documents](../../recipe/table-receipe/task-table-templates/direct-lease-documents.md) | Lease-buyout |
| [Retail Lease Documents](../../recipe/table-receipe/task-table-templates/retail-lease-documents.md) | Lease-buyout |
| [Disbursement](../../recipe/table-receipe/task-table-templates/disbursement.md) | Lease-buyout |

