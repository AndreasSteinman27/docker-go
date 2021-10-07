---
description: Variant of Task Table
---

# Disbursement Table

{% embed url="https://www.figma.com/file/w78ZiMR2USgl1CwXVrcxXv/?node-id=624%3A24574" %}

{% tabs %}
{% tab title="Overview" %}
### Usage

Disbursement is the final section for the lease-buyout flow and signfies to OPS that the application is done. 

### Base Component

{% page-ref page="../../../receips/task-tables/task-table.md" %}
{% endtab %}

{% tab title="Requirements" %}
### Fields

| Name | Drawer Type | Required |
| :--- | :--- | :--- |
| Verbal Insurance Policy Check | Verification | Yes |
| Title Application | Upload/Verify | Yes |
| Bill of Sale | Upload/Verify | Yes |
| Insurance Loss Payee Verification | Upload/Verify | Yes |
| Insurance Coverage Verification | Upload/Verify | No |
| Engagement Letter | Upload/Verify | No |

### Interactions

* The Finish Transaction button should be hidden until all required tasks are completed.
* The Finish Transaction button marks the Asset record status as Ready To Fund.
  * An email is sent to Operations indicating that the vehicle is ready to fund.
  * Another email is sent to the member with an update on their loan status.
* All other Interactions are inherited from the [Task Table component](../../../receips/task-tables/task-table.md).
{% endtab %}

{% tab title="Validation/Error Handling" %}
See [Task Table](../../../receips/task-tables/task-table.md)
{% endtab %}

{% tab title="Designs" %}
### States

### On-loading

{% embed url="https://www.figma.com/file/w78ZiMR2USgl1CwXVrcxXv/?node-id=624%3A24574" %}

### Status Changes

{% embed url="https://www.figma.com/file/w78ZiMR2USgl1CwXVrcxXv/?node-id=2138%3A16338" %}

### Completed 

{% embed url="https://www.figma.com/file/w78ZiMR2USgl1CwXVrcxXv/?node-id=640%3A24741" %}
{% endtab %}

{% tab title="Code" %}

{% endtab %}
{% endtabs %}



